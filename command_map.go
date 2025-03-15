package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"errors"
)



type LocationAreas struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}


func commandMap(config *Config) error {
	var locations LocationAreas

	if data, ok := config.Cache.Get(config.Next); ok{
		fmt.Println("Used a cache to generate responses")
		err := json.Unmarshal(data, &locations)
		if err != nil {
			return err
		}
		
	} else {

		resp, err := http.Get(config.Next)
		if err != nil {
			return err
		}

		defer resp.Body.Close()
		data, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		err = json.Unmarshal([]byte(data), &locations)
		if err != nil {
			return err
		}
		config.Cache.Add(config.Next, []byte(data))
	}

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	config.Previous = locations.Previous
	config.Next = locations.Next
	return nil
}

func commandMapBack(config *Config) error {
	if config.Previous == ""{
		return errors.New("Map Back command must have something to map back to")
	}

	var locations LocationAreas

	if data, ok := config.Cache.Get(config.Previous); ok {
		fmt.Println("Used a cache to generate responses")
		err := json.Unmarshal(data, &locations)
		if err != nil {
			return err
		}

	} else {

		resp, err := http.Get(config.Previous)
		if err != nil {
			return err
		}
		
		defer resp.Body.Close()
		data, err := io.ReadAll(resp.Body)

		err = json.Unmarshal([]byte(data), &locations)
		if err != nil {
			return err
		}
		config.Cache.Add(config.Previous, []byte(data))
	}

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	config.Next = locations.Next
	config.Previous = locations.Previous
	return nil
}
