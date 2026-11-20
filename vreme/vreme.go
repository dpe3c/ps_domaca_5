package vreme

import (
	"fmt"
)

type Postaja struct {
	ime     string
	kraj    string
	meritve []float64
}

func DodajMeritev(postaje map[string]Postaja, oznaka string, temperatura float64) {
	postaja, obstaja := postaje[oznaka]
	if !obstaja {
		fmt.Println("Postaja ne obstaja.")
		return
	}
	if temperatura < -50 || temperatura > 50 {
		fmt.Println("Meritev ni realna. Funkcija je bila prekinjena.")
		return
	}
	postaja.meritve = append(postaja.meritve, temperatura)
	postaje[oznaka] = postaja
}

func IzpisMeritev(postaje map[string]Postaja) {
	fmt.Println("MERITVE:")
	for oznaka, postaja := range postaje {
		fmt.Printf("%s - %s (%s): %v\n", oznaka, postaja.ime, postaja.kraj, postaja.meritve)
	}
}

func IzpisPorocila(postaje map[string]Postaja) {
	fmt.Println("POROČILO:")
	for oznaka, postaja := range postaje {
		najnizja, najvisja, povprecje, ok := statistika(postaje, oznaka)

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

func statistika(postaje map[string]Postaja, oznaka string) (najnizja, najvisja, povprecna float64, ok bool) {
	postaja, ok := postaje[oznaka]

	if !ok || len(postaja.meritve) < 3 {
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

func DodajPostajo(postaje map[string]Postaja, oznaka, ime, kraj string) {
	postaje[oznaka] = Postaja{
		ime:     ime,
		kraj:    kraj,
		meritve: []float64{},
	}
}
