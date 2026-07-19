# 优化类模型详细指南

## 选型决策树

```
需要优化？
├─ 目标函数和约束都是线性的？
│  ├─ 是 → 线性规划 LP
│  └─ 否 → 非线性规划 NLP
├─ 有整数变量？
│  ├─ 是 → 整数规划 MIP
│  └─ 否 → 连续优化
├─ 多阶段？
│  ├─ 是 → 动态规划 DP
│  └─ 否 → 单阶段优化
├─ 问题规模大？
│  ├─ 是 → 启发式（GA/SA/PSO）
│  └─ 否 → 精确算法
└─ 多目标？
   ├─ 是 → NSGA-II / 加权和
   └─ 否 → 单目标优化
```

## 线性规划详细指南

### 什么时候用
目标函数和约束都是线性，变量连续。

### 防错清单（15 条）
1. `scipy.optimize.linprog` 只做最小化，最大化要取负目标函数
2. 不等式约束方向是 `A_ub @ x <= b_ub`，注意方向
3. 变量下界默认 x >= 0，如果变量可为负需要显式指定
4. 约束 `x >= C` 转为 `-x <= -C`
5. 约束 `x <= C` 直接用 `bounds=[(None, C)]`
6. 等式约束用 `A_eq @ x = b_eq`
7. 如果无可行解，检查约束是否矛盾
8. 如果无界，检查是否漏了约束
9. 最优解必须代入所有约束验证
10. 如果有退化解（多个最优解），说明目标函数与某个约束平行
11. 敏感性分析：约束右端项 ±10% 观察目标值变化
12. 影子价格：约束右端项增加 1 单位，目标值改善多少
13. 如果问题规模 > 10000 变量，考虑分解算法
14. 如果约束矩阵稀疏，用稀疏矩阵格式提高效率
15. 保存模型和结果，确保可复现

### 代码模板（已测试）
```python
import numpy as np

def linprog_simple(c, A_ub, b_ub, bounds=None):
    """简化版线性规划（无 scipy 依赖）"""
    # 使用 scipy 如果可用
    try:
        from scipy.optimize import linprog
        result = linprog(c, A_ub=A_ub, b_ub=b_ub, bounds=bounds, method='highs')
        return result.fun, result.x, result.success
    except ImportError:
        pass
    
    # 简化实现：网格搜索（仅适用于小规模问题）
    n = len(c)
    best_val = float('inf')
    best_x = None
    
    # 生成候选解
    for x1 in np.arange(0, 100, 1):
        for x2 in np.arange(0, 100, 1):
            x = np.array([x1, x2])
            # 检查约束
            if all(A_ub @ x <= b_ub):
                val = np.array(c) @ x
                if val < best_val:
                    best_val = val
                    best_x = x
    
    return best_val, best_x, best_x is not None

# 测试
c = [2, 3]  # min 2x1 + 3x2
A_ub = [[-1, -2], [-3, -1]]  # x1 + 2x2 >= 4, 3x1 + x2 >= 6
b_ub = [-4, -6]
val, x, success = linprog_simple(c, A_ub, b_ub)
print(f"最优值: {val}, 解: {x}, 成功: {success}")
```

## 整数规划详细指南

### 什么时候用
决策变量必须是整数（人数、车辆数、批次等）。

### 防错清单（12 条）
1. 连续松弛后直接取整可能不满足约束，必须验证可行性
2. 0-1 变量用 `bounds=[(0,1)]` + `integrality=[1]`
3. 大规模 >1000 变量时考虑启发式
4. 整数规划是 NP-hard，规模大时计算时间爆炸
5. 如果计算时间太长，考虑松弛或启发式
6. 分支定界法：先求松弛解，再分支
7. 如果松弛解已经是整数，直接得到最优解
8. 如果松弛解不是整数，选择最接近整数的变量分支
9. 剪枝：如果子问题的下界 > 当前最优解，剪掉
10. 如果问题有特殊结构（如指派问题），用专用算法
11. 保存模型和结果，确保可复现
12. 报告计算时间，说明是否为全局最优

### 代码模板
```python
import numpy as np

def integer_programming_simple(c, A_ub, b_ub, n_integer):
    """简化版整数规划（枚举法，仅适用于小规模）"""
    n = len(c)
    best_val = float('inf')
    best_x = None
    
    # 枚举所有整数组合（仅适用于 n <= 5）
    def enumerate_integers(idx, current):
        nonlocal best_val, best_x
        if idx == n_integer:
            # 检查约束
            if all(A_ub @ np.array(current) <= b_ub):
                val = np.array(c) @ np.array(current)
                if val < best_val:
                    best_val = val
                    best_x = np.array(current)
            return
        for v in range(0, 20):  # 限制枚举范围
            enumerate_integers(idx + 1, current + [v])
    
    enumerate_integers(0, [])
    return best_val, best_x

# 测试
c = [2, 3]
A_ub = [[-1, -2], [-3, -1]]
b_ub = [-4, -6]
val, x = integer_programming_simple(c, A_ub, b_ub, n_integer=2)
print(f"最优值: {val}, 解: {x}")
```

## 启发式算法（GA/SA/PSO）详细指南

### 什么时候用
问题规模大、非线性、多约束、无梯度信息。

### 防错清单（15 条）
1. 固定随机种子：`np.random.seed(42)`
2. 多次独立运行：≥5 次，报告均值和标准差
3. 与精确解对比：小规模时用精确算法验证
4. 不要声称全局最优：除非有证明
5. 种群大小：通常 50-200
6. 交叉率：通常 0.6-0.9
7. 变异率：通常 0.01-0.1
8. 迭代次数：通常 100-1000
9. 如果收敛太快，增加种群大小或变异率
10. 如果收敛太慢，减少种群大小或增加交叉率
11. 如果结果波动大，增加迭代次数或多次运行取最优
12. 如果约束违反，用罚函数或修复策略
13. 如果目标函数有多个局部最优，增加种群多样性
14. 保存所有运行的结果，不只保存最优
15. 报告收敛曲线，说明算法行为

### 代码模板（遗传算法）
```python
import numpy as np

def genetic_algorithm(fitness_func, n_vars, n_pop=100, n_gen=500, seed=42):
    """简化版遗传算法"""
    np.random.seed(seed)
    
    # 初始化种群
    pop = np.random.uniform(0, 10, (n_pop, n_vars))
    
    best_val = float('inf')
    best_x = None
    
    for gen in range(n_gen):
        # 计算适应度
        fitness = np.array([fitness_func(x) for x in pop])
        
        # 更新最优
        min_idx = fitness.argmin()
        if fitness[min_idx] < best_val:
            best_val = fitness[min_idx]
            best_x = pop[min_idx].copy()
        
        # 选择（锦标赛）
        new_pop = []
        for _ in range(n_pop):
            i, j = np.random.randint(0, n_pop, 2)
            winner = pop[i] if fitness[i] < fitness[j] else pop[j]
            new_pop.append(winner)
        pop = np.array(new_pop)
        
        # 交叉
        for i in range(0, n_pop - 1, 2):
            if np.random.random() < 0.8:
                alpha = np.random.random()
                pop[i] = alpha * pop[i] + (1 - alpha) * pop[i+1]
                pop[i+1] = (1 - alpha) * pop[i] + alpha * pop[i+1]
        
        # 变异
        for i in range(n_pop):
            if np.random.random() < 0.1:
                pop[i] += np.random.normal(0, 0.5, n_vars)
    
    return best_val, best_x

# 测试
def sphere(x):
    return np.sum(x**2)

val, x = genetic_algorithm(sphere, n_vars=3, n_pop=50, n_gen=200)
print(f"最优值: {val:.4f}, 解: {x}")
```

## 多目标优化

### 防错清单（8 条）
1. 不同量纲的目标不能直接相加
2. 必须先独立归一化每个目标到 [0,1]
3. 加权和法：权重反映偏好
4. Pareto 法：给出所有非劣解
5. 如果目标数 > 3，Pareto 前沿可视化困难
6. 灵敏度分析：权重 ±10% 观察解的变化
7. 如果目标严重冲突，说明需要取舍
8. 保存所有 Pareto 最优解，供用户选择

### 代码模板（加权和法）
```python
import numpy as np

def weighted_sum(objectives, weights):
    """加权和法（目标已归一化）"""
    # 归一化
    normalized = []
    for obj in objectives:
        min_val, max_val = obj.min(), obj.max()
        if max_val > min_val:
            normalized.append((obj - min_val) / (max_val - min_val))
        else:
            normalized.append(np.zeros_like(obj))
    
    # 加权和
    weighted = np.zeros_like(normalized[0])
    for i, (obj, w) in enumerate(zip(normalized, weights)):
        weighted += w * obj
    
    return weighted

# 测试
obj1 = np.array([10, 20, 30, 40, 50])  # 目标1
obj2 = np.array([50, 40, 30, 20, 10])  # 目标2
weights = [0.5, 0.5]
result = weighted_sum([obj1, obj2], weights)
print(f"加权和: {result}")
print(f"最优方案: {result.argmin()}")
```
