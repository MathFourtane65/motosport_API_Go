package controller

import (
	"api_go_mathieu_fourtane/config"
	"api_go_mathieu_fourtane/model"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

const secretKey = "rtzbslxnsCKHQSCGQIYAEBDQKJH"

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	db := config.Connect()
	defer db.Close()

	var user model.Utilisateur
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Fatal(err)
	}

	// Vérifiez si l'utilisateur existe déjà dans la base de données
	var existingUser model.Utilisateur
	err = db.QueryRow("SELECT id_utilisateur FROM utilisateurs WHERE identifiant = $1", user.Identifiant).Scan(&existingUser.Id_utilisateur)
	if err != sql.ErrNoRows {
		// L'utilisateur existe déjà
		http.Error(w, "Cet identifiant est déjà utilisé", http.StatusBadRequest)
		return
	}

	// Hash du mot de passe avant de le stocker dans la base de données
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Mot_de_passe), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

	// Insérer le nouvel utilisateur dans la base de données
	_, err = db.Exec("INSERT INTO utilisateurs (identifiant, mot_de_passe) VALUES ($1, $2)", user.Identifiant, string(hashedPassword))
	if err != nil {
		log.Fatal(err)
	}

	// Répondre avec une réponse JSON indiquant que l'utilisateur a été enregistré avec succès
	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{
		"message": "Utilisateur enregistré avec succès",
	}
	json.NewEncoder(w).Encode(response)
}

func AuthenticateUser(w http.ResponseWriter, r *http.Request) {
	db := config.Connect()
	defer db.Close()

	var user model.Utilisateur
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Fatal(err)
	}

	// Vérifier si l'utilisateur existe dans la base de données
	var existingUser model.Utilisateur
	err = db.QueryRow("SELECT id_utilisateur, mot_de_passe FROM utilisateurs WHERE identifiant = $1", user.Identifiant).Scan(&existingUser.Id_utilisateur, &existingUser.Mot_de_passe)
	if err != nil {
		if err == sql.ErrNoRows {
			// L'utilisateur n'existe pas
			http.Error(w, "Identifiant ou mot de passe incorrect", http.StatusUnauthorized)
		} else {
			log.Fatal(err)
		}
		return
	}

	// Vérifier le mot de passe hashé
	err = bcrypt.CompareHashAndPassword([]byte(existingUser.Mot_de_passe), []byte(user.Mot_de_passe))
	if err != nil {
		// Mot de passe incorrect
		http.Error(w, "Identifiant ou mot de passe incorrect", http.StatusUnauthorized)
		return
	}

	// Générer le token JWT pour l'utilisateur authentifié
	token, err := generateJWT(existingUser.Id_utilisateur)
	if err != nil {
		http.Error(w, "Erreur lors de la génération du token JWT", http.StatusInternalServerError)
		return
	}

	// Répondre avec le token JWT en tant que réponse JSON
	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{
		"token": token,
	}
	json.NewEncoder(w).Encode(response)
}

func generateJWT(userId int) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["userId"] = userId
	// Ajoutez d'autres informations au JWT si nécessaire

	// Signez le token avec votre clé secrète
	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func validateJWT(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("Invalid token")
}

func AuthentificationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "Token manquant", http.StatusUnauthorized)
			return
		}

		claims, err := validateJWT(tokenString)
		if err != nil {
			http.Error(w, "Token invalide", http.StatusUnauthorized)
			return
		}
		fmt.Println(claims["userId"])

		// Ajoutez des vérifications supplémentaires si nécessaire, par exemple, vérifiez les autorisations de l'utilisateur.

		// Si toutes les vérifications sont réussies, passez à l'handler suivant
		next.ServeHTTP(w, r)
	})
}