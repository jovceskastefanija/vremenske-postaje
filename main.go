package main

import (
	"context"
	"fmt"
	"github.com/jovceskastefanija/vremenske-postaje/vreme"
	"github.com/urfave/cli/v3"
	"log"
	"os"
)

func izvediProgram(minTemp, maxTemp float64, stMeritev int) {
	postaje := make(map[string]vreme.Postaja)

	oznake := []string{"LJ01", "LJ02", "MB01", "MB02", "NM01", "NM02"}
	imena := []string{"Ljubljana1", "Ljubljana2", "Maribor1", "Maribor2", "NovoMesto1", "NovoMesto2"}
	kraje := []string{"Ljubljana", "Ljubljana", "Maribor", "Maribor", "NovoMesto", "NovoMesto"}
	meritve := [][]float64{{12.5, 13.7, 15.01, 13.4, 14.0},
		{20.5, 21.7, 22.01, 23.4, 24.0},
		{10, 10.7, 11.01, 13.4, 14.0},
		{30.5, 33.1, 34.01, 33.4, 31.0},
		{12.05, 17.7, 18.01, 13.4, 19.0},
		{40, 41},
	}

	for i, oznaka := range oznake {
		vreme.DodajPostajo(postaje, oznaka, imena[i], kraje[i])
		for _, meritev := range meritve[i] {
			vreme.DodajMeritev(postaje, oznaka, meritev, minTemp, maxTemp)
		}
	}

	vreme.IzpisMeritev(postaje)
	vreme.IzpisPorocila(postaje, stMeritev)
}

func main() {
	cmd := &cli.Command{
		Name:  "vremenske-postaje",
		Usage: "Program za obdelavo meritev vremenskih postaj",

		Flags: []cli.Flag{
			&cli.Float64Flag{
				Name:  "minTemp",
				Usage: "Najnižja dovoljena temperatura",
				Value: -50,
			},
			&cli.Float64Flag{
				Name:  "maxTemp",
				Usage: "Najvišja dovoljena temperatura",
				Value: 50,
			},
			&cli.IntFlag{
				Name:  "stMeritev",
				Usage: "Najmanjše število meritev za statistiko",
				Value: 3,
			},
		},

		Action: func(ctx context.Context, cmd *cli.Command) error {
			minTemp := cmd.Float64("minTemp")
			maxTemp := cmd.Float64("maxTemp")
			stMeritev := cmd.Int("stMeritev")

			if minTemp > maxTemp {
				return fmt.Errorf("minTemp ne sme biti večja od maxTemp")
			}

			if stMeritev < 1 {
				return fmt.Errorf("stMeritev mora biti vsaj 1")
			}

			izvediProgram(minTemp, maxTemp, stMeritev)
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
