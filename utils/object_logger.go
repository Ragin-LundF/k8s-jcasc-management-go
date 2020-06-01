package utils

import (
	"encoding/json"
	"log"
)

// pretty output of a struct
func LogStruct(objectName string, class interface{}) {
	log.Printf("Content of %v:\n", objectName)
	log.Println("==================================")
	configValuesJson, _ := json.MarshalIndent(class, "", "\t")
	log.Printf("\n%v", string(configValuesJson))
}
