package cli

import (
	"github.com/urfave/cli"
	"gitlab.com/ilaryonov/fiascli-clean/domain/version/service"
)

func RegisterCliEndpoints(app *cli.App, versionService *service.VersionService) {
	h := NewHandler(*versionService)
	app.Commands = []cli.Command{
		{
			Name:  "version",
			Usage: "fias version",
			Action: func(c *cli.Context) {
				h.GetVersionInfo()
			},
		},
		/*{
			Name:  "checkdelta",
			Usage: "check deltas from fias.nalog.ru",
			Action: func(c *cli.Context) {
				controllers.CheckUpdates()
			},
		},*/
	}
}