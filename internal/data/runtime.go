package data

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// ErrInvalidRuntimeFormat Define an error that our UnmarshalJSON() method can return if we're unable to parse
// or convert the JSON string successfully.
var ErrInvalidRuntimeFormat = errors.New("invalud runtime format")

// Runtime
// Declare a custom Runtime type, which has the underlying type int32 (the same as out
// Movie struct field).
type Runtime int32

// MarshalJSON
// Implement a MarshalJSON() method on the Runtime type so that it satisfies the
// json.Marshaller interface. This should return the JSON-encoded value for the movie
// runtime (in our case, it will return a string in the format "<runtime> mins").
func (r *Runtime) MarshalJSON() ([]byte, error) {
	// Generate a string containing the movie runtime in the required format.
	jsonValue := fmt.Sprintf("%d mins", r)

	// Use the strconv.Quote() function on the string to wrap it in double quotes. It
	// needs to be surrounded by double quotes in order to be a valid *JSON string*.
	quoteJSONValue := strconv.Quote(jsonValue)

	// Convert the quoted string value to a byte slice and return it.
	return []byte(quoteJSONValue), nil
}

func (r *Runtime) UnmarshalJSON(jsonValue []byte) error {
	unquoteJSONValue, err := strconv.Unquote(string(jsonValue))
	if err != nil {
		return ErrInvalidRuntimeFormat
	}

	parts := strings.Split(unquoteJSONValue, " ")

	if len(parts) != 2 || parts[1] != "mins" {
		return ErrInvalidRuntimeFormat
	}

	// Otherwise, parse the string containing the number into an int32. Again, if this
	// fails return the ErrInvalidRuntimeFormat error.
	i, err := strconv.ParseInt(parts[0], 10, 32)
	if err != nil {
		return ErrInvalidRuntimeFormat
	}

	*r = Runtime(i)

	return nil
}
