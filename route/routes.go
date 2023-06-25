package route

import (
	"api_go_mathieu_fourtane/controller"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"io/ioutil"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// Lire le contenu du fichier index.html
	html, err := ioutil.ReadFile("views/index.html")
	if err != nil {
		// Gérer l'erreur de lecture du fichier
		http.Error(w, "Erreur lors de la lecture du fichier", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write(html)
}

func GetAllOfficialCircuitsHandler(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()
	limit := queryValues.Get("limit")
	circuits, err := controller.GetAllOfficialCircuits(limit)
	if err != nil {
		// Gérer l'erreur de récupération des circuits
		http.Error(w, "Erreur lors de la récupération des circuits", http.StatusInternalServerError)
		return
	}

	// Répondre avec les circuits récupérés en tant que réponse JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(circuits)
}

func Router() *mux.Router {
	router := mux.NewRouter()

	// Créer une instance de cors.Handler avec les options appropriées
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
	})

	// Ajouter le middleware CORS au routeur
	router.Use(corsMiddleware.Handler)

	// routes table pilotes
	router.HandleFunc("/pilotes", controller.GetAllPilotes).Methods("GET")
	//router.Handle("/pilotes", controller.AuthentificationMiddleware(http.HandlerFunc(controller.GetAllPilotes))).Methods("GET")
	router.HandleFunc("/pilotes/{id_pilote}", controller.GetOnePilote).Methods("GET")
	router.HandleFunc("/pilotes", controller.CreateOnePilote).Methods("POST")
	router.HandleFunc("/pilotes/{id_pilote}", controller.DeleteOnePilote).Methods("DELETE")
	router.HandleFunc("/pilotes/{id_pilote}", controller.UpdateOnePilote).Methods("PUT")

	// routes tables circuits
	router.HandleFunc("/circuits", controller.GetAllCircuits).Methods("GET")
	router.HandleFunc("/circuits/{id_circuit}", controller.GetOneCircuit).Methods("GET")
	router.HandleFunc("/circuits", controller.CreateOneCircuit).Methods("POST")
	router.HandleFunc("/circuits/{id_circuit}", controller.DeleteOneCircuit).Methods("DELETE")
	router.HandleFunc("/circuits/{id_circuit}", controller.UpdateOneCircuit).Methods("PUT")

	//routes tables chronos
	router.HandleFunc("/chronos", controller.GetAllChronos).Methods("GET")
	router.HandleFunc("/chronos/{id_chrono}", controller.GetOneChrono).Methods("GET")
	router.HandleFunc("/chronos", controller.CreateOneChrono).Methods("POST")
	router.HandleFunc("/chronos/{id_chrono}", controller.DeleteOneChrono).Methods("DELETE")
	router.HandleFunc("/chronos/{id_chrono}", controller.UpdateOneChrono).Methods("PUT")

	//routes offcials circuits (API externe : http://ergast.com/api)
	router.HandleFunc("/official-circuits", GetAllOfficialCircuitsHandler).Methods("GET")

	//routes utilsiateurs
	router.HandleFunc("/register", controller.RegisterUser).Methods("POST")
	router.HandleFunc("/login", controller.AuthenticateUser).Methods("POST")

	// Route de base
	router.HandleFunc("/", HomeHandler).Methods("GET")

	return router

}
