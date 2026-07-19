# 编码规范

## 可复现性清单（15 条）
1. 固定随机种子：`np.random.seed(42)`、`random.seed(42)`、`torch.manual_seed(42)`
2. 记录依赖版本：`requirements.txt` 或 `pyproject.toml`
3. 记录运行环境：Python 版本、OS
4. 如果用 GPU，记录 GPU 型号和 CUDA 版本
5. 如果用多线程，记录线程数
6. 如果用随机算法，多次运行取平均
7. 如果用蒙特卡洛，报告置信区间
8. 如果用启发式，报告多次运行的结果
9. 保存模型参数，确保可复现
10. 保存数据预处理代码，确保可复现
11. 保存特征工程代码，确保可复现
12. 保存评估代码，确保可复现
13. 保存随机种子，确保可复现
14. 保存版本信息，确保可复现
15. 保存运行日志，确保可复现

## 结果记录规范

每个子问题输出 `results/problemN_results.json`：

```json
{
  "problem": "子问题一",
  "model": "使用的模型",
  "parameters": {"关键参数": "值"},
  "results": {"核心数值": "值"},
  "constraints_check": {"约束1": {"satisfied": true, "value": "..."}},
  "seed": 42,
  "runtime": "Python 3.12, Ubuntu 22.04",
  "timestamp": "2025-07-19T22:00:00"
}
```

## 常见错误清单（20 条）

### 优化问题
1. `scipy.optimize.minimize` 只做最小化，最大化要取负
2. 不等式约束方向是 `fun(x) >= 0`，容量约束最容易写反
3. 最优解必须重新代入所有约束检查，不信 solver 的 `success`
4. 整数变量不能停留在连续解，取整后重新验证
5. 如果无可行解，检查约束是否矛盾
6. 如果无界，检查是否漏了约束
7. 如果收敛慢，调整算法或初始点
8. 如果结果不合理，检查目标函数和约束

### 数据问题
9. 读数据后先检查编码、列名、形状、缺失值
10. 预测防泄露：划分后再 fit scaler，时序不能随机打乱
11. 如果数据有异常值，先处理异常值再建模
12. 如果数据有缺失值，先处理缺失值再建模
13. 如果数据有噪声，先平滑再建模
14. 如果数据量太少，考虑数据增强
15. 如果数据量太大，考虑采样

### 数值问题
16. 检查矩阵条件数，条件数太大说明矩阵病态
17. 检查收敛状态，未收敛说明算法有问题
18. 检查迭代次数，迭代太多说明算法有问题
19. 检查步长，步长太大说明算法不稳定
20. 检查精度，精度不够说明算法有问题

## 代码模板（结果记录）
```python
import json
import numpy as np
from datetime import datetime

def save_results(problem_name, model_name, parameters, results, constraints, seed=42):
    """保存结果到 JSON"""
    data = {
        "problem": problem_name,
        "model": model_name,
        "parameters": parameters,
        "results": results,
        "constraints_check": constraints,
        "seed": seed,
        "runtime": f"Python {__import__('sys').version}, {__import__('platform').platform()}",
        "timestamp": datetime.now().isoformat()
    }
    
    with open(f"results/{problem_name}_results.json", "w") as f:
        json.dump(data, f, indent=2, ensure_ascii=False)
    
    print(f"结果已保存到 results/{problem_name}_results.json")

# 测试
save_results(
    problem_name="problem1",
    model_name="Dijkstra",
    parameters={"start": 1, "end": 6},
    results={"path": [1, 4, 5, 6], "distance": 7.5},
    constraints={"path_valid": True, "distance_positive": True}
)
```

## 工具推荐

### 数据处理
- pandas：数据读取、清洗、转换
- numpy：数值计算
- scipy：科学计算

### 机器学习
- scikit-learn：传统机器学习
- xgboost：梯度提升
- lightgbm：梯度提升

### 深度学习
- pytorch：深度学习
- tensorflow：深度学习

### 可视化
- matplotlib：基础绑图
- seaborn：统计可视化
- plotly：交互式可视化

### 优化
- scipy.optimize：优化
- cvxpy：凸优化
- gurobipy：商业优化器
