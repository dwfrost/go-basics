package httpclient

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/oauth2"
)

// DemonstrateAuthentication 展示HTTP认证
func DemonstrateAuthentication() {
	// 基本认证
	fmt.Println("4.1 基本认证")
	basicAuth()

	// Bearer令牌认证
	fmt.Println("\n4.2 Bearer令牌认证")
	bearerTokenAuth()

	// API密钥认证
	fmt.Println("\n4.3 API密钥认证")
	apiKeyAuth()

	// OAuth2认证 (仅展示概念)
	fmt.Println("\n4.4 OAuth2认证 (概念示例)")
	fmt.Println("OAuth2认证流程较为复杂，通常需要以下步骤:")
	fmt.Println("1. 获取授权码")
	fmt.Println("2. 使用授权码获取访问令牌")
	fmt.Println("3. 使用访问令牌调用API")
	fmt.Println("4. 刷新访问令牌")
}

// 基本认证
func basicAuth() {
	// 创建一个HTTP客户端
	client := &http.Client{}

	// 创建请求
	req, err := http.NewRequest("GET", "https://httpbun.com/basic-auth/user/passwd", nil)
	if err != nil {
		fmt.Printf("创建请求失败: %v\n", err)
		return
	}

	// 添加基本认证头
	username := "user"
	password := "passwd"
	auth := username + ":" + password
	encodedAuth := base64.StdEncoding.EncodeToString([]byte(auth))
	req.Header.Add("Authorization", "Basic "+encodedAuth)

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("发送请求失败: %v\n", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("读取响应失败: %v\n", err)
		return
	}

	// 打印响应
	fmt.Printf("状态码: %d\n", resp.StatusCode)
	fmt.Printf("响应体: %s\n", body)
}

// Bearer令牌认证
func bearerTokenAuth() {
	// 创建一个HTTP客户端
	client := &http.Client{}

	// 创建请求
	req, err := http.NewRequest("GET", "https://httpbun.com/bearer/valid-token", nil)
	if err != nil {
		fmt.Printf("创建请求失败: %v\n", err)
		return
	}

	// 添加Bearer令牌头
	token := "valid-token"
	req.Header.Add("Authorization", "Bearer "+token)

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("发送请求失败: %v\n", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("读取响应失败: %v\n", err)
		return
	}

	// 打印响应
	fmt.Printf("状态码: %d\n", resp.StatusCode)
	fmt.Printf("响应体: %s\n", body)
}

// API密钥认证
func apiKeyAuth() {
	// 方法1: 在URL中添加API密钥
	apiKey := "abcdef123456"
	url := fmt.Sprintf("https://httpbun.com/get?api_key=%s", apiKey)

	// 发送请求
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("发送请求失败: %v\n", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("读取响应失败: %v\n", err)
		return
	}

	// 打印响应
	fmt.Printf("方法1 (URL参数) - 状态码: %d\n", resp.StatusCode)
	if len(body) > 100 {
		fmt.Printf("响应体 (前100字节): %s...\n", body[:100])
	} else {
		fmt.Printf("响应体: %s\n", body)
	}

	// 方法2: 在请求头中添加API密钥
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://httpbun.com/get", nil)
	if err != nil {
		fmt.Printf("创建请求失败: %v\n", err)
		return
	}

	// 添加API密钥到请求头
	req.Header.Add("X-API-Key", apiKey)

	// 发送请求
	resp2, err := client.Do(req)
	if err != nil {
		fmt.Printf("发送请求失败: %v\n", err)
		return
	}
	defer resp2.Body.Close()

	// 读取响应
	body2, err := io.ReadAll(resp2.Body)
	if err != nil {
		fmt.Printf("读取响应失败: %v\n", err)
		return
	}

	// 打印响应
	fmt.Printf("\n方法2 (请求头) - 状态码: %d\n", resp2.StatusCode)
	if len(body2) > 100 {
		fmt.Printf("响应体 (前100字节): %s...\n", body2[:100])
	} else {
		fmt.Printf("响应体: %s\n", body2)
	}
}

// OAuth2 客户端示例
type OAuth2Config struct {
	ClientID     string
	ClientSecret string
	RedirectURI  string
	AuthURL      string
	TokenURL     string
	Scopes       []string
}

// 1. 获取授权码
func getAuthorizationCode(config OAuth2Config) string {
	// 构建授权URL
	authURL := fmt.Sprintf("%s?client_id=%s&redirect_uri=%s&response_type=code&scope=%s",
		config.AuthURL,
		config.ClientID,
		url.QueryEscape(config.RedirectURI),
		strings.Join(config.Scopes, " "))

	// 重定向用户到授权页面
	fmt.Printf("请访问此URL进行授权: %s\n", authURL)

	// 等待用户授权并获取授权码
	var code string
	fmt.Print("请输入授权码: ")
	fmt.Scan(&code)
	return code
}

// 2. 使用授权码获取访问令牌
func getAccessToken(config OAuth2Config, code string) (*oauth2.Token, error) {
	// 构建获取令牌的请求
	data := url.Values{}
	data.Set("grant_type", "authorization_code")
	data.Set("code", code)
	data.Set("client_id", config.ClientID)
	data.Set("client_secret", config.ClientSecret)
	data.Set("redirect_uri", config.RedirectURI)

	// 发送请求
	resp, err := http.PostForm(config.TokenURL, data)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 解析响应
	var token oauth2.Token
	if err := json.NewDecoder(resp.Body).Decode(&token); err != nil {
		return nil, err
	}

	return &token, nil
}

// 3. 使用访问令牌调用API
func callAPI(token *oauth2.Token) {
	req, _ := http.NewRequest("GET", "https://api.example.com/user", nil)
	req.Header.Add("Authorization", "Bearer "+token.AccessToken)
}

// OAuth2 的令牌刷新机制
func refreshToken(token *oauth2.Token, config OAuth2Config) (*oauth2.Token, error) {
	if !token.Valid() {
		// 使用刷新令牌获取新的访问令牌
		data := url.Values{}
		data.Set("grant_type", "refresh_token")
		data.Set("refresh_token", token.RefreshToken)
		data.Set("client_id", config.ClientID)
		data.Set("client_secret", config.ClientSecret)

		// 发送刷新请求...
	}
	return token, nil
}
