# 验收流程（9 步）

## Step 1: 文本质量门禁

### 检查项
- 内部文件名泄露：`code/`、`results/`、`reports/`
- AI 痕迹：`AI 生成`、`由.*生成`、`ChatGPT`、`GPT`、`Claude`
- 占位符残留：`TODO`、`FIXME`、`??`、`[待填]`

### 命令
```bash
# 检查内部文件名
grep -r "code/\|results/\|reports/" paper/ 2>/dev/null && echo "发现内部文件名" || echo "无内部文件名"

# 检查 AI 痕迹
grep -r "AI 生成\|由.*生成\|ChatGPT\|GPT\|Claude" paper/ 2>/dev/null && echo "发现 AI 痕迹" || echo "无 AI 痕迹"

# 检查占位符
grep -r "TODO\|FIXME\|??\|\[待填\]" paper/ 2>/dev/null && echo "发现占位符" || echo "无占位符"
```

### 判断标准
- PASS：无内部文件名、无 AI 痕迹、无占位符
- FAIL：发现任何一个

## Step 2: 章节结构

### 检查项
- Typst: `#include("...")` 数量和顺序
- LaTeX: `\input{...}` 数量和顺序
- 每个 section 有明确一级标题

### 命令
```bash
# Typst 检查
grep -c "#include" paper/main.typ
grep "#include" paper/main.typ | sort

# LaTeX 检查
grep -c "\\\\input\|\\\\include" paper/main.tex
grep "\\\\input\|\\\\include" paper/main.tex | sort
```

### 判断标准
- PASS：include 数量与 sections 目录一致，顺序正确
- FAIL：数量不一致或顺序错误

## Step 3: 占位符检测

### 检查项
- `TODO`、`FIXME`、`XXX`、`PLACEHOLDER`
- `??`、`[待填]`、`[占位]`
- 空章节（只有标题没有内容）

### 命令
```bash
# 检查占位符
grep -r "TODO\|FIXME\|XXX\|PLACEHOLDER\|??\|\[待填\]\|\[占位\]" paper/ 2>/dev/null

# 检查空章节
for f in paper/sections/*.typ; do
  lines=$(wc -l < "$f")
  if [ "$lines" -lt 5 ]; then
    echo "警告: $f 只有 $lines 行"
  fi
done
```

### 判断标准
- PASS：无占位符，无空章节
- FAIL：发现占位符或空章节

## Step 4: 图表引用

### 检查项
- 每张图在正文中被引用
- 引用的图存在
- 图片路径正确

### 命令
```bash
# 检查图片文件
ls paper/figures/ 2>/dev/null

# 检查引用
grep -r "figure\|图" paper/sections/*.typ 2>/dev/null
```

### 判断标准
- PASS：所有图片都被引用，所有引用的图片都存在
- FAIL：有未引用的图片或引用了不存在的图片

## Step 5: 数值一致性

### 检查项
- 论文关键数值与 `results/*.json` 一致
- 摘要数值与正文一致
- 表格数值与代码输出一致

### 命令
```bash
# 提取论文中的数值
grep -oP '\d+\.\d+' paper/sections/*.typ | sort

# 提取结果中的数值
python3 -c "
import json
r = json.load(open('results/results.json'))
for key, val in r.items():
  if isinstance(val, dict):
    for k, v in val.items():
      if isinstance(v, (int, float)):
        print(f'{key}.{k} = {v}')
"
```

### 判断标准
- PASS：所有数值一致
- FAIL：发现不一致的数值

## Step 6: 参考文献

### 检查项
- 每篇引用有完整信息
- 引用格式一致
- 无明显编造痕迹

### 命令
```bash
# 检查参考文献文件
cat paper/references.typ

# 检查引用格式
grep -c "\[" paper/references.typ
```

### 判断标准
- PASS：所有引用有完整信息，格式一致
- FAIL：有不完整或格式不一致的引用

## Step 7: 编译检查

### 命令
```bash
# Typst
typst compile paper/main.typ --check

# LaTeX
cd paper && xelatex -interaction=nonstopmode main.tex
```

### 判断标准
- PASS：编译成功，无错误
- FAIL：编译失败或有错误

## Step 8: PDF 检查

### 检查项
- 页数合理
- 图表正常显示
- 公式正确渲染

### 命令
```bash
# 检查 PDF 文件
ls -la paper/main.pdf

# 检查页数
python3 -c "
import subprocess
result = subprocess.run(['pdfinfo', 'paper/main.pdf'], capture_output=True, text=True)
print(result.stdout)
" 2>/dev/null || echo "pdfinfo 未安装"
```

### 判断标准
- PASS：PDF 存在，页数合理，图表正常
- FAIL：PDF 不存在或页数异常

## Step 9: 验收报告

### 报告模板
```markdown
# 验收报告

## 结果：通过 / 未通过

| 检查项 | 结果 | 说明 |
|--------|------|------|
| 文本质量 | PASS/FAIL | ... |
| 章节结构 | PASS/FAIL | ... |
| 占位符 | PASS/FAIL | ... |
| 图表引用 | PASS/FAIL | ... |
| 数值一致 | PASS/FAIL | ... |
| 参考文献 | PASS/FAIL | ... |
| 编译 | PASS/FAIL | ... |
| PDF | PASS/FAIL | ... |

## 硬错误（必须修复）
1. ...

## 警告（建议修复）
1. ...
```
