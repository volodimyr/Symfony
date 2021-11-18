# Symfony

### Simple REST API service written in golang based on Echo framework. Most common features you'd like to see in every standalone project:
1. Support of unit and integration tests.
2. No import cycle trap.
3. Makefile, docker help to run, test, format etc. project rapidly.
4. Support of SQLboiler ORM, but you could write plain sql (whatever you'd prefer).
5. Use Services and domain layer to inject any repo's and therefore to combine/write the business logic you need.
6. Linter out of the box `make lint`.
7. Migration out of the box. Just add another version within migration package and re-run the service to commit your changes. 
8. Security is applicable.

#### Before you create a PR, run 
`make pr`

#### To see a test coverage, run
`make cover`

#### To run unit and integration tests
`make test`

#### To run databases and service itself, run
`make run`

And many other features that you can learn from Makefile.
