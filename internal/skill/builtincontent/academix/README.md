# Academix — 学术副驾驶

> 研究者是驾驶员，AI 是发动机组。

Academix 是一个 Skills 驱动的学术研究助手系统，内置在 [Reasonix](https://github.com/esengine/DeepSeek-Reasonix) 中。

## 功能

- **数学建模**：赛题分析、模型选择、代码实现、论文写作
- **论文写作**：15 套 Typst + 15 套 LaTeX 模板
- **科研绘图**：11 种内置图表模板
- **建模知识库**：7 大类模型详细指南（评价、预测、优化、动力学、图论、统计、ML）
- **自动验收**：9 步验收流程

## 安装

```bash
# 1. 克隆 Reasonix
git clone --branch main-v2 https://github.com/esengine/DeepSeek-Reasonix.git
cd DeepSeek-Reasonix

# 2. 构建
make build

# 3. 提取 Academix 资产
./bin/reasonix academix-extract ~/.reasonix/academix-assets
# 或使用独立工具
go run ./cmd/academix-extract
```

## 使用

```bash
# 启动 Reasonix
reasonix

# 使用 Academix
/academix 帮我做数学建模竞赛

# 或自然语言
帮我做数学建模
```

## 工作流

### 数学建模竞赛
```
启动 → 赛题分析 → 代码实现 → 科研绘图 → 论文写作 → 验收
```

### 期刊论文
```
启动 → 文献调研 → 问题分析 → 实验 → 论文写作 → 审稿 → 验收
```

### 按需调用
用户问什么，答什么。不强制流水线。

## 知识库

| 文件 | 内容 |
|------|------|
| `modeling/catalog.md` | 7 大类模型速查 |
| `modeling/evaluation.md` | 评价类详细指南（AHP/TOPSIS/熵权） |
| `modeling/prediction.md` | 预测类详细指南（ARIMA/LSTM/灰色） |
| `modeling/optimization.md` | 优化类详细指南（LP/GA/PSO） |
| `modeling/dynamics.md` | 动力学详细指南（ODE/蒙特卡洛） |
| `modeling/graph.md` | 图论详细指南（Dijkstra/最大流） |
| `modeling/statistics.md` | 统计与 ML 详细指南 |
| `coding/norms.md` | 编码规范 |
| `writing/norms.md` | 写作规范 |
| `review/checklist.md` | 审稿检查清单 |
| `verify/procedures.md` | 9 步验收流程 |

## 模板

中文：cumcm、huashubei、huaweibei、diangongbei、shuweibei、wuyibei、huazhongbei、mathorcup、changsanjiao、dongsansheng、stats、default

英文：mcm、apmcm、default

每套模板同时提供 Typst 和 LaTeX 两个版本。

## 扩展

在工作目录创建 `academix-ext/` 扩展知识：

```
academix-ext/
├── knowledge/          # 自定义知识
├── templates/          # 自定义模板
└── workflows/          # 自定义工作流
```

## 示例

参见 `test-academix/` 目录中的完整示例：
- 赛题：城市交通流量优化
- 代码：Dijkstra + BPR 流量分配
- 论文：9 个章节 + 代码附录
- 验收：5 项检查全部通过

## 许可证

MIT License
