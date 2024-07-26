# Restaurant Flow

## About

Restaurant Flow (UW Flow for food) is an application built for students and by students. Restaurant Flow features restaurant reviews by students, but it is not just a Yelp or Google Reviews duplicate. The value proposition of Restaurant Flow is that students know these are reviews from other students, who understand their needs and wants.

Functionalities: In addition to the main value proposition of Restaurant Flow, we also have a number of other functionalities we would like to support. These include:
Categorizing meals and restaurants into different categories such as 
- “comfort food” (for stressful study periods)
- “study snacks” (for when you want to reward yourself after a hard study session)
- “cheap eats” (for when you are on a budget)
- “grab and go” (for when you are in a time crunch and need something quick)
- Featuring many on campus restaurants such as those in SLC and student cafeterias, which may not be on Google Maps
- Filtering by different cuisines
- Listing meals a user has bought within each review and popular meals
- Star ratings
- Incorporating Google Maps to retrieve restaurant information and to allow users to see their location in relation to the location of other restaurants
- Enabling people to join groups and eat together


## Setup

### Things to Install

1. [Docker](https://www.docker.com/get-started/)
2. [Docker Compose](https://docs.docker.com/compose/install/)
3. [Node](https://nodejs.org/en/download/package-manager) and [pnpm](https://pnpm.io/installation#using-npm) with `npm i -g pnpm`
4. To interact with the database, install at least one of [MyCLI](https://www.mycli.net/install), [MySQL CLI](https://dev.mysql.com/doc/refman/8.0/en/mysql.html), [MySQL Workbench](https://dev.mysql.com/downloads/workbench/), or [DBeaver](https://dbeaver.io/download/).
5. [Go](https://go.dev/doc/install), preferably 1.22 but other versions _might_ work
6. [SQLC CLI](https://docs.sqlc.dev/en/latest/overview/install.html)

### What to run

Run `docker-compose up` to host the database. See the `README` files of each service for further instructions.

To seed the database with sample data, go to [packages/seed](./packages/seed/).

## Database

Connect with `mycli -P 3306 -u user -p password` or if you don't have MyCLI instaled run `mysql -P 3306 -h 127.0.0.1 -u user --password=password`.

DB migrations can be found at [packages/api/migrations](./packages/api/migrations/).

Queries can be found at [packages/api/queries.sql](./packages/api/queries.sql). Queries are pre-processed with SQLC. See API README for more details.
