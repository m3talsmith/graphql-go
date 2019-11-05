package resolvable

import (
	// "reflect"
	"testing"

	"github.com/graph-gophers/graphql-go/internal/schema"
)

func Test_newMeta(t *testing.T) {
	var err error
	gql := `
type Foo {
	value: String
}

type Query {
	query: Foo
}
	`
	schema := schema.New()
	if err = schema.Parse(gql, false); err != nil {
		t.Errorf("error parsing schema: %s", err)
	}
	// TODO: Finish test
}
