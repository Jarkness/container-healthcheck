package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

var docker_client, _ = client.NewClientWithOpts(client.FromEnv)

func get_container_healthz(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	containers, err := docker_client.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		http.Error(w, "Unable to fetch container list: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(containers)
}

func main() {
	fmt.Printf("Starting container info service")
	http.HandleFunc("/health", get_container_healthz)
	log.Fatal(http.ListenAndServe(":80", nil))
}
