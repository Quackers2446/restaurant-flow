-- SQL queries for SQLC generation go here. Note that these queries are run through a preprocessor, so there may be
-- functions that are not defined in MySQL used here.
-- Docs: https://docs.sqlc.dev/en/latest/tutorials/getting-started-mysql.html

-- name: GetRestaurants :many
select
    Restaurant.restaurant_id,
    Restaurant.created_at,
    Restaurant.updated_at,
    sqlc.embed(GoogleRestaurant),
    sqlc.embed(Location)
from Restaurant
inner join GoogleRestaurant on GoogleRestaurant.google_restaurant_id = Restaurant.google_restaurant_id
inner join `Location` on `Location`.google_restaurant_id = Restaurant.google_restaurant_id;
