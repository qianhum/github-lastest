package internal

import (
	"context"
	"fmt"
	"time"

	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

type release struct {
	Name      githubv4.String
	CreatedAt time.Time
	UpdatedAt time.Time
}

func CheckRelease(token string, owner string, repoName string) error {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	httpClient := oauth2.NewClient(context.Background(), src)

	client := githubv4.NewClient(httpClient)

	variables := map[string]interface{}{
		"owner": githubv4.String(owner),
		"name":  githubv4.String(repoName),
	}

	var query struct {
		Repository struct {
			Releases struct {
				Nodes []release
			} `graphql:"releases(last: 1)"`
		} `graphql:"repository(owner: $owner, name: $name)"`
	}

	err := client.Query(context.Background(), &query, variables)
	if err != nil {
		return err
	}

	for _, node := range query.Repository.Releases.Nodes {
		fmt.Println("Release Name:", node.Name)
		fmt.Println("  Created At:", node.CreatedAt.Local())
		fmt.Println("  Updated At:", node.UpdatedAt.Local())
	}
	return nil
}
