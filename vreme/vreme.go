// Package vreme omogoča upravljanje vremenskih postaj in njihovih
// temperaturnih meritev.
//
// Vsaka vremenska postaja je predstavljena s strukturo Postaja,
// ki vsebuje ime, kraj in seznam temperaturnih meritev.
// Postaje so shranjene v slovarju, kjer je ključ enolična oznaka postaje.
//
// Paket omogoča dodajanje postaj in meritev, preverjanje veljavnosti
// temperatur ter izpis zbranih podatkov in statističnih poročil.
// Temperaturne meritve se sprejmejo le, če so znotraj podanega
// dovoljenega območja.
//
// Statistično poročilo za posamezno postajo vsebuje najnižjo,
// najvišjo in povprečno temperaturo ter temperaturno kategorijo.
// Izračun statistike je mogoč le, če ima postaja zadostno
// število veljavnih meritev.
//
// Paket omogoča nastavitev dovoljenega temperaturnega območja
// in najmanjšega števila meritev prek argumentov funkcij.
package vreme

import "fmt"

// Postaja predstavlja vremensko postajo z imenom, krajem
// in seznamom temperaturnih meritev.
type Postaja struct {
	ime     string
	kraj    string
	meritve []float64
}

// statistika izračuna najnižjo, najvišjo in povprečno temperaturo
// postaje. Vrne ok=false, če postaja ne obstaja ali nima
// dovolj meritev za izračun.
// stMeritev mora biti vsaj 1.
func statistika(postaje map[string]Postaja, oznaka string, stMeritev int) (najnizja, najvisja, povprecna float64, ok bool) {
	_, je := postaje[oznaka]

	x := len(postaje[oznaka].meritve)
	if !je || x < stMeritev {
		return
	}

	ok = true
	najnizja = postaje[oznaka].meritve[0]
	najvisja = postaje[oznaka].meritve[0]
	for _, v := range postaje[oznaka].meritve {
		povprecna += v
		if v < najnizja {
			najnizja = v
		}

		if v > najvisja {
			najvisja = v
		}

	}

	n := float64(x)
	povprecna = povprecna / n

	return najnizja, najvisja, povprecna, ok
}

// IzpisMeritev izpiše oznake, imena, kraje in vse shranjene
// temperaturne meritve vremenskih postaj.
func IzpisMeritev(postaje map[string]Postaja) {
	fmt.Println("MERITVE:")
	for k, p := range postaje {
		fmt.Printf("%s - %s (%s): %v\n", k, p.ime, p.kraj, p.meritve)
	}

}

// IzpisPorocila izpiše statistiko in temperaturno kategorijo
// posamezne postaje. Za izračun je potrebnih vsaj
// stMeritev meritev.
func IzpisPorocila(postaje map[string]Postaja, stMeritev int) {
	for oznaka, p := range postaje {
		najnizja, najvisja, povprecna, ok := statistika(postaje, oznaka, stMeritev)

		if !ok {
			fmt.Printf("%s: premalo meritev\n", oznaka)
			fmt.Println()
			continue
		}

		var s string
		switch {
		case povprecna >= 25:
			s = "Vroče"
		case povprecna >= 15:
			s = "Toplo"
		case povprecna >= 5 && povprecna < 15:
			s = "Hladno"
		default:
			s = "Mrzlo"
		}

		fmt.Printf("%s: min %.2f °C, max %.2f °C, povprečje %.2f °C -> %s\n",
			p.ime, najnizja, najvisja, povprecna, s)
		fmt.Println()
	}
}

// DodajMeritev doda temperaturo postaji z dano oznako.
// Meritev zavrne, če postaja ne obstaja ali temperatura
// ni v dovoljenem območju [minTemp, maxTemp].
func DodajMeritev(postaje map[string]Postaja, oznaka string, temperatura float64, minTemp, maxTemp float64) {
	if temperatura < minTemp || temperatura > maxTemp {
		fmt.Printf("Temperature ja izven dovoljenega območja, postaja: %s\n", oznaka)
		return
	}

	p, ok := postaje[oznaka]

	if !ok {
		fmt.Println("Postaja ne obstaja:", oznaka)
		return
	}

	p.meritve = append(p.meritve, temperatura)
	postaje[oznaka] = p

}

// DodajPostajo doda novo postajo z dano oznako, imenom
// in krajem. Če oznaka že obstaja, obstoječo postajo prepiše.
func DodajPostajo(postaje map[string]Postaja, oznaka string, ime string, kraj string) {
	postaja1 := Postaja{ime, kraj, []float64{}}
	postaje[oznaka] = postaja1

}
