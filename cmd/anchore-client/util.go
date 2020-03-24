package main

import (
	"context"
	"fmt"
	"github.com/anchore/client-go/pkg/anchore/client"
	"github.com/antihax/optional"
	"net/url"
)

func NewInitializedClient(Url string, Username string, Password string) (*client.APIClient, context.Context, error){
	cfg := client.NewConfiguration()

	parsedUrl, err := url.Parse(Url)
	if err != nil {
		fmt.Printf("Invalid url %v\n", Url)
	}

	cfg.Scheme = parsedUrl.Scheme
	cfg.Host = parsedUrl.Host
	cfg.BasePath = parsedUrl.Path
	cfg.UserAgent = "anchore/client-go@some_commit_sha"

	c := client.NewAPIClient(cfg)

	auth := context.WithValue(context.Background(), client.ContextBasicAuth, client.BasicAuth{
		UserName: Username,
		Password: Password,
		})
	return c, auth, nil
}

func ResolveImage(Client *client.APIClient, authCtx context.Context, ImageReference string) (client.AnchoreImage, error) {
	listOpts := client.ListImagesOpts{
		History:         optional.NewBool(false),
		Fulltag:         optional.NewString(ImageReference),
	}
	var imgs []client.AnchoreImage
	imgs, _, err := Client.DefaultApi.ListImages(authCtx, &listOpts)
	if err != nil {
		fmt.Println("Got an error")
		return client.AnchoreImage{}, err
	}

	if len(imgs) > 0 {
		fmt.Printf("Found an image %v", imgs[0].ImageDigest)
	}

	return imgs[0], nil
}
