package portal

// init permet de créer un portail lors de la pression de la touche t
func (p *Portal) Init(x, y int) {

	// si deja 2 dans liste
	if len(p.Portals) == 2 {
		// on supprime le premier
		p.Portals = p.Portals[1:]
	}

	// portal in new placement of the character
	newPortal := Portal{
		XPortal: x,
		YPortal: y,
		Visible: true,
	}

	// add new portal to list
	p.Portals = append(p.Portals, newPortal)

}
