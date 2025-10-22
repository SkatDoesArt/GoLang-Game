package portal

//update() met a jour la position du portail lors de la création de celui ci

func (p *Portal) Update(caracterex, caracterey int) {
	p.XPortal = caracterex
	p.YPortal = caracterey
}
