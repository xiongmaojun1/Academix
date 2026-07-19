# 图论与网络模型详细指南

## 算法选型

| 问题 | 算法 | 复杂度 | 注意 |
|------|------|--------|------|
| 单源最短路（非负权）| Dijkstra | O(E log V) | 负权不可用 |
| 单源最短路（含负权）| Bellman-Ford | O(VE) | 检测负权环 |
| 全对最短路 | Floyd-Warshall | O(V³) | >1000 节点不可接受 |
| 最小生成树 | Prim/Kruskal | O(E log V) | Prim 稠密，Kruskal 稀疏 |
| 最大流 | Dinic | O(V²E) | 容量必须非负 |
| 最小费用最大流 | MCMF | O(VE·SPFA) | 兼顾流量和费用 |
| TSP（精确）| 分支定界 | O(n!) | >20 节点爆炸 |
| TSP（近似）| LKH/GA | — | 近似比 |
| 二分图匹配 | 匈牙利 | O(n³) | 指派问题 |

## 最短路径详细指南

### Dijkstra
- **适用**：非负权图
- **不能用**：有负权边 → 用 Bellman-Ford
- **代码**：`heapq` 实现 O(E log V)

### Bellman-Ford
- **适用**：含负权边
- **检测负权环**：第 V 次迭代仍有松弛 → 存在负权环

### Floyd-Warshall
- **适用**：全对最短路、稠密图
- **不能用**：>1000 节点（O(V³) 太慢）

## 最大流详细指南

### 常见错误
- **无向边处理**：无向边转为两条方向相反、容量相同的有向边
- **流量守恒**：中间节点流入 = 流出
- **容量非负**：所有容量必须 ≥ 0

### 代码模板（Dinic 算法）
```python
from collections import deque

class Dinic:
    def __init__(self, n):
        self.n = n
        self.adj = [[] for _ in range(n)]
    
    def add_edge(self, u, v, cap):
        self.adj[u].append([v, cap, len(self.adj[v])])
        self.adj[v].append([u, 0, len(self.adj[u]) - 1])
    
    def max_flow(self, s, t):
        flow = 0
        while self._bfs(s, t):
            ptr = [0] * self.n
            while True:
                f = self._dfs(s, t, float('inf'), ptr)
                if not f: break
                flow += f
        return flow
```

## TSP 详细指南

### 精确算法（≤20 节点）
- 动态规划：O(n² · 2^n)
- 分支定界

### 近似算法（>20 节点）
- LKH（最近邻改进）
- 遗传算法
- 模拟退火

### 常见错误
- **子回路**：必须检查 Hamiltonian 回路约束
- **不对称 TSP**：d(i,j) ≠ d(j,i) 时需要特殊处理

## 网络流防错

- 流量守恒：中间节点流入 = 流出
- 容量约束：每条边流量 ≤ 容量
- 负权边：检查是否适合 Bellman-Ford
- 最小费用最大流：同时优化流量和费用
