package main

import (
	"Weather-API-minipetproject/my-city"
	"Weather-API-minipetproject/weather"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func getWeather(w http.ResponseWriter, r *http.Request) {
	myCity := r.URL.Query().Get("my_city")
	inCity := r.URL.Query().Get("city")

	var weatherNow *weather.Weather

	switch {
	case myCity == "1":
		city, err := my_city.Definition()
		if err != nil {
			return
		}
		weatherNow, err = weather.Get(city)
		if err != nil {
			return
		}
	case inCity != "":
		var err error
		weatherNow, err = weather.Get(inCity)
		if err != nil {
			return
		}
	default:
		answer, err := json.Marshal("Установите my_city=1 или введите city")
		if err != nil {
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write(answer)
		return
	}

	answer, err := json.Marshal(weatherNow)
	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(answer)

}

func main() {
	fmt.Println("Введите токен от api.openweathermap.org")
	fmt.Scanln(&weather.Token)

	r := chi.NewRouter()

	r.Get("/weather", getWeather)

	fmt.Println("Сервер запущен")
	if err := http.ListenAndServe(":8080", r); err != nil {
		fmt.Printf("Ошибка при запуске сервера: %s", err.Error())
		return
	}
}
