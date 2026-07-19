# 评价类模型详细指南

## 选型决策树

```
需要评价？
├─ 有专家经验？
│  ├─ 有 → AHP
│  └─ 没有 → 熵权法
├─ 指标有层级关系？
│  ├─ 有 → AHP/ANP
│  └─ 没有 → TOPSIS
├─ 指标能量化？
│  ├─ 能 → TOPSIS
│  └─ 不能 → 模糊综合评价
├─ 指标高度相关？
│  ├─ 是 → PCA + TOPSIS
│  └─ 否 → 直接 TOPSIS
└─ 要评价效率？
   ├─ 是 → DEA
   └─ 否 → TOPSIS
```

## AHP 详细指南

### 什么时候用
评价对象有层级结构（目标层→准则层→方案层），且有专家可以给出两两比较判断。

### 完整流程
1. 建立层次结构（目标→准则→方案）
2. 构造判断矩阵（1-9 标度）
3. 计算权重向量（特征值法或几何平均法）
4. 一致性检验（CR < 0.1）
5. 计算综合得分
6. 灵敏度分析

### 防错清单（20 条）
1. 判断矩阵必须是正互反矩阵：a_ij = 1/a_ji
2. 对角线必须全为 1：a_ii = 1
3. 矩阵维度 >9×9 时一致性极难保证，拆成多级层次
4. CR ≥ 0.1 必须调整矩阵，不能强行使用
5. 权重向量归一化后求和必须 = 1
6. 最终得分 = 权重向量 × 方案矩阵，不是简单加法
7. AHP 只能从已有方案中选优，不能自动生成新方案
8. 判断矩阵的值必须在 1-9 之间，不能有 0 或负数
9. 一致性指标 CI = (λ_max - n)/(n - 1)，n 为矩阵维度
10. 随机一致性指标 RI 取值：1→0, 2→0, 3→0.58, 4→0.90, 5→1.12, 6→1.24, 7→1.32, 8→1.41, 9→1.45
11. 几何平均法比特征值法更稳定，推荐使用
12. 专家打分时，先让专家独立打分，再取平均
13. 如果专家打分差异大（标准差 > 2），需要重新讨论
14. 层次结构中，每个准则下的方案数不能超过 9 个
15. 如果准则数超过 7 个，考虑分组
16. 最终得分的排序可能因权重微小变化而改变，需要做灵敏度分析
17. 灵敏度分析：权重 ±10% 观察排序是否变化
18. 如果排序对某个权重特别敏感，需要特别说明
19. AHP 不适合处理定量数据，定量数据用 TOPSIS
20. AHP 的结果是相对排序，不是绝对评分

### 代码模板（已测试）
```python
import numpy as np

def ahp_weights(matrix):
    """计算 AHP 权重向量（几何平均法）"""
    n = matrix.shape[0]
    # 几何平均
    geo_mean = np.prod(matrix, axis=1) ** (1/n)
    # 归一化
    weights = geo_mean / geo_mean.sum()
    return weights

def consistency_ratio(matrix, weights):
    """一致性检验"""
    n = matrix.shape[0]
    RI = {1:0, 2:0, 3:0.58, 4:0.90, 5:1.12, 6:1.24, 7:1.32, 8:1.41, 9:1.45}
    # 计算 λ_max
    lambda_max = (matrix @ weights / weights).mean()
    # 计算 CI
    CI = (lambda_max - n) / (n - 1)
    # 计算 CR
    CR = CI / RI[n]
    return CR, CI, lambda_max

# 测试
matrix = np.array([
    [1, 3, 5],
    [1/3, 1, 3],
    [1/5, 1/3, 1]
])
weights = ahp_weights(matrix)
CR, CI, lambda_max = consistency_ratio(matrix, weights)
print(f"权重: {weights}")
print(f"CR: {CR:.4f} (< 0.1 通过)")
print(f"一致性检验: {'通过' if CR < 0.1 else '不通过'}")
```

## TOPSIS 详细指南

### 什么时候用
需要对多个方案进行综合排序，且指标之间没有层级关系。

### 完整流程
1. 构建决策矩阵（方案 × 指标）
2. 标准化（向量归一化或 min-max 归一化）
3. 加权（权重来自 AHP 或熵权法）
4. 计算正理想解和负理想解
5. 计算距离和贴近度
6. 排序

### 防错清单（15 条）
1. 必须有 2+ 方案才能使用
2. 欧氏距离对量纲敏感，必须先标准化再计算距离
3. 贴近度 C = D⁻/(D⁺+D⁻)，越接近 1 越优
4. C 值本身不是百分制得分，不要直接乘以 100 展示
5. 指标间高度相关时 TOPSIS 会重复计数，用 PCA 降维后再做
6. 正向指标用 (x-min)/(max-min)，负向指标用 (max-x)/(max-min)
7. 标准化后所有值必须在 [0, 1]
8. 如果某列全相同，标准化后全为 0，该列无区分度
9. 权重必须归一化后使用
10. 正理想解是各指标的最优值，负理想解是最差值
11. 对于正向指标，正理想解是最大值，负理想解是最小值
12. 对于负向指标，正理想解是最小值，负理想解是最大值
13. 距离计算用欧氏距离，不用曼哈顿距离
14. 如果方案数 < 3，TOPSIS 的区分度很低
15. 灵敏度分析：权重 ±10% 观察排序是否变化

### 代码模板（已测试）
```python
import numpy as np

def topsis(data, weights, directions):
    """
    data: 方案 × 指标矩阵
    weights: 权重向量（已归一化）
    directions: 指标方向（1=正向，-1=负向）
    """
    n, m = data.shape
    
    # 1. 标准化
    norm = np.zeros_like(data, dtype=float)
    for j in range(m):
        col = data[:, j]
        if directions[j] == 1:  # 正向
            norm[:, j] = (col - col.min()) / (col.max() - col.min())
        else:  # 负向
            norm[:, j] = (col.max() - col) / (col.max() - col.min())
    
    # 2. 加权
    weighted = norm * weights
    
    # 3. 正理想解和负理想解
    ideal_pos = weighted.max(axis=0)
    ideal_neg = weighted.min(axis=0)
    
    # 4. 距离
    d_pos = np.sqrt(((weighted - ideal_pos) ** 2).sum(axis=1))
    d_neg = np.sqrt(((weighted - ideal_neg) ** 2).sum(axis=1))
    
    # 5. 贴近度
    c = d_neg / (d_pos + d_neg)
    
    return c, d_pos, d_neg

# 测试
data = np.array([
    [35, 80, 85, 90, 75],
    [28, 65, 92, 80, 88],
    [42, 95, 70, 85, 82],
    [30, 72, 88, 78, 90],
])
weights = np.array([0.262, 0.167, 0.154, 0.245, 0.171])
directions = np.array([1, -1, 1, 1, 1])

c, d_pos, d_neg = topsis(data, weights, directions)
labels = ['A', 'B', 'C', 'D']
ranking = np.argsort(-c)
for rank, idx in enumerate(ranking):
    print(f"第{rank+1}名: {labels[idx]} (贴近度={c[idx]:.3f})")
```

## 熵权法详细指南

### 什么时候用
希望权重完全由数据决定，不引入主观判断。

### 防错清单（10 条）
1. 所有值相同时熵权 = 0（正常，不是错误）
2. 熵权忽略业务重要性，必要时 AHP+熵权组合
3. 标准化方式要统一正负向
4. 标准化后所有值必须在 [0, 1]
5. 如果某列有 0 值，log(0) 会出错，需要加小常数
6. 熵值 e 的范围是 [0, 1]，越接近 1 越均匀
7. 差异系数 d = 1 - e，越大越有区分度
8. 权重 = d / sum(d)，归一化后求和 = 1
9. 如果所有指标的熵值都接近 1，说明数据无区分度
10. 熵权法不适合样本数 < 5 的情况

### 代码模板（已测试）
```python
import numpy as np

def entropy_weight(data):
    """熵权法计算权重"""
    n, m = data.shape
    
    # 1. 标准化（正向化）
    norm = np.zeros_like(data, dtype=float)
    for j in range(m):
        col = data[:, j]
        norm[:, j] = (col - col.min()) / (col.max() - col.min() + 1e-10)
    
    # 2. 计算比重
    p = norm / (norm.sum(axis=0) + 1e-10)
    
    # 3. 计算熵值
    k = 1 / np.log(n)
    e = -k * (p * np.log(p + 1e-10)).sum(axis=0)
    
    # 4. 计算权重
    d = 1 - e
    w = d / d.sum()
    
    return w, e, d

# 测试
data = np.array([
    [35, 80, 85, 90, 75],
    [28, 65, 92, 80, 88],
    [42, 95, 70, 85, 82],
    [30, 72, 88, 78, 90],
])
w, e, d = entropy_weight(data)
print(f"权重: {w}")
print(f"权重和: {w.sum():.3f} (应为1.000)")
```

## 组合权重（AHP + 熵权）

当既想利用专家经验，又想利用数据信息时：
```
组合权重 = α × AHP权重 + (1-α) × 熵权
```
α 通常取 0.5，或根据专家信心调整。

### 代码模板
```python
def combined_weight(ahp_w, entropy_w, alpha=0.5):
    """组合权重"""
    return alpha * ahp_w + (1 - alpha) * entropy_w
```
