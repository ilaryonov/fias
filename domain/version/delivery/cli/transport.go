package cli

import (
	"github.com/urfave/cli"
	server "github.com/ilaryonov/fiasserver/cli"
)

func RegisterCliEndpoints(app *server.App) {
	h := NewHandler(*app.VersionService)
	app.Server.Commands = []cli.Command{
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