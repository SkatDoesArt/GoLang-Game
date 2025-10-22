package game

import "gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"

// Init initialise les données d'un jeu. Il faut bien
// faire attention à l'ordre des initialisation car elles
// pourraient dépendre les unes des autres.
func (g *Game) Init() {
	g.floor.Init()
	g.character.Init(g.floor.GetContent().TrouvePosition())
	g.camera.Init(g.character.X, g.character.Y)
	//initialisation du tresor dans le cas de la chasse au tresor
	if configuration.Global.Chasseautresor {
		g.tresor.Init(g.floor.GetWidth(), g.floor.GetHeight())
	}
}
