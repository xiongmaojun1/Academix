# 验收流程（9 步）

## Step 1: 文本质量门禁
运行脚本检查内部文件名泄露、AI 痕迹、占位符残留。

## Step 2: 章节结构
- Typst: `#include("...")` 数量和顺序
- LaTeX: `\input{...}` 数量和顺序
- 每个 section 有明确一级标题

## Step 3: 占位符检测
扫描：`TODO`、`FIXME`、`??`、`[待填]`、空章节

## Step 4: 图表引用
- 每张图在正文中被引用
- 引用的图存在
- 图片路径正确

## Step 5: 数值一致性
- 论文关键数值与 `results/*.json` 一致
- 摘要数值与正文一致
- 表格数值与代码输出一致

## Step 6: 参考文献
- 每篇引用有完整信息
- 引用格式一致
- 无明显编造痕迹

## Step 7: 编译检查
```bash
# Typst
typst compile main.typ --check
# LaTeX
xelatex -interaction=nonstopmode main.tex
```

## Step 8: PDF 检查
- 页数合理
- 图表正常显示
- 公式正确渲染

## Step 9: 验收报告
输出 `reports/VERIFY_REPORT.md`，包含通过/未通过状态和详细说明。
