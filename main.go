package main

import (
	"douban-movie/movie"
	"fmt"
	"log"
	"sort"
	"time"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func main() {
	start := time.Now()
	movies := GetAllMovies()
	movie.WriteJsonToFile(movies)
	end := time.Since(start)
	fmt.Println(end)
}

func GetSelectedMoviePage(start int) []*movie.Movie {
	return movie.GetMovie(start)
}

func GetAllMovies() []*movie.Movie {
	var DoubanMovie []*movie.Movie

	pageStart := []int{0, 25, 50, 75, 100, 125, 150, 175, 200, 225}
	result := make(chan []*movie.Movie)

	for _, page := range pageStart {
		go func(page int) {
			result <- movie.GetMovie(page)
		}(page)
	}
	defer close(result)
	for i := 0; i < len(pageStart); i++ {
		select {
		case v := <-result:
			DoubanMovie = append(DoubanMovie, v...)
		}

	}
	sort.Sort(movie.ById(DoubanMovie))
	return DoubanMovie
}
