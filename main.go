package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	res, err := http.Get("https://www.msn.com/zh-tw/news/living/%E5%BF%AB%E8%A8%8A-%E6%9C%AC%E5%9C%9F-56339-%E6%AD%BB%E4%BA%A1%E5%A2%9E115%E4%BA%BA/ar-AAYGIJP?ocid=msedgntp&cvid=1c714afa3d384cc58698d1ffc8eaf8a6")

	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	sitemap, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(sitemap)
	fmt.Println(string(sitemap))
	fmt.Println(strings.Contains(string(sitemap), "新冠肺炎"))
}
