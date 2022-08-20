/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	util "Patato/pcli/cmd/utils"
	"encoding/json"
	"fmt"
	"math"
	"strconv"
	"time"

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
	city := util.GoDotEnvVariable("CITY_NAME")
	key := util.GoDotEnvVariable("API_KEY")
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city, key)
	res := WeatherData{}

	responseBytes := util.FetchJSONData(url)

	if err := json.Unmarshal(responseBytes, &res); err != nil {
		fmt.Printf("Error: Could not unmarshal JSON response: %v", err)
	}

	fmt.Printf("%v's Weather [%v]\nTemperature: %v°C\nFeels like: %v°C\nHumidity: %v%%\n", res.Name, res.Cod, math.Round(((res.Main.Temp-273.15)*100)/100), math.Round(((res.Main.Feels_like-273.15)*100)/100), res.Main.Humidity)
	// Fix this to date time later
	fmt.Printf("\nSunrise/Sunset: %v/%v", ParseTime(res.Sys.Sunrise), ParseTime(res.Sys.Sunset))
}

func ParseTime(timestamp int) int {

	stringStamp := strconv.Itoa(timestamp)

	i, err := strconv.ParseInt(stringStamp, 10, 64)
	if err != nil {
		panic(err)
	}
	tm := time.Unix(i, 0)
	return tm.Hour()
}

type WeatherData struct {
	Main struct {
		Temp       float64 `json:"temp"`
		Feels_like float64 `json:"feels_like"`
		Temp_min   float32 `json:"temp_min"`
		Pressure   float32 `json:"pressure"`
		Humidity   float32 `json:"humidity"`
		Sea_level  float32 `json:"sea_level"`
		Grnd_level float32 `json:"grnd_level"`
	}

	Sys struct {
		Sunrise int `json:"sunrise"`
		Sunset  int `json:"sunset"`
	}

	Timezone int32  `json:"timezone"`
	Name     string `json:"name"`
	Cod      int32  `json:"cod"`
}
