package main

import (
	"douban-movie/movie"
	"log"
)

func main() {
	movies := movie.GetMovie(0)
	err := movie.WriteToFile(movies)
	if err != nil {
		log.Fatal(err)
	}
}
