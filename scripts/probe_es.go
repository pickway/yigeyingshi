package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

const esURL = "http://118.145.113.88:9200"

func main() {
	// 1. ES 健康检查
	fmt.Println("=== 1. ES 健康 ===")
	get(esURL+"/_cluster/health", nil)

	// 2. ES 版本信息（判断是否为 ES 8.x）
	fmt.Println("\n=== 2. ES 版本 ===")
	body := get(esURL, nil)
	var v map[string]interface{}
	if err := json.Unmarshal(body, &v); err == nil {
		if ver, ok := v["version"].(map[string]interface{}); ok {
			if num, ok := ver["number"].(string); ok {
				fmt.Println("  number:", num)
			}
			if bp, ok := ver["build_flavor"].(string); ok {
				fmt.Println("  build_flavor:", bp)
			}
		}
	}

	// 3. 是否启用 xpack（ILM 是 xpack 功能）
	fmt.Println("\n=== 3. X-Pack 状态 ===")
	get(esURL+"/_xpack", nil)

	// 4. 当前 ILM 策略列表
	fmt.Println("\n=== 4. 现有 ILM 策略 ===")
	get(esURL+"/_ilm/policy", nil)

	// 5. 当前索引模板列表
	fmt.Println("\n=== 5. 索引模板 ===")
	get(esURL+"/_index_template", nil)
}

func get(url string, payload map[string]interface{}) []byte {
	var body io.Reader
	if payload != nil {
		data, _ := json.Marshal(payload)
		body = bytes.NewReader(data)
	}
	req, err := http.NewRequest("GET", url, body)
	if err != nil {
		log.Printf("构造请求失败 %s: %v", url, err)
		return nil
	}
	if payload != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("请求失败 %s: %v", url, err)
		return nil
	}
	defer resp.Body.Close()
	b, _ := io.ReadAll(resp.Body)
	fmt.Printf("  [%s] %s -> %d bytes\n", http.MethodGet, url, len(b))
	if len(b) > 0 && len(b) < 2000 {
		fmt.Println("  body:", string(b))
	} else if len(b) >= 2000 {
		fmt.Println("  body: (truncated)", string(b[:2000]), "...")
	}
	return b
}
