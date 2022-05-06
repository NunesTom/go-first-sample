package main

import (
	"fmt"
)

func main() {
	var queryString string = "?valor1=key1&valor2=key2&order=giropops"
	var result = AnaliserQuery(queryString)
	fmt.Println(result)
	fmt.Println(result["valor2"])
}
