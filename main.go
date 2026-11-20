package main

import (
	"context"
	"log"
	"os"

	"github.com/dpe3c/ps_domaca_5/vreme"
	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "Postaje",
		Usage: "Postaje nam povejo min,max in povprečno temperaturo za vse postaje ki jih beležimo.",
		Flags: []cli.Flag{
			&cli.FloatFlag{
				Name:  "minTemp",
				Usage: "Najnižja dovoljena temperatura",
				Value: -50.0,
			},
			&cli.FloatFlag{
				Name:  "maxTemp",
				Usage: "Najvišja dovoljena temperatura",
				Value: 50.0,
			},
			&cli.IntFlag{
				Name:  "stMeritev",
				Usage: "Najnižja dovoljena količina meritev na postaji za izpis",
				Value: 3,
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			minTemp := cmd.Float("minTemp")
			maxTemp := cmd.Float("maxTemp")
			min_st_Mer := cmd.Int("stMeritev")

			var postaje = map[string]vreme.Postaja{}
			vreme.DodajPostajo(postaje, "LJ12", "Ajdovscina", "Ljubljana")
			vreme.DodajPostajo(postaje, "LJ13", "Kolodvor", "Ljubljana")
			vreme.DodajPostajo(postaje, "LJ14", "Center", "Ljubljana")
			vreme.DodajMeritev(postaje, "LJ12", 40.0, minTemp, maxTemp)
			vreme.DodajMeritev(postaje, "LJ12", 30.0, minTemp, maxTemp)
			vreme.DodajMeritev(postaje, "LJ12", 10.0, minTemp, maxTemp)
			vreme.DodajMeritev(postaje, "LJ12", 20.0, minTemp, maxTemp)
			vreme.DodajMeritev(postaje, "LJ13", 40.0, minTemp, maxTemp)
			vreme.DodajMeritev(postaje, "LJ14", 30.0, minTemp, maxTemp)
			vreme.DodajMeritev(postaje, "LJ13", 40.0, minTemp, maxTemp)
			vreme.DodajMeritev(postaje, "LJ14", 30.0, minTemp, maxTemp)
			vreme.DodajMeritev(postaje, "LJ13", 10.0, minTemp, maxTemp)
			vreme.DodajMeritev(postaje, "LJ14", 20.0, minTemp, maxTemp)
			vreme.DodajMeritev(postaje, "LJ13", 10.0, minTemp, maxTemp)
			vreme.DodajMeritev(postaje, "LJ14", 20.0, minTemp, maxTemp)
			vreme.IzpisPorocila(postaje, min_st_Mer)
			return nil
		},
	}
	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}

}
