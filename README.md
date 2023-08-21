# Welsh academy
[![coverage report](https://github.com/casahugo/go-welsh-academy/badges/master/coverage.svg)](https://github.com/casahugo/go-welsh-academy/badges/master/coverage.svg)

"Welsh Academy" is an application dedicated to provide recipes to cheddar lovers around the world.

Your duty here is to provide a backend in order to allow our cheddar experts to:

- Create ingredients
- Create recipes of meals using the previously created ingredients
For the moment, admins will be the one that create users though the API (don't worry about permission management, the cheese world is a kind world)

A user should be able to enjoy the recipes by using the API to:

- list all existing ingredients
- list all possible recipes (with or without ingredient constraints)
- flag/unflag recipes as his favorite ones
- list his favorite recipes

Hints:
- Design and build this API
- Provide tests
- Provide documentation to launch and test this API

Constraints:

language: use either Golang (with whatever web framework you want) or Python (using Flask and SQLAlchemy)
provide us a link to a git repository

we should be able to run the application in a container

## Requirements

- [Docker](https://docs.docker.com/install/#supported-platforms) >= 24.0.4
- [Docker compose](https://docs.docker.com/compose/install) >= 2.5.0

## Install

```bash
make install
```

Open http://localhost:8080 

## Commands

Connect to shell
```bash
make shell
```

Show logs
```bash
make logs
```

Stop application
```bash
make stop
```

## Migration SQL
```bash
go run cmd/migration/migration
```
