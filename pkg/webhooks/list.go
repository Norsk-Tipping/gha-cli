package webhooks

import (
	"context"
	"github.com/google/go-github/v45/github"
)

func GetWebhooks(ctx context.Context, client *github.Client, orgName string, repository string) ([]*github.Hook, error) {
	hooks, _, err := client.Repositories.ListHooks(ctx, orgName, repository, nil)
	return hooks, err
}
