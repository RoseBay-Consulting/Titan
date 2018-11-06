package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("Starting the application...")

	jsonData := map[string]string{"test": "I am Tilak"}
	jsonValue, _ := json.Marshal(jsonData)
	response, err := http.Post("https://dm0fpx25wg.execute-api.ap-southeast-1.amazonaws.com/test/check", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
		var responseObject Response
		json.Unmarshal(responseData, &responseObject)

	}

	fmt.Println("Terminating the application...")
}


/*package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		// Note if you set a break point on the line above you can't reach it using the debugger
		// but if you just run it, then the code is executed.
		fmt.Println("we got a line")
		counts[line]++
		if len(line) == 0 {
			break
		}
	}

	for line, n := range counts {
		fmt.Printf("%d\t%s\n", n, line)
	}
}*/