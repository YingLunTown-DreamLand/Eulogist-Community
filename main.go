package main

import (
	Eulogist "Eulogist/eulogist"
	"fmt"

	"github.com/pterm/pterm"
)

func main() {
	pterm.Println(pterm.Yellow("" +
		"╔═════════════════════════════════════════════════════════════════╗\n" +
		"║ ███████╗██╗   ██╗██╗      ██████╗  ██████╗ ██╗███████╗████████╗ ║\n" +
		"║ ██╔════╝██║   ██║██║     ██╔═══██╗██╔════╝ ██║██╔════╝╚══██╔══╝ ║\n" +
		"║ █████╗  ██║   ██║██║     ██║   ██║██║  ███╗██║███████╗   ██║    ║\n" +
		"║ ██╔══╝  ██║   ██║██║     ██║   ██║██║   ██║██║╚════██║   ██║    ║\n" +
		"║ ███████╗╚██████╔╝███████╗╚██████╔╝╚██████╔╝██║███████║   ██║    ║\n" +
		"║ ╚══════╝ ╚═════╝ ╚══════╝ ╚═════╝  ╚═════╝ ╚═╝╚══════╝   ╚═╝    ║\n" +
		"╚═════════════════════════════════════════════════════════════════╝",
	))
	pterm.DefaultBox.Println(
		pterm.LightCyan("" +
			"                    " +
			"LICENSE | AGPL-v3.0" + "\n" +
			"https://github.com/YingLunTown-DreamLand/Eulogist-Community",
		),
	)

	err := Eulogist.Eulogist()
	if err != nil {
		pterm.Error.Println(err)
	}

	fmt.Println()
	pterm.Info.Println("Program running down, now press enter to exit.")
	fmt.Scanln()
}
