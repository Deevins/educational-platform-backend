package tests

import "gitlab.ozon.dev/daker255/homework-8/tests/pgmock"

var (
	DbInstance *pgmock.TDB
)

func init() {
	DbInstance = pgmock.NewMokeDB()
}
