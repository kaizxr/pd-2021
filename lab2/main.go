package main

import (
	"net/http"
	"strconv"
)

var m = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func romanToInt(s string) int {
	var result int = 0
	for i := 0; i < len(s)-1; i++ {
		if m[string(s[i])] >= m[string(s[i+1])] {
			result = result + m[string(s[i])]
		} else {
			result = result - m[string(s[i])]
		}
	}
	result = result + m[string(s[len(s)-1])]
	return result
}

func main() {
	romanHandler := http.HandlerFunc(romanHandler)
	http.Handle("/task2", romanHandler)
	http.ListenAndServe(":8080", nil)
}

func romanHandler(w http.ResponseWriter, r *http.Request) {
	headerContentTtype := r.Header.Get("Content-Type")
	if headerContentTtype != "application/x-www-form-urlencoded" {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		return
	}
	r.ParseForm()
	roman := r.Form["roman"]
	var res int = romanToInt(roman[0])
	w.WriteHeader(200)
	w.Write([]byte(strconv.Itoa(res)))

}
