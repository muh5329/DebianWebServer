package main

import (
	"log"
	"net/http"
	"os/exec"
)

func main() {
	http.HandleFunc("/webhook", func(w http.ResponseWriter, r *http.Request) {
		// Verify the request is coming from GitHub
		// Parse the webhook payload
		// Check if the event is a merge to the main branch

		// If the event is a merge to the main branch, trigger the build process
		err := triggerBuild()
		if err != nil {
			log.Println("Error triggering build process:", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
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
