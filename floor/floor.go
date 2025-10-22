package floor

import (
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/quadtree"
)

// Floor représente les données du terrain. Pour le moment
// aucun champs n'est exporté.
//
//   - content : partie du terrain qui doit être affichée à l'écran
//   - fullContent : totalité du terrain (utilisé seulement avec le type
//     d'affichage du terrain "fromFileFloor")
//   - quadTreeContent : totalité du terrain sous forme de quadtree (utilisé
//     avec le type d'affichage du terrain "quadtreeFloor")
type Floor struct {
	content         [][]int
	fullContent     [][]int
	quadtreeContent quadtree.Quadtree
}

// types d'affichage du terrain disponibles
const (
	GridFloor int = iota
	FromFileFloor
	QuadTreeFloor
	Random // ajout de la generation aléatoire du terrain
)

// GetHeight retourne la hauteur (en cases) du terrain
// à partir du tableau fullContent, en supposant que
// ce tableau représente un terrain rectangulaire
func (f Floor) GetHeight() (height int) {
	switch configuration.Global.FloorKind {
	case QuadTreeFloor, Random: //ajout de la recuperation de la taille de la map pour les quadtree
		return f.quadtreeContent.GetHeight()
	case GridFloor, FromFileFloor:

		return len(f.fullContent)
	}

	return

}

// GetWidth retourne la largeur (en cases) du terrain
// à partir du tableau fullContent, en supposant que
// ce tableau représente un terrain rectangulaire
func (f Floor) GetWidth() (width int) {

	switch configuration.Global.FloorKind {
	case QuadTreeFloor, Random: //ajout de la recuperation de la taille de la map pour les quadtree
		return f.quadtreeContent.GetWidth()
	case GridFloor, FromFileFloor:
		if len(f.fullContent) > 0 {
			width = len(f.fullContent[0])
		}

	}
	return
}
