/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

var weatherCmd = &cobra.Command{
	Use:   "weather",
	Short: "Gives you the current weather",
	Long:  `Fetches weather data from openweathermap.org and logs it.`,
	Run: func(cmd *cobra.Command, args []string) {
		getWeather()
	},
}

func init() {
	rootCmd.AddCommand(weatherCmd)
}

func getWeather() {
	key := goDotEnvVariable("API_KEY")
	city := goDotEnvVariable("CITY_ID")
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?id=%s&appid=%s", city, key)
	res := WeatherData{}

	reponseBytes := getWeatherData(url)

	if err := json.Unmarshal(reponseBytes, &res); err != nil {
		log.Printf("Error: Could not unmarshal JSON response: %v", err)
	}

	fmt.Printf("%v's Weather\nTemperature: %v\nFeels like: %v\nHumidity: %v\nTimezone: %v\n", res.Name, res.Main.Temp, res.Main.Feels_like, res.Main.Humidity, res.Timezone)
}

func getWeatherData(baseAPI string) []byte {
	request, err := http.NewRequest(
		http.MethodGet,
		baseAPI,
		nil,
	)

	if err != nil {
		log.Printf("Error %v", err)
	}

	request.Header.Add("Accept", "application/json")
	request.Header.Add("User-Agent", "Dadjoke CLI (github.com/PathonScript/CLI)")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Printf("Error %v", err)
	}

	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("Error %v", err)
	}

	return responseBytes
}

func goDotEnvVariable(key string) string {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

type WeatherData struct {
	Main struct {
		Temp       float32 `json:"temp"`
		Feels_like float32 `json:"feels_like"`
		Temp_min   float32 `json:"temp_min"`
		Pressure   float32 `json:"pressure"`
		Humidity   float32 `json:"humidity"`
		Sea_level  float32 `json:"sea_level"`
		Grnd_level float32 `json:"grnd_level"`
	}

	Timezone int32  `json:"timezone"`
	Name     string `json:"name"`
	Cod      int    `json:"cod"`
}
