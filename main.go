package main

import (
	"douban-movie/movie"
	"log"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func main() {
	movies := GetAllMovies()
	movie.WriteJsonToFile(movies)
}

func GetSelectedMoviePage(start int) []movie.Movie {
	return movie.GetMovie(start)
}

func GetAllMovies() []movie.Movie {
	var DoubanMovie []movie.Movie

	pageStart := []int{0, 25, 50, 75, 100, 125, 150, 175, 200, 225}
	for _, page := range pageStart {
		movies := movie.GetMovie(page)
		DoubanMovie = append(DoubanMovie, movies...)
	}
	return DoubanMovie
}
