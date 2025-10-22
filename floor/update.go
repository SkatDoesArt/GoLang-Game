package floor

import (
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
)

// Update se charge de stocker dans la structure interne (un tableau)
// de f une représentation de la partie visible du terrain à partir
// des coordonnées absolues de la case sur laquelle se situe la
// caméra.
//
// On aurait pu se passer de cette fonction et tout faire dans Draw.
// Mais cela permet de découpler le calcul de l'affichage.
func (f *Floor) Update(camXPos, camYPos int) {
	topLeftX := camXPos - configuration.Global.ScreenCenterTileX
	topLeftY := camYPos - configuration.Global.ScreenCenterTileY
	switch configuration.Global.FloorKind {
	case GridFloor:
		f.updateGridFloor(topLeftX, topLeftY)
	case FromFileFloor:
		f.updateFromFileFloor(topLeftX, topLeftY)
	case QuadTreeFloor:
		f.updateQuadtreeFloor(topLeftX, topLeftY)
	case Random:
		f.updateQuadtreeFloor(topLeftX, topLeftY)
	}

	// partie des animations du sol

	if configuration.Global.BlocAnimated == true {

		animationStruct[0].frameAnimation++

		if animationStruct[0].frameAnimation >= 15 {

			if animationStruct[0].stepAnimation < 3 {
				animationStruct[0].stepAnimation++
			} else {
				animationStruct[0].stepAnimation = 0
			}
			animationStruct[0].frameAnimation = 0
		}
	}

	if configuration.Global.BlocAnimated == true {

		animationStruct[4].frameAnimation++

		if animationStruct[4].frameAnimation >= 15 {

			if animationStruct[4].stepAnimation < 3 {
				animationStruct[4].stepAnimation++
			} else {
				animationStruct[4].stepAnimation = 0
			}
			animationStruct[4].frameAnimation = 0
		}
	}

}

// le sol est un quadrillage de tuiles d'herbe et de tuiles de désert
func (f *Floor) updateGridFloor(topLeftX, topLeftY int) {
	for y := 0; y < len(f.content); y++ {
		for x := 0; x < len(f.content[y]); x++ {
			absX := topLeftX
			if absX < 0 {
				absX = -absX
			}
			absY := topLeftY
			if absY < 0 {
				absY = -absY
			}
			f.content[y][x] = ((x + absX%2) + (y + absY%2)) % 2
		}
	}
}

// le sol est récupéré depuis un tableau, qui a été lu dans un fichier
//
// la version actuelle recopie fullContent dans content, ce qui n'est pas
// le comportement attendu dans le rendu du projet
func (f *Floor) updateFromFileFloor(topLeftX, topLeftY int) {
	if configuration.Global.TerreRonde {
		for y := 0; y < len(f.content); y++ {
			for x := 0; x < len(f.content[y]); x++ {
				f.content[y][x] = f.fullContent[mod(topLeftY+y, len(f.fullContent))][mod(topLeftX+x, len(f.fullContent[0]))]
			}
		}
	} else {
		for y := 0; y < len(f.content); y++ {
			for x := 0; x < len(f.content[y]); x++ {
				if y < len(f.fullContent) && y+topLeftY < len(f.fullContent) && x+topLeftX < len(f.fullContent[y]) && y+topLeftY >= 0 && x+topLeftX >= 0 {
					f.content[y][x] = f.fullContent[topLeftY+y][topLeftX+x]
				} else {
					f.content[y][x] = -1
				}
			}
		}
	}
}

// le sol est récupéré depuis un quadtree, qui a été lu dans un fichier
func (f *Floor) updateQuadtreeFloor(topLeftX, topLeftY int) {
	f.quadtreeContent.GetContent(topLeftX, topLeftY, f.content)
}

func mod(a, b int) int {
	return (a%b + b) % b
}

func (f *Floor) Remove() {
	f.content = f.content[:len(f.content)-1]
	for i := 0; i < len(f.content); i++ {
		f.content[i] = append(f.content[i][:len(f.content[i])-1], f.content[i][len(f.content[i]):]...)
	}
}

func (f *Floor) Augment() {
	f.content = append(f.content, make([]int, len(f.content[0])))
	for i := 0; i < len(f.content); i++ {
		f.content[i] = append(f.content[i], 0)
	}
}
