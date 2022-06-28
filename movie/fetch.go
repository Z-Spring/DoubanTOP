package movie

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func GetMovieBodyFromStart(start int) string {
	//url := "https://movie.douban.com/top250"
	url := ChooseStartId(start)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		log.Println(err)

	}
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.77 Safari/537.36")
	// 此cookie是登录豆瓣后从浏览器获得，因为有时间限制，需要勤换
	req.Header.Add("cookie", "ll=\"108288\"; bid=tfmu0VNtswM; douban-fav-remind=1; ct=y; viewed=\"26979890_30357170_30397714_27663285_30389976_25782902_19952400_35410754_35668581_30395692\"; push_noty_num=0; push_doumail_num=0; ap_v=0,6.0; dbcl2=\"156810236:67LMVXkh0Sw\"; ck=1CSY")

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)

	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(err)
		}
	}(resp.Body)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)

	}
	return string(body)
}

func ChooseStartId(start int) string {
	if start >= 0 && start <= 249 {
		url := "https://movie.douban.com/top250?start=%d&filter="
		url2 := fmt.Sprintf(url, start)
		return url2
	}
	return "The start is from 0~249"
}
