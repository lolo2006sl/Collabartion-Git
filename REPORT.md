# Rapport de collaboration – Projet Puissance 4

## Objectif
Création du jeu Puissance 4 en terminal
---

## Conventions Git utilisées

### Nommage des branches
Convention suivie : `<type>/<fonctionnalité_ou_action>`  


Exemples :
- `feat/ft1` : initialisation du jeu (grille vide, joueurs, variables)
- `feat/ft3` : insertion des jetons dans la grille
- `feat/ft4` : détection du match nul et alternance des joueurs

### Messages de commit


Exemples :
- `feat: ajout de la vérification de victoire (FT3 bis)`
- `feat: ajout de fin en cas d'égalité (FT4)`
- `feat: FT1`

---

## Branches créées

| Branche               | Rôle                                      |
|-----------------------|-------------------------------------------|
| `main`                | Version stable du projet                  |
| `feat/ft1`            | Initialisation du jeu                     |
| `feat/ft3`            | Placement des jetons                      |
| `feat/ft4`            | Match nul et alternance des joueurs       |

---

## Commandes Git utilisées

### Création de branche
```bash
git switch -c feat/ft4
