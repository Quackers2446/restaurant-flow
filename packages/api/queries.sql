-- SQL queries for SQLC generation go here. Note that these queries are run through a preprocessor, so there may be
-- functions that are not defined in MySQL used here.
-- Docs: https://docs.sqlc.dev/en/latest/tutorials/getting-started-mysql.html
-- Note: we cannot just left join and expect SQLC to handle one-to-many relationships properly
--       see https://github.com/sqlc-dev/sqlc/issues/2348
--       and https://github.com/sqlc-dev/sqlc/discussions/2643

-- name: GetRestaurants :many
select
    Restaurant.*,
    sqlc.embed(GoogleRestaurant),
    sqlc.embed(Location)
from Restaurant
inner join GoogleRestaurant on GoogleRestaurant.google_restaurant_id = Restaurant.google_restaurant_id
inner join `Location` on `Location`.google_restaurant_id = Restaurant.google_restaurant_id
order by
    case when sqlc.arg("order") = "desc" then (
        case
            when sqlc.arg("order_by") = "name" then `GoogleRestaurant`.`name`
            when sqlc.arg("order_by") = "updated_at" then `GoogleRestaurant`.`updated_at`
            when sqlc.arg("order_by") = "created_at" then `created_at`
            when sqlc.arg("order_by") = "avg_rating" then `avg_rating`
            else `GoogleRestaurant`.`name`
        end
    ) end desc,
    case when sqlc.arg("order") != "desc" then (
        case
            when sqlc.arg("order_by") = "name" then `GoogleRestaurant`.`name`
            when sqlc.arg("order_by") = "updated_at" then `GoogleRestaurant`.`updated_at`
            when sqlc.arg("order_by") = "created_at" then `created_at`
            when sqlc.arg("order_by") = "avg_rating" then `avg_rating`
            else `GoogleRestaurant`.`name`
        end
    ) end asc,
    Restaurant.restaurant_id asc
limit ?, ?;

-- name: GetRestaurant :one
select
    Restaurant.*,
    sqlc.embed(GoogleRestaurant),
    sqlc.embed(Location)
from Restaurant
inner join GoogleRestaurant on GoogleRestaurant.google_restaurant_id = Restaurant.google_restaurant_id
inner join `Location` on `Location`.google_restaurant_id = Restaurant.google_restaurant_id
where Restaurant.restaurant_id = ?
limit 1;

-- name: GetTags :many
select Tag.* from Tag where Tag.restaurant_id in (sqlc.slice("restaurant_ids"));

-- name: GetOpeningHours :many
select OpeningHours.*, sqlc.embed(OpeningPeriod)
from OpeningHours
inner join OpeningPeriod on OpeningPeriod.opening_hours_id = OpeningHours.opening_hours_id
where OpeningHours.google_restaurant_id in (sqlc.slice("google_restaurant_ids"));
