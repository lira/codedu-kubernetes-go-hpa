package main

import (
	"fmt"
	"math"
	"net/http"
)

func sqrt(x float64) float64 {
	return math.Sqrt(x)
}

func handler(w http.ResponseWriter, r *http.Request) {
	var x float64 = 0.0001

	for i := 0; i < 1000000; i++ {
		x += sqrt(x)
	}

	fmt.Fprintf(w, "Code.education Rocks!")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8000", nil)
}
