package auth

import "testing"

func Test_hash(t *testing.T) {
	password := "my_pass"
	hash, err := HashPassword(password)
	if err != nil {
		t.Fail()
	}

	err = CompareHashPassword(hash, password)
	if err != nil {
		t.Fail()
	}
}
