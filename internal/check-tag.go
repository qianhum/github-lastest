package internal

import (
	"context"
	"fmt"

	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

type node struct {
	Name   githubv4.String
	Target struct {
		Commit `graphql:"... on Commit"`
	}
}

func CheckTag(token string, owner string, repoName string) error {
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
			Name githubv4.String
			Refs struct {
				Nodes []node
			} `graphql:"refs(refPrefix: \"refs/tags/\", last: 1)"`
		} `graphql:"repository(owner: $owner, name: $name)"`
	}

	err := client.Query(context.Background(), &query, variables)
	if err != nil {
		return err
	}

	for _, node := range query.Repository.Refs.Nodes {
		fmt.Println("      Tag Name:", node.Name)
		fmt.Println(" Git Object ID:", node.Target.Commit.Oid)
		fmt.Println(" Authored Date:", node.Target.Commit.AuthoredDate.Local())
		fmt.Println("Committed Date:", node.Target.Commit.CommittedDate.Local())
		fmt.Println("   Pushed Date:", node.Target.Commit.PushedDate.Local())
	}
	return nil
}
