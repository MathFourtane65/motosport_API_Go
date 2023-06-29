# Projet API Go

Ce projet est une API écrite en langage Go qui gère la gestion des pilotes de course, des circuits et des chronos.

## Fonctionnalités

L'API offre les fonctionnalités suivantes :

- Récupération de la liste des pilotes
- Récupération d'un pilote spécifique par son ID
- Création d'un nouveau pilote
- Suppression d'un pilote existant
- Mise à jour des informations d'un pilote
- Récupération de la liste des circuits
- Récupération d'un circuit spécifique par son ID
- Création d'un nouveau circuit
- Suppression d'un circuit existant
- Mise à jour des informations d'un circuit
- Récupération de la liste des chronos
- Récupération d'un chrono spécifique par son ID
- Création d'un nouveau chrono
- Suppression d'un chrono existant
- Mise à jour des informations d'un chrono
- Récupération des circuits officiels via une API externe

## Installation

1. Assurez-vous d'avoir Go installé sur votre machine.
2. Clonez ce dépôt : `git clone https://github.com/votre-utilisateur/api-go`
3. Accédez au répertoire du projet : `cd api-go`
4. Installez les dépendances en exécutant la commande suivante : `go mod download`

## Configuration

1. Assurez-vous d'avoir une base de données PostgreSQL installée et configurée.
2. Créez une base de données pour l'application.
3. Configurez les paramètres de connexion à la base de données dans le fichier `config.go`.

## Utilisation

1. Exécutez l'application en utilisant la commande suivante : `go run main.go`.
2. L'API sera accessible à l'adresse `http://localhost:8000`.

## Documentation de l'API

Vous pouvez accéder à la documentation en ouvrant le fichier `views/index.html` dans votre navigateur. Elle est égaelment accessible sur la route "/" une fois le serveur en marche.
