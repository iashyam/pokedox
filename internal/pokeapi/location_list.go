package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)



func GetLocationResponse(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return []byte{}, err
	}
	return bytes, nil
}

func ParseLocationResponse(bytes []byte) (locationAreaObj, error) {

	var result locationAreaObj
	err := json.Unmarshal(bytes, &result)
	if err != nil {
		fmt.Println("json data", string(bytes))
		fmt.Println("Error unmarshalling JSON:", err)
		return locationAreaObj{}, err
	}

	return result, nil
}

func GetLocationList(url string) (locationAreaObj, error) {
	resp, err := GetLocationResponse(url)
	if err != nil {
		return locationAreaObj{}, err
	}

	results, err := ParseLocationResponse(resp)
	if err != nil {
		return locationAreaObj{}, err
	}
	return results, nil
}

// func main() {
// 	url := "https://pokeapi.co/api/v2/location-area"

// 	results, err := GetLocationList(url)
// 	if err != nil {
// 		fmt.Println("Error parsing location JSON:", err)
// 		return
// 	}
// 	for _, location := range results.Results {
// 		fmt.Printf("- %x", location["name"])
// 	}
// }
