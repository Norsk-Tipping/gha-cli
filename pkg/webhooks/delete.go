package webhooks

import (
	"context"
	"github.com/google/go-github/v45/github"
)

func DeleteWebhook(ctx context.Context, client *github.Client, orgName string, repository string, webhookId int64) error {
	_, err := client.Repositories.DeleteHook(ctx, orgName, repository, webhookId)
	return err
}
