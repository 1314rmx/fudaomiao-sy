package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// 定义一个结构体，用于匹配JSON文件的格式
type DirStructure struct {
	Text []string `json:"text"`
}

func main() {
	// JSON文件路径
	jsonFilePath := "test/dir.json"

	// 读取JSON文件的内容
	fileContent, err := ioutil.ReadFile(jsonFilePath)
	if err != nil {
		fmt.Printf("Error reading JSON file: %s\n", err)
		return
	}

	// 解析JSON内容
	var dirs DirStructure
	if err := json.Unmarshal(fileContent, &dirs); err != nil {
		fmt.Printf("Error parsing JSON file: %s\n", err)
		return
	}

	// 创建目录
	for _, dir := range dirs.Text {
		if err := os.Mkdir(dir, 0755); err != nil {
			fmt.Printf("Error creating directory '%s': %s\n", dir, err)
		} else {
			fmt.Printf("Directory '%s' created successfully.\n", dir)
		}
	}
}
