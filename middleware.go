package main

import (
	"crypto/rand"
	"encoding/hex"
)

// authentication 認証
func (cfg *apiConfig) middleware() {

}
func generateAPIKey() (string, error) {
	bytes := make([]byte, 16) // 16 bytes = 128-bit key
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
