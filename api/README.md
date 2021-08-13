# Recipe API

This is a simple API that allows the user to sign up and currently just create and remove ingredients via endpoints

## Technology

* [Ent](https://entgo.io/docs/getting-started) an entity framework, that handles the database. It creates and updates tables the database based on a defined schema
* [Gin](https://github.com/gin-gonic/gin) a web framework used to define and handle the endpoints in the app

## Schema updates

The schema is defined in [./ent/schema/](./ent/schema/), whenever the definition is changed here you need to run `go generate ./ent` or `make generate` to update the generated code that ent provides.
