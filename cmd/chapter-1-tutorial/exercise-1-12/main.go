/*
	Modify the Lissajous server to read parameter values from the URL.
	For example, you might arrange it so that a URL like http://localhost:8000/?cycles=20 sets
	the number of cycles to 20 instead of the default
	5. Use the strconv.Atoi function to convert the string parameter into an integer.
	You can see its documentation with go doc strconv.Atoi.
*/
package main

import (
	"fmt"
	"github.com/AloisCRR/go-exercises/cmd/chapter-1-tutorial/exercise-1-6"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

//!+handler
// handler echoes the HTTP request.
func handler(w http.ResponseWriter, r *http.Request) {

	cyclesQuery := r.URL.Query().Get("cycles")

	cycles, err := strconv.Atoi(cyclesQuery)

	if err != nil {
		fmt.Fprint(w, "Error, query parameter \"cycles\" is required an must be an integer")
		return
	}

	exercise_1_6.Lissajous(w, cycles)
}
