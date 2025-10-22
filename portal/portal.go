package portal

type Portal struct {
	XPortal int      //position x du portail
	YPortal int      //position y du portail
	Visible bool     //indique la visibilité du portail sur la map
	Portals []Portal //liste des portails (ne dépasse jamais deux éléments)
}

var P = Portal{}
