# 图论与网络模型详细指南

## 选型决策树

```
图论问题？
├─ 最短路径？
│  ├─ 非负权 → Dijkstra
│  ├─ 含负权 → Bellman-Ford
│  └─ 全对最短路 → Floyd-Warshall
├─ 最大流？
│  ├─ 是 → Dinic
│  └─ 最小费用最大流 → MCMF
├─ 匹配？
│  ├─ 二分图 → 匈牙利
│  └─ 一般图 → 带花树
├─ 遍历？
│  ├─ TSP → 分支定界/LKH
│  └─ VRP → 启发式
└─ 连通性？
   ├─ 强连通 → Tarjan
   └─ 桥/割点 → Tarjan
```

## Dijkstra 详细指南

### 防错清单（12 条）
1. 只能用于非负权图，有负权边用 Bellman-Ford
2. 时间复杂度 O(E log V)，用优先队列实现
3. 如果用邻接矩阵，时间复杂度 O(V²)
4. 如果图不连通，未访问的节点距离为无穷
5. 如果有多条最短路径，只返回其中一条
6. 如果需要所有最短路径，需要修改算法
7. 如果边权为 0，仍然可以用 Dijkstra
8. 如果图有自环，忽略自环
9. 如果图有重边，取最小权重
10. 如果起点和终点相同，距离为 0
11. 如果起点或终点不存在，返回错误
12. 保存路径（不只保存距离），用 prev 数组回溯

### 代码模板（已测试）
```python
import heapq

def dijkstra(graph, start, end):
    """
    graph: 邻接表 {node: [(neighbor, weight), ...]}
    start: 起点
    end: 终点
    返回: (最短距离, 最短路径)
    """
    dist = {start: 0}
    prev = {start: None}
    pq = [(0, start)]
    
    while pq:
        d, u = heapq.heappop(pq)
        if d > dist.get(u, float('inf')):
            continue
        if u == end:
            break
        for v, w in graph.get(u, []):
            new_dist = d + w
            if new_dist < dist.get(v, float('inf')):
                dist[v] = new_dist
                prev[v] = u
                heapq.heappush(pq, (new_dist, v))
    
    # 回溯路径
    if end not in dist:
        return float('inf'), []
    
    path = []
    node = end
    while node is not None:
        path.append(node)
        node = prev[node]
    path.reverse()
    
    return dist[end], path

# 测试
graph = {
    1: [(2, 2), (4, 1.5)],
    2: [(3, 3), (5, 2.5)],
    3: [(6, 1)],
    4: [(5, 2)],
    5: [(6, 1.5)],
}
dist, path = dijkstra(graph, 1, 6)
print(f"最短距离: {dist}")
print(f"最短路径: {path}")
```

## 最大流详细指南

### 防错清单（10 条）
1. 无向边转为两条方向相反、容量相同的有向边
2. 流量守恒：中间节点流入 = 流出
3. 容量必须非负
4. 最大流 = 最小割（最大流最小割定理）
5. 如果有多条增广路，选择任意一条都可以
6. 如果容量为浮点数，需要设置精度阈值
7. 如果图不连通，最大流为 0
8. 如果源点和汇点相同，最大流为无穷
9. 保存每条边的流量（不只保存最大流值）
10. 如果需要最小费用最大流，需要额外处理费用

### 代码模板（Dinic 算法 - 完整版）
```python
from collections import deque

class Dinic:
    def __init__(self, n):
        self.n = n
        self.adj = [[] for _ in range(n)]
    
    def add_edge(self, u, v, cap):
        """添加边（有向）"""
        self.adj[u].append([v, cap, len(self.adj[v])])
        self.adj[v].append([u, 0, len(self.adj[u]) - 1])
    
    def _bfs(self, s, t):
        """BFS 构建层次图"""
        self.level = [-1] * self.n
        self.level[s] = 0
        q = deque([s])
        while q:
            u = q.popleft()
            for v, cap, _ in self.adj[u]:
                if cap > 0 and self.level[v] == -1:
                    self.level[v] = self.level[u] + 1
                    q.append(v)
        return self.level[t] != -1
    
    def _dfs(self, u, t, f, ptr):
        """DFS 寻找增广路"""
        if u == t:
            return f
        for i in range(ptr[u], len(self.adj[u])):
            ptr[u] = i
            v, cap, rev = self.adj[u][i]
            if cap > 0 and self.level[v] == self.level[u] + 1:
                pushed = self._dfs(v, t, min(f, cap), ptr)
                if pushed > 0:
                    self.adj[u][i][1] -= pushed
                    self.adj[v][rev][1] += pushed
                    return pushed
        return 0
    
    def max_flow(self, s, t):
        """计算最大流"""
        flow = 0
        while self._bfs(s, t):
            ptr = [0] * self.n
            while True:
                pushed = self._dfs(s, t, float('inf'), ptr)
                if pushed == 0:
                    break
                flow += pushed
        return flow

# 测试
dinic = Dinic(6)
dinic.add_edge(0, 1, 10)  # 源点 → 节点1
dinic.add_edge(0, 2, 10)  # 源点 → 节点2
dinic.add_edge(1, 3, 5)   # 节点1 → 节点3
dinic.add_edge(1, 4, 5)   # 节点1 → 节点4
dinic.add_edge(2, 3, 5)   # 节点2 → 节点3
dinic.add_edge(2, 4, 5)   # 节点2 → 节点4
dinic.add_edge(3, 5, 10)  # 节点3 → 汇点
dinic.add_edge(4, 5, 10)  # 节点4 → 汇点
max_flow = dinic.max_flow(0, 5)
print(f"最大流: {max_flow}")
```

## TSP 详细指南

### 防错清单（10 条）
1. 精确算法只适用于 ≤20 节点
2. 近似算法适用于 >20 节点
3. 必须检查 Hamiltonian 回路约束
4. 不对称 TSP：d(i,j) ≠ d(j,i) 时需要特殊处理
5. 如果图不连通，TSP 无解
6. 如果有负权边，TSP 可能无界
7. 近似比：最近邻算法的近似比为 O(log n)
8. LKH 算法通常能得到高质量解
9. 遗传算法需要多次运行取最优
10. 保存所有运行的结果，不只保存最优

### 代码模板（动态规划 - 精确算法）
```python
import numpy as np

def tsp_dp(dist_matrix):
    """TSP 动态规划（精确算法，O(n² · 2^n)）"""
    n = len(dist_matrix)
    
    # dp[mask][i] = 从起点出发，经过 mask 中所有节点，最后到达 i 的最短距离
    INF = float('inf')
    dp = [[INF] * n for _ in range(1 << n)]
    dp[1][0] = 0  # 起点为 0
    
    for mask in range(1 << n):
        for u in range(n):
            if dp[mask][u] == INF:
                continue
            if not (mask & (1 << u)):
                continue
            for v in range(n):
                if mask & (1 << v):
                    continue
                new_mask = mask | (1 << v)
                new_dist = dp[mask][u] + dist_matrix[u][v]
                if new_dist < dp[new_mask][v]:
                    dp[new_mask][v] = new_dist
    
    # 找最优解
    full_mask = (1 << n) - 1
    best = INF
    last = -1
    for u in range(n):
        total = dp[full_mask][u] + dist_matrix[u][0]
        if total < best:
            best = total
            last = u
    
    return best

# 测试
dist = np.array([
    [0, 2, 3, 4],
    [2, 0, 1, 3],
    [3, 1, 0, 2],
    [4, 3, 2, 0]
])
best = tsp_dp(dist)
print(f"TSP 最优解: {best}")
```
