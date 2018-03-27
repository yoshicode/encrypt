package encrypt

import "fmt"

func EncStr(str string) (txt string, err error) {

	fmt.Println("start enctyption...")

	txt = "hey !!" + str

	return txt, err
}