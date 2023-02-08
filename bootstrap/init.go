package bootstrap

import (
	my_errors "ipcve/app/utils/errors"
	"log"
	"os"
)

func checkRequiredFolders() {
	//1.检查配置文件是否存在
	if _, err := os.Stat("./config/config.yml"); err != nil {
		log.Fatal(my_errors.ErrorsConfigYamlNotExists + err.Error())
	}
}

func init() {
	// 检查文件是否存在，不存在则创建
	checkRequiredFolders()
	log.Println("Test.....")
}
