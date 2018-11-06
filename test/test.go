package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"bytes"
)

// A Response struct to map the Entire Response
type Response struct {
	Status    string    `json:"status"`
	Test    string    `json:"test"`
	Message    string    `json:"message"`
//	Pokemon []Pokemon `json:"pokemon_entries"`
}



// A Pokemon Struct to map every pokemon to.
/*type Pokemon struct {
	EntryNo int            `json:"entry_number"`
	Species PokemonSpecies `json:"pokemon_species"`
}
*/
// A struct to map our Pokemon's Species which includes it's name
/*type PokemonSpecies struct {
	Name string `json:"name"`
}*/

func main() {
	jsonData := map[string]string{"test": "I am Tilak"}
	jsonValue, _ := json.Marshal(jsonData)
	response, err := http.Post("https://dm0fpx25wg.execute-api.ap-southeast-1.amazonaws.com/test/check", "application/json", bytes.NewBuffer(jsonValue))

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject Response
	json.Unmarshal(responseData, &responseObject)

	fmt.Println(responseObject.Test)
	fmt.Println(len(responseObject.Status))
//	fmt.Println(len(responseObject.Message))

	/*for i := 0; i < len(responseObject.Pokemon); i++ {
		fmt.Println(responseObject.Pokemon[i].Species.Name)
	}
*/
}