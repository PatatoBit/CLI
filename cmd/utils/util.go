package util

import (
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func FetchJSONData(baseAPI string) []byte {
	request, err := http.NewRequest(
		http.MethodGet,
		baseAPI,
		nil,
	)

	if err != nil {
		fmt.Printf("Error %v", err)
	}

	request.Header.Add("Accept", "application/json")
	request.Header.Add("User-Agent", "Dadjoke CLI (github.com/PathonScript/CLI)")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Printf("Error %v", err)
	}

	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Error %v", err)
	}

	return responseBytes
}

func GoDotEnvVariable(key string) string {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	return os.Getenv(key)
}

func ParseTemp(temp float64) float64 {
	ans := math.Round(((temp - 273.15) * 100) / 100)
	return ans
}
