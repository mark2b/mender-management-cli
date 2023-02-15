package cli

import (
	"github.com/urfave/cli"
	"mender-management-cli/actions"
	"mender-management-cli/app"
	"mender-management-cli/conf"
	"os"
)

func ProcessCommandLine() {
	cl := cli.NewApp()
	cl.Name = app.Name
	cl.Usage = app.Name
	cl.Version = app.Version
	cl.Commands = []cli.Command{
		{
			Name: "upload-artifact",
			Action: func(*cli.Context) {
				conf.Config.Init()
				actions.UploadArtifact()
			},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "artifact",
					Destination: &actions.UploadArtifactContext.ArtifactFilepath,
				},
			},
		},
	}
	cl.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:        conf.VerboseFlag,
			Destination: &conf.Config.Verbose,
		},
		cli.BoolFlag{
			Name:        conf.DebugFlag,
			Destination: &conf.Config.Debug,
		},
		cli.StringFlag{
			Name:        conf.UserFlag,
			Destination: &conf.Config.User,
		},
		cli.StringFlag{
			Name:        conf.PasswordFlag,
			Destination: &conf.Config.Password,
		},
		cli.StringFlag{
			Name:        conf.EndpointFlag,
			Destination: &conf.Config.Endpoint,
			Value:       "https://hosted.mender.io",
		},
	}
	cl.Run(os.Args)
}
