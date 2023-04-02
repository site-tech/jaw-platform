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

	ctx = context.WithValue(context.Background(), "dbClient", clientDB)

	return &res, nil
}

// BuildReport is the resolver for the buildReport field.
func (r *queryResolver) BuildReport(ctx context.Context, clause *model.ReportClause) (string, error) {
	log.Println("clause: ", clause.Selections)
	clientDB, ok := ctx.Value("dbClient").(*sql.DB)
	if !ok {
		return "500", fmt.Errorf("couldn't get db out of context")
	}
	sql := "SELECT * FROM routes_flat WHERE "
	var inputs []string
	for i, v := range clause.Selections {
		sql = fmt.Sprintf("%s%s", sql, v.Field)
		switch v.Operator {
		case "equals":
			sql = fmt.Sprintf("%s %s", sql, "=")
		default:
			log.Println("default case")
		}
		sql = fmt.Sprintf("%s %s%d", sql, "$", i)
		inputs = append(inputs, v.Value)
	}

	log.Println("sql query: ", sql)
	log.Println("inputs: ", inputs)

	clientDB.Query(sql, inputs)
	return "200", nil
}
