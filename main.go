package main

import (
	Eulogist "Eulogist/eulogist"
	_ "embed"
	"fmt"

	"github.com/pterm/pterm"
)

//go:embed version
var versionInfo []byte

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
			pterm.Sprintf("Bedrock - v%s", string(versionInfo)) + "\n" +
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
