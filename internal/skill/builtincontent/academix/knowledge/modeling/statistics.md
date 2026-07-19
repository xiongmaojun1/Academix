# 统计与机器学习详细指南

## 统计检验速查

| 目的 | 方法 | 适用条件 |
|------|------|---------|
| 两组均值差异 | t 检验 | 正态、方差齐 |
| 多组均值差异 | ANOVA | 正态、方差齐 |
| 非参数两组 | Mann-Whitney U | 不要求正态 |
| 非参数多组 | Kruskal-Wallis | 不要求正态 |
| 相关性 | Pearson（线性）/ Spearman（非线性）| Pearson 要求正态 |
| 独立性 | 卡方检验 | 期望频数 ≥ 5 |
| 平稳性 | ADF 单位根检验 | — |
| 正态性 | Shapiro-Wilk（小样本）/ K-S（大样本）| — |

## 机器学习通用防错

### 数据划分
- **训练/验证/测试三分**：验证集调参，测试集只评估最终模型
- **时序数据**：用 TimeSeriesSplit，不能随机 K-Fold
- **不能用测试集调参**：这是最常见的错误

### 特征工程
- **scaler 在训练集 fit**：不能在全集 fit，否则数据泄露
- **类别编码在划分后做**：One-Hot、Label Encoding 都要在划分后
- **特征要有业务含义**：不要无脑堆特征

### 类别不平衡
- **Accuracy 失效**：99% 负样本时，全预测负也有 99% 准确率
- **用 F1、AUC-ROC、Precision-Recall**
- **处理方法**：过采样（SMOTE）、欠采样、类别权重

### 交叉验证
- **时序不能随机 K-Fold**：必须用 TimeSeriesSplit
- **分层交叉验证**：分类任务用 StratifiedKFold

## 模型速查

| 模型 | 优点 | 缺点 | 用途 |
|------|------|------|------|
| 随机森林 | 抗过拟合、特征重要性 | 外推弱 | 分类/回归/特征筛选 |
| XGBoost | 精度高、处理缺失值 | 需调参 | 竞赛首选 |
| SVM | 小样本好、高维有效 | 大数据慢 | 分类/文本 |
| K-means | 简单快速 | 需指定 K | 聚类 |
| DBSCAN | 自动发现簇数 | 密度变化差 | 异常检测 |
| Logistic 回归 | 可解释、概率输出 | 线性边界 | 二分类 |
| PCA | 保留主成分 | 线性降维 | 降维 |

## 评价指标

| 任务 | 指标 | 说明 |
|------|------|------|
| 分类（平衡）| Accuracy、F1 | — |
| 分类（不平衡）| F1、AUC-ROC | Accuracy 失效 |
| 回归 | RMSE、MAE、R² | RMSE 对异常值敏感 |
| 聚类 | 轮廓系数 | 越接近 1 越好 |
| 预测 | RMSE、MAE、MAPE | MAPE 有百分比含义 |

## SHAP 值

SHAP 比随机森林的 `feature_importances_` 更可靠：
- 后者对相关特征有偏
- SHAP 基于博弈论，理论上更严谨
- SHAP 可以解释单个预测

```python
import shap
explainer = shap.TreeExplainer(model)
shap_values = explainer.shap_values(X_test)
shap.summary_plot(shap_values, X_test)
```
