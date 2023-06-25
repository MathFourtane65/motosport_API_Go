package controller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type OfficialCircuit struct {
	Id_official_circuit string `json:"circuitId"`
	Url                 string `json:"url"`
	CircuitName         string `json:"circuitName"`
	Location            struct {
		Lat      string `json:"lat"`
		Long     string `json:"long"`
		Locality string `json:"locality"`
		Country  string `json:"country"`
	} `json:"Location"`
}

func GetAllOfficialCircuits(limit string) ([]OfficialCircuit, error) {
	url := "http://ergast.com/api/f1/circuits.json"
	if limit != "" {
		url += "?limit=" + limit
	}
	// Faites une requête GET à l'API pour récupérer la liste des circuits
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// Vérifiez si la réponse de l'API est réussie (code d'état 200)
	if response.StatusCode != http.StatusOK {
		log.Fatalf("La requête a échoué avec le code d'état %d", response.StatusCode)
	}

	// Lire le corps de la réponse
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Analysez le corps de la réponse JSON dans une structure de données
	var result struct {
		MRData struct {
			CircuitTable struct {
				Circuits []OfficialCircuit `json:"Circuits"`
			} `json:"CircuitTable"`
		} `json:"MRData"`
	}
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Fatal(err)
	}

	// Récupérer la liste des circuits
	circuits := result.MRData.CircuitTable.Circuits

	return circuits, nil
}
