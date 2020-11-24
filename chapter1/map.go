package main

import "fmt"

func main() {

	countryCapitalMap := make(map[string]string)

	countryCapitalMap [ "Italy" ] = "罗马"
	countryCapitalMap [ "Japan" ] = "东京"
	countryCapitalMap [ "India " ] = "新德里"
	countryCapitalMap [ "France" ] = "巴黎"

	//删除元素map中元素
	delete(countryCapitalMap, "France")

	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap [country])
	}
}