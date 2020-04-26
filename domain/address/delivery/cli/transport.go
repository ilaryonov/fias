package cli

import (
	"github.com/urfave/cli"
	"gitlab.com/ilaryonov/fiascli-clean/domain/address/service"
	server "gitlab.com/ilaryonov/fiascli-clean/server/cli"
)

func RegisterCliEndpoints(app *server.App, addressService *service.AddressService) {
	h := NewHandler(*addressService)
	app.Server.Commands = []cli.Command{
		{
			Name:  "checkupdates",
			Usage: "fias run full import or delta's",
			Action: func(c *cli.Context) {
				h.CheckUpdates()
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