/**
 * This example is using single table "users"
 *
 * CREATE TABLE users (
 * id serial PRIMARY KEY,
 * name VARCHAR(50),
 * email VARCHAR(50)
 * );
 *
 * Endpoints:
 *
 * GET /users
 * GET /users/{id}
 * POST /users
 * PUT /users/{id}
 * DELETE /users/{id}
 *
 * Example:
 *
 * GET http://localhost:8001/users/1
 * POST http://localhost:8001/users?name=John&email=doe@john.com
 */

package crema

import (
	"database/sql"
	"log"
	"net/http"
	//import crema
	//. "github.com/gadp22/
)

func GetUser(conditions map[string]string) (*sql.Rows, error) {
	q := GetGenericSelectQuery("users", conditions)

	return ExecuteQuery(q.QueryString)
}

func PostUser(tx *sql.Tx, values map[string]string) *sql.Row {
	q := GetGenericInsertQuery("users", values)

	return ExecuteQueryRow(tx, q.QueryString)
}

func PutUser(tx *sql.Tx, values map[string]string) (sql.Result, error) {
	q := GetGenericUpdateQuery("users", values)

	return ExecuteNonQuery(q.QueryString)
}

func DeleteUser(tx *sql.Tx, conditions map[string]string) (sql.Result, error) {
	q := GetGenericDeleteQuery("users", conditions)

	return ExecuteNonQuery(q.QueryString)
}

func main_example() {
	server := InitServer()

	server.AddRoutes(http.MethodGet, "/users", MakeGenericGetHandler(GetUser))
	server.AddRoutes(http.MethodGet, "/users/{id}", MakeGenericGetHandler(GetUser))

	server.AddRoutes(http.MethodPost, "/users", MakeGenericPostHandler(PostUser))

	server.AddRoutes(http.MethodPut, "/users/{id}", MakeGenericPutHandler(PutUser))

	server.AddRoutes(http.MethodDelete, "/users/{id}", MakeGenericDeleteHandler(DeleteUser))

	LogPrintf("server is started, listening to 8001")
	log.Fatal(http.ListenAndServe(":8001", server.Router))
}
