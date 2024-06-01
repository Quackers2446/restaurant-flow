# Restaurant Flow

## Setup

### Things to Install

1. [Docker](https://www.docker.com/get-started/)
2. [Docker Compose](https://docs.docker.com/compose/install/)
3. [Node](https://nodejs.org/en/download/package-manager) and [pnpm](https://pnpm.io/installation#using-npm) with `npm i -g pnpm`
4. To interact with the database, install at least one of [MyCLI](https://www.mycli.net/install), [MySQL CLI](https://dev.mysql.com/doc/refman/8.0/en/mysql.html), or [MySQL Workbench](https://dev.mysql.com/downloads/workbench/)
5. [Go](https://go.dev/doc/install), preferably 1.22 but other versions _might_ work

### What to run

Run `docker-compose up` to host the database. See the `README` files of each service for further instructions.

## Database

Connect with `mycli -P 3306 -u user -p password` or if you don't have MyCLI instaled run `mysql -P 3306 -u user -p` and enter password `password`.
