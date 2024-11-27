package db

import (
	sq "github.com/Masterminds/squirrel"
)

var psql sq.StatementBuilderType

func GetQueryBuilder() sq.StatementBuilderType {
	return psql
}
