package internal

import (
	"database/sql"
	"github.com/ernesto2108/APIRest_Business/internal/logs"
	_ "github.com/lib/pq"
)

type SqlClient struct {
	*sql.DB
}

func getConn(source string) *SqlClient {
	db, err := sql.Open("postgres", source)
	if err != nil {
		logs.Log().Error("Cannot connect with postgres")
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		logs.Log().Warn("Failed to ping postgres")
		panic(err)
	}

	return &SqlClient{}
}

func (c SqlClient) viewState() sql.DBStats {
	return c.Stats()
}
