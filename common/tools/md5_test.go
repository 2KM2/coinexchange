package tools

import (
	"fmt"
	"testing"
)

func TestGenerateSalt(t *testing.T) {
	const saltLength = 16
	salt := generateSalt(saltLength)
	if len(salt) != saltLength {
		t.Errorf("Expected salt length of %d, but got %d", saltLength, len(salt))
	}
}

func TestEncodeAndVerifyWithDefaultOptions(t *testing.T) {
	rawPassword := "testpassword"
	salt, encodedPwd := Encode(rawPassword, nil)
	fmt.Println("salt len ", len(salt), "encode len ", len(encodedPwd))
	if len(salt) != defaultSaltLen {
		t.Errorf("Expected salt length of %d, but got %d", defaultSaltLen, len(salt))
	}

	if len(encodedPwd) == 0 {
		t.Error("Expected encoded password to be non-empty")
	}

	if !Verify(rawPassword, salt, encodedPwd, nil) {
		t.Error("Verification failed for the correct password")
	}

	if Verify("wrongpassword", salt, encodedPwd, nil) {
		t.Error("Verification succeeded for an incorrect password")
	}
}

func TestEncodeAndVerifyWithCustomOptions(t *testing.T) {
	customOptions := &Options{
		SaltLen:      32,
		Iterations:   5000,
		KeyLen:       64,
		HashFunction: defaultHashFunction,
	}
	rawPassword := "custompassword"
	salt, encodedPwd := Encode(rawPassword, customOptions)

	if len(salt) != customOptions.SaltLen {
		t.Errorf("Expected salt length of %d, but got %d", customOptions.SaltLen, len(salt))
	}

	if len(encodedPwd) == 0 {
		t.Error("Expected encoded password to be non-empty")
	}

	if !Verify(rawPassword, salt, encodedPwd, customOptions) {
		t.Error("Verification failed for the correct password with custom options")
	}

	if Verify("wrongpassword", salt, encodedPwd, customOptions) {
		t.Error("Verification succeeded for an incorrect password with custom options")
	}
}
