package main

import "fmt"

type Movie struct {
	Actors		[]string
	Rating		float32							// Exported field
	releaseYear	int 							// Unexported field
	Title		string
}
// A Movie type method
func (movie Movie) DisplayTitle() string {		// Declare a reciever type of (movie Movie)
	return fmt.Sprintf("%s (%d)", movie.Title, movie.releaseYear)
}

type Counter struct {							// Having first letter Cap make the type/constant/field/function exported
												// i.e. available outside the package
	Count int
}
func (c Counter) increment() {					// Unexported method
	c.Count++
}
func (c *Counter) IncrementWithPointer() {		// Exported method
	c.Count++
}

func main() {
	episodeIV := Movie{
		Title:			"Star Wars: A New Hope",
		Rating:			5.0,
		releaseYear:	1977,
	}

	episodeV := Movie{
		Title:			"Star Wars: The Empire Strikes Back",
		releaseYear:	1980,
	}

	episodeIV.Actors = []string{
		"Mark Hamill",
		"Harrison Ford",
		"Carrie Fisher",
	}
	fmt.Println(episodeIV.Title, "has rating", episodeIV.Rating)
	fmt.Println(episodeV.DisplayTitle())

	counter := &Counter{}
	fmt.Println(counter.Count)		// 0

	counter.increment() 			// We can also call the normal Increment with the address of Counter
	fmt.Println(counter.Count)		// 0

	counter.IncrementWithPointer()
	fmt.Println(counter.Count)		// 1
}