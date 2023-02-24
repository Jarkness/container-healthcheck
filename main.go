package main

import (
	"context"
	"encoding/json"
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
		http.Error(w, "Unable to fetch container list", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(containers)
}

func main() {
	http.HandleFunc("/healthz", get_container_healthz)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
