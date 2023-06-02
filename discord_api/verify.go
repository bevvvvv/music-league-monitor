// Package discord_api provides wrapper functions for the Discord API.
package discord_api

import (
	"crypto/ed25519"
	"encoding/hex"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

type DiscordInteraction struct {
	Type int `json:"type"`
}

type DiscordInteractionResponse struct {
	Type int `json:"type"`
}

func HandleChallenge(w http.ResponseWriter, r *http.Request) {
	var interaction DiscordInteraction

	// read the request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error reading body", http.StatusBadRequest)
		return
	}

	log.Println(string(body))

	// Decode the JSON body of the request
	err = json.Unmarshal(body, &interaction)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error decoding body", http.StatusBadRequest)
		return
	}

	// Verify the signature
	if !VerifySignature(r, body) {
		http.Error(w, "Invalid signature", http.StatusUnauthorized)
		return
	}

	// If it's a ping event (type 1), respond with a pong event (type 1)
	if interaction.Type == 1 {
		response := DiscordInteractionResponse{Type: 1}
		responseBody, err := json.Marshal(response)
		if err != nil {
			log.Println("Error writing response")
			http.Error(w, "Error writing response", http.StatusInternalServerError)
		}
		w.Write(responseBody)
	}
}

func getPublicKey() ed25519.PublicKey {
	// Replace this with your actual public key.
	keyHex := os.Getenv("DISCORD_PUBLIC_KEY")
	keyBytes, _ := hex.DecodeString(keyHex)
	return keyBytes
}

func VerifySignature(r *http.Request, body []byte) bool {
	signatureHeader := r.Header.Get("X-Signature-Ed25519")
	timestampHeader := r.Header.Get("X-Signature-Timestamp")

	if signatureHeader == "" || timestampHeader == "" {
		return false
	}

	signature, _ := hex.DecodeString(signatureHeader)
	data := []byte(timestampHeader + string(body))

	return ed25519.Verify(getPublicKey(), data, signature)
}

// usage
//	if !verifySignature(r, body) {
//		http.Error(w, "Invalid signature", http.StatusUnauthorized)
//		return
//	}
