# 预测类模型详细指南

## 选型决策树

```
需要预测？
├─ 数据量多少？
│  ├─ 4~15 个 → 灰色预测
│  ├─ 10~50 个 → 指数平滑
│  ├─ 50+ 个 → ARIMA/SARIMA
│  └─ 1000+ 个 → LSTM
├─ 有季节性？
│  ├─ 有 → SARIMA
│  └─ 没有 → ARIMA
├─ 非线性？
│  ├─ 是 → XGBoost/随机森林
│  └─ 否 → 线性回归
└─ 多变量？
   ├─ 是 → VAR/XGBoost
   └─ 否 → ARIMA
```

## ARIMA 详细指南

### 什么时候用
50+ 个数据点，单变量时序，无明显季节性。

### 完整流程
1. 平稳性检验：ADF 单位根检验。p < 0.05 → 平稳
2. 差分：不平稳则差分，d 通常为 0、1、2
3. 定阶：ACF/PACF 图，或 AIC/BIC 准则
4. 拟合：`statsmodels.tsa.arima.ARIMA`
5. 诊断：残差应为白噪声（Ljung-Box 检验）
6. 预测：给出预测值和置信区间

### 防错清单（20 条）
1. 时序数据必须按时间顺序划分训练/测试，不能随机打乱
2. 标准化 scaler 必须在训练集上 fit，再 transform 测试集
3. 差分阶数 d 由 ADF 检验确定，不要凭感觉选
4. d 通常为 0、1、2，不要超过 2
5. p、q 由 ACF/PACF 图或 AIC/BIC 准则确定
6. 残差应为白噪声，Ljung-Box p-value > 0.05
7. 残差应正态分布，Shapiro-Wilk p-value > 0.05
8. 残差应无自相关，ACF 图应在置信区间内
9. 季节性必须用 SARIMA，不是 ARIMA
10. SARIMA 的季节周期必须预先指定（月度=12，季度=4）
11. 预测值必须检查物理边界（人口 ≥ 0，概率 ≤ 1）
12. 外推超出训练数据范围时必须声明风险
13. 如果 AIC/BIC 不收敛，尝试其他 (p,d,q) 组合
14. 如果拟合后残差有趋势，说明 d 不够，增加差分阶数
15. 如果拟合后残差有季节性，用 SARIMA
16. 如果数据有异常值，先处理异常值再拟合
17. 如果数据有缺失值，先插值再拟合
18. 预测步数越多，置信区间越宽，不确定性越大
19. 不要用 ARIMA 预测非平稳数据，会得到虚假回归
20. 保存模型参数，确保可复现

### 代码模板（已测试）
```python
import numpy as np
import warnings
warnings.filterwarnings('ignore')

# 纯 Python 实现 ARIMA 预测（无 statsmodels 依赖）
def arima_predict_simple(data, steps=10):
    """简化版 ARIMA 预测（使用移动平均）"""
    # 1. 计算一阶差分
    diff = np.diff(data)
    
    # 2. 计算均值和标准差
    mean_diff = diff.mean()
    std_diff = diff.std()
    
    # 3. 预测
    last_value = data[-1]
    predictions = []
    for i in range(steps):
        pred = last_value + mean_diff * (i + 1)
        predictions.append(pred)
    
    # 4. 置信区间
    confidence = []
    for i in range(steps):
        margin = 1.96 * std_diff * np.sqrt(i + 1)
        confidence.append((predictions[i] - margin, predictions[i] + margin))
    
    return predictions, confidence

# 测试
data = np.array([100, 102, 105, 103, 108, 110, 112, 115, 113, 118])
predictions, confidence = arima_predict_simple(data, steps=5)
print("预测值:", [f"{p:.1f}" for p in predictions])
print("95%置信区间:", [(f"{c[0]:.1f}", f"{c[1]:.1f}") for c in confidence])
```

## 灰色预测 GM(1,1) 详细指南

### 什么时候用
只有 4~15 个数据点，且数据呈单调递增或近似指数增长趋势。

### 防错清单（15 条）
1. 数据必须单调递增或近似指数增长
2. 震荡、递减数据不能用灰色预测
3. 数据量 >15 时用 ARIMA 更好
4. 数据量 <4 时无法使用
5. 用后验差比 C 和小误差概率 P 检验精度
6. C < 0.35 为好，C < 0.5 为合格，C < 0.65 为勉强
7. P > 0.95 为好，P > 0.8 为合格
8. 预测步数越多，精度越低
9. 不要预测超过原始数据长度的 2 倍
10. 如果原始数据有明显趋势变化，不要用灰色预测
11. 累加生成序列必须单调递增
12. 最小二乘求参数时，矩阵可能奇异，需要检查
13. 参数 a 的符号决定趋势：a < 0 递增，a > 0 递减
14. 如果 |a| > 1，模型不稳定，不要使用
15. 保存模型参数，确保可复现

### 代码模板（已测试）
```python
import numpy as np

def gm11_predict(x0, n_predict):
    """GM(1,1) 灰色预测"""
    # 1. 累加生成
    x1 = np.cumsum(x0)
    
    # 2. 紧邻均值序列
    z1 = 0.5 * (x1[:-1] + x1[1:])
    
    # 3. 最小二乘求参数
    B = np.vstack([-z1, np.ones(len(z1))]).T
    Y = x0[1:].reshape(-1, 1)
    params = np.linalg.lstsq(B, Y, rcond=None)[0]
    a, b = params[0, 0], params[1, 0]
    
    # 4. 检查参数稳定性
    if abs(a) > 1:
        print(f"警告: |a| = {abs(a):.3f} > 1，模型不稳定")
    
    # 5. 预测
    result = []
    for k in range(len(x0) + n_predict):
        result.append((x0[0] - b/a) * np.exp(-a*k) + b/a)
    
    # 6. 还原（差分）
    predicted = np.diff(result)
    
    return predicted, a, b

# 测试
x0 = np.array([100, 105, 110, 118, 125])
predicted, a, b = gm11_predict(x0, n_predict=3)
print(f"参数 a = {a:.4f}, b = {b:.4f}")
print(f"预测值: {[f'{p:.1f}' for p in predicted]}")
```

## 时序预测通用防错

### 数据处理
- 按时间划分，不能随机打乱
- scaler 在训练集 fit，再 transform 测试集
- 检查物理边界（人口 ≥ 0，概率 ≤ 1）
- 外推声明风险

### 模型验证
- 残差应为白噪声
- 残差应正态分布
- 残差应无自相关
- 预测值应在合理范围内

### 代码模板（数据划分）
```python
import numpy as np

def time_series_split(data, train_ratio=0.8):
    """时序数据划分（不打乱顺序）"""
    n = len(data)
    train_size = int(n * train_ratio)
    train = data[:train_size]
    test = data[train_size:]
    return train, test

# 测试
data = np.array([100, 102, 105, 103, 108, 110, 112, 115, 113, 118])
train, test = time_series_split(data, train_ratio=0.8)
print(f"训练集: {train}")
print(f"测试集: {test}")
```
