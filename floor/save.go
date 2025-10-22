package floor

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func (f *Floor) Save(x, y int) {

	//récupération de la carte
	var carte [][]int

	for i := 0; i < y; i++ {
		ligne := make([]int, 0)
		for j := 0; j < x; j++ {
			ligne = append(ligne, 0)
		}
		carte = append(carte, ligne)
	}

	f.quadtreeContent.GetContent(0, 0, carte)

	//sauvegarde de la carte
	var myFile *os.File
	var err error

	time := time.Now()
	titre := "../floor-files/randomfloor" + string(time.Format("2006-01-02 15h05m05s"))
	myFile, err = os.Create(titre)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(carte); i++ {
		var ligne string
		for j := 0; j < len(carte[i]); j++ {
			ligne += strconv.Itoa(carte[i][j])
		}
		_, err = fmt.Fprintln(myFile, ligne)
		if err != nil {
			log.Fatal(err)
		}
	}

	err = myFile.Close()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("sauvegarde terminée")
}
