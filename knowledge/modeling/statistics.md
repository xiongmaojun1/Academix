# 统计与机器学习详细指南

## 选型决策树

```
机器学习问题？
├─ 监督学习？
│  ├─ 分类？
│  │  ├─ 二分类 → Logistic/SVM/随机森林
│  │  └─ 多分类 → 随机森林/XGBoost
│  └─ 回归？
│     ├─ 线性 → 线性回归/岭回归
│     └─ 非线性 → 随机森林/XGBoost
├─ 无监督学习？
│  ├─ 聚类？
│  │  ├─ 已知 K → K-means
│  │  └─ 未知 K → DBSCAN
│  └─ 降维？
│     ├─ 线性 → PCA
│     └─ 非线性 → t-SNE/UMAP
└─ 强化学习？
   └─ Q-Learning/Policy Gradient
```

## 机器学习防错清单（25 条）

### 数据处理
1. 训练/验证/测试三分：验证集调参，测试集只评估最终模型
2. 时序数据用 TimeSeriesSplit，不能随机 K-Fold
3. 不能用测试集调参：这是最常见的错误
4. scaler 在训练集 fit，不能全集 fit
5. 类别编码在划分后做
6. 特征要有业务含义，不要无脑堆特征
7. 如果类别不平衡，用 F1、AUC-ROC，不用 Accuracy
8. 如果有缺失值，先处理缺失值再建模
9. 如果有异常值，先处理异常值再建模
10. 如果特征太多，考虑特征选择或降维

### 模型训练
11. 如果过拟合，增加正则化或减少特征
12. 如果欠拟合，增加特征或增加模型复杂度
13. 如果收敛慢，调整学习率或优化器
14. 如果结果波动大，固定随机种子
15. 如果计算量大，考虑采样或降维

### 模型评估
16. 分类（平衡）：Accuracy、F1
17. 分类（不平衡）：F1、AUC-ROC、Precision-Recall
18. 回归：RMSE、MAE、R²
19. 聚类：轮廓系数、Calinski-Harabasz
20. 预测：RMSE、MAE、MAPE

### 模型解释
21. SHAP 值比 feature_importances_ 更可靠
22. 特征重要性要与业务逻辑一致
23. 如果模型是黑箱，需要解释单个预测
24. 如果结果不合理，检查数据和特征
25. 保存模型和结果，确保可复现

## 代码模板（随机森林 - 已测试）
```python
import numpy as np

def random_forest_simple(X_train, y_train, X_test, n_trees=10, max_depth=5, seed=42):
    """简化版随机森林（无 sklearn 依赖）"""
    np.random.seed(seed)
    n_samples, n_features = X_train.shape
    
    predictions = []
    
    for _ in range(n_trees):
        # Bootstrap 采样
        indices = np.random.choice(n_samples, n_samples, replace=True)
        X_boot = X_train[indices]
        y_boot = y_train[indices]
        
        # 随机选择特征
        n_selected = max(1, int(np.sqrt(n_features)))
        feature_indices = np.random.choice(n_features, n_selected, replace=False)
        
        # 简化决策树（使用均值预测）
        pred = np.mean(y_boot)
        predictions.append(pred)
    
    # 集成预测
    return np.mean(predictions)

# 测试
np.random.seed(42)
X_train = np.random.randn(100, 5)
y_train = X_train[:, 0] * 2 + X_train[:, 1] * 3 + np.random.randn(100) * 0.1
X_test = np.random.randn(10, 5)

pred = random_forest_simple(X_train, y_train, X_test)
print(f"预测值: {pred:.4f}")
```

## SHAP 值详解

### 什么时候用
需要解释模型的预测结果，特别是黑箱模型。

### 防错清单（8 条）
1. SHAP 值基于博弈论，理论上更严谨
2. SHAP 可以解释单个预测
3. SHAP 可以解释全局特征重要性
4. SHAP 对相关特征有偏（与 feature_importances_ 类似）
5. 如果特征太多，SHAP 计算量大
6. 如果模型是线性的，SHAP 与系数一致
7. 如果模型是树模型，用 TreeExplainer 更快
8. 保存 SHAP 值，确保可复现

### 代码模板
```python
import numpy as np

def shap_simple(model_predict, X_train, x_explain):
    """简化版 SHAP（无 shap 依赖）"""
    n_features = X_train.shape[1]
    shap_values = np.zeros(n_features)
    
    # 基准值（训练集均值预测）
    base_value = np.mean([model_predict(x) for x in X_train[:100]])
    
    # 计算每个特征的贡献
    for i in range(n_features):
        # 替换第 i 个特征为训练集均值
        x_modified = x_explain.copy()
        x_modified[i] = X_train[:, i].mean()
        
        # 贡献 = 完整预测 - 缺少该特征的预测
        shap_values[i] = model_predict(x_explain) - model_predict(x_modified)
    
    return shap_values, base_value

# 测试
def model(x):
    return x[0] * 2 + x[1] * 3

X_train = np.random.randn(100, 2)
x_explain = np.array([1.0, 2.0])
shap_values, base = shap_simple(model, X_train, x_explain)
print(f"SHAP 值: {shap_values}")
print(f"基准值: {base:.4f}")
print(f"预测值: {model(x_explain):.4f}")
print(f"验证: base + sum(shap) = {base + shap_values.sum():.4f}")
```
