package main

import (
	"fmt"
	"log"
	"sort"
	"time"

	"douban-movie/movie"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func main() {
	start := time.Now()
	movies := GetAllMovies()
	movie.WriteJsonToFile(movies)
	// movie.WriteMdToFile(movies)
	end := time.Since(start)
	fmt.Printf("movies have added to your file! run time: %v", end)
}

// GetSelectedMoviePage you can get specify page info with this func
func GetSelectedMoviePage(start int) []*movie.Movie {
	return movie.GetMovie(start)
}

func GetAllMovies() []*movie.Movie {
	DoubanMovie := make([]*movie.Movie, 0, 250)

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
