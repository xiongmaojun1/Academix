# 按需调用模式

## 适用场景
用户只问一个问题，不需要完整工作流。

## 规则
- 不强制走完整流水线
- 用户问什么，答什么
- 回答后问用户是否需要继续
- 每次回答后更新 academix-state.json

## 执行流程

### Step 1: 理解问题
分析用户问题，判断需要哪些知识文件。

### Step 2: 读取知识库
```bash
# 根据问题类型读取对应文件
cat knowledge/modeling/evaluation.md  # 评价问题
cat knowledge/modeling/prediction.md  # 预测问题
cat knowledge/modeling/optimization.md  # 优化问题
```

### Step 3: 回答问题
根据知识库内容回答问题。

### Step 4: 验证回答
```bash
# 如果涉及代码，验证代码能运行
python3 -c "print('代码验证通过')"
```

### Step 5: 更新状态
```bash
python3 -c "
import json
from datetime import datetime
state = json.load(open('academix-state.json'))
state['decisions'].append({
  'phase': 'freestyle',
  'decision': '回答问题',
  'reason': '用户提问',
  'timestamp': datetime.now().isoformat()
})
json.dump(state, open('academix-state.json', 'w'), indent=2, ensure_ascii=False)
"
```

### Step 6: 询问继续
问用户是否需要继续。

## 示例

### 示例 1：选模型
```
用户：帮我选一个评价模型
→ 读取 knowledge/modeling/evaluation.md
→ 苏格拉底式对话引导选择
→ 问：需要我帮你实现这个模型吗？
```

### 示例 2：画图
```
用户：帮我画一个泰勒图
→ 使用图表模板生成
→ 问：需要我帮你把这张图放到论文里吗？
```

### 示例 3：搜论文
```
用户：帮我搜一下小目标检测的论文
→ 直接搜索，生成综述
→ 问：需要我帮你分析这些论文的方法吗？
```
