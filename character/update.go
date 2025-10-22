package character

import (
	"github.com/hajimehoshi/ebiten/v2"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/portal"
)

// Update met à jour la position du personnage, son orientation
// et son étape d'animation (si nécessaire) à chaque pas
// de temps, c'est-à-dire tous les 1/60 secondes.
func (c *Character) Update(blocking [4]bool, portails []portal.Portal) {
	if !c.moving {
		if ebiten.IsKeyPressed(ebiten.KeyRight) {
			c.orientation = orientedRight
			if !blocking[1] {
				c.xInc = 1
				c.moving = true
			}
		} else if ebiten.IsKeyPressed(ebiten.KeyLeft) {
			c.orientation = orientedLeft
			if !blocking[3] {
				c.xInc = -1
				c.moving = true
			}
		} else if ebiten.IsKeyPressed(ebiten.KeyUp) {
			c.orientation = orientedUp
			if !blocking[0] {
				c.yInc = -1
				c.moving = true
			}
		} else if ebiten.IsKeyPressed(ebiten.KeyDown) {
			c.orientation = orientedDown
			if !blocking[2] {
				c.yInc = 1
				c.moving = true
			}
		}
	} else {
		c.animationFrameCount++
		if c.animationFrameCount >= configuration.Global.NumFramePerCharacterAnimImage {
			c.animationFrameCount = 0
			shiftStep := configuration.Global.TileSize / configuration.Global.NumCharacterAnimImages
			c.shift += shiftStep
			c.animationStep = -c.animationStep
			if c.shift > configuration.Global.TileSize-shiftStep {
				c.shift = 0
				c.moving = false
				c.X += c.xInc
				c.Y += c.yInc
				c.xInc = 0
				c.yInc = 0
			}
		}
		if configuration.Global.Portal && !c.moving {
			if len(portails) == 2 {
				if c.X == portails[0].XPortal && c.Y == portails[0].YPortal {
					c.X = portails[1].XPortal
					c.Y = portails[1].YPortal
				} else if c.X == portails[1].XPortal && c.Y == portails[1].YPortal {
					c.X = portails[0].XPortal
					c.Y = portails[0].YPortal
				}
			}
		}
	}

}
