package encrypt

import "testing"

func TestEncStr(t *testing.T) {
	str := "abbc"
	key := 1

	txt, err := EncStr(str, key)
	if err != nil {
		t.Errorf("Somethins happen here: %s", err.Error())
	}

	expectStr := "bcca"

	if txt != expectStr {
		t.Error("Return text is not expected.")
	}
}