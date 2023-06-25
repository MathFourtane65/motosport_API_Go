package model

// STRUCTURES DE DONNEES POUR LES TABLES DE LA BDD
type Pilote struct {
	Id_pilote         int    `json:"id_pilote"`
	Nom               string `json:"nom"`
	Prenom            string `json:"prenom"`
	Date_naissance    string `json:"date_naissance"`
	Categorie         string `json:"categorie"`
	Annees_experience int    `json:"annees_experience"`
}

type Circuit struct {
	Id_circuit     int     `json:"id_circuit"`
	Nom            string  `json:"nom"`
	Ville          string  `json:"ville"`
	Pays           string  `json:"pays"`
	Url            string  `json:"url"`
	Longueur       float64 `json:"longueur"`
	Nombre_virages int     `json:"nombre_virages"`
}

type Chrono struct {
	Id_chrono      int    `json:"id_chrono"`
	Id_pilote      int    `json:"id_pilote"`
	Id_circuit     int    `json:"id_circuit"`
	Chrono_realise string `json:"chrono_realise"`
	Etat_piste     string `json:"etat_piste"`
	Type_voiture   string `json:"type_voiture"`
	Date_du_chrono string `json:"date_du_chrono"`
}

type Utilisateur struct {
	Id_utilisateur int    `json:"id_utilisateur"`
	Identifiant    string `json:"identifiant"`
	Mot_de_passe   string `json:"mot_de_passe"`
}
