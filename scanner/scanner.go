package scanner

import (
	"log"
	"os/exec"
)

//func checkTools() {
// 检查lsb_release 工具是否安装
//}

func cveScan() {
	// ubuntu上安装openSCAP: sudo apt install libopenscap8 -y
	// 安装基线库文件 wget https://security-metadata.canonical.com/oval/com.ubuntu.$(lsb_release -cs).usn.oval.xml.bz2

	getCodeNameCommand := exec.Command("lsb_release", "-cs")

	stdoutStderr, err := getCodeNameCommand.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%s\n", stdoutStderr)

	log.Println("openSCAP bash")
	cmd := exec.Command("crontab", "-h")

	cmd.Run()
}

func init() {
	cveScan()
}
