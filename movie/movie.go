package movie

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"os"
	"strings"
	"time"
)

type Movie struct {
	Name  string `json:"name,omitempty"`
	Rate  string `json:"rate,omitempty"`
	Quote string `json:"quote,omitempty"`
	Info  string `json:"info,omitempty"`
	Type  string `json:"type,omitempty"`
}

const (
	Header = `# 豆瓣 TOP Movie 250

> use goquery to achieve this.

🎁 you can also read [Douban-Movie250](https://github.com/Z-Spring/Douban-Movie250) which achieves the same features but native html to parse.


| Title | Rate | Type | Info | Quote |
| ----- | ---- | ---- | ---- | ----- |
`
	Footer = "\n*Last update Time: %v*"
)

func GetMovie(start int) []Movie {
	body := GetMovieBodyFromStart(start)

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(body))
	if err != nil {
		log.Fatal(err)
	}
	var doubanMovie []Movie
	doc.Find("div.info ").Each(func(i int, selection *goquery.Selection) {
		rawTitle := selection.Find("div.hd > a > span.title").Text()
		title := strings.ReplaceAll(rawTitle, " / ", " ")

		var info string
		info = selection.Find(`div.bd > p[class=""]`).Text()
		info = strings.ReplaceAll(info, " ", "	")
		info = strings.ReplaceAll(info, "\n", " ")
		info = strings.TrimSpace(info)
		// or substr can change to "\t\t\t"
		index := strings.Index(info, "                             ")
		info2 := info[:index]

		movieType := info[index:]
		movieType2 := strings.TrimSpace(movieType)

		star := selection.Find("div.bd > div.star > span.rating_num").Text()
		quote := selection.Find("div.bd > p.quote > span").Text()

		movies := Movie{
			Name:  title,
			Info:  info2,
			Rate:  star,
			Quote: quote,
			Type:  movieType2,
		}
		doubanMovie = append(doubanMovie, movies)

	})
	return doubanMovie
}

var ids []int64

func WriteMdToFile(movie []Movie) error {
	// change path here
	file, err := os.OpenFile("README.md", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0666)
	//file, err := os.OpenFile("README.md", os.O_RDWR|os.O_TRUNC, 0666)
	//file, err := os.OpenFile("C:\\Users\\Murphy\\Desktop\\Movie.md", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf(Header))
	if err != nil {
		return err
	}
	for _, movie := range movie {
		_, err := file.WriteString(fmt.Sprintf("| %s | %s | %s | %s | %s |\n", movie.Name, movie.Rate, movie.Type, movie.Info, movie.Quote))
		if err != nil {
			return err
		}
	}
	_, err = file.WriteString(fmt.Sprintf(Footer, time.Now().Format("2006-01-02 15:04:05")))
	if err != nil {
		return err
	}
	return nil
}

func WriteJsonToFile(movie []Movie) {
	bytes, err := json.MarshalIndent(movie, "\t", " ")
	if err != nil {
		log.Fatal(err)
	}
	file, err := os.OpenFile("README.md", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0666)
	if err != nil {
		log.Println(err)
	}

	_, err = file.WriteString(fmt.Sprintf("*Last update Time: %v*\n\n", time.Now().Format("2006-01-02 15:04:05")))
	if err != nil {
		log.Fatal(err)
	}

	_, err = file.WriteString("🎁 you can also read [Douban-Movie250](https://github.com/Z-Spring/Douban-Movie250) which achieves the same features but native html to parse.\n")
	if err != nil {
		log.Fatal(err)
	}

	_, err = file.WriteString(fmt.Sprintf("```json\n%v```", string(bytes)))
	if err != nil {
		log.Fatal(err)
	}

}
