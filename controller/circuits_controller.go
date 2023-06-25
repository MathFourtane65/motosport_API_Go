package controller

import (
	"api_go_mathieu_fourtane/config"
	"api_go_mathieu_fourtane/model"
	"encoding/json"

	//"fmt"
	"net/http"
	//"strconv"
	"log"

	"github.com/gorilla/mux"
)

// Récupérer tous les circuits
func GetAllCircuits(w http.ResponseWriter, r *http.Request) {
	db := config.Connect()
	defer db.Close()
	var circuits []model.Circuit
	rows, err := db.Query("SELECT * FROM circuits")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var circuit model.Circuit
		err := rows.Scan(&circuit.Id_circuit, &circuit.Nom, &circuit.Ville, &circuit.Pays, &circuit.Url, &circuit.Longueur, &circuit.Nombre_virages)
		if err != nil {
			log.Fatal(err)
		}
		circuits = append(circuits, circuit)
	}
	defer rows.Close()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(circuits)
}

// Récupérer un circuit spécifique
func GetOneCircuit(w http.ResponseWriter, r *http.Request) {
	db := config.Connect()
	defer db.Close()
	var circuit model.Circuit
	vars := mux.Vars(r)
	id_circuit := vars["id_circuit"]
	rows, err := db.Query("SELECT * FROM circuits WHERE id_circuit = $1", id_circuit)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		err := rows.Scan(&circuit.Id_circuit, &circuit.Nom, &circuit.Ville, &circuit.Pays, &circuit.Url, &circuit.Longueur, &circuit.Nombre_virages)
		if err != nil {
			log.Fatal(err)
		}
	}
	defer rows.Close()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(circuit)
}

// Créer un circuit
func CreateOneCircuit(w http.ResponseWriter, r *http.Request) {
	db := config.Connect()
	defer db.Close()
	var circuit model.Circuit
	err := json.NewDecoder(r.Body).Decode(&circuit)
	if err != nil {
		log.Fatal(err)
	}
	err = db.QueryRow("INSERT INTO circuits (nom, ville, pays, url, longueur, nombre_virages) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id_circuit, nom, ville, pays, url, longueur, nombre_virages", circuit.Nom, circuit.Ville, circuit.Pays, circuit.Url, circuit.Longueur, circuit.Nombre_virages).Scan(&circuit.Id_circuit, &circuit.Nom, &circuit.Ville, &circuit.Pays, &circuit.Url, &circuit.Longueur, &circuit.Nombre_virages)
	if err != nil {
		log.Fatal(err)
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(circuit)
}

// Supprimer un circuit
func DeleteOneCircuit(w http.ResponseWriter, r *http.Request) {
	db := config.Connect()
	defer db.Close()
	vars := mux.Vars(r)
	id_circuit := vars["id_circuit"]
	var circuit model.Circuit
	err := db.QueryRow("DELETE FROM circuits WHERE id_circuit = $1 RETURNING id_circuit, nom, ville, pays, url, longueur, nombre_virages", id_circuit).Scan(&circuit.Id_circuit, &circuit.Nom, &circuit.Ville, &circuit.Pays, &circuit.Url, &circuit.Longueur, &circuit.Nombre_virages)
	if err != nil {
		log.Fatal(err)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(circuit)
}


// Mettre à jour un circuit
func UpdateOneCircuit(w http.ResponseWriter, r *http.Request) {
	db := config.Connect()
	defer db.Close()
	vars := mux.Vars(r)
	id_circuit := vars["id_circuit"]
	var circuit model.Circuit
	err := json.NewDecoder(r.Body).Decode(&circuit)
	if err != nil {
		log.Fatal(err)
	}
	err = db.QueryRow("UPDATE circuits SET nom = $1, ville = $2, pays = $3, url = $4, longueur = $5, nombre_virages = $6 WHERE id_circuit = $7 RETURNING id_circuit, nom, ville, pays, url, longueur, nombre_virages", circuit.Nom, circuit.Ville, circuit.Pays, circuit.Url, circuit.Longueur, circuit.Nombre_virages, id_circuit).Scan(&circuit.Id_circuit, &circuit.Nom, &circuit.Ville, &circuit.Pays, &circuit.Url, &circuit.Longueur, &circuit.Nombre_virages)
	if err != nil {
		log.Fatal(err)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(circuit)
}
