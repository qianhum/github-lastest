package internal

import (
	"context"
	"fmt"
	"time"

	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

type Commit struct {
	Oid           githubv4.GitObjectID
	AuthoredDate  time.Time
	CommittedDate time.Time
	PushedDate    time.Time
}

func CheckCommit(token string, owner string, repoName string, branchName string) error {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	httpClient := oauth2.NewClient(context.Background(), src)

	client := githubv4.NewClient(httpClient)

	variables := map[string]interface{}{
		"owner":      githubv4.String(owner),
		"name":       githubv4.String(repoName),
		"branchName": githubv4.String(branchName),
	}

	var query struct {
		Repository struct {
			Name githubv4.String
			Ref  struct {
				Name   githubv4.String
				Target struct {
					Commit `graphql:"... on Commit"`
				}
			} `graphql:"ref(qualifiedName: $branchName)"`
		} `graphql:"repository(owner: $owner, name: $name)"`
	}

	err := client.Query(context.Background(), &query, variables)
	if err != nil {
		return err
	}

	fmt.Println("   Branch Name:", query.Repository.Ref.Name)
	fmt.Println(" Git Object ID:", query.Repository.Ref.Target.Commit.Oid)
	fmt.Println(" Authored Date:", query.Repository.Ref.Target.Commit.AuthoredDate.Local())
	fmt.Println("Committed Date:", query.Repository.Ref.Target.Commit.CommittedDate.Local())
	fmt.Println("   Pushed Date:", query.Repository.Ref.Target.Commit.PushedDate.Local())
	return nil
}

func CheckDefaultCommit(token string, owner string, repoName string) error {
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
			Name             githubv4.String
			DefaultBranchRef struct {
				Name   githubv4.String
				Target struct {
					Commit `graphql:"... on Commit"`
				}
			}
		} `graphql:"repository(owner: $owner, name: $name)"`
	}

	err := client.Query(context.Background(), &query, variables)
	if err != nil {
		return err
	}

	fmt.Println("   Branch Name:", query.Repository.DefaultBranchRef.Name)
	fmt.Println(" Git Object ID:", query.Repository.DefaultBranchRef.Target.Commit.Oid)
	fmt.Println(" Authored Date:", query.Repository.DefaultBranchRef.Target.Commit.AuthoredDate.Local())
	fmt.Println("Committed Date:", query.Repository.DefaultBranchRef.Target.Commit.CommittedDate.Local())
	fmt.Println("   Pushed Date:", query.Repository.DefaultBranchRef.Target.Commit.PushedDate.Local())
	return nil
}
