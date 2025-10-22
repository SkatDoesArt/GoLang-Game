package portal

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/assets"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
)

// Ajouter pour animation
var IndexFrame = 0   // indice of the actual frame
var CounterFrame = 0 // counter for the animation
//

// Draw() dessinne le ou les portails sur la map
func (p Portal) Draw(screen *ebiten.Image, CamX, CamY int) {
	// if activated then visible
	if p.Visible {
		// take position to add portal
		xTileForDisplay := p.XPortal - CamX + configuration.Global.ScreenCenterTileX
		yTileForDisplay := p.YPortal - CamY + configuration.Global.ScreenCenterTileY
		PosAbsX := xTileForDisplay * configuration.Global.TileSize
		PosAbsY := yTileForDisplay*configuration.Global.TileSize - configuration.Global.TileSize/2 + 8

		//ajouter pour animation
		CounterFrame++           // increment of the counter
		if CounterFrame >= 300 { // change every 10 iterations
			IndexFrame = (IndexFrame + 1) % 5 // 5 frame in total
			CounterFrame = 0
		}

		frameX := IndexFrame * 16 // calcul of the position of the actual frame

		// create option
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(PosAbsX), float64(PosAbsY))

		// draw portal
		screen.DrawImage(assets.PortalImage.SubImage(
			image.Rect(frameX, 0, frameX+configuration.Global.TileSize, configuration.Global.TileSize),
		).(*ebiten.Image), op)
	}
}
