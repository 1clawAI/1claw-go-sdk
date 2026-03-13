package cmek

import (
	"testing"
)

func TestGenerateKey(t *testing.T) {
	key, err := GenerateKey()
	if err != nil {
		t.Fatal(err)
	}
	if len(key) != 32 {
		t.Errorf("key length = %d, want 32", len(key))
	}
}

func TestFingerprint(t *testing.T) {
	key := make([]byte, 32)
	for i := range key {
		key[i] = byte(i)
	}
	fp := Fingerprint(key)
	if len(fp) != 64 {
		t.Errorf("fingerprint length = %d, want 64 (hex)", len(fp))
	}
}

func TestEncryptDecryptRoundTrip(t *testing.T) {
	key, err := GenerateKey()
	if err != nil {
		t.Fatal(err)
	}
	plaintext := "my-secret-value"
	encrypted, err := Encrypt(key, plaintext)
	if err != nil {
		t.Fatal(err)
	}
	if encrypted == plaintext {
		t.Error("encrypted should differ from plaintext")
	}
	decrypted, err := Decrypt(key, encrypted)
	if err != nil {
		t.Fatal(err)
	}
	if decrypted != plaintext {
		t.Errorf("decrypted = %q, want %q", decrypted, plaintext)
	}
}

func TestDecrypt_WrongKey(t *testing.T) {
	key1, _ := GenerateKey()
	key2, _ := GenerateKey()
	encrypted, _ := Encrypt(key1, "secret")
	_, err := Decrypt(key2, encrypted)
	if err == nil {
		t.Error("expected error when decrypting with wrong key")
	}
}
