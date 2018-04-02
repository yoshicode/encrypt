package encrypt

import (
	"fmt"
	"strings"
)

func EncStr(str string, key int) (txt string, err error) {

	fmt.Println("start enctyption...")

	list := strings.Split(str, "")

	var a []string

	for _, v := range list {
		position := getKeyByValue(v)
		position += key
		val := getValueByKey(position)
		a = append(a, val)
		txt = strings.Join(a, "")
	}

	fmt.Println("finish enctyption...")

	return txt, err
}

func getKeyByValue (value string) int {

	AlphabetList := map[string]int{"a": 0, "b": 1, "c": 2, "d": 3}

	return AlphabetList[value]
}

func getValueByKey (key int) string {

	InverseAlphabetList := map[int]string{1: "b", 2: "c", 3: "d", 4: "a"}

	return InverseAlphabetList[key]
}

