package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetLocationAreaEncounterResponse(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}
	return bytes, nil
}

func ParseLocationAreaEncounterResponse(bytes []byte) (LocationAreaEncounter, error) {

	var result LocationAreaEncounter
	err := json.Unmarshal(bytes, &result)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return LocationAreaEncounter{}, err
	}

	return result, nil
}

func GetLocationAreaEncounter(url string) (LocationAreaEncounter, error) {
	resp, err := GetLocationAreaEncounterResponse(url)
	if err != nil {
		return LocationAreaEncounter{}, err
	}

	results, err := ParseLocationAreaEncounterResponse(resp)
	if err != nil {
		return LocationAreaEncounter{}, err
	}
	return results, nil
}
