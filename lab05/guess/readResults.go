package guess

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		fmt.Println(e)
		panic(e)
	}
}

func readResults() map[Date][]Result {
	m := make(map[Date][]Result)
	// check if file exists
	if _, err := os.Stat("hallOfFame.json"); err == nil {
		data, err := ioutil.ReadFile("hallOfFame.json")
		check(err)

		errj := json.Unmarshal(data, &m)
		check(errj)
	}
	return m
}
