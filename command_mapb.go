package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func commandMapb(cfg *config) error {
	resp, err := http.Get(cfg.Previous)
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
