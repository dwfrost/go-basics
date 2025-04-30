package httpclient

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/joho/godotenv"
)

// DemonstrateThirdPartyAPIs 展示调用第三方API
func DemonstrateThirdPartyAPIs() {
	// 加载环境变量
	godotenv.Load()

	// 调用天气API
	// fmt.Println("5.1 调用天气API")
	// callWeatherAPI()

	// // 调用汇率API
	// fmt.Println("\n5.2 调用汇率API")
	// callExchangeRateAPI()

	// 调用GitHub API
	// fmt.Println("\n5.3 调用GitHub API")
	// callGitHubAPI()

	// // 调用新闻API
	fmt.Println("\n5.4 调用新闻API")
	callNewsAPI()
}

// 调用天气API
func callWeatherAPI() {
	// 创建请求URL
	baseURL := "https://api.openweathermap.org/data/2.5/weather"

	// 添加查询参数
	params := url.Values{}
	params.Add("q", "Beijing,cn")
	params.Add("units", "metric")
	params.Add("lang", "zh_cn")

	// 添加API密钥
	apiKey := os.Getenv("OPENWEATHER_API_KEY")
	fmt.Printf("当前使用的 API Key: %s\n", apiKey)
	if apiKey == "" {
		fmt.Println("未设置OPENWEATHER_API_KEY环境变量，使用模拟数据")
		fmt.Println("模拟数据: 北京，温度25°C，天气晴朗，湿度45%")
		return
	}
	params.Add("appid", apiKey)

	// 构建完整URL
	requestURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

	// 创建HTTP客户端
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	// 发送请求
	resp, err := client.Get(requestURL)
	if err != nil {
		fmt.Printf("请求天气API失败: %v\n", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("读取响应失败: %v\n", err)
		return
	}

	// 检查状态码
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("API返回错误状态码: %d, 响应: %s\n", resp.StatusCode, body)
		return
	}

	// 解析JSON响应
	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Printf("解析JSON失败: %v\n", err)
		return
	}

	// 提取天气信息
	if result["name"] != nil {
		fmt.Printf("城市: %s\n", result["name"])
	}

	if main, ok := result["main"].(map[string]interface{}); ok {
		if temp, ok := main["temp"].(float64); ok {
			fmt.Printf("温度: %.1f°C\n", temp)
		}
		if humidity, ok := main["humidity"].(float64); ok {
			fmt.Printf("湿度: %.0f%%\n", humidity)
		}
	}

	if weather, ok := result["weather"].([]interface{}); ok && len(weather) > 0 {
		if weatherMap, ok := weather[0].(map[string]interface{}); ok {
			if description, ok := weatherMap["description"].(string); ok {
				fmt.Printf("天气: %s\n", description)
			}
		}
	}
}

// 调用汇率API
func callExchangeRateAPI() {
	// 创建请求URL
	baseURL := "https://api.exchangerate.host/live"

	// 添加查询参数
	params := url.Values{}
	params.Add("source", "CNY")
	params.Add("currencies", "USD,EUR,JPY,GBP")
	params.Add("access_key", os.Getenv("EXCHANGERATE_API_KEY"))

	// 构建完整URL
	requestURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

	// 创建HTTP客户端
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	// 发送请求
	resp, err := client.Get(requestURL)
	if err != nil {
		fmt.Printf("请求汇率API失败: %v\n", err)
		fmt.Println("使用模拟数据:")
		fmt.Println("基准货币: CNY (人民币)")
		fmt.Println("USD: 0.1400")
		fmt.Println("EUR: 0.1300")
		fmt.Println("JPY: 15.5000")
		fmt.Println("GBP: 0.1100")
		return
	}
	defer resp.Body.Close()

	// 读取响应
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("读取响应失败: %v\n", err)
		return
	}

	// 检查状态码
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("API返回错误状态码: %d, 响应: %s\n", resp.StatusCode, body)
		return
	}

	// 解析JSON响应
	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Printf("解析JSON失败: %v\n", err)
		return
	}

	// 提取汇率信息
	fmt.Printf("基准货币: %s\n", result["source"])
	// 将时间戳转换为本地日期格式
	if timestamp, ok := result["timestamp"].(float64); ok {
		// 将时间戳转换为time.Time
		t := time.Unix(int64(timestamp), 0)
		// 格式化为本地日期时间
		fmt.Printf("日期: %s\n", t.Format("2006-01-02 15:04:05"))
	} else {
		fmt.Printf("日期: %v\n", result["timestamp"])
	}

	if rates, ok := result["quotes"].(map[string]interface{}); ok {
		for currency, rate := range rates {
			fmt.Printf("%s: %.4f\n", currency, rate)
		}
	}
}

// 调用GitHub API
func callGitHubAPI() {
	// 创建请求URL
	requestURL := "https://api.github.com/repos/golang/go"

	// 创建HTTP客户端
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	// 创建请求
	req, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		fmt.Printf("创建请求失败: %v\n", err)
		return
	}

	// 添加请求头
	req.Header.Add("Accept", "application/vnd.github.v3+json")
	req.Header.Add("User-Agent", "Go-HTTP-Client")

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("请求GitHub API失败: %v\n", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("读取响应失败: %v\n", err)
		return
	}

	// 检查状态码
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("API返回错误状态码: %d, 响应: %s\n", resp.StatusCode, body)
		return
	}

	// 解析JSON响应
	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Printf("解析JSON失败: %v\n", err)
		return
	}

	// 提取仓库信息
	fmt.Printf("仓库名称: %s\n", result["full_name"])
	fmt.Printf("描述: %s\n", result["description"])
	fmt.Printf("星标数: %.0f\n", result["stargazers_count"])
	fmt.Printf("Fork数: %.0f\n", result["forks_count"])
	fmt.Printf("开放的Issue数: %.0f\n", result["open_issues_count"])
	fmt.Printf("语言: %s\n", result["language"])
}

// 调用新闻API
func callNewsAPI() {
	// 创建请求URL
	baseURL := "https://newsapi.org/v2/top-headlines"

	// 添加查询参数
	params := url.Values{}
	params.Add("country", "us")
	params.Add("category", "technology")

	// 添加API密钥
	apiKey := os.Getenv("NEWS_API_KEY")
	if apiKey == "" {
		fmt.Println("未设置NEWS_API_KEY环境变量，使用模拟数据")
		fmt.Println("模拟数据:")
		fmt.Println("标题: 中国科技创新取得重大突破")
		fmt.Println("来源: 科技日报")
		fmt.Println("发布时间: 2023-05-15")
		fmt.Println("---")
		fmt.Println("标题: 人工智能在医疗领域的应用")
		fmt.Println("来源: 中国科学报")
		fmt.Println("发布时间: 2023-05-14")
		return
	}
	params.Add("apiKey", apiKey)

	// 构建完整URL
	requestURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

	// 创建HTTP客户端
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	// 发送请求
	resp, err := client.Get(requestURL)
	if err != nil {
		fmt.Printf("请求新闻API失败: %v\n", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("读取响应失败: %v\n", err)
		return
	}

	// 检查状态码
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("API返回错误状态码: %d, 响应: %s\n", resp.StatusCode, body)
		return
	}

	// 解析JSON响应
	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Printf("解析JSON失败: %v\n", err)
		return
	}

	// 提取新闻信息
	if articles, ok := result["articles"].([]interface{}); ok {
		fmt.Printf("获取到 %d 条新闻\n", len(articles))

		// 显示前3条新闻
		count := 0
		for _, article := range articles {
			if count >= 3 {
				break
			}

			if articleMap, ok := article.(map[string]interface{}); ok {
				fmt.Printf("标题: %s\n", articleMap["title"])

				if source, ok := articleMap["source"].(map[string]interface{}); ok {
					fmt.Printf("来源: %s\n", source["name"])
				}

				fmt.Printf("发布时间: %s\n", articleMap["publishedAt"])
				fmt.Println("---")

				count++
			}
		}
	}
}
