-- SQL queries for SQLC generation go here. Note that these queries are run through a preprocessor, so there may be
-- functions that are not defined in MySQL used here.
-- Docs: https://docs.sqlc.dev/en/latest/tutorials/getting-started-mysql.html
-- Note: we cannot just left join and expect SQLC to handle one-to-many relationships properly
--       see https://github.com/sqlc-dev/sqlc/issues/3144
--       and https://github.com/sqlc-dev/sqlc/issues/2348
--       and https://github.com/sqlc-dev/sqlc/discussions/2643

-- name: GetRestaurants :many
select
    restaurant.*,
    sqlc.embed(google_restaurant),
    sqlc.embed(location)
from restaurant
inner join google_restaurant on google_restaurant.google_restaurant_id = restaurant.google_restaurant_id
inner join `location` on `location`.google_restaurant_id = restaurant.google_restaurant_id
order by
    case when sqlc.arg("order") = "desc" then (
        case
            when sqlc.arg("order_by") = "name" then `google_restaurant`.`name`
            when sqlc.arg("order_by") = "updated_at" then `google_restaurant`.`updated_at`
            when sqlc.arg("order_by") = "created_at" then `created_at`
            when sqlc.arg("order_by") = "avg_rating" then `avg_rating`
            else `google_restaurant`.`name`
        end
    ) end desc,
    case when sqlc.arg("order") != "desc" then (
        case
            when sqlc.arg("order_by") = "name" then `google_restaurant`.`name`
            when sqlc.arg("order_by") = "updated_at" then `google_restaurant`.`updated_at`
            when sqlc.arg("order_by") = "created_at" then `created_at`
            when sqlc.arg("order_by") = "avg_rating" then `avg_rating`
            else `google_restaurant`.`name`
        end
    ) end asc,
    restaurant.restaurant_id asc
limit ?, ?;

-- name: GetRestaurantsInArea :many
select
    restaurant.*,
    sqlc.embed(google_restaurant),
    sqlc.embed(location),
    ST_Distance(
        ST_SRID(Point(`location`.lat, `location`.lng), 4326),
        ST_SRID(Point(sqlc.arg("lat"), sqlc.arg("lng")), 4326),
        "metre"
    ) as distance
from restaurant
inner join google_restaurant on google_restaurant.google_restaurant_id = restaurant.google_restaurant_id
inner join `location` on `location`.google_restaurant_id = restaurant.google_restaurant_id
where ST_Distance(
    ST_SRID(Point(`location`.lat, `location`.lng), 4326),
    ST_SRID(Point(sqlc.arg("lat"), sqlc.arg("lng")), 4326),
    "metre"
) < sqlc.arg("radius")  -- Unfortunately there's not a great way to DRY this out
order by distance asc
limit 100; -- Probably a reasonable default (can change later if needed)

-- TODO: combine this with GetRestaurants later
-- name: GetRestaurantsSearch :many
select restaurant.*,
    sqlc.embed(google_restaurant),
    sqlc.embed(location),
    concat(sqlc.arg("search")) as search_query
from restaurant
inner join google_restaurant on google_restaurant.google_restaurant_id = restaurant.google_restaurant_id
inner join `location` on `location`.google_restaurant_id = restaurant.google_restaurant_id
having match (google_restaurant.name, google_restaurant.description) against (search_query in natural language mode);
-- TODO: see if there's a way to use where instead of having because having is slower
-- Not trivial: see https://github.com/sqlc-dev/sqlc/issues/3091

-- name: GetRestaurant :one
select
    restaurant.*,
    sqlc.embed(google_restaurant),
    sqlc.embed(location)
from restaurant
inner join google_restaurant on google_restaurant.google_restaurant_id = restaurant.google_restaurant_id
inner join `location` on `location`.google_restaurant_id = restaurant.google_restaurant_id
where restaurant.restaurant_id = ?
limit 1;

-- name: GetTags :many
select tag.* from tag where tag.restaurant_id in (sqlc.slice("restaurant_ids"));

-- name: GetOpeningHours :many
select opening_hours.*, sqlc.embed(opening_period)
from opening_hours
inner join opening_period on opening_period.opening_hours_id = opening_hours.opening_hours_id
where opening_hours.google_restaurant_id in (sqlc.slice("google_restaurant_ids"));

-- REVIEW CRUD

-- name: CreateReview :execlastid
insert into review set
    rating=?,
    comments=sqlc.narg("comments"),
    restaurant_id=?,
    user_id=unhex(replace(sqlc.arg("user_id"),'-','')), -- Accept textual form of UUID
    is_anonymous=sqlc.narg("is_anonymous");

-- name: CreateTag :execlastid
insert into tag set
    restaurant_id=?,
    name=sqlc.narg("name");

-- name: GetRestaurantReviews :many
select review.*, user.username
from review
inner join user on user.user_id = review.user_id
where restaurant_id = ?
order by
    case when sqlc.arg("order") = "desc" then (
        case
            when sqlc.arg("order_by") = "rating" then `rating`
            when sqlc.arg("order_by") = "created_at" then review.`created_at`
            else `rating`
        end
    ) end desc,
    case when sqlc.arg("order") != "desc" then (
        case
            when sqlc.arg("order_by") = "rating" then `rating`
            when sqlc.arg("order_by") = "created_at" then review.`created_at`
            else `rating`
        end
    ) end asc,
    review_id asc
limit ?, ?;

-- name: GetUserReviews :many
select review.*
from review
where user_id = unhex(replace(sqlc.arg("user_id"),'-','')) -- Accept textual form of UUID
order by
    case when sqlc.arg("order") = "desc" then (
        case
            when sqlc.arg("order_by") = "rating" then `rating`
            when sqlc.arg("order_by") = "created_at" then review.`created_at`
            else `rating`
        end
    ) end desc,
    case when sqlc.arg("order") != "desc" then (
        case
            when sqlc.arg("order_by") = "rating" then `rating`
            when sqlc.arg("order_by") = "created_at" then review.`created_at`
            else `rating`
        end
    ) end asc,
    review_id asc
limit ?, ?;

-- name: GetReview :one
select review.*, user.username
from review
inner join user on user.user_id = review.user_id
where review_id = ?;

-- name: GetUpdatedReview :one
select review.* from review where restaurant_id=? and user_id=unhex(replace(sqlc.arg("user_id"),'-',''));

-- name: DeleteReview :exec
delete from review where restaurant_id=? and user_id=unhex(replace(sqlc.arg("user_id"),'-',''));

-- name: GetTag :one
select tag.* from tag where tag_id = ?;

-- name: UpdateReview :exec
update review set
    rating=?,
    comments=sqlc.narg("comments"),
    is_anonymous=sqlc.narg("is_anonymous")
where restaurant_id=? and user_id=unhex(replace(sqlc.arg("user_id"),'-',''));

-- name: CreateParty :execlastid
insert into party set
    max_members=?,
    `description`=?,
    restaurant_id=?,
    `time`=?;

-- name: JoinParty :execlastid
insert into party_members set
    party_id=?,
    user_id=unhex(replace(sqlc.arg("user_id"),'-',''));

-- name: GetParties :many
select party.*
from party
where party.`time` >= current_timestamp
order by
    case when sqlc.arg("order") = "desc" then (
        case
            when sqlc.arg("order_by") = "time" then `time`
            else `time`
        end
    ) end desc,
    case when sqlc.arg("order") != "desc" then (
        case
            when sqlc.arg("order_by") = "time" then `time`
            else `time`
        end
    ) end asc,
    party_id asc
limit ?, ?;

-- name: GetPartyDetails :one
select party.*
from party
where party_id=?;

-- name: GetPartyMembers :many
select party_members.*, user.username, (insert(
        insert(
            insert(
                insert(hex(party_members.user_id),9,0,'-'),
                14,0,'-'
            ),
            19,0,'-'
        ),
        24,0,'-')
    ) as user_id_text
from party_members
inner join user on user.user_id = party_members.user_id
where party_id in (sqlc.slice("party_ids"));

-- name: GetPartySize :one
select COUNT(*)
from party_members
where party_id=?;

-- name: LeaveParty :exec
delete from party_members where party_id=? and user_id=unhex(replace(sqlc.arg("user_id"),'-',''));

-- name: CreateUser :exec
insert into user set
    user_id=unhex(replace(sqlc.arg("user_id"),'-','')),
    name=?,
    username=?,
    email=?;
