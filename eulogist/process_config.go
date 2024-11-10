package Eulogist

import (
	"Eulogist/core/tools/skin_process"
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	"github.com/pterm/pterm"
)

// FileExist 检查 path 对应路径的文件是否存在。
// 如果不存在，或该路径指向一个文件夹，则返回假，否则返回真
func FileExist(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false, nil
	}
	if err != nil {
		return false, fmt.Errorf("FileExist: %v", err)
	}
	return !fileInfo.IsDir(), nil
}

// WriteJsonFile 将 content 以 JSON
// 形式写入到 path 指代的文件处
func WriteJsonFile(path string, content any) error {
	contentBytes, _ := json.Marshal(content)
	// 将内容转换为 JSON 格式
	buffer := bytes.NewBuffer([]byte{})
	json.Indent(buffer, contentBytes, "", "	")
	// 格式化 JSON
	err := os.WriteFile(path, buffer.Bytes(), 0600)
	if err != nil {
		return fmt.Errorf("WriteJsonFile: %v", err)
	}
	// 将 JSON 写入文件
	return nil
}

// ReadEulogistConfig 在当前目录读取赞颂者的配置文件。
// 如果没有对应的文件，则将尝试生成默认配置文件。
//
// 生成默认配置文件期间需要从控制台读取用户输入，
// 读取的内容包括需要进入的租赁服号及其密码，
// 以及 FastBuilder 原生验证服务器的 Token。
func ReadEulogistConfig() (*EulogistConfig, error) {
	var cfg EulogistConfig

	eulogistConfigExist, err := FileExist("eulogist_config.json")
	if err != nil {
		return nil, fmt.Errorf("ReadEulogistConfig: %v", err)
	}
	if !eulogistConfigExist {
		config, err := GenerateEulogistConfig()
		if err != nil {
			return nil, fmt.Errorf("ReadEulogistConfig: %v", err)
		}
		return config, nil
	}

	fileBytes, err := os.ReadFile("eulogist_config.json")
	if err != nil {
		return nil, fmt.Errorf("ReadEulogistConfig: %v", err)
	}

	err = json.Unmarshal(fileBytes, &cfg)
	if err != nil {
		return nil, fmt.Errorf("ReadEulogistConfig: %v", err)
	}

	return &cfg, nil
}

// GenerateEulogistConfig 在当前目录生成赞颂者的默认配置文件，
// 并返回该配置文件。此函数会从控制台读取用户输入，
// 读取的内容包括需要进入的租赁服号及其密码，
// 以及 FastBuilder 原生验证服务器的 Token
func GenerateEulogistConfig() (config *EulogistConfig, err error) {
	cfg := DefaultEulogistConfig()

	pterm.Info.Printf("Type your rental server code: ")
	fmt.Scanln(&cfg.RentalServerCode)

	pterm.Info.Printf("Type your rental server password: ")
	fmt.Scanln(&cfg.RentalServerPassword)

	pterm.Info.Printf("Type your token of PhoenixBuilder's or Liliya233's Auth Server: ")
	fmt.Scanln(&cfg.FBToken)

	err = WriteJsonFile("eulogist_config.json", cfg)
	if err != nil {
		return nil, fmt.Errorf("GenerateEulogistConfig: %v", err)
	}

	return &cfg, nil
}

// GenerateNetEaseConfig 根据皮肤路径 skinPath、
// 皮肤手臂是否纤细 skinIsSlim 及赞颂者开放的服务器地址，
// 在当前目录下生成用于启动 Minecraft 客户端的配置文件，
// 并返回该配置文件的绝对路径
func GenerateNetEaseConfig(
	skinPath string,
	skinIsSlim bool,
	ip string,
	port int,
) (configPath string, err error) {
	cfg := DefaultNetEaseConfig()

	cfg.RoomInfo.IP = ip
	cfg.RoomInfo.Port = port
	cfg.SkinInfo.Slim = skinIsSlim

	skinIsExist, err := FileExist(skinPath)
	if err != nil {
		return "", fmt.Errorf("GenerateNetEaseConfig: %v", err)
	}

	if !skinIsExist {
		currentPath, _ := os.Getwd()
		cfg.SkinInfo.SkinPath = fmt.Sprintf(`%s\steve.png`, currentPath)
		err = os.WriteFile("steve.png", skin_process.SteveSkin, 0600)
		if err != nil {
			return "", fmt.Errorf("GenerateNetEaseConfig: %v", err)
		}
	} else {
		cfg.SkinInfo.SkinPath = skinPath
	}

	err = WriteJsonFile("netease.cppconfig", cfg)
	if err != nil {
		return "", fmt.Errorf("GenerateNetEaseConfig: %v", err)
	}

	configPath, _ = os.Getwd()
	configPath = fmt.Sprintf(`%s\netease.cppconfig`, configPath)

	return
}
