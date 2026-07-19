# 数学建模竞赛工作流

## Phase 0: 启动

```bash
# 创建目录骨架
mkdir -p reports code results figures paper/sections

# 提取资产（首次使用）
academix-extract ~/.reasonix/academix-assets

# 询问用户：竞赛类型、语言、排版引擎、子问题数量
# 生成 plan.md 和 todo.md
```

产出：`plan.md`、`todo.md`

## Phase 1: 赛题分析与建模

```bash
# 1. 读取题面
# 2. 子问题拆解（只认明确编号的顶层问题）
# 3. 假设敏感性预检（至少两种解释）
# 4. 读取 knowledge/modeling/catalog.md 选择模型
# 5. 读取 knowledge/modeling/pitfalls.md 检查防错
# 6. 读取对应模型的详细指南（如 evaluation.md）
```

检查点：
- [ ] 子问题数量明确
- [ ] 每个假设有理由
- [ ] 模型选择有依据
- [ ] 防错清单已核对

产出：`reports/ANALYSIS_MODELING_REPORT.md`

## Phase 2: 代码实现

```bash
# 1. 读取 knowledge/coding/norms.md
# 2. 按建模报告实现代码
python3 code/problem1.py
python3 code/problem2.py

# 3. 验证结果
python3 -c "import json; r=json.load(open('results/problem1_results.json')); print(r['constraints_check'])"
```

检查点：
- [ ] 固定随机种子
- [ ] 结果 JSON 包含约束检查
- [ ] 最优解重新代入约束验证
- [ ] 整数变量正确处理

产出：`code/*.py`、`results/*.json`

## Phase 3: 科研绘图

```bash
# 1. 根据 results JSON 生成图表
python3 code/problem1.py  # 如果绘图逻辑在问题代码中

# 2. 或使用内置模板
ASSETS="${HOME}/.reasonix/academix-assets"
python3 "$ASSETS/figures/render_template.py" taylor-diagram
```

检查点：
- [ ] 图表数据来自 results JSON
- [ ] 高 DPI 输出（300+）
- [ ] 图中文字与论文语言一致

产出：`figures/*.png`

## Phase 4: 论文写作

```bash
# 1. 读取 knowledge/writing/norms.md
# 2. 复制模板
ASSETS="${HOME}/.reasonix/academix-assets"
cp -r "$ASSETS/templates/zh/cumcm/" paper/

# 3. 填充内容（数值与 results JSON 一致）
# 4. 编译
typst compile paper/main.typ
# 或
cd paper && xelatex main.tex && bibtex main && xelatex main.tex
```

检查点：
- [ ] 摘要覆盖所有子问题
- [ ] 数值与 results JSON 一致
- [ ] 公式解释符号
- [ ] 图表有引导和解释
- [ ] 参考文献真实存在
- [ ] 编译成功

产出：`paper/main.pdf`

## Phase 5: 验收

```bash
# 1. 读取 knowledge/verify/procedures.md
# 2. 运行文本检查
ASSETS="${HOME}/.reasonix/academix-assets"
bash "$ASSETS/scripts/writing_check.sh" --paper-dir paper --main paper/main.typ

# 3. 检查数值一致性
python3 -c "
import json
results = json.load(open('results/problem1_results.json'))
print('结果数值:', results['results'])
# 与论文中引用的数值对比
"
```

检查点：
- [ ] 文本质量门禁通过
- [ ] 章节结构正确
- [ ] 无占位符残留
- [ ] 图表引用完整
- [ ] 数值一致
- [ ] 编译成功
- [ ] PDF 正常

产出：`reports/VERIFY_REPORT.md`
