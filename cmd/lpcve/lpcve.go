package main

import (
	"fmt"
	"lpcve/app/routers"
	_ "lpcve/bootstrap"
	_ "lpcve/scanner"
)

func main() {
	fmt.Println("main 函数入口")

	routers.InitRouter().Run()
}
