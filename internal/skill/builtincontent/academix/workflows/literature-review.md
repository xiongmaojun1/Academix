# 文献综述工作流

## 阶段间契约

```
Phase 0 (启动)
  输出: plan.md, academix-state.json
  ↓
Phase 1 (搜索)
  输入: 用户描述的研究问题
  输出: reports/search_results.md, academix-state.json
  ↓
Phase 2 (筛选)
  输入: reports/search_results.md
  输出: reports/filtered_papers.md, academix-state.json
  ↓
Phase 3 (阅读分类)
  输入: reports/filtered_papers.md
  输出: reports/classified_papers.md, academix-state.json
  ↓
Phase 4 (撰写综述)
  输入: reports/classified_papers.md
  输出: reports/LITERATURE_REVIEW.md, academix-state.json
  ↓
Phase 5 (引用检查)
  输入: reports/LITERATURE_REVIEW.md
  输出: reports/LITERATURE_REVIEW.md (修正版), academix-state.json
```

## Phase 1: 搜索

### 输入
- 用户描述的研究问题

### 输出
- `reports/search_results.md`：搜索结果
- `academix-state.json`：更新状态

### 命令
```bash
# 使用 WebSearch 搜索
# 关键词：中英文各一组
# 时间范围：近 5 年优先
# 来源：arXiv、Semantic Scholar、Google Scholar
```

### 检查点
- [ ] 关键词覆盖主要方向
- [ ] 搜索结果数量足够（>20 篇）
- [ ] 搜索结果相关性高

### 错误处理
- 如果搜索无结果 → 调整关键词
- 如果结果太少 → 扩大搜索范围

### 回退机制
- 如果搜索质量差 → 重新搜索

## Phase 2: 筛选

### 输入
- `reports/search_results.md`：搜索结果

### 输出
- `reports/filtered_papers.md`：筛选后的论文
- `academix-state.json`：更新状态

### 命令
```bash
# 按被引次数排序
# 关注高被引论文和最新论文
# 通过引用关系扩展
```

### 检查点
- [ ] 筛选标准明确
- [ ] 筛选结果数量合理（10-30 篇）
- [ ] 筛选结果覆盖主要方向

### 错误处理
- 如果筛选结果太少 → 放宽标准
- 如果筛选结果太多 → 收紧标准

### 回退机制
- 如果筛选质量差 → 重新筛选

## Phase 3: 阅读分类

### 输入
- `reports/filtered_papers.md`：筛选后的论文

### 输出
- `reports/classified_papers.md`：分类后的论文
- `academix-state.json`：更新状态

### 命令
```bash
# 按方法分类（方法一、方法二、...）
# 按时间分类（早期、近期、最新）
# 记录每篇论文的核心思想、优点、局限
```

### 检查点
- [ ] 分类标准明确
- [ ] 每篇论文有核心思想、优点、局限
- [ ] 分类覆盖主要方向

### 错误处理
- 如果分类不清晰 → 重新分类
- 如果论文太多 → 选择代表性论文

### 回退机制
- 如果分类质量差 → 重新分类

## Phase 4: 撰写综述

### 输入
- `reports/classified_papers.md`：分类后的论文

### 输出
- `reports/LITERATURE_REVIEW.md`：文献综述
- `academix-state.json`：更新状态

### 命令
```bash
# 撰写综述
# 结构：摘要→研究背景→研究现状→研究趋势→参考文献
```

### 检查点
- [ ] 结构清晰
- [ ] 每个方法有代表论文、核心思想、优点、局限
- [ ] 研究趋势明确
- [ ] 参考文献完整

### 错误处理
- 如果综述结构不清晰 → 重新组织
- 如果内容不完整 → 补充信息

### 回退机制
- 如果综述质量差 → 重新撰写

## Phase 5: 引用检查

### 输入
- `reports/LITERATURE_REVIEW.md`：文献综述

### 输出
- `reports/LITERATURE_REVIEW.md`：修正版文献综述
- `academix-state.json`：更新状态

### 命令
```bash
# 检查每篇引用有完整信息
# 检查引用格式一致
# 检查不编造文献
```

### 检查点
- [ ] 每篇引用有完整信息
- [ ] 引用格式一致
- [ ] 无编造痕迹

### 错误处理
- 如果引用不完整 → 补充信息
- 如果格式不一致 → 统一格式

### 回退机制
- 如果引用有误 → 修正后重新检查
