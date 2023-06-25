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

// Récupérer tous les chronos
func GetAllChronos(w http.ResponseWriter, r *http.Request) {
	db := config.Connect()
	defer db.Close()
	var chronos []model.Chrono
	rows, err := db.Query("SELECT * FROM chronos")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var chrono model.Chrono
		err := rows.Scan(&chrono.Id_chrono, &chrono.Id_pilote, &chrono.Id_circuit, &chrono.Chrono_realise, &chrono.Etat_piste, &chrono.Type_voiture, &chrono.Date_du_chrono)
		if err != nil {
			log.Fatal(err)
		}
		chronos = append(chronos, chrono)
	}
	defer rows.Close()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(chronos)
}

// Récupérer un chrono spécifique
func GetOneChrono(w http.ResponseWriter, r *http.Request) {
	db := config.Connect()
	defer db.Close()
	var chrono model.Chrono
	vars := mux.Vars(r)
	id_chrono := vars["id_chrono"]
	rows, err := db.Query("SELECT * FROM chronos WHERE id_chrono = $1", id_chrono)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		err := rows.Scan(&chrono.Id_chrono, &chrono.Id_pilote, &chrono.Id_circuit, &chrono.Chrono_realise, &chrono.Etat_piste, &chrono.Type_voiture, &chrono.Date_du_chrono)
		if err != nil {
			log.Fatal(err)
		}
	}
	defer rows.Close()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(chrono)
}

// Ajouter un chrono
func CreateOneChrono(w http.ResponseWriter, r *http.Request) {
	db := config.Connect()
	defer db.Close()
	var chrono model.Chrono
	err := json.NewDecoder(r.Body).Decode(&chrono)
	if err != nil {
		log.Fatal(err)
	}
	err = db.QueryRow("INSERT INTO chronos (id_pilote, id_circuit, chrono_realise, etat_piste, type_voiture, date_du_chrono) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id_chrono, id_pilote, id_circuit, chrono_realise, etat_piste, type_voiture, date_du_chrono", chrono.Id_pilote, chrono.Id_circuit, chrono.Chrono_realise, chrono.Etat_piste, chrono.Type_voiture, chrono.Date_du_chrono).Scan(&chrono.Id_chrono, &chrono.Id_pilote, &chrono.Id_circuit, &chrono.Chrono_realise, &chrono.Etat_piste, &chrono.Type_voiture, &chrono.Date_du_chrono)
	if err != nil {
		log.Fatal(err)
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(chrono)
}

// Modifier un chrono
func UpdateOneChrono(w http.ResponseWriter, r *http.Request) {
	db := config.Connect()
	defer db.Close()
	var chrono model.Chrono
	err := json.NewDecoder(r.Body).Decode(&chrono)
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec("UPDATE chronos SET id_pilote = $1, id_circuit = $2, chrono_realise = $3, etat_piste = $4, type_voiture = $5, date_du_chrono = $6 WHERE id_chrono = $7", chrono.Id_pilote, chrono.Id_circuit, chrono.Chrono_realise, chrono.Etat_piste, chrono.Type_voiture, chrono.Date_du_chrono, chrono.Id_chrono)
	if err != nil {
		log.Fatal(err)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(chrono)
}

// Supprimer un chrono
func DeleteOneChrono(w http.ResponseWriter, r *http.Request) {
	db := config.Connect()
	defer db.Close()
	var chrono model.Chrono
	vars := mux.Vars(r)
	id_chrono := vars["id_chrono"]
	err := db.QueryRow("DELETE FROM chronos WHERE id_chrono = $1 RETURNING id_chrono, id_pilote, id_circuit, chrono_realise, etat_piste, type_voiture, date_du_chrono", id_chrono).Scan(&chrono.Id_chrono, &chrono.Id_pilote, &chrono.Id_circuit, &chrono.Chrono_realise, &chrono.Etat_piste, &chrono.Type_voiture, &chrono.Date_du_chrono)
	if err != nil {
		log.Fatal(err)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(chrono)
}