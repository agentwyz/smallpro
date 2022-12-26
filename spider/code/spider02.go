package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "https://www.thepaper.cn/"
	resp, err := http.Get(url)

	if err != nil {
		fmt.Printf("featch url error:%v", err)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error status code: %v", resp.StatusCode)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Printf("read content failed: %v", err)
		return
	}

	fmt.Println("body:", string(body))
}