package my_city

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Location struct {
	Country  string  `json:"country"`
	City     string  `json:"city"`
	Lat      float64 `json:"lat"`
	Lon      float64 `json:"lon"`
	Timezone string  `json:"timezone"`
	Ip       string  `json:"query"`
}

func Definition() (string, error) {
	url := "http://ip-api.com/json/"

	response, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("ошибка при выполнении запроса %v", err)
	}
	defer response.Body.Close()

	var location Location
	if err := json.NewDecoder(response.Body).Decode(&location); err != nil {
		return "", fmt.Errorf("ошибка при декодировании JSON %v", err)
	}

	return location.City, nil
}
