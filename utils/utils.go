package utils

import (
	"fmt"
	"strconv"
	"strings"
)

func StringToMapInt(value string) map[string]int {

	var result = map[string]int{}

	for _, val := range strings.Split(value, ",") {
		key_value := strings.Split(val, ":")
		str_int, err := strconv.Atoi(key_value[1])
		if err != nil {
			fmt.Println("Error occured")
			panic(err)
		} else {
			result[key_value[0]] = str_int
		}
	}
	return result
}

func StringToMapString(value string) map[string]string {

	var result = map[string]string{}

	for _, val := range strings.Split(value, ",") {
		key_value := strings.Split(val, ":")
		result[key_value[0]] = key_value[1]
	}
	return result
}
