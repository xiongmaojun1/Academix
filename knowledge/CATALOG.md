# Academix 资产目录

本文件是你的资产索引。读取本文件后，你就知道能做什么、怎么做。

## 知识库

| 文件 | 内容 | 何时读取 |
|------|------|---------|
| `modeling/catalog.md` | 7 大类模型速查表 | 用户问"用什么模型" |
| `modeling/evaluation.md` | 评价类详细指南（AHP/TOPSIS/熵权/DEA 防错） | 涉及综合评价 |
| `modeling/prediction.md` | 预测类详细指南（ARIMA/LSTM/灰色 防错+检验） | 涉及预测 |
| `modeling/optimization.md` | 优化类详细指南（LP/MIP/GA/PSO 防错） | 涉及优化 |
| `modeling/dynamics.md` | 动力学详细指南（ODE/蒙特卡洛/马尔可夫） | 涉及物理过程 |
| `modeling/graph.md` | 图论详细指南（最短路/最大流/TSP 防错） | 涉及路径/网络 |
| `modeling/statistics.md` | 统计与 ML 详细指南（防错+指标） | 涉及数据分析 |
| `modeling/pitfalls.md` | 题型防错速查（各类型常见错误） | 任何建模任务 |
| `coding/norms.md` | 编码规范（可复现性、结果记录、常见错误） | 写代码时 |
| `writing/norms.md` | 写作规范（结构、引用、禁忌） | 写论文时 |
| `writing/mcm-icm.md` | 美赛专项规范 | MCM/ICM 竞赛 |
| `review/checklist.md` | 审稿检查清单 | 审稿/验收时 |
| `verify/procedures.md` | 9 步验收流程 | 最终验收时 |

## 工作流模板

| 文件 | 场景 | 何时读取 |
|------|------|---------|
| `workflows/math-competition.md` | 数学建模竞赛完整流程 | 用户说"做数学建模" |
| `workflows/journal-paper.md` | 期刊论文写作流程 | 用户说"写期刊论文" |
| `workflows/literature-review.md` | 文献综述流程 | 用户说"做文献综述" |
| `workflows/freestyle.md` | 按需调用模式 | 用户只问一个问题 |

## 图表模板（11 种）

Python/matplotlib 脚本，提取后在 `~/.reasonix/academix-assets/figures/` 目录中。

| 模板 ID | 用途 |
|---------|------|
| `correlation-pairgrid` | 多变量相关性 |
| `cv-roc-ci` | 分类模型 ROC |
| `grouped-circular-heatmap` | 多维对比 |
| `grouped-corr-split-violin` | 分布对比 |
| `multiclass-shap-combo` | 特征重要性 |
| `nature-chord-diagram` | 关系网络 |
| `paired-raincloud` | 配对数据分布 |
| `prediction-marginal-grid` | 预测 vs 真实 |
| `rf-tpe-surface` | 超参数调优 |
| `taylor-diagram` | 模型性能对比 |
| `urban-park-cooling-combo` | 多指标组合 |

## 论文模板（15 套 Typst + 15 套 LaTeX）

在 `~/.reasonix/academix-assets/templates/` 目录中。

中文 Typst：cumcm、huashubei、huaweibei、diangongbei、shuweibei、wuyibei、huazhongbei、mathorcup、changsanjiao、dongsansheng、stats、default

中文 LaTeX：cumcm-latex、huashubei-latex、huaweibei-latex、diangongbei-latex、shuweibei-latex、wuyibei-latex、huazhongbei-latex、mathorcup-latex、changsanjiao-latex、dongsansheng-latex、stats-latex、default-latex

英文 Typst：mcm、apmcm、default

英文 LaTeX：mcm-latex、apmcm-latex、default-latex

## 使用规则

1. **先读目录，再读内容** — 不要凭记忆工作
2. **按需读取** — 只读与当前任务相关的文件
3. **组合使用** — 一个任务可能需要多个知识文件
4. **记录引用** — 在 plan.md 中记录读取了哪些文件
