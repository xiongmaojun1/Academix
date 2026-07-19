// Academix — 学术副驾驶
// 独立应用，不依赖 Reasonix
package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const version = "1.0.0"

func main() {
	fmt.Println("Academix — 学术副驾驶 v" + version)
	fmt.Println("输入 /help 查看可用命令，输入 /quit 退出")
	fmt.Println()

	// 获取 Academix 根目录
	exe, _ := os.Executable()
	root := filepath.Dir(exe)

	// 检查 Skills 目录
	skillsDir := filepath.Join(root, "skills")
	if _, err := os.Stat(skillsDir); os.IsNotExist(err) {
		// 尝试当前目录
		skillsDir = "skills"
	}

	// 检查知识库目录
	knowledgeDir := filepath.Join(root, "knowledge")
	if _, err := os.Stat(knowledgeDir); os.IsNotExist(err) {
		knowledgeDir = "knowledge"
	}

	// 检查资产目录
	assetsDir := filepath.Join(root, "assets")
	if _, err := os.Stat(assetsDir); os.IsNotExist(err) {
		assetsDir = "assets"
	}

	// 初始化状态文件
	stateFile := "academix-state.json"
	if _, err := os.Stat(stateFile); os.IsNotExist(err) {
		initState(stateFile)
	}

	// REPL
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		if !reader.Scan() {
			break
		}

		input := strings.TrimSpace(reader.Text())
		if input == "" {
			continue
		}

		// 处理命令
		if strings.HasPrefix(input, "/") {
			if handleCommand(input, skillsDir, knowledgeDir, assetsDir, stateFile) {
				return
			}
			continue
		}

		// 处理普通消息
		handleMessage(input, skillsDir, knowledgeDir, assetsDir, stateFile)
	}
}

// initState 初始化状态文件
func initState(path string) {
	state := `{
  "current_phase": "init",
  "completed_phases": [],
  "decisions": [],
  "results": {},
  "blockers": []
}`
	os.WriteFile(path, []byte(state), 0644)
}

// handleCommand 处理斜杠命令，返回 true 表示退出
func handleCommand(cmd, skillsDir, knowledgeDir, assetsDir, stateFile string) bool {
	parts := strings.Fields(cmd)
	command := parts[0]

	switch command {
	case "/quit", "/exit":
		fmt.Println("再见！")
		return true

	case "/help":
		showHelp()

	case "/status":
		showStatus(stateFile)

	case "/skills":
		showSkills(skillsDir)

	case "/knowledge":
		showKnowledge(knowledgeDir)

	case "/export-log":
		exportLog(stateFile)

	case "/clear":
		initState(stateFile)
		fmt.Println("[INFO] 状态已重置")

	default:
		fmt.Printf("[ERROR] 未知命令: %s\n", command)
	}
	return false
}

// showHelp 显示帮助信息
func showHelp() {
	help := `
可用命令:
  /help        显示此帮助
  /status      显示当前状态
  /skills      列出可用 Skills
  /knowledge   列出知识库文件
  /export-log  导出研究日志
  /clear       重置状态
  /quit        退出

直接输入文字与 Academix 对话。
`
	fmt.Print(help)
}

// showStatus 显示当前状态
func showStatus(stateFile string) {
	data, err := os.ReadFile(stateFile)
	if err != nil {
		fmt.Println("[ERROR] 无法读取状态文件")
		return
	}
	fmt.Println("当前状态:")
	fmt.Println(string(data))
}

// showSkills 列出可用 Skills
func showSkills(skillsDir string) {
	fmt.Println("可用 Skills:")
	entries, err := os.ReadDir(skillsDir)
	if err != nil {
		fmt.Println("[ERROR] 无法读取 Skills 目录")
		return
	}
	for _, entry := range entries {
		if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".md") {
			fmt.Printf("  - %s\n", entry.Name())
		}
	}
}

// showKnowledge 列出知识库文件
func showKnowledge(knowledgeDir string) {
	fmt.Println("知识库文件:")
	filepath.Walk(knowledgeDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".md") {
			rel, _ := filepath.Rel(knowledgeDir, path)
			fmt.Printf("  - %s\n", rel)
		}
		return nil
	})
}

// exportLog 导出研究日志
func exportLog(stateFile string) {
	data, err := os.ReadFile(stateFile)
	if err != nil {
		fmt.Println("[ERROR] 无法读取状态文件")
		return
	}

	logFile := "academix-research-log.md"
	content := "# Academix 研究日志\n\n"
	content += "## 状态\n\n```json\n" + string(data) + "\n```\n"

	os.WriteFile(logFile, []byte(content), 0644)
	fmt.Printf("[INFO] 日志已导出到 %s\n", logFile)
}

// handleMessage 处理用户消息
func handleMessage(input, skillsDir, knowledgeDir, assetsDir, stateFile string) {
	// 记录用户消息
	fmt.Printf("[Academix] 收到: %s\n", input)

	// 分析问题类型
	questionType := analyzeQuestion(input)

	// 读取对应的知识库
	knowledge := readKnowledge(questionType, knowledgeDir)

	// 生成响应
	response := generateResponse(input, questionType, knowledge)

	// 输出响应
	fmt.Printf("\nAcademix: %s\n", response)

	// 更新状态
	updateState(stateFile, input, response)
}

// analyzeQuestion 分析问题类型
func analyzeQuestion(input string) string {
	input = strings.ToLower(input)

	if strings.Contains(input, "评价") || strings.Contains(input, "排序") || strings.Contains(input, "ahp") || strings.Contains(input, "topsis") {
		return "evaluation"
	}
	if strings.Contains(input, "预测") || strings.Contains(input, "时序") || strings.Contains(input, "arima") {
		return "prediction"
	}
	if strings.Contains(input, "优化") || strings.Contains(input, "规划") || strings.Contains(input, "调度") {
		return "optimization"
	}
	if strings.Contains(input, "微分方程") || strings.Contains(input, "动力学") || strings.Contains(input, "蒙特卡洛") {
		return "dynamics"
	}
	if strings.Contains(input, "图论") || strings.Contains(input, "最短路径") || strings.Contains(input, "网络") {
		return "graph"
	}
	if strings.Contains(input, "统计") || strings.Contains(input, "机器学习") || strings.Contains(input, "分类") {
		return "statistics"
	}
	if strings.Contains(input, "论文") || strings.Contains(input, "写作") || strings.Contains(input, "摘要") {
		return "writing"
	}
	if strings.Contains(input, "代码") || strings.Contains(input, "编程") || strings.Contains(input, "python") {
		return "coding"
	}
	return "general"
}

// readKnowledge 读取对应的知识库文件
func readKnowledge(questionType, knowledgeDir string) string {
	var filePath string
	switch questionType {
	case "evaluation":
		filePath = filepath.Join(knowledgeDir, "modeling", "evaluation.md")
	case "prediction":
		filePath = filepath.Join(knowledgeDir, "modeling", "prediction.md")
	case "optimization":
		filePath = filepath.Join(knowledgeDir, "modeling", "optimization.md")
	case "dynamics":
		filePath = filepath.Join(knowledgeDir, "modeling", "dynamics.md")
	case "graph":
		filePath = filepath.Join(knowledgeDir, "modeling", "graph.md")
	case "statistics":
		filePath = filepath.Join(knowledgeDir, "modeling", "statistics.md")
	case "writing":
		filePath = filepath.Join(knowledgeDir, "writing", "norms.md")
	case "coding":
		filePath = filepath.Join(knowledgeDir, "coding", "norms.md")
	default:
		filePath = filepath.Join(knowledgeDir, "CATALOG.md")
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		return ""
	}
	return string(data)
}

// generateResponse 生成响应
func generateResponse(input, questionType, knowledge string) string {
	// 基于问题类型生成响应
	switch questionType {
	case "evaluation":
		return fmt.Sprintf("这是一个评价类问题。根据知识库，我可以帮你选择合适的评价模型（AHP、TOPSIS、熵权法等）。你想评价什么对象？有多少个指标？")
	case "prediction":
		return fmt.Sprintf("这是一个预测类问题。根据知识库，我可以帮你选择合适的预测模型（ARIMA、LSTM、灰色预测等）。你有多少数据点？数据有什么特征？")
	case "optimization":
		return fmt.Sprintf("这是一个优化问题。根据知识库，我可以帮你选择合适的优化方法（线性规划、整数规划、遗传算法等）。你的目标函数和约束是什么？")
	case "dynamics":
		return fmt.Sprintf("这是一个动力学问题。根据知识库，我可以帮你建立微分方程模型或蒙特卡洛仿真。你的系统有什么状态变量？")
	case "graph":
		return fmt.Sprintf("这是一个图论问题。根据知识库，我可以帮你选择合适的算法（Dijkstra、最大流、TSP等）。你的图有什么特征？")
	case "statistics":
		return fmt.Sprintf("这是一个统计/机器学习问题。根据知识库，我可以帮你选择合适的模型和评估指标。你的数据是什么类型的？")
	case "writing":
		return fmt.Sprintf("这是一个写作问题。根据知识库，我可以帮你规范论文结构、引用格式、图表规范。你想写什么类型的论文？")
	case "coding":
		return fmt.Sprintf("这是一个编码问题。根据知识库，我可以帮你规范代码、记录结果、避免常见错误。你想实现什么模型？")
	default:
		return fmt.Sprintf("你好！我是 Academix，你的学术副驾驶。关于「%s」，请告诉我你想研究什么问题，或者你现在处于研究的哪个阶段？", input)
	}
}

// updateState 更新状态文件
func updateState(stateFile, input, response string) {
	// 读取现有状态
	data, err := os.ReadFile(stateFile)
	if err != nil {
		return
	}

	// 简单追加决策记录
	// 实际实现应该解析 JSON 并更新
	_ = data
	_ = input
	_ = response
}
