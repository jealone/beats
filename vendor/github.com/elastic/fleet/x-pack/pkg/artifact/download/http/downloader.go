// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package http

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"

	"github.com/pkg/errors"

	"github.com/elastic/fleet/pkg/release"
	"github.com/elastic/fleet/x-pack/pkg/artifact"
)

const (
	packagePermissions = 0660
)

var headers = map[string]string{
	"User-Agent": fmt.Sprintf("Beat agent v%s", release.Version()),
}

// Downloader is a downloader able to fetch artifacts from elastic.co web page.
type Downloader struct {
	config *artifact.Config
	client http.Client
}

// NewDownloader creates and configures Elastic Downloader
func NewDownloader(config *artifact.Config) *Downloader {
	client := http.Client{Timeout: config.Timeout}
	rt := withHeaders(client.Transport, headers)
	client.Transport = rt
	return NewDownloaderWithClient(config, client)
}

// NewDownloaderWithClient creates Elastic Downloader with specific client used
func NewDownloaderWithClient(config *artifact.Config, client http.Client) *Downloader {
	return &Downloader{
		config: config,
		client: client,
	}
}

// Download fetches the package from configured source.
// Returns absolute path to downloaded package and an error.
func (e *Downloader) Download(programName, version string) (string, error) {
	// download from source to dest
	path, err := e.download(e.config.OS(), programName, version)
	if err != nil {
		os.Remove(path)
	}

	return path, err
}

func (e *Downloader) composeURI(programName, packageName string) (string, error) {
	upstream := e.config.BeatsSourceURI
	if !strings.HasPrefix(upstream, "http") && !strings.HasPrefix(upstream, "file") && !strings.HasPrefix(upstream, "/") {
		// always default to https
		upstream = fmt.Sprintf("https://%s", upstream)
	}

	// example: https://artifacts.elastic.co/downloads/beats/filebeat/filebeat-7.1.1-x86_64.rpm
	uri, err := url.Parse(upstream)
	if err != nil {
		return "", errors.Wrap(err, "invalid upstream URI")
	}

	uri.Path = path.Join(uri.Path, programName, packageName)
	return uri.String(), nil
}

func (e *Downloader) download(operatingSystem, programName, version string) (string, error) {
	filename, err := artifact.GetArtifactName(programName, version, operatingSystem, e.config.Arch())
	if err != nil {
		return "", errors.Wrap(err, "generating package name failed")
	}

	fullPath, err := artifact.GetArtifactPath(programName, version, operatingSystem, e.config.Arch(), e.config.TargetDirectory)
	if err != nil {
		return "", errors.Wrap(err, "generating package path failed")
	}

	sourceURI, err := e.composeURI(programName, filename)
	if err != nil {
		return "", err
	}

	resp, err := e.client.Get(sourceURI)
	if err != nil {
		return "", errors.Wrap(err, "fetching package failed")
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("call to '%s' returned unsuccessful status code: %d", sourceURI, resp.StatusCode)
	}

	destinationFile, err := os.OpenFile(fullPath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, packagePermissions)
	if err != nil {
		return "", errors.Wrap(err, "creating package file failed")
	}
	defer destinationFile.Close()

	_, err = io.Copy(destinationFile, resp.Body)
	return fullPath, nil
}