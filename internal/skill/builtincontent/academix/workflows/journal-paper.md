# 期刊论文工作流

## 适用场景
投稿 IEEE、ACM、Springer、Elsevier 等期刊。

## Phase 1: 文献调研
```bash
# 搜索相关论文
# 使用 WebSearch 搜索关键词
# 使用 WebFetch 获取论文详情
# 生成综述报告
```

检查点：
- [ ] 搜索覆盖主要方向
- [ ] 综述分类清晰
- [ ] 引用真实存在

产出：`reports/LITERATURE_REVIEW.md`

## Phase 2: 问题分析
```bash
# 明确研究问题和贡献点
# 与已有工作的区别
```

检查点：
- [ ] 研究问题明确
- [ ] 贡献点清晰
- [ ] 与已有工作有区别

产出：`reports/PROBLEM_ANALYSIS.md`

## Phase 3: 建模/实验
```bash
# 读取 knowledge/modeling/catalog.md 选择模型
# 读取对应模型的详细指南
# 设计实验方案
```

检查点：
- [ ] 模型选择有依据
- [ ] 实验方案可复现

产出：`reports/MODEL_DESIGN.md`

## Phase 4: 代码实现
```bash
# 读取 knowledge/coding/norms.md
# 实现实验代码
python3 code/experiment.py

# 验证结果
python3 -c "import json; r=json.load(open('results/results.json')); print(r)"
```

检查点：
- [ ] 固定随机种子
- [ ] 结果 JSON 包含约束检查
- [ ] 可复现

产出：`code/*.py`、`results/*.json`

## Phase 5: 科研绘图
```bash
# 根据 results JSON 生成图表
ASSETS="${HOME}/.reasonix/academix-assets"
python3 "$ASSETS/figures/render_template.py" <template-id>
```

检查点：
- [ ] 图表数据来自 results JSON
- [ ] 高 DPI 输出
- [ ] 图中文字与论文语言一致

产出：`figures/*.png`

## Phase 6: 论文写作
```bash
# 读取 knowledge/writing/norms.md
# 使用通用模板
ASSETS="${HOME}/.reasonix/academix-assets"
cp -r "$ASSETS/templates/en/default/" paper/

# 编译
typst compile paper/main.typ
# 或
cd paper && xelatex main.tex && bibtex main && xelatex main.tex
```

检查点：
- [ ] 摘要覆盖主要贡献
- [ ] 数值与 results JSON 一致
- [ ] 公式解释符号
- [ ] 参考文献真实存在
- [ ] 编译成功

产出：`paper/main.pdf`

## Phase 7: 审稿
```bash
# 读取 knowledge/review/checklist.md
# 模拟审稿人审查
```

检查点：
- [ ] 方法合理性
- [ ] 代码正确性
- [ ] 论文质量
- [ ] 创新性

产出：`reports/REVIEW_REPORT.md`

## Phase 8: 验收
```bash
# 读取 knowledge/verify/procedures.md
ASSETS="${HOME}/.reasonix/academix-assets"
bash "$ASSETS/scripts/writing_check.sh" --paper-dir paper --main paper/main.typ
```

检查点：
- [ ] 文本质量门禁通过
- [ ] 数值一致
- [ ] 编译成功
- [ ] PDF 正常

产出：`reports/VERIFY_REPORT.md`
