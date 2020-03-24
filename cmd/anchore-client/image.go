package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	log "github.com/sirupsen/logrus"
)

var (
	ImageCommand = &cli.Command{
		Name:                   "images",
		Aliases: []string{"image"},
		Description:            "Image operations working with images",
		Subcommands: []*cli.Command{
			{
				Name:        "list",
				Description: "List images in the system",
				Aliases: []string{"l"},
				Action: ListImagesAction,
				Flags: []cli.Flag{
					&cli.BoolFlag{Name:"show-all"},
					&cli.BoolFlag{Name:"full"},
				},
			},
			{
				Name:        "get",
				Description: "Get information on a single image",
				Aliases:     []string{"g"},
				Action:      GetImageAction,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name: "image",
						Required: true,
					},
					&cli.BoolFlag{Name:"show-history"},
				},
			},
		},
	}
)

func ListImagesAction(c *cli.Context) error {
	host := c.String("url")
	user := c.String("username")
	pass := c.String("password")
	log.Debugf("Vars user %v, pass %v, url %v", user, pass, host)
	apiClient, authCtx, _ := NewInitializedClient(host, user, pass)
	images, _, err := apiClient.DefaultApi.ListImages(authCtx, nil)
	if err != nil {
		return err
	}

	cfg := NewOutputConfiguration(false, true)
	renderer, err := NewTableRenderer(cfg)
	err = renderer.RenderImages(&images)
	if err != nil {
		return err
	}

	return nil
}

func GetImageAction(context *cli.Context) error {
	fmt.Println("Getting images")
	host := context.Value("url").(string)
	user := context.Value("username").(string)
	pass := context.Value("password").(string)
	c, authCtx, _ := NewInitializedClient(host, user, pass)

	var imgString string
	imgString = context.Value("image").(string)
	img, err := ResolveImage(c, authCtx, imgString)
	if err != nil {
		return err
	}
	fmt.Printf("Got image %v", img)
	return nil
}

