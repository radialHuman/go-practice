package main

import (
	"encoding/json"
	"fmt"
)

// need to create a struct that can fit the json elements
type kvPairs struct { // this is a new thing for json to match key words
	One   string `json:"k1"`
	Two   string `json:"k2"`
	Three int    `json:"k3"`
}

func main() {
	// creatuing a json or a dict
	// #doubt is ` for multiline string?
	newJson := `[ 
	{
		"k1" : "v1",
		"k2" : "v1",
		"k3" : 1
	},
	{
		"k1" : "v2",
		"k2" : "v2",
		"k3" : 2
	}
]`
	fmt.Printf("This is the json \n%s\n", newJson)
	// ASSUMING we get this json from external api and we know the structur eof the kvs we can unmarshall the json using struct defined above
	var kvs []kvPairs
	err := json.Unmarshal([]byte(newJson), &kvs) // it doesn like string, we need ti in bytes. also we need to pass the reference of variable
	if err != nil {
		fmt.Printf("The result of unmarshlling is %s", err)
	}
	fmt.Println(kvs) // this gives all the values of all the elemetns in json slice

	// writting from struct to json
	// once a new struct is created with its slcie type []Dict, then
	// newJson, er := json.MarshalIndent(dcitName, "", "   ") // "" is the prefix, "   " is the indentation size
	// print string(newJson)

}
