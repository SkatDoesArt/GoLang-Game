package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
)

// Update met à jour les données du jeu à chaque 1/60 de seconde.
// Il faut bien faire attention à l'ordre des mises-à-jour car elles
// dépendent les unes des autres (par exemple, pour le moment, la
// mise-à-jour de la caméra dépend de celle du personnage et la définition
// du terrain dépend de celle de la caméra).
func (g *Game) Update() error {

	if inpututil.IsKeyJustPressed(ebiten.KeyD) {
		configuration.Global.DebugMode = !configuration.Global.DebugMode
	}

	//ajout pour le zoom
	if inpututil.IsKeyJustPressed(ebiten.KeyNumpadAdd) && configuration.Global.Zoom {
		configuration.Global.NumTileX++
		configuration.Global.NumTileY++
		g.floor.Augment()
		configuration.Global.ScreenWidth = configuration.Global.NumTileX * configuration.Global.TileSize
		configuration.Global.ScreenHeight = configuration.Global.NumTileY * configuration.Global.TileSize
		configuration.Global.ScreenCenterTileX = configuration.Global.NumTileX / 2
		configuration.Global.ScreenCenterTileY = configuration.Global.NumTileY / 2
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyNumpadSubtract) && configuration.Global.Zoom && configuration.Global.NumTileX > 3 && configuration.Global.NumTileY > 3 {
		configuration.Global.NumTileX--
		configuration.Global.NumTileY--
		g.floor.Remove()
		configuration.Global.ScreenWidth = configuration.Global.NumTileX * configuration.Global.TileSize
		configuration.Global.ScreenHeight = configuration.Global.NumTileY * configuration.Global.TileSize
		configuration.Global.ScreenCenterTileX = configuration.Global.NumTileX / 2
		configuration.Global.ScreenCenterTileY = configuration.Global.NumTileY / 2
	}

	g.character.Update(g.floor.Blocking(g.character.X, g.character.Y, g.camera.X, g.camera.Y), g.portal.Portals)
	g.camera.Update(g.character.X, g.character.Y, g.floor.GetHeight(), g.floor.GetWidth())
	g.floor.Update(g.camera.X, g.camera.Y)

	if configuration.Global.Portal {
		// ajouté pour portail ( premiere partie ) (touche 't' pour créer un portail)
		if inpututil.IsKeyJustPressed(ebiten.KeyT) {
			g.portal.Init(g.character.X, g.character.Y)
		}

		// mis a jour portail
		for i := range g.portal.Portals {
			g.portal.Portals[i].Update(g.portal.Portals[i].XPortal, g.portal.Portals[i].YPortal)
		}

	}

	//mise en place de la sauvegarde de map (touche 's' pour sauvegarder)

	if configuration.Global.FloorKind == 3 {
		if configuration.Global.Sauvegarde {
			if inpututil.IsKeyJustPressed(ebiten.KeyS) {
				g.floor.Save(configuration.Global.RandomMapXSize, configuration.Global.RandomMapYSize)
			}
		}
	}

	return nil
}
