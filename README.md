# Academix — 学术副驾驶

> 研究者是驾驶员，AI 是发动机组。

Academix 是一个独立的学术研究助手系统，帮助研究者完成数学建模、论文写作、科研绘图等学术任务。

## 安装

### 从 GitHub Releases 下载

https://github.com/xiongmaojun1/Academix/releases

### 从源码编译

```bash
git clone git@github.com:xiongmaojun1/Academix.git
cd Academix
go build -o bin/academix ./cmd/academix
```

## 使用

```bash
# 启动
./bin/academix

# 交互
> 帮我做数学建模竞赛
> 帮我选一个评价模型
> 帮我画一个泰勒图
> 帮我写论文摘要
```

## 功能

| 功能 | 说明 |
|------|------|
| 数学建模 | 赛题分析、模型选择、代码实现 |
| 论文写作 | 30 套 Typst/LaTeX 模板 |
| 科研绘图 | 11 种内置图表模板 |
| 建模知识库 | 7 大类模型详细指南 |
| 自动验收 | 9 步验收流程 |

## 命令

```
/help        显示帮助
/status      显示当前状态
/skills      列出可用 Skills
/knowledge   列出知识库文件
/export-log  导出研究日志
/clear       重置状态
/quit        退出
```

## 目录结构

```
academix/
├── academix              # 可执行文件
├── skills/               # Skills 定义
├── knowledge/            # 知识库
├── workflows/            # 工作流模板
└── assets/               # 模板、图表脚本
```

## 知识库

| 文件 | 内容 |
|------|------|
| `knowledge/modeling/evaluation.md` | 评价类模型（AHP/TOPSIS/熵权） |
| `knowledge/modeling/prediction.md` | 预测类模型（ARIMA/LSTM/灰色） |
| `knowledge/modeling/optimization.md` | 优化类模型（LP/GA/PSO） |
| `knowledge/modeling/dynamics.md` | 动力学模型（ODE/蒙特卡洛） |
| `knowledge/modeling/graph.md` | 图论模型（Dijkstra/最大流） |
| `knowledge/modeling/statistics.md` | 统计与机器学习 |
| `knowledge/coding/norms.md` | 编码规范 |
| `knowledge/writing/norms.md` | 写作规范 |

## 许可证

MIT License
