package banner

import (
	"fmt"

	"github.com/krishpranav/wpscan/internal/pkg/update"
	"github.com/krishpranav/wpscan/internal/pkg/version"
	"github.com/krishpranav/wpscan/pkg/printer"
)

func Banner() {
	printer.Println("——————————————————————————————————————————————————————————————————")
	fmt.Print("___       ______________________________________________   __\n__ |     / /__  __ \\__  __ \\__  ____/_  ____/_  __ \\__  | / /\n__ | /| / /__  /_/ /_  /_/ /_  __/  _  /    _  / / /_   |/ / \n__ |/ |/ / _  ____/_  _, _/_  /___  / /___  / /_/ /_  /|  /  \n____/|__/  /_/     /_/ |_| /_____/  \\____/  \\____/ /_/ |_/   \n\n")
	printer.Println("Github: ", "https://github.com/krishpranav/wpscan")

	if newVersion := update.CheckUpdate(); newVersion != "" {
		printer.Println("Version:", version.Version, "(New Version: "+newVersion+")")
	} else {
		printer.Println("Version:", version.Version)
	}

	printer.Println("——————————————————————————————————————————————————————————————————")
}
