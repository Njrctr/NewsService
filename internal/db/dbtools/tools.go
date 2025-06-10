package dbtools

import (
	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
)

func Count1() exp.LiteralExpression {
	return goqu.L(`count(1)`)
}

func TabName(tab exp.AliasedExpression) string {
	return tab.Aliased().(exp.IdentifierExpression).GetTable()
}

func ColName(c exp.IdentifierExpression) string {
	return c.GetCol().(string)
}

func Now() exp.SQLFunctionExpression {
	return goqu.Func(`now`)
}

func DelCol(col exp.IdentifierExpression) exp.IdentifierExpression {
	return goqu.C(col.GetCol().(string))
}

func AnyIDInt(ids []int) exp.SQLFunctionExpression {
	return goqu.Func(`any`, ids)
}

var DC = DelCol
