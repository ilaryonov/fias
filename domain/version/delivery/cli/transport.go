package cli

import (
	"github.com/urfave/cli"
	"gitlab.com/ilaryonov/fiascli-clean/domain/version/service"
	server "gitlab.com/ilaryonov/fiascli-clean/server/cli"
)

func RegisterCliEndpoints(app *server.App, versionService *service.VersionService) {
	h := NewHandler(*versionService)
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