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

// Récupérer tous les pilotes
func GetAllPilotes(w http.ResponseWriter, r *http.Request) {
	db := config.Connect()
	defer db.Close()
	var pilotes []model.Pilote
	rows, err := db.Query("SELECT * FROM pilotes")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var pilote model.Pilote
		err := rows.Scan(&pilote.Id_pilote, &pilote.Nom, &pilote.Prenom, &pilote.Date_naissance, &pilote.Categorie, &pilote.Annees_experience)
		if err != nil {
			log.Fatal(err)
		}
		pilotes = append(pilotes, pilote)
	}
	defer rows.Close()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(pilotes)
}

// Récupérer un pilote spécifique
func GetOnePilote(w http.ResponseWriter, r *http.Request) {
	db := config.Connect()
	defer db.Close()
	var pilote model.Pilote
	vars := mux.Vars(r)
	id_pilote := vars["id_pilote"]
	rows, err := db.Query("SELECT * FROM pilotes WHERE id_pilote = $1", id_pilote)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		err := rows.Scan(&pilote.Id_pilote, &pilote.Nom, &pilote.Prenom, &pilote.Date_naissance, &pilote.Categorie, &pilote.Annees_experience)
		if err != nil {
			log.Fatal(err)
		}
	}
	defer rows.Close()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(pilote)
}

// Ajouter un pilote
func CreateOnePilote(w http.ResponseWriter, r *http.Request) {
	db := config.Connect()
	defer db.Close()
	var pilote model.Pilote
	err := json.NewDecoder(r.Body).Decode(&pilote)
	if err != nil {
		log.Fatal(err)
	}
	err = db.QueryRow("INSERT INTO pilotes (nom, prenom, date_naissance, categorie, annees_experience) VALUES ($1, $2, $3, $4, $5) RETURNING id_pilote, nom, prenom, date_naissance, categorie, annees_experience", pilote.Nom, pilote.Prenom, pilote.Date_naissance, pilote.Categorie, pilote.Annees_experience).Scan(&pilote.Id_pilote, &pilote.Nom, &pilote.Prenom, &pilote.Date_naissance, &pilote.Categorie, &pilote.Annees_experience)
	if err != nil {
		log.Fatal(err)
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(pilote)
}

// Supprimer un pilote
func DeleteOnePilote(w http.ResponseWriter, r *http.Request) {
	db := config.Connect()
	defer db.Close()
	vars := mux.Vars(r)
	id_pilote := vars["id_pilote"]
	var pilote model.Pilote
	err := db.QueryRow("DELETE FROM pilotes WHERE id_pilote = $1 RETURNING id_pilote, nom, prenom, date_naissance, categorie, annees_experience", id_pilote).Scan(&pilote.Id_pilote, &pilote.Nom, &pilote.Prenom, &pilote.Date_naissance, &pilote.Categorie, &pilote.Annees_experience)
	if err != nil {
		log.Fatal(err)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(pilote)
}

// Modifier un pilote
func UpdateOnePilote(w http.ResponseWriter, r *http.Request) {
	db := config.Connect()
	defer db.Close()
	vars := mux.Vars(r)
	id_pilote := vars["id_pilote"]
	var pilote model.Pilote
	err := json.NewDecoder(r.Body).Decode(&pilote)
	if err != nil {
		log.Fatal(err)
	}
	err = db.QueryRow("UPDATE pilotes SET nom = $1, prenom = $2, date_naissance = $3, categorie = $4, annees_experience = $5 WHERE id_pilote = $6 RETURNING id_pilote, nom, prenom, date_naissance, categorie, annees_experience", pilote.Nom, pilote.Prenom, pilote.Date_naissance, pilote.Categorie, pilote.Annees_experience, id_pilote).Scan(&pilote.Id_pilote, &pilote.Nom, &pilote.Prenom, &pilote.Date_naissance, &pilote.Categorie, &pilote.Annees_experience)
	if err != nil {
		log.Fatal(err)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(pilote)
}
