package bearToSN

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func ReadStandardNotesFile() {
	raw, err := ioutil.ReadFile("./sample_notes.json")

	var result map[string][]interface{}

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	json.Unmarshal(raw, &result)

	items := result["items"]

	// fmt.Println(items)

	for i := 0; i < len(items)-1; i++ {
		fmt.Println(items[i])
	}
}
