package data

import (
	"encoding/json"
	"fmt"
	"greenlight.alexedwards.net/internal/validator"
	"time"
)

type Movie struct {
	ID        int64     `json:"id"`               // Unique integer ID for the movie
	CreatedAt time.Time `json:"-"`                // Timestamp for when the movie is added to our database
	Title     string    `json:"title"`            // Movie title
	Year      int32     `json:"year,omitempty"`   // Movie release year
	Runtime   Runtime   `json:"-"`                // Movie runtime (in minutes)
	Genres    []string  `json:"genres,omitempty"` // Slice of genres for the movie (romance, comedy, etc.)
	Version   int32     `json:"version"`          // The version number starts at 1 and will be incremented each time the movie information is updated
}

func ValidateMovie(v *validator.Validator, movie *Movie) {
	v.Check(movie.Title != "", "title", "must be provided")
	v.Check(len(movie.Title) <= 500, "title", "must not be more than 500 bytes long")

	v.Check(movie.Year != 0, "year", "must be provided")
	v.Check(movie.Year >= 1888, "year", "must be greater than 1888")
	v.Check(movie.Year <= int32(time.Now().Year()), "year", "must not be in the future")

	v.Check(movie.Runtime != 0, "runtime", "must be provided")
	v.Check(movie.Runtime > 0, "runtime", "must be positive integer")

	v.Check(movie.Genres != nil, "genres", "must be provided")
	v.Check(len(movie.Genres) >= 1, "genres", "must contain at least 1 genre")
	v.Check(len(movie.Genres) <= 5, "genres", "must not contain more than 5 genres")
	v.Check(validator.Unique(movie.Genres), "genres", "must not contain duplicate values")
}

func (m Movie) MarshalJSON() ([]byte, error) {
	// Create a variable holding the custom runtime string, just like before.
	var runtime string

	if m.Runtime != 0 {
		runtime = fmt.Sprintf("%d mins", m.Runtime)
	}

	// Define a MovieAlias type which has the underlying type Movie. Due to the way that
	// Go handles type definitions (https://golang.org/ref/spec#Type_definitions) the
	// MovieAlias type will contain all the fields that our Movie struct has but,
	// importantly, none of the methods.
	type MovieAlias Movie

	// Embed the MovieAlias type inside the anonymous struct, along with a Runtime field
	// that has the type string and the necessary struct tags. It's important that we
	// embed the MovieAlias type here, rather than the Movie type directly, to avoid
	// inheriting the MarshalJSON() method of the Movie type (which would result in an
	// infinite loop during encoding).
	aux := struct {
		MovieAlias
		Runtime string `json:"runtime,omitempty"`
	}{
		// Set the values for the anonymous struct.
		MovieAlias: MovieAlias(m),
		Runtime:    runtime,
	}

	// Encode the anonymous struct to JSON, and return it.
	return json.Marshal(aux)
}
