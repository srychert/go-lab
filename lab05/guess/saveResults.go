package guess

import (
	"encoding/json"
	"io/ioutil"
)

func saveResults(results map[Date][]Result) {
	r_json, _ := json.MarshalIndent(results, "", "  ")
	ioutil.WriteFile("hallOfFame.json", r_json, 0644)
}
