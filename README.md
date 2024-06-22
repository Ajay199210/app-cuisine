
# Application de Cuisine

Ce projet est une application full-stack qui permet aux utilisateurs d'explorer des recettes de diverses cuisines du monde entier. Le frontend est développé en Angular, tandis que l'API backend est propulsée par Go.

*L'API TheMealDB a été utilisée pour l'affichage des recettes.
Vous pouvez consulter leur site en visitant le lien suivant :
https://www.themealdb.com/*

![image](https://github.com/Ajay199210/app-cuisine/assets/46723178/c12f1136-df05-425e-8cf7-176481d469b0)

![image](https://github.com/Ajay199210/app-cuisine/assets/46723178/39a26225-afe9-4d5f-a0a5-f4f75130289b)

## Sommaire

- [Fonctionnalités](#fonctionnalités)
- [Technologies utilisées](#technologies-utilisées)
- [Installation](#installation)

## Fonctionnalités

- **Authentification** : Inscription et connexion.
- **Gestion des recettes** : Ajout, visualisation et mise à jour. 
- **Recherche et filtrage** : Recherche de recettes par nom.
- **Favoris** : Sauvegardez vos recettes préférées pour un accès facile.

## Technologies Utilisées

### Frontend

- **Angular** : Un puissant framework pour créer des applications web dynamiques.
- **RxJS** : Pour gérer les opérations asynchrones et la gestion d'état.

### Backend

- **Go** : Un langage de programmation compilé, typé statiquement, conçu pour créer des services backend évolutifs et efficaces.
- **Gin** : Un framework web pour Go qui offre un ensemble de fonctionnalités robustes pour créer des applications web et des microservices.

## Installation

### Prérequis

- Node.js et npm
- Go (version 1.16 ou ultérieure)

### Installation du frontend

1. Clonez le dépôt :
   ```
   git clone https://github.com/ajay199210/app-cuisine.git
   cd app-cuisine/src/Frontend
   ```
2. Installez les dépendances 
   ```
   npm install
   ```
3. Démarrez le serveur de développement 
   ```
   ng serve
   ```

### Installation du backend

1. Accédez au répertoire backend :
   ```
   cd app-cuisine/src/Backend
   ```

2. Exécutez le serveur
   ```
   go run .
   ```
   
  ***N.B***: *la commande `go run .` télécharge les dépendances automatiquement.*

## Contributions
Merci à [aaguerre1105](https://github.com/aaguerre1105) et [amarboutiti](https://github.com/amarboutiti) pour leurs contributions au projet !
