package main

import (
	"fmt"
	"github.com/anchore/client-go/pkg/anchore/client"
	"github.com/olekukonko/tablewriter"
	"os"
	"time"
)

type OutputFormat int
const (
	JSON = iota
	TABLE // Default
)

type OutputConfiguration struct {
	Format OutputFormat
	Verbose bool
}

func NewOutputConfiguration(isJson bool, isDebug bool) *OutputConfiguration {
	var outFmt OutputFormat = TABLE
	if isJson {
		outFmt = JSON
	}

	return &OutputConfiguration{
		Format:  outFmt,
		Verbose: isDebug,
	}
}

// Rendering interface to convert CLI objects to consistent outputs
// Expect implementations for each output type: json, table, etc
type OutputRenderer interface {
	RenderImages(image *[]client.AnchoreImage) error
}

type TableRenderer struct {
	Configuration *OutputConfiguration
	Table *tablewriter.Table
}

func (t *TableRenderer) RenderImages(images *[]client.AnchoreImage) error {
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

	t.Table.SetHeader([]string{"Tag", "Digest", "Status", "Created At"})
	t.Table.SetBorder(false)                                // Set Border to false
	for _, itm := range *images {
		r := []string{itm.ImageDetail[0].Fulltag, itm.ImageDigest, itm.AnalysisStatus, itm.CreatedAt.Format(time.RFC3339)}
		t.Table.Append(r)
	}
	t.Table.Render()
	return nil
}

func NewTableRenderer(outputConfig *OutputConfiguration) (*TableRenderer, error) {
	if outputConfig.Format != TABLE {
		return nil, fmt.Errorf("Invalid input for this renderer. Should be TABLE")
	}

	r := &TableRenderer{
		Configuration: outputConfig,
		Table:         tablewriter.NewWriter(os.Stdout),
	}

	return r, nil
}

