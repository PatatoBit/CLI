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

	"github.com/spf13/cobra"
)

// randomCmd represents the random command
var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "Gives you a random dad joke",
	Long:  `This will fetch a random dad joke from the icanhazdad dad joke api`,
	Run: func(cmd *cobra.Command, args []string) {
		getJoke()
	},
}

func init() {
	rootCmd.AddCommand(randomCmd)
}

type Joke struct {
	ID     string `json:"id"`
	Joke   string `json:"joke"`
	Status int    `json:"status"`
}

func getJoke() {
	url := "https://icanhazdadjoke.com"
	reponseBytes := getJokeData(url)
	joke := Joke{}

	if err := json.Unmarshal(reponseBytes, &joke); err != nil {
		log.Printf("Error: Could not unmarshal JSON response: %v", err)
	}

	fmt.Println(string(joke.Joke))
}

func getJokeData(baseAPI string) []byte {
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
