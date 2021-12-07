package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var Port string

func main() {
	Flag()
	Parse(Port)
	webServ(Port)
}

// httpListen & Print
func webServ(Port string) {

	helloHandler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("-------------------- START --------------------")
		fmt.Printf("RemoteAddr   ：%v\n\n", r.RemoteAddr) // 获取请求源的Addr
		fmt.Printf("UrlPath      ：%v%v\n", r.Host, r.RequestURI)
		fmt.Printf("Method       ：%v\n", r.Method)
		fmt.Printf("Proto        ：%v\n", r.Proto)
		for key, value := range r.Header {
			fmt.Printf("%s : %s\n", key, value)
		}
		// fmt.Printf("Header     is：%v\n\n",r.Header)

		body, _ := ioutil.ReadAll(r.Body)
		if body != nil {
			fmt.Print("\n")
			fmt.Printf("Body       is：%v\n", string(body))
		}
		fmt.Println("--------------------  END  --------------------")
		fmt.Println("	持续监听中---")
		io.WriteString(w, "<title>Request Success</title><h1>请求成功</h1>\nRequest Success!\n\n")
	}
	http.HandleFunc("/", helloHandler) // 匹配请求路径
	getPort := ":" + Port
	log.Fatal(http.ListenAndServe(getPort, nil)) //   getPort = ":7777"
}

// banner
func Banner() {
	banner := `

		█ █ ███ ███ ███ █  █ ███ ███ ███ █   █ 
		█ █  █   █  █ █ █    █    █  █   ██  █ 
		███  █   █  ███ █  █ ███  █  ███ █ █ █ 
		█ █  █   █  █   █  █   █  █  █   █  ██ 
		█ █  █   █  █   ██ █ ███  █  ███ █   █ HttpListen
							Version: 1.0.0
							Author: 0x25w
							Date: 2021/12/06
`
	print(banner)
}

// input & usage
func Flag() {
	Banner()
	flag.StringVar(&Port, "p", "", "设置端口,建议使用不常用的端口,端口范围: 1-65535")
	flag.Parse()
}

// 校验空值
func Parse(Port string) {
	if Port == "" {
		fmt.Println("         请先设置端口！ 示例： httplisten0.1.exe -p 8888             ")
		os.Exit(1)
	}
	fmt.Println("         端口设置成功，现在可以进行请求了！             ")
	fmt.Println("         本地测试URL=HTTP://127.0.0.1:" + Port)
}
