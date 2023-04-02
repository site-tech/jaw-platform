package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/site-tech/jaw-platform/cmd"
	"github.com/site-tech/jaw-platform/ent"
	"github.com/site-tech/jaw-platform/graph/model"
	"github.com/site-tech/jaw-platform/pkg"
)

// Jaw is the resolver for the jaw field.
func (r *queryResolver) Jaw(ctx context.Context) (*ent.User, error) {
	go cmd.Run()
	return r.client.User.Query().First(ctx)
}

// DbConnection is the resolver for the dbConnection field.
func (r *queryResolver) DbConnection(ctx context.Context, cred *model.DBConnection) (*model.FlatTable, error) {
	dbDSN := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=%v",
		cred.User, cred.Password, cred.Host, cred.Port, cred.Dbname, cred.Sslmode)

	clientDB, err := sql.Open("postgres", dbDSN)
	if err != nil {
		log.Println("opening database: ", err)
		return nil, err
	}

	// Execute the query to get the columns information
	log.Println("performing query...")
	rows, err := clientDB.Query(`
		SELECT column_name, data_type
		FROM information_schema.columns
		WHERE table_name = 'routes_flat';
	`)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	defer clientDB.Close()

	// Fetch the results
	var columnName, dataType string
	res := model.FlatTable{}
	res.Name = "Airline Routes Data"
	for rows.Next() {
		err := rows.Scan(&columnName, &dataType)
		if err != nil {
			log.Println(err)
		}
		res.Columns = append(res.Columns, &model.Column{Name: columnName, Type: dataType})
	}

	pkg.InsertDBContext(ctx, clientDB)

	return &res, nil
}
