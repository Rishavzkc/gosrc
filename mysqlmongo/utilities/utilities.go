package utilities

import (
	"encoding/json"
	"mysqlmongo/model"
	"os"
)

func GetConfiguration() (model.Configuration, error) {
	config := model.Configuration{}
	file, err := os.Open("./configuration.json")
	if err != nil {
		return config, err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return config, err
	}

	return config, nil
}
