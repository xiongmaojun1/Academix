# 动力学模型详细指南

## 常见微分方程模型

| 模型 | 方程 | 应用 |
|------|------|------|
| Logistic | dx/dt = rx(1-x/K) | 种群增长、市场渗透、S 曲线 |
| SIR | dS/dt=-βSI, dI/dt=βSI-γI, dR/dt=γI | 疾病传播、谣言扩散 |
| SEIR | 增加潜伏期 E 仓室 | 有潜伏期的传染病 |
| Lotka-Volterra | dx/dt=ax-bxy, dy/dt=-cy+dxy | 捕食者-猎物、市场竞争 |
| 牛顿运动 | ma = F | 力学、弹道、振动 |

## 微分方程建模流程

1. **定义状态变量**：明确物理含义和量纲
2. **建立方程**：基于物理定律或业务逻辑
3. **确定初始条件**：t=0 时各变量取值
4. **选择求解器**：刚性系统用 Radau/BDF，非刚性用 RK45
5. **验证**：检查守恒量、步长收敛性
6. **参数拟合**：分离训练期和验证期

## 常见错误

- **状态变量无量纲**：必须标注单位（m、kg、s 等）
- **初始条件缺失**：没有初始条件无法求解
- **刚性系统用错求解器**：`solve_ivp(method='RK45')` 在刚性系统会很慢或失败，用 `'Radau'` 或 `'BDF'`
- **不检查守恒量**：SIR 模型 S+I+R 必须恒等于 N
- **步长不收敛**：步长减半后结果变化应 <1%
- **参数拟合数据泄露**：用全部数据拟合再在同期验证 = 数据泄露

## SIR 模型代码模板

```python
import numpy as np
from scipy.integrate import solve_ivp

def sir_model(t, y, beta, gamma):
    S, I, R = y
    dS = -beta * S * I
    dI = beta * S * I - gamma * I
    dR = gamma * I
    return [dS, dI, dR]

# 参数
beta, gamma = 0.3, 0.1
N = 1000
y0 = [N-1, 1, 0]  # 初始：999易感，1感染，0康复

# 求解（刚性系统用 Radau）
sol = solve_ivp(sir_model, [0, 160], y0, args=(beta, gamma),
                method='Radau', dense_output=True, max_step=0.5)

# 验证守恒量
total = sol.y[0] + sol.y[1] + sol.y[2]
print(f"守恒检查: {np.allclose(total, N)}")  # 应为 True
```

## 蒙特卡洛仿真

### 什么时候用
- 系统有随机性
- 需要估计概率分布
- 解析解困难

### 常见错误
- **采样次数不够**：通常需要 ≥10,000 次
- **不报告置信区间**：只报均值不够，要报 95% CI
- **随机种子不固定**：结果不可复现

### 代码模板
```python
import numpy as np

np.random.seed(42)
n_sim = 10000

# 示例：估计圆周率
x = np.random.uniform(-1, 1, n_sim)
y = np.random.uniform(-1, 1, n_sim)
inside = (x**2 + y**2) <= 1
pi_estimate = 4 * inside.mean()
pi_std = 4 * np.sqrt(inside.mean() * (1-inside.mean()) / n_sim)
print(f"π ≈ {pi_estimate:.4f} ± {pi_std:.4f} (95% CI)")
```

## 马尔可夫链

### 常见错误
- **转移矩阵每行和 ≠ 1**：必须归一化
- **稳态不存在**：链必须不可约且非周期
- **吸收态和瞬态混淆**：明确区分
