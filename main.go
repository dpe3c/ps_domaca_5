package main

import (
	"github.com/dpe3c/ps_domaca_5/vreme"
)

func main() {
	var postaje = map[string]vreme.Postaja{}
	vreme.DodajPostajo(postaje, "LJ12", "Ajdovscina", "Ljubljana")
	vreme.DodajPostajo(postaje, "LJ13", "Ajdovscina", "Ljubljana")
	vreme.DodajPostajo(postaje, "LJ14", "Ajdovscina", "Ljubljana")

}
