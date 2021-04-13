package app

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (a *App) Root(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "Hello from GoLanggggg!")
}

func (a* App) createBackendServer(w http.ResponseWriter, r *http.Request){
	var server Server
	json.NewDecoder(r.Body).Decode(&server)
	a.DB.Create(&server)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(server)
}

func (a* App) getBackendServer(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	serverID := params["serverID"]
	var servers []Server
	a.DB.Preload("ServerEndpoints").Find(&servers, serverID)
	json.NewEncoder(w).Encode(servers)
}

func (a* App) updateBackendServer(w http.ResponseWriter, r *http.Request){

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["serverID"])

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	var updatedServer Server

	if err:= json.NewDecoder(r.Body).Decode(&updatedServer); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload.")
		return
	}

	updatedServer.ServerID = uint(id)

	var inInterface map[string]interface{}
	inrec, err := json.Marshal(updatedServer)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload.")
		return
	}

	err = json.Unmarshal(inrec, &inInterface)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Something went wrong.")
		return
	}

	a.DB.Model(&updatedServer).Updates(inInterface)


	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(updatedServer); err != nil {
		respondWithError(w, http.StatusInternalServerError, "Something went wrong")
		return
	}
}