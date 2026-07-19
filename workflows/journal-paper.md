# 期刊论文工作流

## 阶段间契约

```
Phase 0 (启动)
  输出: plan.md, todo.md, academix-state.json
  ↓
Phase 1 (文献调研)
  输入: 用户描述的研究问题
  输出: reports/LITERATURE_REVIEW.md, academix-state.json
  ↓
Phase 2 (问题分析)
  输入: reports/LITERATURE_REVIEW.md
  输出: reports/PROBLEM_ANALYSIS.md, academix-state.json
  ↓
Phase 3 (建模/实验)
  输入: reports/PROBLEM_ANALYSIS.md
  输出: reports/MODEL_DESIGN.md, academix-state.json
  ↓
Phase 4 (代码实现)
  输入: reports/MODEL_DESIGN.md
  输出: code/*.py, results/results.json, academix-state.json
  ↓
Phase 5 (科研绘图)
  输入: results/results.json
  输出: figures/*.png, academix-state.json
  ↓
Phase 6 (论文写作)
  输入: results/results.json, figures/*.png
  输出: paper/main.pdf, academix-state.json
  ↓
Phase 7 (审稿)
  输入: paper/main.pdf
  输出: reports/REVIEW_REPORT.md, academix-state.json
  ↓
Phase 8 (验收)
  输入: paper/main.pdf, results/results.json
  输出: reports/VERIFY_REPORT.md, academix-state.json
```

## Phase 1: 文献调研

### 输入
- 用户描述的研究问题

### 输出
- `reports/LITERATURE_REVIEW.md`：文献综述
- `academix-state.json`：更新状态

### 命令
```bash
# 读取知识库
cat knowledge/writing/norms.md

# 使用 WebSearch 搜索
# 使用 WebFetch 获取论文详情
```

### 检查点
- [ ] 搜索覆盖主要方向
- [ ] 综述分类清晰
- [ ] 引用真实存在
- [ ] 引用格式一致

### 错误处理
- 如果搜索无结果 → 调整关键词
- 如果引用不完整 → 补充信息

### 回退机制
- 如果综述质量差 → 重新搜索和整理

## Phase 6: 论文写作

### 输入
- `results/results.json`：结果文件
- `figures/*.png`：图表文件

### 输出
- `paper/main.pdf`：论文 PDF
- `academix-state.json`：更新状态

### 命令
```bash
# 读取写作规范
cat knowledge/writing/norms.md

# 使用通用模板
ASSETS="${HOME}/.reasonix/academix-assets"
cp -r "$ASSETS/templates/en/default/" paper/

# 编译
typst compile paper/main.typ
# 或
cd paper && xelatex main.tex && bibtex main && xelatex main.tex
```

### 检查点
- [ ] 摘要覆盖主要贡献
- [ ] 数值与 results JSON 一致
- [ ] 公式解释符号
- [ ] 参考文献真实存在
- [ ] 编译成功

### 错误处理
- 如果编译失败 → 检查语法错误
- 如果数值不一致 → 检查 results.json

### 回退机制
- 如果编译失败 → 修复后重新编译
- 如果数值不对 → 回到 Phase 4 重新运行

## Phase 7: 审稿

### 输入
- `paper/main.pdf`：论文 PDF

### 输出
- `reports/REVIEW_REPORT.md`：审稿报告
- `academix-state.json`：更新状态

### 命令
```bash
# 读取审稿清单
cat knowledge/review/checklist.md

# 模拟审稿人审查
```

### 检查点
- [ ] 方法合理性
- [ ] 代码正确性
- [ ] 论文质量
- [ ] 创新性

### 错误处理
- 如果审稿发现问题 → 记录并建议修复

### 回退机制
- 如果审稿不通过 → 回到 Phase 6 修复
