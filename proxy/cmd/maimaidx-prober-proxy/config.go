package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Diving-Fish/maimaidx-prober/proxy/lib"
)

type workingMode int

const (
	workingModeUpdate workingMode = 0
	workingModeExport workingMode = 1 // only for debug or other
)

type config struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	Mode     string `json:"mode,omitempty"`
}

func (c *config) getWorkingMode() workingMode {
	if c.Mode == "export" {
		return workingModeExport
	}
	return workingModeUpdate
}

func initConfig(path string) (config, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		// First run
		lib.GenerateCert()
		os.WriteFile(path, []byte("{\"username\": \"\", \"password\": \"\"}"), 0644)
		return config{}, fmt.Errorf("初次使用请填写 %s 文件，并依据教程完成根证书的安装。", path)
	}

	var obj config
	err = json.Unmarshal(b, &obj)
	if err != nil {
		return config{}, fmt.Errorf("配置文件格式有误，无法解析：%w。请检查 %s 文件的内容", err, path)
	}

	return obj, nil
}
