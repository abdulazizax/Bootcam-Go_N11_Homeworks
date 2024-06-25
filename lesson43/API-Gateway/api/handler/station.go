package handler

import (
	"bytes"
	"encoding/json"
	m "api-gateway/models"
	"io"
	"net/http"
)

func (h *handler) CreateStation(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	resp, err := http.Post("http://localhost:8082/station/post", "application/json", bytes.NewBuffer(body))
	if err != nil {
		http.Error(w, "Failed to send POST request", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read response body", http.StatusInternalServerError)
		return
	}

	var book m.StationResponse
	if err := json.Unmarshal(respBody, &book); err != nil {
		http.Error(w, "Failed to unmarshal response body", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(book); err != nil {
		http.Error(w, "Failed to encode response body", http.StatusInternalServerError)
		return
	}
}

func (h *handler) GetStations(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://localhost:8082/station/get")
	if err != nil {
		http.Error(w, "Failed to send GET request", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read response body", http.StatusInternalServerError)
		return
	}

	var books []m.StationResponse
	if err := json.Unmarshal(body, &books); err != nil {
		http.Error(w, "Failed to unmarshal response body", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(books); err != nil {
		http.Error(w, "Failed to encode response body", http.StatusInternalServerError)
		return
	}
}

func (h *handler) GetStationById(w http.ResponseWriter, r *http.Request) {
	id := GetIDFromURL(r.URL.Path)
	if id == "" {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	resp, err := http.Get("http://localhost:8082/station/get/" + id)
	if err != nil {
		http.Error(w, "Failed to send GET request", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read response body", http.StatusInternalServerError)
		return
	}

	var books m.StationResponse
	if err := json.Unmarshal(body, &books); err != nil {
		http.Error(w, "Failed to unmarshal response body", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(books); err != nil {
		http.Error(w, "Failed to encode response body", http.StatusInternalServerError)
		return
	}
}

func (h *handler) GetStationByName(w http.ResponseWriter, r *http.Request) {
	name := GetIDFromURL(r.URL.Path)
	if name == "" {
		http.Error(w, "Invalid Name", http.StatusBadRequest)
		return
	}

	resp, err := http.Get("http://localhost:8082/station/get/name/" + name)
	if err != nil {
		http.Error(w, "Failed to send GET request", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read response body", http.StatusInternalServerError)
		return
	}

	var books m.StationResponse
	if err := json.Unmarshal(body, &books); err != nil {
		http.Error(w, "Failed to unmarshal response body", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(books); err != nil {
		http.Error(w, "Failed to encode response body", http.StatusInternalServerError)
		return
	}
}

func (h *handler) UpdateStationById(w http.ResponseWriter, r *http.Request) {
	id := GetIDFromURL(r.URL.Path)
	if id == "" {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	req, err := http.NewRequest("PUT", "http://localhost:8082/station/put/"+id, bytes.NewBuffer(body))
	if err != nil {
		http.Error(w, "Failed to create new request", http.StatusInternalServerError)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := h.client.Do(req)
	if err != nil {
		http.Error(w, "Failed to send PUT request", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		http.Error(w, "No book found with the given ID", http.StatusNotFound)
		return
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read response body", http.StatusInternalServerError)
		return
	}

	var book m.StationResponse
	if err := json.Unmarshal(respBody, &book); err != nil {
		http.Error(w, "Failed to unmarshal response body", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(book); err != nil {
		http.Error(w, "Failed to encode response body", http.StatusInternalServerError)
		return
	}
}

func (h *handler) DeleteStationByID(w http.ResponseWriter, r *http.Request) {
	id := GetIDFromURL(r.URL.Path)

	if id == "" {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if h.client == nil {
		http.Error(w, "HTTP client is not initialized", http.StatusInternalServerError)
		return
	}

	req, err := http.NewRequest("DELETE", "http://localhost:8082/station/delete/"+id, nil)
	if err != nil {
		http.Error(w, "Failed to create new request", http.StatusInternalServerError)
		return
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := h.client.Do(req)
	if err != nil {
		http.Error(w, "Failed to send DELETE request", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		http.Error(w, "No book found with the given ID", http.StatusNotFound)
		return
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read response body", http.StatusInternalServerError)
		return
	}

	var result string
	if err := json.Unmarshal(respBody, &result); err != nil {
		http.Error(w, "Failed to unmarshal response body", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(result); err != nil {
		http.Error(w, "Failed to encode response body", http.StatusInternalServerError)
		return
	}
}
