package floor

import (
	"bufio"
	"log"
	"os"
	"strconv"

	"math/rand"

	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/quadtree"
)

// Init initialise les structures de données internes de f.
func (f *Floor) Init() {
	f.content = make([][]int, configuration.Global.NumTileY)
	for y := 0; y < len(f.content); y++ {
		f.content[y] = make([]int, configuration.Global.NumTileX)
	}

	switch configuration.Global.FloorKind {
	case FromFileFloor:
		f.fullContent = readFloorFromFile(configuration.Global.FloorFile)
	case QuadTreeFloor:
		f.quadtreeContent = quadtree.MakeFromArray(readFloorFromFile(configuration.Global.FloorFile))
	case Random:
		f.quadtreeContent = quadtree.MakeFromArray(RandomMapGen(configuration.Global.RandomMapXSize, configuration.Global.RandomMapYSize))
	}
}

// lecture du contenu d'un fichier représentant un terrain
// pour le stocker dans un tableau
func readFloorFromFile(fileName string) (floorContent [][]int) {
	// TODO lire fichier, ligne toutes égales, parcourir et rajouter chaque ligne en sous liste de la liste floorContent
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal("erreur ouverture du fichier!")
	}

	fileScanner := bufio.NewScanner(file)
	var l int = -1
	for fileScanner.Scan() {
		if l == -1 {
			l = len(fileScanner.Text())
		} else if len(fileScanner.Text()) != l {
			return nil
		}

		tabtemp := make([]int, 0)

		for i := 0; i < len(fileScanner.Text()); i++ {
			ligne := fileScanner.Text()
			elt, err := strconv.Atoi(string(ligne[i]))
			if err != nil {
				log.Fatal()
			}
			tabtemp = append(tabtemp, elt)
		}
		floorContent = append(floorContent, tabtemp)
	}
	return
}

//création de la fonction permettant la génération aléatoire de la map

func RandomMapGen(x, y int) (carte [][]int) {

	for i := 0; i < y; i++ {
		ligne := make([]int, 0)
		for j := 0; j < x; j++ {
			ligne = append(ligne, rand.Intn(5))
		}
		carte = append(carte, ligne)
	}

	return carte
}

func (f Floor) GetContent() (quad quadtree.Quadtree) {
	return f.quadtreeContent
}
