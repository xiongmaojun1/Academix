# 数学建模竞赛工作流

## 阶段间契约

每个阶段的输入必须来自上一个阶段的输出：

```
Phase 0 (启动)
  输出: plan.md, todo.md, academix-state.json
  ↓
Phase 1 (分析建模)
  输入: PROBLEM.md (题面)
  输出: reports/ANALYSIS_MODELING_REPORT.md, academix-state.json
  ↓
Phase 2 (代码实现)
  输入: reports/ANALYSIS_MODELING_REPORT.md
  输出: code/*.py, results/results.json, academix-state.json
  ↓
Phase 3 (科研绘图)
  输入: results/results.json
  输出: figures/*.png, academix-state.json
  ↓
Phase 4 (论文写作)
  输入: results/results.json, figures/*.png
  输出: paper/main.pdf, academix-state.json
  ↓
Phase 5 (验收)
  输入: paper/main.pdf, results/results.json
  输出: reports/VERIFY_REPORT.md, academix-state.json
```

## Phase 0: 启动

### 输入
- 用户描述的问题

### 输出
- `plan.md`：研究计划
- `todo.md`：待办事项
- `academix-state.json`：初始状态

### 命令
```bash
# 创建目录骨架
mkdir -p reports code results figures paper/sections

# 提取资产（首次使用）
ASSETS="${HOME}/.reasonix/academix-assets"
if [ ! -d "$ASSETS" ]; then
  academix-extract "$ASSETS"
fi

# 初始化状态
echo '{"current_phase":"init","completed_phases":[],"decisions":[],"results":{},"blockers":[]}' > academix-state.json
```

### 检查点
- [ ] 目录结构创建成功
- [ ] 资产目录存在
- [ ] 状态文件初始化

### 错误处理
- 如果 `academix-extract` 不存在 → 提示用户安装 Reasonix
- 如果资产目录创建失败 → 检查权限

## Phase 1: 赛题分析与建模

### 输入
- `PROBLEM.md`：题面文件

### 输出
- `reports/ANALYSIS_MODELING_REPORT.md`：建模报告
- `academix-state.json`：更新状态

### 命令
```bash
# 读取题面
cat PROBLEM.md

# 读取知识库
cat knowledge/modeling/catalog.md
cat knowledge/modeling/pitfalls.md
cat knowledge/modeling/evaluation.md  # 如果涉及评价
cat knowledge/modeling/prediction.md  # 如果涉及预测
cat knowledge/modeling/optimization.md  # 如果涉及优化
```

### 检查点
- [ ] 子问题数量明确
- [ ] 每个假设有理由
- [ ] 模型选择有依据
- [ ] 防错清单已核对
- [ ] 建模报告格式正确

### 错误处理
- 如果题面不清晰 → 向用户提问
- 如果模型选择不确定 → 列出备选方案
- 如果知识库文件不存在 → 使用通用知识

### 回退机制
- 如果建模报告有误 → 回到 Phase 1 重新分析
- 如果假设不合理 → 修改假设后重新建模

## Phase 2: 代码实现

### 输入
- `reports/ANALYSIS_MODELING_REPORT.md`：建模报告

### 输出
- `code/*.py`：代码文件
- `results/results.json`：结果文件
- `academix-state.json`：更新状态

### 命令
```bash
# 读取建模报告
cat reports/ANALYSIS_MODELING_REPORT.md

# 读取编码规范
cat knowledge/coding/norms.md

# 运行代码
python3 code/solve.py

# 验证结果
python3 -c "
import json
r = json.load(open('results/results.json'))
# 检查数值合理性
for key, val in r.items():
  if isinstance(val, dict):
    for k, v in val.items():
      if isinstance(v, (int, float)):
        if v < 0 and 'time' in k.lower():
          print(f'❌ {key}.{k} = {v} (时间不能为负)')
"
```

### 检查点
- [ ] 固定随机种子
- [ ] 结果 JSON 包含约束检查
- [ ] 最优解重新代入约束验证
- [ ] 整数变量正确处理
- [ ] 代码可复现

### 错误处理
- 如果代码报错 → 检查依赖、语法、逻辑
- 如果结果不合理 → 检查约束、参数、算法
- 如果约束违反 → 修改代码后重新运行

### 回退机制
- 如果代码有 bug → 修复后重新运行
- 如果结果不合理 → 回到 Phase 1 重新建模

## Phase 3: 科研绘图

### 输入
- `results/results.json`：结果文件

### 输出
- `figures/*.png`：图表文件
- `academix-state.json`：更新状态

### 命令
```bash
# 读取结果
python3 -c "import json; r=json.load(open('results/results.json')); print(r)"

# 使用内置模板
ASSETS="${HOME}/.reasonix/academix-assets"
python3 "$ASSETS/figures/render_template.py" taylor-diagram

# 或自定义绘图
python3 code/plot.py
```

### 检查点
- [ ] 图表数据来自 results JSON
- [ ] 高 DPI 输出（300+）
- [ ] 图中文字与论文语言一致
- [ ] 图表有标题和标签

### 错误处理
- 如果图表脚本报错 → 检查数据格式
- 如果图表不清晰 → 调整参数
- 如果图表无数据 → 检查 results.json

### 回退机制
- 如果图表有误 → 修复后重新生成
- 如果数据不对 → 回到 Phase 2 重新运行

## Phase 4: 论文写作

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

# 复制模板
ASSETS="${HOME}/.reasonix/academix-assets"
cp -r "$ASSETS/templates/zh/cumcm/" paper/

# 填充内容（数值与 results JSON 一致）

# 编译
typst compile paper/main.typ
# 或
cd paper && xelatex main.tex && bibtex main && xelatex main.tex
```

### 检查点
- [ ] 摘要覆盖所有子问题
- [ ] 数值与 results JSON 一致
- [ ] 公式解释符号
- [ ] 图表有引导和解释
- [ ] 参考文献真实存在
- [ ] 编译成功

### 错误处理
- 如果编译失败 → 检查语法错误
- 如果数值不一致 → 检查 results.json
- 如果格式不对 → 检查模板

### 回退机制
- 如果编译失败 → 修复后重新编译
- 如果数值不对 → 回到 Phase 2 重新运行

## Phase 5: 验收

### 输入
- `paper/main.pdf`：论文 PDF
- `results/results.json`：结果文件

### 输出
- `reports/VERIFY_REPORT.md`：验收报告
- `academix-state.json`：更新状态

### 命令
```bash
# 读取验收流程
cat knowledge/verify/procedures.md

# 运行文本检查
ASSETS="${HOME}/.reasonix/academix-assets"
bash "$ASSETS/scripts/writing_check.sh" --paper-dir paper --main paper/main.typ

# 检查数值一致性
python3 -c "
import json
r = json.load(open('results/results.json'))
print('结果数值:', r)
"
```

### 检查点
- [ ] 文本质量门禁通过
- [ ] 章节结构正确
- [ ] 无占位符残留
- [ ] 图表引用完整
- [ ] 数值一致
- [ ] 编译成功
- [ ] PDF 正常

### 错误处理
- 如果验收失败 → 修复后重新验收
- 如果有硬错误 → 回到对应阶段修复

### 回退机制
- 如果验收失败 → 回到 Phase 4 修复
- 如果有硬错误 → 回到对应阶段修复
