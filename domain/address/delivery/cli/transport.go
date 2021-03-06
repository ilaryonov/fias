package cli

import (
	"github.com/urfave/cli"
	server "github.com/ilaryonov/fias/server/cli"
)

func RegisterCliEndpoints(app *server.App) {
	h := NewHandler(app.ImportService, app.Logger)
	app.Server.Commands = []cli.Command{
		{
			Name:  "checkupdates",
			Usage: "fias run full import or delta's",
			Action: func(c *cli.Context) {
				h.CheckUpdates(app.FiasApiService, app.VersionService)
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