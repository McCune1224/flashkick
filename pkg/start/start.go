package start

import (
	"context"

	"github.com/hasura/go-graphql-client"
	"golang.org/x/oauth2"
)

const baseURL = "https://api.start.gg/gql/alpha"

type StartGGClient struct {
	Client *graphql.Client
}

func NewStartGGClient(graphQLToken string) *StartGGClient {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: graphQLToken},
	)
	httpClient := oauth2.NewClient(context.Background(), src)
	startGGClient := graphql.NewClient("https://api.start.gg/gql/alpha", httpClient)

	return &StartGGClient{
		Client: startGGClient,
	}
}
