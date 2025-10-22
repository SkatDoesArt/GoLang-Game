package tresor

import (
	"math/rand"

	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
)

// Init met en place un tresor
func (t *Tresor) Init(w, h int) {
	if configuration.Global.TerreRonde {
		t.X = rand.Intn(100) //distance max de 100 blocs en mode terre ronde
		t.Y = rand.Intn(100)
	} else {
		t.X = rand.Intn(w)
		t.Y = rand.Intn(h)
	}

}
