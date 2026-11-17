# Vremenske postaje

Projekt je izdelan v programskem jeziku Go v okviru 5. domače naloge pri predmetu Porazdeljeni sistemi na Fakulteti za računalništvo in informatiko Univerze v Ljubljani.

## Opis

Program omogoča upravljanje vremenskih postaj, dodajanje temperaturnih meritev ter izračun najnižje, najvišje in povprečne temperature.

Projekt je razdeljen na dva paketa:
- `vreme` – vsebuje strukturo `Postaja` in funkcije za upravljanje vremenskih postaj.
- `main` – uporablja paket `vreme` in omogoča konfiguracijo programa prek ukazne vrstice.

## Zagon programa

Za zagon uporabimo:

`go run .`

Program podpira naslednja stikala:

| Stikalo | Opis | Privzeta vrednost |
|---|---|---|
| `--minTemp` | Najnižja dovoljena temperatura | -50 °C |
| `--maxTemp` | Najvišja dovoljena temperatura | 50 °C |
| `--stMeritev` | Najmanjše število meritev za statistiko | 3 |

Primer zagona s spremenjenimi nastavitvami:

`go run . --minTemp -20 --maxTemp 40 --stMeritev 2`

Za pregled vseh možnosti:

`go run . --help`

## Funkcionalnosti

- Dodajanje vremenskih postaj in temperaturnih meritev.
- Preverjanje veljavnosti temperatur.
- Izpis vseh meritev.
- Izračun statističnih podatkov.
- Razvrščanje temperatur v kategorije Mrzlo, Hladno, Toplo in Vroče.

## Dokumentacija

Dokumentacija paketa `vreme` je dostopna z ukazom:

`go doc ./vreme`