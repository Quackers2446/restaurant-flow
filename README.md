# Restaurant Flow

## Setup

### Things to Install

1. [Docker](https://www.docker.com/get-started/)
2. [Docker Compose](https://docs.docker.com/compose/install/)
3. [Node](https://nodejs.org/en/download/package-manager) and [pnpm](https://pnpm.io/installation#using-npm)
4. To interact with the database, install at least one of [MyCLI](https://www.mycli.net/install), [MySQL CLI](https://dev.mysql.com/doc/refman/8.0/en/mysql.html), or [MySQL Workbench](https://dev.mysql.com/downloads/workbench/)

### What to run

1. In `client`, run `pnpm install`
    - We will use `pnpm` because it is better and faster in every way
2. Run `docker-compose up`
    - For the moment, this just gets the database running, but in the future it will encompass all services

## Database

Connect with `mycli -P 3306 -u user -p password` or if you don't have MyCLI instaled run `mysql -P 3306 -u user -p` and enter password `password`.
