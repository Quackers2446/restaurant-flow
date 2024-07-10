-- Call with:
-- mysql -P 3306 -h 127.0.0.1 -u user --password=password restaurantFlow  < test-production.sql > test-production.out

set @order = "";
set @orderCol1 = "name";
set @start = 0;
set @limit = 50;

prepare getRestaurants from 'select
    restaurant.restaurant_id, restaurant.created_at, restaurant.updated_at, restaurant.google_restaurant_id,
    google_restaurant.google_restaurant_id, google_restaurant.description, google_restaurant.phone, google_restaurant.website, google_restaurant.google_url, google_restaurant.business_status, google_restaurant.price_level, google_restaurant.supports_curbside_pickup, google_restaurant.supports_delivery, google_restaurant.supports_dine_in, google_restaurant.supports_reservations, google_restaurant.place_id, google_restaurant.updated_at, google_restaurant.supports_takeout, google_restaurant.serves_breakfast, google_restaurant.serves_brunch, google_restaurant.serves_dinner, google_restaurant.serves_lunch, google_restaurant.serves_vegetarian_food, google_restaurant.serves_wine, google_restaurant.serves_beer, google_restaurant.serves_cocktails, google_restaurant.serves_coffee, google_restaurant.serves_dessert, google_restaurant.good_for_groups, google_restaurant.good_for_watching_sports, google_restaurant.has_outdoor_seating, google_restaurant.has_restroom, google_restaurant.accepts_credit_cards, google_restaurant.accepts_debit_cards, google_restaurant.accepts_cash_only, google_restaurant.accepts_nfc, google_restaurant.wheelchair_accessible_entrance, google_restaurant.wheelchair_accessible_seating, google_restaurant.name, google_restaurant.avg_rating,
    location.location_id, location.address, location.lat, location.lng, location.viewport_high_lat, location.viewport_high_lng, location.viewport_low_lat, location.viewport_low_lng, location.google_restaurant_id
from restaurant
inner join google_restaurant on google_restaurant.google_restaurant_id = restaurant.google_restaurant_id
inner join `location` on `location`.google_restaurant_id = restaurant.google_restaurant_id
order by
    case when ? = "desc" then (
        case
            when ? = "name" then `google_restaurant`.`name`
            when ? = "updated_at" then `google_restaurant`.`updated_at`
            when ? = "created_at" then `created_at`
            when ? = "avg_rating" then `avg_rating`
            else `google_restaurant`.`name`
        end
    ) end desc,
    case when ? != "desc" then (
        case
            when ? = "name" then `google_restaurant`.`name`
            when ? = "updated_at" then `google_restaurant`.`updated_at`
            when ? = "created_at" then `created_at`
            when ? = "avg_rating" then `avg_rating`
            else `google_restaurant`.`name`
        end
    ) end asc,
    restaurant.restaurant_id asc
limit ?, ?';


execute getRestaurants using @order, @orderCol1, @orderCol1, @orderCol1, @orderCol1, @order, @orderCol1, @orderCol1, @orderCol1, @orderCol1, @start, @limit;

set @id1 = 34;
set @id2 = 17;
set @id3 = 50;
set @id4 = 45;
set @id5 = 18;
set @id6 = 57;
set @id7 = 26;
set @id8 = 54;
set @id9 = 9;
set @id10 = 47;
set @id11 = 30;
set @id12 = 38;
set @id13 = 11;
set @id14 = 15;
set @id15 = 5;
set @id16 = 29;
set @id17 = 23;
set @id18 = 60;
set @id19 = 42;
set @id20 = 13;

prepare getTags from 'select tag.tag_id, tag.name, tag.restaurant_id from tag where tag.restaurant_id in (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)';

prepare getOpeningHours from 'select opening_hours.opening_hours_id, opening_hours.type, opening_hours.google_restaurant_id, opening_period.opening_period_id, opening_period.close_day, opening_period.opening_hours_id, opening_period.open_day, opening_period.open_time, opening_period.close_time
from opening_hours
inner join opening_period on opening_period.opening_hours_id = opening_hours.opening_hours_id
where opening_hours.google_restaurant_id in (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)';

execute getTags using @id1, @id2, @id3, @id4, @id5, @id6, @id7, @id8, @id9, @id10, @id11, @id12, @id13, @id14, @id15, @id16, @id17, @id18, @id19, @id20;
execute getOpeningHours using @id1, @id2, @id3, @id4, @id5, @id6, @id7, @id8, @id9, @id10, @id11, @id12, @id13, @id14, @id15, @id16, @id17, @id18, @id19, @id20;

set @reviewUserId = "00000000-0000-0000-0000-000000000000";

insert into review (`rating`, `comments`, `restaurant_id`, `user_id`)
select
    rand() * 10,
    concat("COMMENT ", restaurant.restaurant_id),
    restaurant.restaurant_id,
    unhex(replace(@reviewUserId, '-', ''))
from restaurant
on duplicate key update review_id = review_id;

set @orderCol2 = "created_at";

prepare getUserReviews from 'select review.*
from review
where user_id = unhex(replace(?,\'-\',\'\')) -- Accept textual form of UUID
order by
    case when ? = "desc" then (
        case
            when ? = "rating" then `rating`
            when ? = "created_at" then `created_at`
            else `rating`
        end
    ) end desc,
    case when ? != "desc" then (
        case
            when ? = "rating" then `rating`
            when ? = "created_at" then `created_at`
            else `rating`
        end
    ) end asc,
    review_id asc
limit ?, ?';

execute getUserReviews using @reviewUserId, @order, @orderCol2, @orderCol2, @order, @orderCol2, @orderCol2, @start, @limit;

set @reviewRestaurantId = 1;

prepare getReviews from 'select review.review_id, review.comments, review.restaurant_id, review.user_id, review.created_at, review.updated_at, review.rating, review.is_anonymous
from review
where restaurant_id = ?
order by
	case when ? = "desc" then (
    	case
        	when ? = "rating" then `rating`
        	when ? = "created_at" then `created_at`
        	else `rating`
    	end
	) end desc,
	case when ? != "desc" then (
    	case
        	when ? = "rating" then `rating`
        	when ? = "created_at" then `created_at`
        	else `rating`
    	end
	) end asc,
	review_id asc
limit ?, ?';

execute getReviews using @reviewRestaurantId, @order, @orderCol2, @orderCol2, @order, @orderCol2, @orderCol2, @start, @limit;

set @lat = 43.472587;
set @lng = -80.537681;
set @rad = 50;

prepare getRestaurantsInArea from 'select
	restaurant.restaurant_id, restaurant.created_at, restaurant.updated_at, restaurant.google_restaurant_id,
	google_restaurant.google_restaurant_id, google_restaurant.description, google_restaurant.phone, google_restaurant.website, google_restaurant.google_url, google_restaurant.business_status, google_restaurant.price_level, google_restaurant.supports_curbside_pickup, google_restaurant.supports_delivery, google_restaurant.supports_dine_in, google_restaurant.supports_reservations, google_restaurant.place_id, google_restaurant.updated_at, google_restaurant.supports_takeout, google_restaurant.serves_breakfast, google_restaurant.serves_brunch, google_restaurant.serves_dinner, google_restaurant.serves_lunch, google_restaurant.serves_vegetarian_food, google_restaurant.serves_wine, google_restaurant.serves_beer, google_restaurant.serves_cocktails, google_restaurant.serves_coffee, google_restaurant.serves_dessert, google_restaurant.good_for_groups, google_restaurant.good_for_watching_sports, google_restaurant.has_outdoor_seating, google_restaurant.has_restroom, google_restaurant.accepts_credit_cards, google_restaurant.accepts_debit_cards, google_restaurant.accepts_cash_only, google_restaurant.accepts_nfc, google_restaurant.wheelchair_accessible_entrance, google_restaurant.wheelchair_accessible_seating, google_restaurant.name, google_restaurant.avg_rating,
	location.location_id, location.address, location.lat, location.lng, location.viewport_high_lat, location.viewport_high_lng, location.viewport_low_lat, location.viewport_low_lng, location.google_restaurant_id,
	ST_Distance(
    	ST_SRID(Point(`location`.lat, `location`.lng), 4326),
    	ST_SRID(Point(?, ?), 4326),
    	"metre"
	) as distance
from restaurant
inner join google_restaurant on google_restaurant.google_restaurant_id = restaurant.google_restaurant_id
inner join `location` on `location`.google_restaurant_id = restaurant.google_restaurant_id
where ST_Distance(
	ST_SRID(Point(`location`.lat, `location`.lng), 4326),
	ST_SRID(Point(?, ?), 4326),
	"metre"
) < ?
order by distance asc
limit 100';

execute getRestaurantsInArea using @lat, @lng, @lat, @lng, @rad;

deallocate prepare getRestaurants;
deallocate prepare getTags;
deallocate prepare getOpeningHours;
deallocate prepare getReviews;
deallocate prepare getUserReviews;
