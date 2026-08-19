// Package main: 初始化 Elasticsearch 用于接收 yige 接口请求日志
// 一次性执行脚本，幂等：
//   1. 创建/更新 ILM 策略 yige-request-logs-policy（hot + delete 7d，无需 rollover，因为按天建索引）
//   2. 创建/更新 index template yige-request-logs（ECS 规范字段类型 + 关联 ILM）
//   3. 在 Kibana 创建 index pattern yige-request-logs-*（带 @timestamp 时间字段）
//   4. 清理旧 filebeat data stream 索引和 rollover 模式的 000001 索引
//
// filebeat 通过 output.elasticsearch.index: "yige-request-logs-%{+yyyy.MM.dd}"
// 直接按天建索引，每个索引创建 7 天后被 ILM 自动删除。
//
// 用法（环境变量未配置时给出占位符提示）：
//   ES_HOST=http://127.0.0.1:9200 KIBANA_HOST=http://127.0.0.1:5601 go run scripts/init_es.go
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

var (
	esHost     = getEnv("ES_HOST", "ES_HOST_PLACEHOLDER")
	kibanaHost = getEnv("KIBANA_HOST", "KIBANA_HOST_PLACEHOLDER")
)

func getEnv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

func httpDo(method, url, body string, headers map[string]string) (int, string) {
	req, _ := http.NewRequest(method, url, bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return -1, fmt.Sprintf("请求失败: %v", err)
	}
	defer resp.Body.Close()
	out, _ := io.ReadAll(resp.Body)
	return resp.StatusCode, truncate(out, 300)
}

func truncate(b []byte, n int) string {
	s := string(b)
	if len(s) > n {
		return s[:n] + "..."
	}
	return s
}

func main() {
	// 0. 探活
	fmt.Println("=== 0. 探活 Elasticsearch ===")
	code, body := httpDo(http.MethodGet, esHost, "", nil)
	fmt.Printf("  [%d] GET / → %s\n", code, body)

	// 1. ILM 策略：hot 不 rollover（按天建索引模式）+ delete 7d
	fmt.Println("\n=== 1. 创建 ILM 策略 yige-request-logs-policy ===")
	code, body = httpDo(http.MethodPut, esHost+"/_ilm/policy/yige-request-logs-policy", `{
  "policy": {
    "phases": {
      "hot": {
        "min_age": "0ms",
        "actions": {
          "set_priority": { "priority": 100 }
        }
      },
      "delete": {
        "min_age": "7d",
        "actions": {
          "delete": {}
        }
      }
    }
  }
}`, nil)
	fmt.Printf("  [%d] PUT _ilm/policy → %s\n", code, body)

	// 2. 索引模板：ECS 规范字段名（与服务器上 yige-server 输出对齐）
	// service.name / client.ip / user_agent.original 是 ECS 规范
	fmt.Println("\n=== 2. 创建 index template yige-request-logs（ECS 字段）===")
	code, body = httpDo(http.MethodPut, esHost+"/_index_template/yige-request-logs", `{
  "index_patterns": ["yige-request-logs-*"],
  "template": {
    "settings": {
      "number_of_shards": 1,
      "number_of_replicas": 0,
      "index.lifecycle.name": "yige-request-logs-policy"
    },
    "mappings": {
      "dynamic": false,
      "properties": {
        "@timestamp":          { "type": "date" },
        "service.name":        { "type": "keyword" },
        "method":              { "type": "keyword" },
        "path":                { "type": "keyword" },
        "query":               { "type": "text" },
        "status":              { "type": "integer" },
        "latency_ms":          { "type": "double" },
        "client.ip":           { "type": "ip" },
        "user_agent.original": { "type": "text" },
        "request_id":          { "type": "keyword" },
        "bytes_out":           { "type": "integer" },
        "error":               { "type": "text" }
      }
    }
  },
  "priority": 300
}`, nil)
	fmt.Printf("  [%d] PUT _index_template → %s\n", code, body)

	// 3. Kibana index pattern（先尝试创建，已存在会 409 但不影响）
	fmt.Println("\n=== 3. 创建 Kibana index pattern yige-request-logs-* ===")
	patternBody, _ := json.Marshal(map[string]map[string]string{
		"attributes": {
			"title":         "yige-request-logs-*",
			"timeFieldName": "@timestamp",
		},
	})
	code, body = httpDo(http.MethodPost, kibanaHost+"/api/saved_objects/index-pattern", string(patternBody), map[string]string{"kbn-xsrf": "true"})
	fmt.Printf("  [%d] POST kibana saved_objects → %s\n", code, body)
	if code == 409 {
		fmt.Println("  (已存在，跳过)")
	}

	// 4. 清理旧的 rollover 模式 000001 索引（之前 init_es.go 创建的）
	fmt.Println("\n=== 4. 删除旧的 yige-request-logs-000001 索引（rollover 模式遗留）===")
	code, body = httpDo(http.MethodDelete, esHost+"/yige-request-logs-000001", "", nil)
	fmt.Printf("  [%d] DELETE yige-request-logs-000001 → %s\n", code, body)

	// 5. 删除旧的 .ds-filebeat-* data stream（脏数据，全是 MySQL 报错请求）
	fmt.Println("\n=== 5. 删除旧 filebeat data stream（MySQL 报错期间的脏数据）===")
	code, body = httpDo(http.MethodDelete, esHost+"/_data_stream/filebeat-8.19.20", "", nil)
	fmt.Printf("  [%d] DELETE _data_stream/filebeat-8.19.20 → %s\n", code, body)

	// 6. 验收
	fmt.Println("\n=== 6. 验收 ===")
	fmt.Println("--- ILM policy yige-request-logs-policy ---")
	code, body = httpDo(http.MethodGet, esHost+"/_ilm/policy/yige-request-logs-policy", "", nil)
	fmt.Printf("  [%d] %s\n", code, body)
	fmt.Println("--- index template ---")
	code, body = httpDo(http.MethodGet, esHost+"/_index_template/yige-request-logs", "", nil)
	fmt.Printf("  [%d] %s\n", code, body)
	fmt.Println("--- 当前 yige-request-logs-* 索引 ---")
	code, body = httpDo(http.MethodGet, esHost+"/_cat/indices/yige-request-logs-*?v", "", nil)
	fmt.Printf("  [%d]\n%s\n", code, body)
	fmt.Println("--- filebeat data stream 是否还在 ---")
	code, body = httpDo(http.MethodGet, esHost+"/_data_stream/filebeat-8.19.20?v", "", nil)
	fmt.Printf("  [%d] %s\n", code, body)

	fmt.Println("\n=== 完成 ✓ ===")
	fmt.Println("下一步：SSH 上服务器改 /root/filebeat/filebeat.yml 加 output.elasticsearch.index + setup.ilm.enabled=false + setup.template.enabled=false，然后重启 filebeat 容器")
}
