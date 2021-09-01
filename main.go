package main

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
)

func main() {
	phones := []string{
		"Cesar 1234654",
		"Robson 88889997",
		"Maria 22223344",
		"Cesar 56676654",
		"Cesar 98776553",
		"Maria 32324444",
		"Cesar 76654444",
		"Luana 89765433",
		"Robson 22334432",
	}
	sort.Strings(phones)

	m := phoneNumbersMap(phones)
	j, _ := json.Marshal(&m)
	fmt.Println(string(j))
}

func phoneNumbersMap(phones []string) map[string][]string {
	m := make(map[string][]string)
	for _, phone := range phones {
		splitStrings := strings.Split(phone, " ")
		name := splitStrings[0]
		var numbersList []string

		for _, item := range phones {
			if strings.Contains(item, name) {
				numbersList = append(numbersList, strings.Split(item, " ")[1])
			}
		}

		m[name] = numbersList
	}
	return m
}
