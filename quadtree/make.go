package quadtree

// MakeFromArray construit un quadtree représentant un terrain
// étant donné un tableau représentant ce terrain.

func MakeFromArray(floorContent [][]int) (q Quadtree) {

	var w int = len(floorContent[0])
	var h int = len(floorContent)
	return Quadtree{width: w, height: h, root: MakeRoot(floorContent, w, h, 0, 0)}
}

//MakeRoot créer l'ensemble des noeuds de l'arbre de manière récursive
//entrée : un tableau floorcontent pour récupérer le contenue de chaque noeud (type de sol) , width et height la taille determinant la zone de floorcontent sur laquelle on travail,
//topleftX et topLeftY le point de départ de la zone
//sortie : un noeud racine qui permettra de créer un arbre avec make from array

func MakeRoot(floorContent [][]int, width, height, topLeftX, topLeftY int) (n *node) {
	if width == 1 && height == 1 {
		node := &node{topLeftX: topLeftX, topLeftY: topLeftY, width: width, height: height, content: floorContent[topLeftY][topLeftX], isLeaf: true}
		return node
	}

	if width == 0 || height == 0 {
		node := &node{topLeftX: topLeftX, topLeftY: topLeftY, width: width, height: height, content: 0, isLeaf: true}
		return node
	}

	var w2 int
	var h2 int

	if width > 1 {
		w2 = width / 2
	}
	if height > 1 {
		h2 = height / 2
	}
	topLeft := MakeRoot(floorContent, w2, h2, topLeftX, topLeftY)
	topRight := MakeRoot(floorContent, width-w2, h2, topLeftX+w2, topLeftY)
	botLeft := MakeRoot(floorContent, w2, height-h2, topLeftX, topLeftY+h2)
	botRight := MakeRoot(floorContent, width-w2, height-h2, topLeftX+w2, topLeftY+h2)
	var egaux bool = true
	var valcont int

	content := []*node{topLeft, topRight, botLeft, botRight}

	for _, elt := range content { //pour avoir egaux qui me dit si je peux réduire
		if elt.width != 0 && elt.height != 0 { // vérifie que les nodes existe
			valcont = elt.content
			for _, elt2 := range content {
				if elt2.width != 0 && elt2.height != 0 && elt.content != elt2.content {
					egaux = false
					break
				}
			}
		}
	}

	var nodes *node
	if topLeft.isLeaf && topRight.isLeaf && botLeft.isLeaf && botRight.isLeaf && egaux { //verification de la condition de réduction
		nodes = &node{topLeftX: topLeftX, topLeftY: topLeftY, width: width, height: height, content: valcont, isLeaf: true}
	} else {
		nodes = &node{topLeftX: topLeftX, topLeftY: topLeftY, width: width, height: height,
			topLeftNode:     topLeft,
			topRightNode:    topRight,
			bottomLeftNode:  botLeft,
			bottomRightNode: botRight}
	}
	return nodes
}
