package names

import (
	"encoding/json"
	"fmt"
)

type AgeData struct {
	Name  string `json:"name"`
	Age   int `json:"age"`
	Count int `json:"count"`
}

/*
Sample
{
    "name": "pepe",
    "age": 51,
    "count": 48187
}
*/

func GetAgeData(body []byte) AgeData {
	AgeData := AgeData{}
	err := json.Unmarshal(body, &AgeData)
	if err != nil {
		fmt.Println("Unable to convert Json")
		panic(err)
	}
	return AgeData
}

