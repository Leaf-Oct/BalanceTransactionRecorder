package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

type Transaction struct {
	Date        string
	Description string
	IsExpand    bool
	Number      float64
	Balance     float64
}

func main() {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	fmt.Println("可执行文件:", ex)

	exec_path := filepath.Dir(ex)
	fmt.Println("路径:", exec_path)

	conf_file := exec_path + "/balance_file"
	fmt.Println("配置文件:", conf_file)

	var balance float64
	content, err := os.ReadFile(conf_file)
	if err != nil {
		fmt.Println(err)
		fmt.Println("未检测到上一次文件，将创建")

		fmt.Println("输入初始余额")
		n, err := fmt.Scan(&balance)
		if n == 0 && err != nil {
			panic("你寄吧输入了什么东西")
			os.Exit(-1)
		}
		file, err := os.Create(conf_file)
		if err != nil {
			panic("配置文件创建失败，没有权限?")
		}
		//TODO 转成字节数组，写入文件中
		file.Write(balance)
	} else {
		var balance_string string = string(content)
		fmt.Println("识别到上一次文件，内容为:", balance_string)
		balance, err = strconv.ParseFloat(balance_string, 64)
		if err != nil {
			panic("你几把在文件里写了什么东西? ")
			os.Exit(-1)
		}
		fmt.Println("余额为", balance)
	}
	// a := Transaction{"2023-10-5", "阿巴阿巴", true, 12.26, 233.3}
	// transactionToJsonObj(a)
	for true {
		var date string
		var description string
		var is_expand bool
		var number float64

	}
}

func transactionToJsonObj(transaction_item Transaction) string {
	v, err := json.Marshal(transaction_item)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(v))
	return string(v)
}
