#**Rapport de projet – Puissance 4**

*Ce projet a pour but de développer un Puissance 4 jouable dans le terminal. On a travaillé en équipe et utilisé les Git pour organiser notre projet*

---

**Branches utilisées dans le projet :**
_______________________________________________________
| Branche      | Rôle                                 |
|--------------|--------------------------------------|
| main         | Version stable du projet             |
| feat/ft1     | Initialisation du jeu                |
| feat/ft3     | Placement des jetons                 |
| feat/ft4     | Match nul et alternance des joueurs  |
-------------------------------------------------------

---

**Conventions Git qu’on a suivies :**

Pour nommer les branches, on a utilisé le format : `type/nom_fonctionnalité`. Par exemple :
- `feat/ft1` pour démarrer le jeu avec une grille vide et les variables de base
- `feat/ft3` pour gérer l’insertion des jetons
- `feat/ft4` pour détecter les égalités et alterner les joueurs

Les messages de commit étaient simples et clairs, comme :
- `feat: ajout de la vérification de victoire`
- `feat: ajout de fin en cas d'égalité`

---

**Commandes Git qu’on a utilisées régulièrement :**

- `git clone https://github.com/nicol1234/collaboration-Git.git`  
  → Nous sert à récupérer le projet en local

- `git switch -c feat/ft3`  
  → Nous à permis de crée une nouvelle branche

- `git checkout main`  
  → Nous avons utilisé cette commande pour revenir sur la branche principale main

- `git pull`  
  → Nous avons utilisé git pull pour mètre a jour la copie locale avec les dernières modifications du dépôt Github

- `git branch`  
  → git branch nous a permis de voir toutes les branches existantes du project et de vérifié sur quel branche nous nous trovons.

- `git stash`  
  → git stash est super utile quand on veut changer de branche sans perdre les modifs en cours

- `git rebase -i develop`  
  → Cella permet de nettoyer l’historique des commits et rendre le projet plus lisible

---

*Ce projet nous a permis de mieux comprendre le fonctionnement de Git en équipe, tout en développant un jeu classique de manière structurée. On a bien collaboré et chaque branche représentait une étape claire du développement.*
