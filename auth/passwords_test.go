package auth

import "testing"

func TestPasswords(t *testing.T) {
	pw := "somesecurepassword123"
	hash, err := hashAndSalt([]byte(pw))
	if err != nil {
		t.Error("password hashing should not have resulted in error")
	}

	match := comparePasswords(hash, []byte(pw))
	if !match {
		t.Error("password was expected to have matched")
	}

	match = comparePasswords(hash, []byte("wrongpassword"))
	if match {
		t.Error("passwords should not have matched")
	}

	match = comparePasswords(hash, []byte(""))
	if match {
		t.Error("empty password should not match")
	}

	_, err = hashAndSalt(nil)
	if err == nil {
		t.Error("hash function should not have hashed a nil password")
	}

	_, err = hashAndSalt([]byte(""))
	if err == nil {
		t.Error("hash function should not have hashed an empty string password")
	}
}

func TestGeneratePassword(t *testing.T) {
	pw := "abcdefg"
	hash, err := hashAndSalt([]byte(pw))
	if err != nil {
		t.Error("password hashing should not have resulted in error")
	}

	t.Log(hash)
}
