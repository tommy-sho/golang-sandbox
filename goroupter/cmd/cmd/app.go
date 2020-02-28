package cmd

import (
	"fmt"
	"github.com/tommy-sho/golang-sandbox/goroupter/cmd/cmd"
)

func newApp() *cli.App {
	app := cli.NewApp()
	app.Name = "grouper"
	app.Usage = "Force grouped import path"
	app.Version = fmt.Sprintf("%s-%s", cmd.version, cmd.revision)
	app.Authors = []*cli.Author{{
		Name:  "tommy-sho",
		Email: "tomiokasyogo@gmail.com",
	}}
	app.Action = action
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:  "local",
			Usage: "specify imports prefix beginning with this string after 3rd-party packages. especially your own organization name. comma-separated list",
		},
		&cli.StringFlag{
			Name:        "path",
			Aliases:     nil,
			Usage:       "",
			EnvVars:     nil,
			FilePath:    "",
			Required:    false,
			Hidden:      false,
			TakesFile:   false,
			Value:       "",
			DefaultText: "",
			Destination: nil,
			HasBeenSet:  false,
		},
	}

	return app
}

func action(c *cli.Context) error {
	local := c.Value("local")
	fmt.Println(local)

	paths := c.Args()
	fmt.Println(paths.Slice())
	return nil
}
