# Recipe API

This is a simple API that allows the user to create and modify recipes and ingredients via endpoints

## Sharding

This service lives behind a gateway that is responsible with sending recipe requests to the right instance of this service.
Different shards are identified via the `SHARD` environment variable. This is mostly used to identify the return queue for
ingredient lookup.

## Technology

* [Ent](https://entgo.io/docs/getting-started) an entity framework, that handles the database. It creates and updates tables the database based on a defined schema
* [Gin](https://github.com/gin-gonic/gin) a web framework used to define and handle the endpoints in the app
* [AMQP](https://github.com/streadway/amqp) to connect to a queue to have another service lookup ingredient information that get returned via another queue

## Schema updates

The schema is defined in [./ent/schema/](./ent/schema/), whenever the definition is changed here you need to run `go generate ./ent` or `make generate` to update the generated code that ent provides.
