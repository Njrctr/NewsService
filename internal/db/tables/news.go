package tables

import (
	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
)

var news = goqu.T(`news`).As(`n`)

type newsTable struct {
	T           exp.AliasedExpression
	ID          exp.IdentifierExpression
	CategoryID  exp.IdentifierExpression
	Title       exp.IdentifierExpression
	Foreword    exp.IdentifierExpression
	Content     exp.IdentifierExpression
	Author      exp.IdentifierExpression
	CreatedAt   exp.IdentifierExpression
	PublishedAt exp.IdentifierExpression
	TagIDs      exp.IdentifierExpression
	StatusID    exp.IdentifierExpression
}

var News = &newsTable{
	T:           news,
	ID:          news.Col(`newsId`),
	CategoryID:  news.Col(`categoryId`),
	Title:       news.Col(`title`),
	Foreword:    news.Col(`foreword`),
	Content:     news.Col(`content`),
	Author:      news.Col(`author`),
	CreatedAt:   news.Col(`createdAt`),
	PublishedAt: news.Col(`publishedAt`),
	TagIDs:      news.Col(`tagIds`),
	StatusID:    news.Col(`statusId`),
}
