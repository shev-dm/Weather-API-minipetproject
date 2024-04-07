package weather

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const baseUrl = "https://api.openweathermap.org/data/2.5/weather?"

var Token string

type Condition struct {
	Main        string `json:"main"`
	Description string `json:"description"`
}

type Main struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	Pressure  int     `json:"pressure"`
	Humidity  int     `json:"humidity"`
}

type Wind struct {
	Speed float64 `json:"speed"`
}

type Weather struct {
	Name      string      `json:"name"`
	Condition []Condition `json:"weather"`
	Main      Main        `json:"main"`
	Wind      Wind        `json:"wind"`
}

func Get(a string) (*Weather, error) {
	url := fmt.Sprintf("%sq=%s&appid=%s&units=metric&lang=ru", baseUrl, a, Token)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("ошибка при выполнении запроса %v", err)
	}
	defer resp.Body.Close()

	var weather Weather

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("ошибка API: статус %v", resp.Status)
	}

	err = json.NewDecoder(resp.Body).Decode(&weather)
	if err != nil {
		return nil, fmt.Errorf("ошибка при декодировании JSON %v", err)
	}

	return &weather, nil
}
