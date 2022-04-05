package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

const (
	fizzBuzz = "fizzBuzz"
	fizz     = "fizz"
	buzz     = "buzz"
)

func FizzBuzz(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Get metodu dışında başka metoda izin vermiyoruz
	var fizzBuzzArray []string
	count, err := strconv.Atoi(ps.ByName("count"))
	// count değeri 0'dan küçük veya int'e çevrilemez ise
	if err != nil {
		http.NotFound(w, r)
		return
	}
	if count <= 0 {
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	// fizzBuzz Algorithm
	for i := 1; i <= count; i++ {
		if i%3 == 0 && i%5 == 0 {
			fizzBuzzArray = append(fizzBuzzArray, fizzBuzz)
		} else if i%5 == 0 {
			fizzBuzzArray = append(fizzBuzzArray, buzz)
		} else if i%3 == 0 {
			fizzBuzzArray = append(fizzBuzzArray, fizz)
		} else {
			fizzBuzzArray = append(fizzBuzzArray, strconv.Itoa(i))
		}
	}
	// Array'i json'a çevirip header ile birlikte yazdırmak.
	jsonStr, err := json.Marshal(fizzBuzzArray)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonStr)
}

func main() {
	router := httprouter.New()
	router.GET("/fizzBuzz/:count", FizzBuzz)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}
