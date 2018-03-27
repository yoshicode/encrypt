package encrypt

import "testing"

func TestEncStr(t *testing.T) {
	str := "It is apple."

	txt, err := EncStr(str)
	if err != nil {
		t.Errorf("Somethins happen here: %s", err.Error())
	}

	expectStr := "hey !!" + str

	if txt != expectStr {
		t.Error("Return text is not expected.")
	}
}