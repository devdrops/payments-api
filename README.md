# Payments API

A [RESTful](https://en.wikipedia.org/wiki/REST) API that processes transactions.

## Index

- [**Concepts**](#concepts)
  - [**Accounts**](#accounts)
  - [**Transactions**](#transactions)
  - [**Extras**](#extras)
- [**Execution**](#execution)
  - [**Prerequisites**](#prerequisites)
  - [**How to Run**](#how-to-run)
  - [**How to Stop**](#how-to-stop)
  - [**How to Run Tests**](#how-to-run-tests)

## Concepts

This project uses:

- [Go](https://go.dev/) as the coding language.
- [PostgreSQL](https://www.postgresql.org/) as the database.

```mermaid
flowchart LR
    api[Go app\nHTTP API]
    db[(PostgreSQL\ndatabase)]
    api <-... Write/Read Operations ...-> db
```

The expected inputs for the application should come as HTTP requests, for each specific resource/endpoint:

### Accounts

Create an Account:

```
POST /accounts

Request body:
{
  "document": "12345678900"
}
```

Reads an Account, based on an ID:

```
GET /accounts/{:accountID}
```

### Transactions

Create a Transaction:

```
POST /transactions

Request body:
{
  "account_id": 1,
  "operation_id": 1,
  "amount": 543.21
}
```

### Extras

Some notes about what I'd thought would be nice to the project.

- In the [extras](./extras) folder, you can find 2 files to use with [Postman](https://www.postman.com/): a Collection
and an Environment.
- I've noticed several flaws in the choice of how to work with databases: breaking interface/abstraction rules, usage of
transactions/commits/rollbacks, etc. In a next scenario, it will be a better choice to use something that already exists
instead of trying to reinvent the wheel :wink:

## Execution

### Prerequisites

This project offers a local environment to get this project up and running. This environment should be used **only** for
local development/execution.

In order to use this environment, first make sure you have the tools below correctly installed:

|TOOL|VERSION\*|
|:---|:---|
|[Git](https://git-scm.com/)\*\*|`2.34.1`|
|[Docker](https://www.docker.com/)|`24.0.6 (linux/amd64)`|
|[Docker Compose](https://docs.docker.com/compose/)|`Docker Compose version v2.21.0`|
|[Make](https://www.gnu.org/software/make/)|`GNU Make 4.3 (x86_64-pc-linux-gnu)`|

\*: these are suggested versions, for a Linux environment. It will (probably) have no issues in Mac OS, but It may (or
not) work using other versions/OSs.
\*\*: required to clone this repository :wink:

**Also important**: the local environment will try to use ports **8080** (API) and **5432** (database). Before start,
check if these ports are available. If not, you can select other ports by changing them in the
[docker-compose.yml](./docker-compose.yml) file.

### How to Run

Once you're OK with the steps on [**Prerequisites**](#prerequisites), you can follow the actions below:

1. Using your terminal, navigate to this project's root folder.
2. In the same terminal windown, copy the file [`env/.env.sample`](./env/.env.sample) to `env/.env`. You can run the
command below in your terminal:

```shell
cp env/.env.sample env/.env
```

3. Also in the same terminal window, run the command below:

```shell
make start
```

This command will:

1. Pull/build the required images.
2. Start the database server.
3. Create the application's database.
4. Start the application at [0.0.0.0:8080](http://0.0.0.0:8080).

### How to Stop

1. Using your terminal, navigate to this project's root folder.
2. In the same terminal window, run the command below:

```shell
make stop
```

### How to Run Tests

1. Using your terminal, navigate to this project's root folder.
2. In the same terminal window, run the command below:

```shell
make tests
```
