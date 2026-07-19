# 编码规范

## 可复现性
- 固定随机种子：`np.random.seed(42)`、`random.seed(42)`、`torch.manual_seed(42)`
- 记录依赖版本：`requirements.txt` 或 `pyproject.toml`
- 记录运行环境：Python 版本、OS

## 结果记录
每个子问题输出 `results/problemN_results.json`：
```json
{
  "problem": "子问题一",
  "model": "使用的模型",
  "parameters": {"关键参数": "值"},
  "results": {"核心数值": "值"},
  "constraints_check": {"约束1": {"satisfied": true, "value": "..."}},
  "seed": 42
}
```

## 常见错误
- `scipy.optimize.minimize` 只做最小化，最大化要取负
- 不等式约束方向是 `fun(x) >= 0`，容量约束最容易写反
- 最优解必须重新代入所有约束检查，不信 solver 的 `success`
- 整数变量不能停留在连续解，取整后重新验证
- 读数据后先检查编码、列名、形状、缺失值
- 预测防泄露：划分后再 fit scaler，时序不能随机打乱
- 检查矩阵条件数、收敛状态、迭代次数
