package main

import (
	"fmt"
	"lpcve/app/routers"
	_ "lpcve/bootstrap"
)

func main() {
	fmt.Println("main 函数入口")

	routers.InitRouter().Run()
}
