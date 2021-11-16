package pkg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func ReadAndSaveJSON(filename string, structure interface{}) {
	jsonFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	err = jsonFile.Close()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = json.Unmarshal(byteValue, &structure)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}