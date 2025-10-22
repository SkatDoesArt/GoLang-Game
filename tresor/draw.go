package tresor

import (
	"image"
	"math"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/assets"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
)

// Draw() dessine le tresor sur la map
func (t Tresor) Draw(screen *ebiten.Image, CamX, CamY, CarX, CarY int) {

	// position du tresor
	xTileForDisplay := t.X - CamX + configuration.Global.ScreenCenterTileX
	yTileForDisplay := t.Y - CamY + configuration.Global.ScreenCenterTileY
	PosAbsX := xTileForDisplay * configuration.Global.TileSize
	PosAbsY := yTileForDisplay*configuration.Global.TileSize - configuration.Global.TileSize/2 + 8

	// create option
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(PosAbsX), float64(PosAbsY))

	// draw tresor
	screen.DrawImage(assets.TresorImage.SubImage(
		image.Rect(0, 0, configuration.Global.TileSize, configuration.Global.TileSize),
	).(*ebiten.Image), op)

	//draw distance blocs à l'aide de pythagore
	var dX int
	var dY int
	if CarX < 0 && CarY < 0 {
		dX = t.X - CarX
		dY = t.Y - CarY
	} else if CarX < 0 {
		dX = t.X - CarX
		dY = CarY - t.Y
	} else if CarY < 0 {
		dX = CarX - t.X
		dY = t.Y - CarY
	} else {
		dX = CarX - t.X
		dY = CarY - t.Y
	}

	var distance int = int(math.Sqrt(float64(dX)*float64(dX) + float64(dY)*float64(dY)))

	d := strconv.Itoa(distance)

	affichage := "Distance du trésor :" + d

	//affichage du message selon la postion du personnage
	if distance == 1 || distance == 0 {
		ebitenutil.DebugPrintAt(screen, "Victoire trésor trouvé !", 0, 0)
	} else {
		ebitenutil.DebugPrintAt(screen, affichage, 0, 0)
	}

}
