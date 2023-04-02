package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	"github.com/site-tech/jaw-platform/cmd"
	"github.com/site-tech/jaw-platform/ent"
	"github.com/site-tech/jaw-platform/graph/model"
)

var ClientDB *sql.DB

// Jaw is the resolver for the jaw field.
func (r *queryResolver) Jaw(ctx context.Context) (*ent.User, error) {
	go cmd.Run()
	return r.client.User.Query().First(ctx)
}

// DbConnection is the resolver for the dbConnection field.
func (r *queryResolver) DbConnection(ctx context.Context, cred *model.DBConnection) (*model.FlatTable, error) {
	dbDSN := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=%v",
		cred.User, cred.Password, cred.Host, cred.Port, cred.Dbname, cred.Sslmode)

	db, err := sql.Open("postgres", dbDSN)
	if err != nil {
		log.Println("opening database: ", err)
		return nil, err
	}
	ClientDB = db

	// Execute the query to get the columns information
	log.Println("performing query...")
	rows, err := ClientDB.Query(`
		SELECT column_name, data_type
		FROM information_schema.columns
		WHERE table_name = 'routes_flat';
	`)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

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

	//ctx = context.WithValue(context.Background(), "dbClient", ClientDB)

	return &res, nil
}

// BuildReport is the resolver for the buildReport field.
func (r *queryResolver) BuildReport(ctx context.Context, clause *model.ReportClause) (string, error) {
	log.Println("clause: ", clause.Selections)
	//clientDB, ok := ctx.Value("dbClient").(*sql.DB)
	//if !ok {
	//return "500", fmt.Errorf("couldn't get db out of context")
	//}
	sql := "SELECT * FROM routes_flat WHERE"
	var inputs []interface{}
	for i, v := range clause.Selections {
		if i > 0 {
			sql = fmt.Sprintf("%s AND", sql)
		}
		sql = fmt.Sprintf("%s %s", sql, v.Field)
		switch v.Operator {
		case "equal":
			sql = fmt.Sprintf("%s %s", sql, "=")
		case "greater":
			sql = fmt.Sprintf("%s %s", sql, ">")
		case "less":
			sql = fmt.Sprintf("%s %s", sql, "<")
		default:
			log.Println("default case")
		}
		sql = fmt.Sprintf("%s %s%d", sql, "$", i+1)
		inputs = append(inputs, v.Value)
	}

	log.Println("sql query: ", sql)
	log.Println("inputs: ", inputs)

	rows, err := ClientDB.Query(sql, inputs...)
	if err != nil {
		return "", err
	}
	defer rows.Close()

	var result []map[string]interface{}

	cols, _ := rows.Columns()
	colNum := len(cols)

	for rows.Next() {
		columns := make([]interface{}, colNum)
		columnPtrs := make([]interface{}, colNum)

		for i := 0; i < colNum; i++ {
			columnPtrs[i] = &columns[i]
		}

		err = rows.Scan(columnPtrs...)
		if err != nil {
			return "500", err
		}

		rowData := make(map[string]interface{})
		for i, colName := range cols {
			var v interface{}
			val := columns[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			rowData[colName] = v
		}

		result = append(result, rowData)
	}

	jsonResult, err := json.Marshal(result)
	if err != nil {
		return "500", err
	}

	return string(jsonResult), nil
}
