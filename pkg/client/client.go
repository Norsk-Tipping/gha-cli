package client

import (
	"fmt"
	"github.com/bradleyfalzon/ghinstallation/v2"
	"github.com/google/go-github/v45/github"
	"net/http"
	"os"
)

func CreateGitHubClientWithPrivateKeyFile(appId int64, installationId int64, privateKeyFile string) *github.Client {
	itr, err := ghinstallation.NewKeyFromFile(http.DefaultTransport, appId, installationId, privateKeyFile)
	if err != nil {
		fmt.Println("Client creation failed. Error:", err)
		os.Exit(1)
	}
	client := github.NewClient(&http.Client{Transport: itr})
	return client
}
