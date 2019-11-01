package learning_go

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://ms.jr.jd.com/gw/generic/hj/h5/m/latestPrice?reqData=%7B%7D", nil)
	if err != nil {
		log.Fatal(err)
	}
	/*
		req.Header.Set("Connection", "keep-alive")
		req.Header.Set("Cache-Control", "max-age=0")
		req.Header.Set("Upgrade-Insecure-Requests", "1")
		req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/77.0.3865.120 Safari/537.36")
		req.Header.Set("Sec-Fetch-Mode", "navigate")
		req.Header.Set("Sec-Fetch-User", "?1")
		req.Header.Set("Sec-Fetch-Site", "none")
		req.Header.Set("Accept-Encoding", "gzip, deflate, br")
		req.Header.Set("Accept-Language", "en-US,en;q=0.9,zh-CN;q=0.8,zh;q=0.7,zh-TW;q=0.6")
		req.Header.Set("Cookie", "shshshfpa=8f3e1180-6fcb-a30e-b017-84ea3a8f38b8-1563337386; shshshfpb=ckJSJwUXLuwDHKxBM6bZ6fw%3D%3D; pinId=zSd16ScyHeIiENP0aEOy5LV9-x-f3wj7; ceshi3.com=201; TrackID=1r2vkLUbWmu505zMBkIdobjyjxWFyrRKIPoR8vcrMvCPvLuQef1emIe-b4wgb9b5DY5HQAo884t6wGaRDfY6rAfWJy_xhIuxr3p9a9G2Aa8Hz0mpMtG0HUdwdkbyYs1kP; mba_muid=1563337362468466449662; shshshfp=7aedc1ccd16cc4ff907b6fcafc88a3da; 3AB9D23F7A4B3C9B=SXFGZJBS2KVOUNR5IPV4R2VSQAUXL7UVV6NAFSILZ6R7ATH3BS6DZWF4YIZI2EULYLCQWWHQYO4FDCTDY7QJUJE6LA; __jda=111857887.1563337362468466449662.1563337362.1570589272.1571811945.10; __jdc=111857887")
	*/
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", bodyText)
}
