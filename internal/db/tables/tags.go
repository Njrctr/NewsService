package tables

import (
	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
)

var tags = goqu.T(`tags`).As(`t`)

type tagsTable struct {
	T        exp.AliasedExpression
	ID       exp.IdentifierExpression
	Title    exp.IdentifierExpression
	StatusID exp.IdentifierExpression
}

var Tags = &tagsTable{
	T:        tags,
	ID:       tags.Col(`tagId`),
	Title:    tags.Col(`title`),
	StatusID: tags.Col(`statusId`),
}
