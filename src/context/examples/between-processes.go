package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// BETWEEN PROCESSES == communicate BETWEEN DIFFERENT processes

func main() {
	// TODO: comprehend from here
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Context transporta información BETWEEN processes via HTTP
	callExternalAPI(ctx, "user-456")
}

func callExternalAPI(ctx context.Context, userID string) {
	// Crear request HTTP con context (va a otro proceso)
	req, _ := http.NewRequestWithContext(ctx, "GET", "https://jsonplaceholder.typicode.com/users/1", nil)

	// Headers transportan valores entre procesos
	req.Header.Set("X-User-ID", userID)
	req.Header.Set("X-Request-ID", "req-789")

	client := &http.Client{Timeout: 3 * time.Second}

	fmt.Println("Calling external API...")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error calling external process:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Printf("External process responded: %s\n", resp.Status)
	fmt.Printf("Context deadline transported between processes\n")
}
