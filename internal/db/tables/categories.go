package tables

import (
	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
)

var categories = goqu.T(`categories`).As(`c`)

type categoriesTable struct {
	T           exp.AliasedExpression
	ID          exp.IdentifierExpression
	Title       exp.IdentifierExpression
	OrderNumber exp.IdentifierExpression
	StatusID    exp.IdentifierExpression
}

var Categories = &categoriesTable{
	T:           categories,
	ID:          categories.Col(`categoryId`),
	Title:       categories.Col(`title`),
	OrderNumber: categories.Col(`orderNumber`),
	StatusID:    categories.Col(`statusId`),
}
