package tools

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func ReadJSON(stringPath string) (Config, error) {

	var config Config

	file, err := os.Open(stringPath)
	if err != nil {
		fmt.Println("Error opening config file:", err)
		return config, err
	}
	defer file.Close()

	// 读取文件内容
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading config file:", err)
		return config, err
	}
	// 解析 JSON 数据

	err = json.Unmarshal(data, &config)
	if err != nil {
		fmt.Println("Error parsing config file:", err)
		return config, err
	}

	return config, err

}

type Config struct {
	ProxyIP       string `json:"proxyIP"`
	LocalListenIP string `json:"localListenIP"`
}
