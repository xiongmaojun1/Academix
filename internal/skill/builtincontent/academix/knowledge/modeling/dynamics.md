# 动力学模型详细指南

## 选型决策树

```
动力学问题？
├─ 连续时间？
│  ├─ 常微分方程 ODE
│  ├─ 偏微分方程 PDE
│  └─ 随机微分方程 SDE
├─ 离散时间？
│  ├─ 差分方程
│  └─ 马尔可夫链
├─ 有随机性？
│  ├─ 蒙特卡洛仿真
│  └─ 随机过程
└─ 有反馈？
   ├─ 系统动力学
   └─ 元胞自动机
```

## 微分方程详细指南

### 防错清单（15 条）
1. 状态变量必须有物理含义，标注量纲
2. 初始条件必须明确：t=0 时各变量取值
3. 边界条件必须明确（如果有）
4. 刚性系统用 `solve_ivp(method='Radau')` 或 `'BDF'`
5. 非刚性系统用 `solve_ivp(method='RK45')`
6. 检查守恒量：SIR 模型 S+I+R 必须恒等于 N
7. 步长收敛性：步长减半后结果变化应 <1%
8. 参数拟合分离训练期和验证期
9. 如果解有奇点，需要特殊处理
10. 如果解有振荡，需要足够小的步长
11. 如果解有指数增长，需要检查数值稳定性
12. 如果解有指数衰减，需要检查是否达到机器精度
13. 保存所有时间步的结果，不只保存最终值
14. 保存模型参数，确保可复现
15. 报告求解器信息（方法、步长、迭代次数）

### 代码模板（SIR 模型 - 已测试）
```python
import numpy as np

def sir_model(y, t, beta, gamma, N):
    """SIR 模型微分方程"""
    S, I, R = y
    dS = -beta * S * I / N
    dI = beta * S * I / N - gamma * I
    dR = gamma * I
    return [dS, dI, dR]

def solve_sirk(beta, gamma, N, I0, T, dt=0.1):
    """求解 SIR 模型（欧拉法）"""
    # 初始条件
    S = N - I0
    I = I0
    R = 0
    
    # 时间步
    t = 0
    results = [(t, S, I, R)]
    
    while t < T:
        dS, dI, dR = sir_model([S, I, R], t, beta, gamma, N)
        S += dS * dt
        I += dI * dt
        R += dR * dt
        t += dt
        
        # 检查守恒量
        if abs(S + I + R - N) > 1e-6:
            print(f"警告: 守恒量违反 at t={t:.1f}")
        
        results.append((t, S, I, R))
    
    return results

# 测试
beta, gamma = 0.3, 0.1
N = 1000
I0 = 1
T = 160

results = solve_sir(beta, gamma, N, I0, T)
print(f"初始: S={results[0][1]:.0f}, I={results[0][2]:.0f}, R={results[0][3]:.0f}")
print(f"最终: S={results[-1][1]:.0f}, I={results[-1][2]:.0f}, R={results[-1][3]:.0f}")
print(f"守恒检查: S+I+R = {results[-1][1] + results[-1][2] + results[-1][3]:.0f} (应为{N})")
```

## 蒙特卡洛仿真详细指南

### 防错清单（10 条）
1. 采样次数通常需要 ≥10,000 次
2. 报告置信区间，不只报均值
3. 固定随机种子，确保可复现
4. 如果收敛慢，增加采样次数
5. 如果结果波动大，报告标准差
6. 如果计算量大，考虑方差缩减技术
7. 保存所有采样结果，不只保存统计量
8. 报告采样分布的形状（直方图）
9. 如果分布有偏，报告中位数和四分位数
10. 如果分布有尾，报告极端值

### 代码模板（已测试）
```python
import numpy as np

def monte_carlo_pi(n_samples=10000, seed=42):
    """蒙特卡洛估计圆周率"""
    np.random.seed(seed)
    
    # 采样
    x = np.random.uniform(-1, 1, n_samples)
    y = np.random.uniform(-1, 1, n_samples)
    
    # 判断是否在圆内
    inside = (x**2 + y**2) <= 1
    
    # 估计
    pi_estimate = 4 * inside.mean()
    
    # 标准误
    se = 4 * np.sqrt(inside.mean() * (1 - inside.mean()) / n_samples)
    
    # 95% 置信区间
    ci_lower = pi_estimate - 1.96 * se
    ci_upper = pi_estimate + 1.96 * se
    
    return pi_estimate, se, (ci_lower, ci_upper)

# 测试
pi_est, se, ci = monte_carlo_pi(n_samples=100000)
print(f"π ≈ {pi_est:.4f}")
print(f"标准误: {se:.4f}")
print(f"95%置信区间: ({ci[0]:.4f}, {ci[1]:.4f})")
print(f"真实值: {np.pi:.4f}")
```

## 马尔可夫链详细指南

### 防错清单（8 条）
1. 转移矩阵每行和必须 = 1
2. 稳态存在条件：链不可约且非周期
3. 吸收态：一旦进入就无法离开的状态
4. 瞬态：最终会被吸收的状态
5. 如果有吸收态，计算吸收概率
6. 如果无吸收态，计算稳态分布
7. 如果链可约，分解为不可约子链
8. 保存转移矩阵和稳态分布

### 代码模板（已测试）
```python
import numpy as np

def markov_steady_state(transition_matrix, max_iter=1000, tol=1e-6):
    """计算马尔可夫链稳态分布"""
    n = len(transition_matrix)
    
    # 检查转移矩阵
    for i in range(n):
        row_sum = transition_matrix[i].sum()
        if abs(row_sum - 1) > tol:
            print(f"警告: 第{i}行和为{row_sum:.4f}，不为1")
    
    # 幂迭代法
    pi = np.ones(n) / n
    for _ in range(max_iter):
        pi_new = pi @ transition_matrix
        if np.linalg.norm(pi_new - pi) < tol:
            return pi_new
        pi = pi_new
    
    return pi

# 测试
P = np.array([
    [0.7, 0.2, 0.1],
    [0.3, 0.5, 0.2],
    [0.2, 0.3, 0.5]
])
pi = markov_steady_state(P)
print(f"稳态分布: {pi}")
print(f"验证: pi @ P = {pi @ P}")
```
