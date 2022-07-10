package main

import (
	"douban-movie/movie"
	"log"
	"sort"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func main() {
	movies := GetAllMovies()
	movie.WriteJsonToFile(movies)
}

func GetSelectedMoviePage(start int) []*movie.Movie {
	return movie.GetMovie(start)
}

func GetAllMovies() []*movie.Movie {
	var DoubanMovie []*movie.Movie
	result := make(chan []*movie.Movie)

	pageStart := []int{0, 25, 50, 75, 100, 125, 150, 175, 200, 225}
	for _, page := range pageStart {
		go func(page int) {
			result <- movie.GetMovie(page)
		}(page)

	}
	for i := 0; i < len(pageStart); i++ {
		select {
		case movie := <-result:
			DoubanMovie = append(DoubanMovie, movie...)
		}

	}
	sort.Sort(movie.ById(DoubanMovie))
	return DoubanMovie
}
