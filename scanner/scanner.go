package scanner

import (
	"io"
	"log"
	"lpcve/app/variable"
	"net/http"
	"os"
	"os/exec"
)

func getSystemCodeName() (string, error) {
	getCodeNameCommand := exec.Command("lsb_release", "-cs")
	stdoutStderr, err := getCodeNameCommand.CombinedOutput()

	return string(stdoutStderr), err
}

func downloadXML(path string, fielUrl string) {
	log.Println("下载文件")

	req, err := http.NewRequest("GET", fielUrl, *(new(io.Reader)))
	if err != nil {
		log.Fatal(err)
	}

	httpProxyAddr, err := http.ProxyFromEnvironment(req)
	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(httpProxyAddr),
		},
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	// 创建xml文件
	out, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}

	// 这个函数并不好，后面需要更改
	io.Copy(out, resp.Body)
}

func checkCompressedXML() {

	// 判断文件是否存在，如果不存在则需要下载
	codeName, err := getSystemCodeName()
	if err != nil {
		log.Fatal(err)
	}
	compressedXMLName := variable.CompressedXMLPrefix + codeName + variable.CompressedXMLSuffix
	xmlFullPath := variable.OVALLcationPath + "/" + compressedXMLName
	_, err = os.Stat(xmlFullPath)
	if err != nil {
		// 文件不存在，需要下载
		url := variable.UbuntuOVALDownloadUrl + "/" + compressedXMLName
		downloadXML(xmlFullPath, url)
	}
}

func cveScan() {
	// ubuntu上安装openSCAP: sudo apt install libopenscap8 -y
	// 安装基线库文件 wget https://security-metadata.canonical.com/oval/com.ubuntu.$(lsb_release -cs).usn.oval.xml.bz2

	log.Println("openSCAP bash")
	cmd := exec.Command("crontab", "-h")

	cmd.Run()
}

func init() {
	checkCompressedXML()
	cveScan()
}
