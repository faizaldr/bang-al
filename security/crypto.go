package security

import (
	"crypto/aes"              //jenis algoritma
	"crypto/cipher"           //encriptor
	crypto_rand "crypto/rand" //menggenerate nilai random
	"encoding/base64"         //encoding (mengubah bentuk) standar base64
	"errors"
)

// base64 URL-safe
var b64 = base64.RawURLEncoding

// fungsi untuk melakukan Encrypt url bytes -> string url-safe
// plainttext : nip , key : environment variable secret-key
func EncryptURLSafe(plainttext []byte, key []byte) (string, error) {
	if len(key) != 32 {
		return "", errors.New("kunci harus berupa 32 byte (AES-256)")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize()) // 12 bytes
	if _, err := crypto_rand.Read(nonce); err != nil {
		return "", err
	}

	ciphertext := gcm.Seal(nil, nonce, plainttext, nil)
	out := append(nonce, ciphertext...)
	return b64.EncodeToString(out), nil
}
