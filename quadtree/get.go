package quadtree

import (
	"fmt"
	"log"

	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
)

// GetContent remplit le tableau contentHolder (qui représente
// un terrain dont la case le plus en haut à gauche a pour coordonnées
// (topLeftX, topLeftY)) à partir du qadtree q.
func (q Quadtree) GetContent(topLeftX, topLeftY int, contentHolder [][]int) {
	if configuration.Global.TerreRonde {
		for i := 0; i < len(contentHolder); i++ {
			for j := 0; j < len(contentHolder[0]); j++ {
				var cont int
				cont = parcours_arbre(mod(topLeftX+j, q.width), mod(topLeftY+i, q.height), q.root)
				contentHolder[i][j] = cont
			}
		}
	} else {
		for i := 0; i < len(contentHolder); i++ {
			for j := 0; j < len(contentHolder[0]); j++ {

				if topLeftX+j >= q.width || topLeftX+j < 0 || topLeftY+i < 0 || topLeftY+i >= q.height {

					contentHolder[i][j] = -1

				} else {
					cont := parcours_arbre(topLeftX+j, topLeftY+i, q.root)

					contentHolder[i][j] = cont
				}

			}
		}
	}
}

func parcours_arbre(x, y int, n *node) int {

	if (n.topLeftX == x && n.topLeftY == y && n.isLeaf) || (x >= n.topLeftX && x < n.topLeftX+n.width && y >= n.topLeftY && y < n.topLeftY+n.height && n.isLeaf && n.width >= 1 && n.height >= 1) {
		return n.content
	}

	if !(x >= n.topLeftX && x < n.topLeftX+n.width && y >= n.topLeftY && y < n.topLeftY+n.height) {

		return 0
	}

	if n.isLeaf {
		return 0
	}
	if x >= n.topRightNode.topLeftX {
		if y >= n.bottomRightNode.topLeftY {
			return parcours_arbre(x, y, n.bottomRightNode)
		} else {
			return parcours_arbre(x, y, n.topRightNode)
		}
	} else {
		if y >= n.bottomLeftNode.topLeftY {
			return parcours_arbre(x, y, n.bottomLeftNode)
		} else {
			return parcours_arbre(x, y, n.topLeftNode)
		}
	}

}

// fonction modulo car le modulo de golang ne marche pas comme je le souhaiterais avec les négatifs
func mod(a, b int) int {
	return (a%b + b) % b
}

// Donne les coordonnées où peut démarrer le personnage
func (q Quadtree) TrouvePosition() (x, y int) {
	x = q.width / 2
	y = q.height / 2
	if configuration.Global.PasMarcheEau {
		if q.root.isLeaf && q.root.content == 4 { //cas où la map est faite seulement d'eau
			log.Fatal("Pas de point de spawn disponible (case autre que eau)")
		} else {
			var Passe = make(map[string]bool) // Passe permet de garder en mémoire les cases que l'on a regardé
			x, y = q.TrouvePositionPasEau(x, y, Passe)
		}
	}
	return
}

// TrouvePositionPasEau() permet d'éviter le spawn du personnage sur une tuile d'eau dans le cas ou l'eau est bloquante
func (q Quadtree) TrouvePositionPasEau(x, y int, Passe map[string]bool) (newX, newY int) {
	var key string = fmt.Sprint(x) + "," + fmt.Sprint(y)
	Passe[key] = true

	var contenu int = parcours_arbre(x, y, q.root)
	if contenu != 4 { //si on a une trouvé une case sans eau on retourne ses coordonnées
		return x, y
	} else {
		//parcours les cases sans répétition grace à Passe
		var key2 string
		key2 = fmt.Sprint(x) + "," + fmt.Sprint(y+1)
		if y+1 < q.height && !(Passe[key2]) {
			return q.TrouvePositionPasEau(x, y+1, Passe)
		}
		key2 = fmt.Sprint(x+1) + "," + fmt.Sprint(y)
		if x+1 < q.width && !(Passe[key2]) {
			return q.TrouvePositionPasEau(x+1, y, Passe)
		}
		key2 = fmt.Sprint(x) + "," + fmt.Sprint(y-1)
		if y-1 >= 0 && !(Passe[key2]) {
			return q.TrouvePositionPasEau(x, y-1, Passe)
		}
		key2 = fmt.Sprint(x-1) + "," + fmt.Sprint(y)
		if x-1 >= 0 && !(Passe[key2]) {
			return q.TrouvePositionPasEau(x-1, y, Passe)
		}
	}
	log.Fatal("y'a un problème, n'a pas trouvé de point de spawn valide")
	return

}
