package variable

import "log"

var (
	//BasePath           string

	// 这里先写成固定值，测试代码
	OVALLcationPath       = "/home/jiang"
	UbuntuOVALDownloadUrl string
	CompressedXMLPrefix   = "com.ubuntu."
	CompressedXMLSuffix   = ".usn.oval.xml.bz2"
	XMLUrl                string

	//系统环境变量
)

func init() {
	log.Println("this a new file")
}
