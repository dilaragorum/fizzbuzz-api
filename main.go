package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)



func FizzBuzz(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Get metodu dışında başka metoda izin vermiyoruz
	var fizzBuzzArray []string
	count, err := strconv.Atoi(ps.ByName("count"))
	// count değeri 0'dan küçük veya int'e çevrilemez ise  
	if err != nil || count <= 0 {
		http.NotFound(w, r)
		return
	}
	// fizzBuzz Algorithm
	for i := 1; i <= count; i++ {
		if i%3 == 0 && i%5 == 0 {
			fizzBuzzArray = append(fizzBuzzArray, "fizzBuzz")
		} else if i%5 == 0 {
			fizzBuzzArray = append(fizzBuzzArray, "buzz")
		} else if i%3 == 0 {
			fizzBuzzArray = append(fizzBuzzArray, "fizz")
		} else {
			fizzBuzzArray = append(fizzBuzzArray, strconv.Itoa(i))
		}
	}
	// Array'i json'a çevirip header ile birlikte yazdırmak.
	jsonStr, _ := json.Marshal(fizzBuzzArray)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonStr)
}

func main() {
	router := httprouter.New()
	router.GET("/fizzbuzz/:count", FizzBuzz)

	err := http.ListenAndServe(":8080", router)
	log.Fatal(err)
}
