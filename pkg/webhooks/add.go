package webhooks

import (
	"context"
	"github.com/google/go-github/v45/github"
)

func AddWebhook(ctx context.Context, client *github.Client, orgName string, repository string, webhookUrl string, contentType string, events []string) error {

	_, _, err := client.Repositories.CreateHook(ctx, orgName, repository, &github.Hook{
		Config: map[string]interface{}{
			"url":          webhookUrl,
			"content_type": contentType,
		},
		Events: events,
	})

	return err
}
