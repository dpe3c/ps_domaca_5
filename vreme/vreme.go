// Package vreme help us with saving and displaying statistics for tempreature measurments
package vreme

import (
	"fmt"
)

// Postaja is a struct that saves its name, place and the measurments for it
type Postaja struct {
	ime     string
	kraj    string
	meritve []float64
}

// DodajMeritev is a func that adds a measurment to a Postaja
func DodajMeritev(postaje map[string]Postaja, oznaka string, temperatura float64, minTemp, maxTemp float64) {
	postaja, obstaja := postaje[oznaka]
	if !obstaja {
		fmt.Println("Postaja ne obstaja.")
		return
	}
	if temperatura < minTemp || temperatura > maxTemp {
		fmt.Println("Meritev ni realna. Funkcija je bila prekinjena.")
		return
	}
	postaja.meritve = append(postaja.meritve, temperatura)
	postaje[oznaka] = postaja
}

// IzpisMeritev displays all measurments for that Postaja
func IzpisMeritev(postaje map[string]Postaja) {
	fmt.Println("MERITVE:")
	for oznaka, postaja := range postaje {
		fmt.Printf("%s - %s (%s): %v\n", oznaka, postaja.ime, postaja.kraj, postaja.meritve)
	}
}

// IzpisPorocila displays min, max and avg for a whole map of Postaja's
func IzpisPorocila(postaje map[string]Postaja, min_st_Mer int) {
	fmt.Println("POROČILO:")
	for oznaka, postaja := range postaje {
		najnizja, najvisja, povprecje, ok := statistika(postaje, oznaka, min_st_Mer)

		if !ok {
			fmt.Printf("%s: premalo meritev\n", postaja.ime)
			continue
		}

		stanje := ""
		switch {
		case povprecje >= 25:
			stanje = "Vroče"
		case povprecje >= 15:
			stanje = "Toplo"
		case povprecje >= 5:
			stanje = "Hladno"
		default:
			stanje = "Mrzlo"
		}
		fmt.Printf("%s: min %.1f °C, max %.1f °C, povprečje %.1f °C -> %s\n", postaja.ime, najnizja, najvisja, povprecje, stanje)
	}
}

// statistika is a private func that calculates the min,max and avg
func statistika(postaje map[string]Postaja, oznaka string, min_st_Mer int) (najnizja, najvisja, povprecna float64, ok bool) {
	postaja, ok := postaje[oznaka]

	if !ok || len(postaja.meritve) < min_st_Mer {
		return 0, 0, 0, false
	}

	meritve := postaja.meritve
	najvisja = meritve[0]
	najnizja = meritve[0]
	var sestevek = 0.0

	for _, temp := range meritve {
		if temp < najnizja {
			najnizja = temp
		}
		if temp > najvisja {
			najvisja = temp
		}
		sestevek += temp
	}
	var povprecje = sestevek / float64(len(meritve))
	return najnizja, najvisja, povprecje, true
}

// DodajPostajo is a func that adds a new Postaja to a map of Postaja's
func DodajPostajo(postaje map[string]Postaja, oznaka, ime, kraj string) {
	postaje[oznaka] = Postaja{
		ime:     ime,
		kraj:    kraj,
		meritve: []float64{},
	}
}
