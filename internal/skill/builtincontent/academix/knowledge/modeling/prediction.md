# 预测类模型详细指南

## 选型原则

先看数据量，再看数据特征：
- 4~15 个数据点 → 灰色预测
- 10~50 个 → 指数平滑
- 50+ 个、单变量 → ARIMA
- 50+ 个、有季节性 → SARIMA
- 100+ 个、非线性 → XGBoost/随机森林
- 1000+ 个、长序列 → LSTM

| 模型 | 数据量 | 场景 | 限制 |
|------|--------|------|------|
| GM(1,1) 灰色 | 4~15 | 短期指数增长 | 不适合长期/震荡 |
| 线性回归 | 30+ | 线性关系 | 无法非线性 |
| 岭/Lasso | 30+ | 多重共线性 | Lasso 强相关不稳定 |
| ARIMA | 50+ | 平稳单变量时序 | 不能多变量 |
| SARIMA | 50+ | 季节性时序 | 周期需预指定 |
| 指数平滑 | 10+ | 中短期少数据 | 长期/非线性差 |
| 随机森林 | 100+ | 非线性多特征 | 外推弱 |
| XGBoost | 100+ | 非线性表格数据 | 需调参 |
| LSTM | 1000+ | 长序列非线性 | 数据少效果差 |

## ARIMA 详细指南

### 怎么做
1. **平稳性检验**：ADF 单位根检验。不平稳 → 差分
2. **定阶**：ACF/PACF 图，或 AIC/BIC 准则
3. **拟合**：`statsmodels.tsa.arima.ARIMA`
4. **诊断**：残差应为白噪声（Ljung-Box 检验）
5. **预测**：给出预测值和置信区间

### 常见错误
- **不检验平稳性**：直接拟合非平稳数据会得到虚假回归
- **差分过度**：d 通常为 0、1、2，不要超过 2
- **不检查残差**：残差不是白噪声说明模型没捕捉到信息
- **忽略季节性**：有季节性必须用 SARIMA，不是 ARIMA

### 代码模板
```python
from statsmodels.tsa.arima.model import ARIMA
from statsmodels.tsa.stattools import adfuller

# 1. 平稳性检验
adf_result = adfuller(data)
print(f"ADF p-value: {adf_result[1]}")
# p < 0.05 → 平稳；否则差分

# 2. 拟合
model = ARIMA(data, order=(p, d, q))
fitted = model.fit()

# 3. 诊断
print(fitted.summary())
# 检查 Ljung-Box p-value > 0.05

# 4. 预测
forecast = fitted.forecast(steps=10)
```

## 灰色预测 GM(1,1) 详细指南

### 什么时候用
只有 4~15 个数据点，且数据呈单调递增或近似指数增长趋势。

### 常见错误
- **数据不适合**：震荡、递减数据不能用灰色预测
- **数据量太多**：>15 个数据点时用 ARIMA 更好
- **不检验精度**：用后验差比 C 和小误差概率 P 检验

### 代码模板
```python
import numpy as np

def gm11_predict(x0, n_predict):
    """GM(1,1) 灰色预测"""
    # 累加生成
    x1 = np.cumsum(x0)
    # 紧邻均值序列
    z1 = 0.5 * (x1[:-1] + x1[1:])
    # 最小二乘求参数
    B = np.vstack([-z1, np.ones(len(z1))]).T
    Y = x0[1:].reshape(-1, 1)
    params = np.linalg.lstsq(B, Y, rcond=None)[0]
    a, b = params[0, 0], params[1, 0]
    # 预测
    result = []
    for k in range(len(x0) + n_predict):
        result.append((x0[0] - b/a) * np.exp(-a*k) + b/a)
    # 还原
    return np.diff(result)
```

## 时序预测通用防错

- **按时间划分**：不能随机打乱，必须按时间顺序划分训练/测试
- **scaler 在训练集 fit**：不能在全集 fit，否则数据泄露
- **检查物理边界**：人口 ≥ 0，概率 ≤ 1，库存不能超上限
- **外推声明风险**：超出训练数据范围时必须在论文中声明

## 回归防错

- **VIF > 10**：严重共线，用岭回归或删除高相关特征
- **残差检查**：正态、方差齐、无自相关
- **外推风险**：超出训练范围时可信度大幅下降
