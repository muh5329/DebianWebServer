package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
)

var secret = []byte("your-github-webhook-secret")

func main() {
	http.HandleFunc("/webhook", func(w http.ResponseWriter, r *http.Request) {
		// Verify the request is coming from GitHub
		signature := r.Header.Get("X-Hub-Signature")
		payload, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("Error reading request body:", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		if !verifySignature(signature, payload) {
			log.Println("Signature verification failed")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// Parse the webhook payload
		// Check if the event is a merge to the main branch

		// If the event is a merge to the main branch, trigger the build process
		err = triggerBuild()
		if err != nil {
			log.Println("Error triggering build process:", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func verifySignature(signature string, payload []byte) bool {
	mac := hmac.New(sha1.New, secret)
	mac.Write(payload)
	expectedMAC := hex.EncodeToString(mac.Sum(nil))

	return "sha1="+expectedMAC == signature
}

func triggerBuild() error {
	// Pull latest changes from GitHub
	cmd := exec.Command("git", "pull", "origin", "main")
	err := cmd.Run()
	if err != nil {
		return err
	}

	// Build Next.js project as a production app
	cmd = exec.Command("npm", "run", "build")
	err = cmd.Run()
	if err != nil {
		return err
	}

	// Optionally, deploy the built app to your production environment

	return nil
}
