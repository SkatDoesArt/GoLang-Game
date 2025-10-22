# GoLang-Game

Jeu 2D en Go basé sur Ebitengine (Ebiten) illustrant un rendu de terrain par quadtree, la lecture de cartes depuis des fichiers, une génération aléatoire de cartes, plusieurs modes de caméra et diverses extensions ludiques (portails, zoom, sauvegarde, chasse au trésor, etc.).

<img src=goGamePic.png alt="GoGamePic">


## Aperçu rapide

- Moteur: Ebitengine v2
- Rendu du sol: quadrillage simple, depuis fichier, via quadtree, ou génération aléatoire (quadtree)
- Déplacements: flèches du clavier, animation paramétrable
- Caméra: statique, suivie du personnage, ou limitée aux bordures
- Options: monde torique (terre ronde), blocage sur l’eau, tuiles animées, zoom, portails, sauvegarde de carte aléatoire, chasse au trésor
- Overlay debug: grille, coordonnées caméra/personnage, etc. (touche D)


## Prérequis

- Go 1.21+ (testé avec Go 1.23.2)
- Windows, macOS ou Linux. Les commandes ci-dessous sont données pour Windows PowerShell.


## Installation des dépendances

Depuis la racine du dépôt :

```powershell
go mod tidy
```


## Lancer le jeu

Plusieurs options selon votre dossier courant :

- Depuis la racine du projet (en ciblant l’appli dans `cmd/`) :

```powershell
go run ./cmd -config ./cmd/config.json
```

- Ou compilez puis exécutez :

```powershell
go build ./cmd
./cmd.exe -config ./cmd/config.json
```

Remarques :
- Le paramètre `-config` est optionnel. Sans précision, l’exécutable cherche `config.json` dans le dossier courant.
- Les chemins relatifs (par ex. vers une carte) sont résolus depuis votre dossier courant au moment de l’exécution.


## Commandes clavier

- Flèches directionnelles: déplacer le personnage
- D: activer/désactiver l’affichage debug
- Pavé numérique +: dézoomer (afficher plus de tuiles) si `Zoom` est true
- Pavé numérique -: zoomer (afficher moins de tuiles) si `Zoom` est true
- T: poser un portail (si `Portal` est true). Deux portails max; marcher sur l’un téléporte à l’autre
- S: sauvegarder la carte générée aléatoirement (si `FloorKind` = 3 et `Sauvegarde` est true)


## Configuration

Le fichier `cmd/config.json` contient tous les réglages. Exemple fourni :

```json
{
	"DebugMode": false,
	"NumTileX": 9,
	"NumTileY": 9,
	"TileSize": 16,
	"NumCharacterAnimImages": 5,
	"NumFramePerCharacterAnimImage": 5,
	"NumTileForDebug": 6,
	"CameraMode": 1,
	"FloorKind": 2,
	"FloorFile": "../floor-files/beaupasbeau",
	"RandomMapXSize": 15,
	"RandomMapYSize": 15,
	"TerreRonde": false,
	"PasMarcheEau": false,
	"BlocAnimated": false,
	"Zoom": false,
	"Portal": false,
	"Sauvegarde": false,
	"Chasseautresor": false
}
```

Description des champs principaux :

- DebugMode (bool): active l’overlay debug (peut aussi être togglé via D)
- NumTileX, NumTileY (int): nombre de tuiles visibles à l’écran (hors zone debug)
- TileSize (int): taille (px) d’une tuile
- NumCharacterAnimImages (int): nombre d’images par pas d’animation du personnage
- NumFramePerCharacterAnimImage (int): nombre d’updates entre deux frames d’anim
- NumTileForDebug (int): largeur en tuiles de la zone d’infos debug à droite/haut
- CameraMode (int): 0=statique, 1=suivi du personnage, 2=limité aux bords
	- En mode 2, si `TerreRonde` est true, la caméra suit le personnage (pas de blocage sur bords)
- FloorKind (int): 0=quadrillage, 1=depuis fichier, 2=quadtree (depuis fichier), 3=génération aléatoire (quadtree)
- FloorFile (string): chemin vers une carte (utilisé pour `FloorKind` 1 ou 2)
- RandomMapXSize, RandomMapYSize (int): dimensions de la carte aléatoire (pour `FloorKind` 3)
- TerreRonde (bool): monde torique (wrap horizontal/vertical) pour la lecture quadtree/fichier
- PasMarcheEau (bool): empêche la marche sur l’eau (tuile de type 4)
- BlocAnimated (bool): active l’animation de certaines tuiles du sol
- Zoom (bool): autorise les touches +/− du pavé numérique
- Portal (bool): autorise la pose et la téléportation via portails (touche T)
- Sauvegarde (bool): autorise la sauvegarde d’une carte aléatoire (touche S)
- Chasseautresor (bool): active le mode chasse au trésor

Notes :
- Types de tuile: les indices de tuiles du sol proviennent de `assets/floor.png`. La valeur 4 correspond à l’eau (utilisée par `PasMarcheEau`).
- En mode « chasse au trésor », un trésor aléatoire est placé ; la distance en blocs s’affiche. À distance ≤ 1, le message « Victoire trésor trouvé ! » apparaît.


## Structure du projet

```
assets/         # images embarquées (go:embed) et crédits
camera/         # modes de caméra (statique, suivi, limité aux bords)
character/      # logique d’animation et déplacement du personnage
cmd/            # point d’entrée (main), config par défaut
configuration/  # lecture/stockage de la configuration globale
floor/          # gestion du terrain (grid, fichier, quadtree, random) + dessin + collisions
floor-files/    # exemples de cartes (fichiers texte de chiffres)
game/           # boucle de jeu Ebiten (Init, Update, Draw, Layout)
portal/         # portails: création, update, dessin, téléportation
quadtree/       # structure quadtree + construction / requêtes + tests
tresor/         # mode chasse au trésor (placement, dessin, distance)
```


## Tests

Des tests couvrent la construction et les requêtes quadtree ainsi que la lecture de cartes.

Exécuter tous les tests :

```powershell
go test ./...
```


## Crédits assets

Voir `assets/licence` :
- character.png — CC-BY 3.0 — Lanea Zimmerman, Clint Bellanger, Charles Gabriel, basxto — https://opengameart.org/content/tiny-16-more-character-animations
- floor.png — CC-BY 4.0 — Ivan Voirol — https://opengameart.org/content/tinyslates-16x16px-orthogonal-tileset-by-ivan-voirol
- portail.png, tresor.png — inclus dans le dépôt (utilisés pour les fonctionnalités associées)


## Dépannage

- Fenêtre noire/écran partiel: vérifiez que `assets.Load()` est bien appelé (c’est le cas dans `cmd/main.go`) et que `TileSize`, `NumTileX/Y` donnent une fenêtre affichable.
- Carte introuvable: ajustez `FloorFile` selon votre dossier courant. Depuis la racine, utilisez un chemin comme `./floor-files/nom-de-carte`.
- Le personnage « bloque » au bord: basculez `CameraMode` sur 2 pour limiter la caméra aux bords, ou activez `TerreRonde` pour du wrap.
- Zoom inactif: mettez `Zoom` à true et utilisez le pavé numérique (+ pour dézoomer, − pour zoomer).
- Portails inactifs: mettez `Portal` à true. Deux portails max; le 1er est remplacé si un 3e est posé.
- Sauvegarde inactive: nécessite `FloorKind` = 3 (carte aléatoire) et `Sauvegarde` = true. Le fichier est écrit dans `floor-files/` avec un timestamp.


## Licence

Code : ce dépôt inclut un fichier `LICENSE`. Les assets ont leurs licences propres (voir section Crédits assets).
