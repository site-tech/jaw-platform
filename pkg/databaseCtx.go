package pkg

import (
	"context"
	"database/sql"
)

type DBContextKey struct{}

func InsertDBContext(ctx context.Context, db *sql.DB) context.Context {
	return context.WithValue(ctx, DBContextKey{}, db)
}

func GetClientDBFromContext(ctx context.Context) (*sql.DB, bool) {
	db, ok := ctx.Value(DBContextKey{}).(*sql.DB)
	if ok {
		return db, ok
	}
	return nil, false
}
