# Projet-Red-Mario#🍄 

Un **RPG en ligne de commande (CLI)** développé en **Go**, dans le cadre du **Projet RED** de l'Ymmer sion Ynov.

Incarnez **Mario**, **Luigi**, **Princesse Peach** ou créez votre propre héros, puis partez à l'aventure dans le Royaume Champignon : combattez des monstres, forgez votre équipement, apprenez des sorts et triomphez de **Bowser** !

---

## 🎮 Présentation

**Mario RPG Adventure** est un jeu de rôle **tour par tour** qui mobilise l'ensemble des compétences vues durant l'Ymmer sion :

- Création et gestion d'un personnage (nom, classe, statistiques, équipement)
- Système d'inventaire avec capacité limitée et améliorable
- Système d'économie (pièces d'or) et de **fabrication** d'équipement
- Système de **combat tour par tour** avec initiative, sorts, mana et expérience
- Interface CLI colorée avec **ASCII art** et **animations** en combat
- Ambiance sonore (musique de fond, musique de combat, effets spéciaux)

### 🎭 Thème choisi

Le thème est basé sur **l'univers de Mario / Nintendo** :

| 🎭 Héros | 💀 Ennemis |
|---|---|
| 🍄 Mario — Équilibré | 🟢 Goomba d'entraînement |
| 🟢 Luigi — Mage agile | 🔵 Thwomp |
| 👑 Princesse Peach — Soutien | 🐢 **Bowser** (BOSS final) |
| ✨ Personnage personnalisé (Humain / Toad / Koopa) | |

---

## ⚔️ Fonctionnalités

### 🧙 Création de personnage
- Choix parmi **3 héros prédéfinis** (Mario, Luigi, Peach) avec statistiques propres
- OU **création personnalisée** : nom + classe (Humain, Toad, Koopa)
- **ASCII art** propre à chaque personnage affiché lors de la sélection
- Affichage des **PV max et Mana max** de chaque classe pour aider au choix

### 🎒 Inventaire
- Capacité limitée à **10 objets** au départ
- Améliorable **jusqu'à 3 fois** (+10 places à chaque fois)
- **Agrandissement automatique** à chaque montée de niveau (+2 places)
- Utilisation d'objets hors combat et **en combat**

### 🛒 Marchand
- Achat d'objets avec des pièces d'or
- Catalogue : **Champignon Super**, **Champignon Poison**, **Potion de Mana**, **Fleur de Feu**, **matériaux de craft** (Fourrure de Loup, Peau de Troll, Cuir de Sanglier, Plume de Corbeau), **Augmentation d'inventaire**
- 🎁 **Cadeau de bienvenue** : la première visite offre un **Champignon Super gratuit**
- **Limite** de 3 achats d'augmentation d'inventaire

### ⚒ Forgeron
- Fabrication d'équipement à partir de **matériaux**
- **Casquette Mario** (+10 PV max)
- **Salopette Mario** (+25 PV max)
- **Bottes Kuribo** (+15 PV max)
- Coût : **5 pièces d'or** + matériaux requis

### 🛡 Équipement
- 3 emplacements : **Tête**, **Torse**, **Pieds**
- Bonus de PV max selon l'équipement porté
- Remplacement automatique : l'ancien équipement retourne dans l'inventaire

### ⚔️ Combat tour par tour
- **Système d'initiative** : le plus rapide commence le combat
- **Attaque physique** ou **lancement de sorts**
- **Système de mana** : les sorts consomment du mana
- **Système d'expérience** : gagner de l'XP en tuant les monstres, monter de niveau et augmenter ses stats
- **Patterns d'attaque** : les monstres frappent x2 tous les 3 tours
- **Animations ASCII** : attaquant qui avance, flash de dégâts, barres de vie colorées

### 🎁 Missions bonus réalisées
- ✅ **Mission 1** — Système d'initiative
- ✅ **Mission 2** — Système d'expérience et niveaux
- ✅ **Mission 3** — Combat magique (Boule de Feu, Étoile)
- ✅ **Mission 4** — Ressource de mana
- ✅ **Mission 5** — Enrichissement du contenu (nouveaux boss, sons, animations)
- ✅ **Mission 6** — Qui sont-ils ? (easter egg 🎵 ABBA + 🎬 Spielberg)

### 🔊 Sons et musique
- Musique de fond du menu
- Musique de combat
- Effets sonores : mort, victoire, pièce, marteau, sortie
- Bips de secours si les fichiers audio sont absents

---

## 🛠️ Installation

### Prérequis

- **[Go](https://go.dev/dl/)** version 1.20 ou supérieure
- Un **terminal** (Terminal macOS, iTerm2, Windows Terminal...)
- (Optionnel) **afplay** (macOS), **mpg123** ou **aplay** (Linux), **PowerShell** (Windows) pour le son

### Cloner le dépôt

```bash
git clone https://github.com/s9bwpvyvj4-ship-it/Projet-red.git
cd Projet-red
```

### Installer les dépendances

```bash
go mod tidy
```

---

## 🚀 Lancer le jeu

Depuis la racine du projet :

```bash
go run .
```

Ou compiler un exécutable :

```bash
go build -o mario-rpg
./mario-rpg
```

---

## 🎮 Comment jouer

1. **Choisis ton personnage** : Mario, Luigi, Peach ou crée le tien
2. **Navigue dans le menu principal** :

| Choix | Action |
|---|---|
| 1 | 👤 Afficher les informations du personnage |
| 2 | 🎒 Accéder à l'inventaire |
| 3 | 🛒 Marchand |
| 4 | ⚒ Forgeron |
| 5 | 🟢 Entraînement (Goomba) |
| 6 | 🔵 Combat contre Thwomp |
| 7 | 🐢 ★ BOSS : Bowser ★ |
| 8 | ❓ Qui sont-ils ? (easter egg) |
| 0 | 🚪 Quitter |

3. **Combats** : attaque, lance des sorts ou utilise des objets pour triompher
4. **Gagne de l'XP** et des pièces pour progresser
5. **Forge ton équipement** pour devenir plus fort
6. **Bats Bowser** pour terminer le jeu 🐢

---

## 📁 Structure du projet

```
Projet-red/
├── README.md              # Ce fichier
├── go.mod                 # Module Go
├── sounds/                # Fichiers audio (optionnel)
│   ├── theme.mp3
│   ├── battle.mp3
│   ├── victory.mp3
│   ├── mario-death.mp3
│   ├── coin.mp3
│   ├── hammer.mp3
│   └── quit.mp3
└── src/                   # (ou racine) Code source
    ├── main.go            # Point d'entrée + menu principal + WhoAreThey
    ├── character.go       # Structure Character + création + progression
    ├── combat.go          # Combat tour par tour + animations
    ├── monster.go         # Monstres + patterns + art
    ├── inventory.go       # Gestion d'inventaire + potions + équipement
    ├── merchant.go        # Marchand + catalogue + achat
    ├── blacksmith.go      # Forgeron + recettes + fabrication
    ├── sound.go           # Système audio
    └── ui.go              # Couleurs, ASCII art, helpers, menus
```

---

## 🎨 Aperçu visuel

Le jeu utilise :
- Des **couleurs ANSI** (rouge, vert, jaune, bleu, etc.)
- Des **cadres ASCII** (`╔═╗`, `║`, `╚═╝`)
- Des **art ASCII** pour chaque personnage et monstre
- Des **barres de vie et de mana** visuelles `[████████░░░░░░]`

Exemple d'interface de combat :

```
╔══════════════════════════════════════════════════════════╗
║                  ⚔  TOUR 3 ⚔                             ║
╚══════════════════════════════════════════════════════════╝

                    ╔════════════════╗
                    ║  Bowser        ║
                    ╚════════════════╝

       Bowser  [████████████░░░░░░░░]  90/150

──────────────────────────────────────────────────────────

       Mario   [██████████████░░░░░░]  50/100 PV
       Mana    [████████░░░░░░░░░░░░]  20/50
       Pièces : 100   Exp : 40/100
```

---

## 👥 Équipe

- **[Sherine]** — Développeur
- **[Alicia]** — Développeur
- **[Antonin]** — Développeur

Projet créé dans le cadre du **Projet RED** — Ynov Campus
Sujet original créé par **Alan PHILIPIERT**

---

## 📚 Documentation

- Sujet complet du Projet RED : voir `docs/`
- Document de gestion de projet : voir `docs/`

---

## 🏆 Crédits

- **Sujet du projet** : Alan PHILIPIERT
- **ASCII art** : généré et adapté par l'équipe
- **Thème** : Mario / Nintendo (utilisation à but pédagogique)
- **Sons** : effets sonores libres de droits inspirés de l'univers Mario

---

## 📝 Licence

Projet pédagogique — Ynov Campus 2024-2025.
