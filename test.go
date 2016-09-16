package main

import (
	"fmt"
	"io/ioutil"
	h "net/http"
	"os"
)

func main() {
	// var a int
	// a=10
	// a=a+1
	// fmt.Println("hello world")
	//循环

	file, err := os.Create("1.html")
	req, err := h.Get("http://www.wljx.net/thread-21578-1-1.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer req.Body.Close()
	body, _ := ioutil.ReadAll(req.Body)
	file.Write(body)

}
