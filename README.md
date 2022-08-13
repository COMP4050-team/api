# API

[![Test and Deploy API to dev](https://github.com/COMP4050-team/api/actions/workflows/deploy.yml/badge.svg)](https://github.com/COMP4050-team/api/actions/workflows/deploy.yml)

[![Coverage Status](https://coveralls.io/repos/github/COMP4050-team/api/badge.svg?branch=main&t=Olikwl)](https://coveralls.io/github/COMP4050-team/api?branch=main)

## Development

When you make changes to the GraphQL schema you will need to generate the new resolvers:

```
make generate
```

In order to run the API:

```
make run
```

Then you can visit the GraphQL playground at http://localhost:8080