package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/joho/godotenv"
)

type PullRequestEvent struct {
	Action      string `json:"action"`
	PullRequest struct {
		Number int  `json:"number"`
		Merged bool `json:"merged"`
		Base   struct {
			Ref string `json:"ref"`
		} `json:"base"`
	} `json:"pull_request"`
	Repository struct {
		Name  string `json:"name"`
		Owner struct {
			Login string `json:"login"`
		} `json:"owner"`
	} `json:"repository"`
}

func main() {

	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)

	// .env file : loading secrets
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error in loading env file")
	}

	ghSecret := []byte(os.Getenv("GH_WEBHOOK_SECRET"))

	http.HandleFunc("/webhook", func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		// Verify the request is coming from GitHub
		signature := r.Header.Get("X-Hub-Signature")
		payload, err := io.ReadAll(r.Body)
		if err != nil {
			log.Println("Error reading request body:", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		if !verifySignature(signature, payload, ghSecret) {
			log.Println("Signature verification failed")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		var event PullRequestEvent
		if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
			log.Println("Failed to decode webhook payload: \n", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if event.Action != "closed" || !event.PullRequest.Merged {
			// Ignore if the action is not "closed" or the pull request is not merged
			return
		}

		if event.PullRequest.Base.Ref == "main" {
			// If the event is a merge to the main branch, trigger the build process
			err = triggerBuild()
			if err != nil {
				log.Println("Error triggering build process:", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}

		w.WriteHeader(http.StatusOK)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func verifySignature(signature string, payload []byte, secret []byte) bool {
	mac := hmac.New(sha1.New, secret)
	mac.Write(payload)
	expectedMAC := hex.EncodeToString(mac.Sum(nil))

	return "sha1="+expectedMAC == signature
}

func triggerBuild() error {
	// Pull latest changes from GitHub
	log.Println("Pulling from main")
	cmd := exec.Command("git", "pull", "origin", "main")
	err := cmd.Run()
	if err != nil {
		return err
	}

	// Build Next.js project as a production app
	log.Println("Running pnpm build")
	cmd = exec.Command("pnpm", "run", "build")
	err = cmd.Run()
	if err != nil {
		return err
	}

	// Optionally, deploy the built app to your production environment

	return nil
}
