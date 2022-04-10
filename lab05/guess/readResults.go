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
	if _, err := os.Stat(filePath); err == nil {
		data, err := ioutil.ReadFile(filePath)
		check(err)

		errj := json.Unmarshal(data, &m)
		check(errj)
	} else {
		filePath = "./hallOfFame.json"
	}
	return m
}
