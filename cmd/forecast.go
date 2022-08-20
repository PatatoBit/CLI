/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	util "Patato/pcli/cmd/utils"

	// env "Patato/pcli/cmd/funcs/env"

	// fetch "Patato/pcli/cmd/funcs/fetch"
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

// forecastCmd represents the forecast command
var forecastCmd = &cobra.Command{
	Use:   "forecast",
	Short: "Tells you the forcast of a city.",
	Long:  `Fetches the forecast from openweathermap.org api and tells you the forcast.`,
	Run: func(cmd *cobra.Command, args []string) {
		getForecast()
	},
}

func init() {
	rootCmd.AddCommand(forecastCmd)
}

func getForecast() {
	// Reads env file
	key := util.GoDotEnvVariable("API_KEY")
	lat := util.GoDotEnvVariable("LAT")
	lon := util.GoDotEnvVariable("LON")
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/forecast?lat=%s&lon=%s&appid=%s", lat, lon, key)
	res := ForcastData{}

	responseBytes := util.FetchJSONData(url)
	if err := json.Unmarshal(responseBytes, &res); err != nil {
		fmt.Printf("Error: Could not unmarshal JSON response: %v", err)
	}

	fmt.Println(res.List[0])
}

type ForcastData struct {
	Cod  string
	List []struct {
		Main struct {
			Temp       float64 `json:"temp"`
			Feels_like float64 `json:"feels_like"`
			Humidity   float64 `json:"humidity"`
		}
		Weather []WeatherForcasts `json:"weather"`
		Clouds  struct {
			All int `json:"all"`
		}
		Wind struct {
			Speed float32 `json:"speed"`
			Deg   int     `json:"deg"`
			Gust  float32 `json:"gust"`
		}
	}
}

type WeatherForcasts struct {
	Main        string `json:"main"`
	Description string `json:"description"`
}
