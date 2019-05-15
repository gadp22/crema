

# Crema
Crema is a Simple Go-Based RESTful API Framework originally written by [Universitas Pertamina ICT teams](https://tki.universitaspertamina.ac.id) to support its Microservices environment development. It is written in pure Go and completely free.

## RESTful
Playing with RESTful API using Crema is fun and easy

* Create HTTP routes

```go
func main() {
	server := crema.InitServer()

	server.AddRoutes(http.MethodGet, "/hello", hello)
	server.AddRoutes(http.MethodGet, "/users", crema.MakeGenericGetHandler(getUser))
  
	crema.LogPrintf("[MAIN] Server is running, listening to port 8001 ....")

	log.Fatal(http.ListenAndServe(":8001", server.Router))
}
```

* Create Handlers
```go
func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

func getUser(conditions map[string]string) (*sql.Rows, error) {
	q := crema.Select().All().From("users")

	return crema.ExecuteQuery(q.QueryString)
}
```
* Try it using cURL
```sh
$ curl http://localhost:8001/hello
Hello World!

$ curl http://localhost:8001/users
{"status":1,"message":"Success 200","data":[{"email":"john.doe@email.com","id":16,"name":"John Doe"},{"email":"bob@email.com","id":17,"name":"Bob"}]}
```


## What's Inside

* A Gorilla Mux-based HTTP router
* Customizable generic request handler
* Simplified Data Access Object mechanism
* PostgreSQL and MySQL database driver
* Personalized system logger

## Installation
1. With a correctly configured Go toolchain:

```sh
$ go get github.com/gadp22/crema
```

2. Import in your code

```go
import "github.com/gadp22/crema"
```

## Example
See [here](https://github.com/gadp22/crema/blob/master/example.go) to see basic example to build a CRUD API. 
