package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetPokemonResponse(url string) ([]byte, error) {
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

func ParsePokemonResponse(bytes []byte) (Pokemon, error) {

	var result Pokemon
	err := json.Unmarshal(bytes, &result)
	if err != nil {
		fmt.Println("values: ", string(bytes))
		fmt.Println("Error unmarshalling JSON:", err)
		return Pokemon{}, err
	}

	return result, nil
}

func GetPokemon(url string) (Pokemon, error) {
	resp, err := GetPokemonResponse(url)
	if err != nil {
		return Pokemon{}, err
	}

	results, err := ParsePokemonResponse(resp)
	if err != nil {
		return Pokemon{}, err
	}
	return results, nil
}
