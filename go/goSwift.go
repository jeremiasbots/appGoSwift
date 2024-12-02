package goSwift

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
)

//export Add
func Add(number1 int, number2 int) int {
	return number1 + number2
}

//export Multiply
func Multiply(number1 int, number2 int) int {
	return number1 * number2
}

//export TheNewMovie
func TheNewMovie() string {
	url := "https://whenisthenextmcufilm.com/api"
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		return "Something went wrong"
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
		return "Something went wrong"
	}
	var theNewMovieStruct map[string]interface{}
	err = json.Unmarshal(body, &theNewMovieStruct)
	if err != nil {
		log.Fatal(err)
		return "Something went wrong"
	}
	return theNewMovieStruct["title"].(string)
}

//export TheNewMovieImage
func TheNewMovieImage() string {
	url := "https://whenisthenextmcufilm.com/api"
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		return "Something went wrong"
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
		return "Something went wrong"
	}
	var theNewMovieStruct map[string]interface{}
	err = json.Unmarshal(body, &theNewMovieStruct)
	if err != nil {
		log.Fatal(err)
		return "Something went wrong"
	}
	return theNewMovieStruct["poster_url"].(string)
}

//export TheNewMovieDays
func TheNewMovieDays() string {
	url := "https://whenisthenextmcufilm.com/api"
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		return "Something went wrong"
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
		return "Something went wrong"
	}
	var theNewMovieStruct map[string]interface{}
	err = json.Unmarshal(body, &theNewMovieStruct)
	if err != nil {
		log.Fatal(err)
		return "Something went wrong"
	}

	return strconv.Itoa(int(theNewMovieStruct["days_until"].(float64)))
}
