package main

import (
	"context"
	"example/sola-example/restful-api/controller"
	"strings"

	"github.com/ddosakura/sola/v2/middleware/graphql"
	"github.com/ddosakura/sola/v2/middleware/router"

	"github.com/ddosakura/sola/v2"
)

var s = `
type Project {
	name: String!
	version: String!
}

type Query {
	hello: String!
	projects(name: String!): [Project]
}
`

type project struct {
	name    string
	version string
}

// Name Handler
func (p *project) Name() string {
	return p.name
}

// Version Handler
func (p *project) Version() string {
	return p.version
}

type query struct{}

// Hello Handler
func (*query) Hello() string {
	return "Hello, world!"
}

func (*query) Projects(ctx context.Context, args struct{ Name string }) *[]*project {
	if strings.HasPrefix("sola", args.Name) {
		return &[]*project{&project{"sola", sola.Version}}
	}
	return &[]*project{}
}

func main() {
	app := sola.New()
	r := router.New(nil)

	h := graphql.New(s, &query{})
	r.Bind("/graphql", h)

	controller.UserC(r)

	app.Use(r.Routes())
	sola.Listen("127.0.0.1:3000", app)
	sola.Keep()
}
