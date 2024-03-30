package start

const baseURL = "https://api.start.gg/gql/alpha"

type (
	StartGraphQLClient struct {
		token string
	}
	GraphQLConfig struct{}
)

func NewStartGraphQLClient(token string, config ...GraphQLConfig) *StartGraphQLClient {
	// TODO: implement config settings for GraphQL

	// Test query at https://api.start.gg/gql/alpha with token to know if it works
	gqlClient := graphql.pa
	return &StartGraphQLClient{token: token}
}
