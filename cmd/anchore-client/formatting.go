package main

import (
	"github.com/anchore/client-go/pkg/anchore/client"
	"github.com/olekukonko/tablewriter"
	"os"
	"time"
)


func RenderAnchoreImageToTable(Images *[]client.AnchoreImage) error {
	/*
	Image Digest: sha256:3936fb3946790d711a68c58be93628e43cbca72439079e16d154b5db216b58da
	Parent Digest: sha256:2539d4344dd18e1df02be842ffc435f8e1f699cfc55516e2cf2cb16b7a9aea0b
	Analysis Status: analyzing
	Image Type: docker
	Analyzed At: None
	Image ID: 6678c7c2e56c970388f8d5a398aa30f2ab60e85f20165e101053c3d3a11e6663
	Dockerfile Mode: None
	Distro: None
	Distro Version: None
	Size: None
	Architecture: None
	Layer Count: None

	Full Tag: docker.io/nginx:latest
	Tag Detected At: 2020-03-24T16:31:58Z
	*/
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Tag", "Digest", "Status", "Created At"})
	table.SetBorder(false)                                // Set Border to false
	for _, itm := range *Images {
		r := []string{itm.ImageDetail[0].Fulltag, itm.ImageDigest, itm.AnalysisStatus, itm.CreatedAt.Format(time.RFC3339)}
		table.Append(r)
	}
	table.Render()
	return nil
}

// Rendering interface to convert CLI objects to consistent outputs
// Expect implementations for each output type: json, table, etc
type OutputRenderer interface {
	RenderImage(image *client.AnchoreImage) error
}

type TableRenderer struct {
	Table *tablewriter.Table
}

func NewTableRenderer(outputConfig *OutputConfiguration) (*OutputRenderer, error) {
	return nil, nil
}