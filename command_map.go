package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Location struct {
	Count    int       `json:"count"`
	Next     string    `json:"next"`
	Previous string    `json:"previous"`
	Results  []Results `json:"results"`
}

type Results struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Config struct {
	Next     string `json:"next"`
	Previous string `json:"previous"`
}

func commandMap(cfg *Config) error {
	resp, err := http.Get(cfg.Next)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	locationData, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	location := Location{}
	err = json.Unmarshal(locationData, &location)
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(locationData, &cfg)
	if err != nil {
		fmt.Println(err)
	}

	for _, r := range location.Results {
		fmt.Println(r.Name)
	}

	return nil
}
