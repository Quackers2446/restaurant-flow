-- Call with:
-- mysql -P 3306 -h 127.0.0.1 -u user --password=password restaurantFlow  < test-sample.sql > test-sample.out

set @order = "";
set @orderCol1 = "name";
set @start = 0;
set @limit = 20;

prepare getRestaurants from 'select
    restaurant.restaurant_id, restaurant.created_at, restaurant.updated_at, restaurant.google_restaurant_id,
    googlerestaurant.google_restaurant_id, googlerestaurant.description, googlerestaurant.phone, googlerestaurant.website, googlerestaurant.google_url, googlerestaurant.business_status, googlerestaurant.price_level, googlerestaurant.supports_curbside_pickup, googlerestaurant.supports_delivery, googlerestaurant.supports_dine_in, googlerestaurant.supports_reservations, googlerestaurant.place_id, googlerestaurant.updated_at, googlerestaurant.supports_takeout, googlerestaurant.serves_breakfast, googlerestaurant.serves_brunch, googlerestaurant.serves_dinner, googlerestaurant.serves_lunch, googlerestaurant.serves_vegetarian_food, googlerestaurant.serves_wine, googlerestaurant.serves_beer, googlerestaurant.serves_cocktails, googlerestaurant.serves_coffee, googlerestaurant.serves_dessert, googlerestaurant.good_for_groups, googlerestaurant.good_for_watching_sports, googlerestaurant.has_outdoor_seating, googlerestaurant.has_restroom, googlerestaurant.accepts_credit_cards, googlerestaurant.accepts_debit_cards, googlerestaurant.accepts_cash_only, googlerestaurant.accepts_nfc, googlerestaurant.wheelchair_accessible_entrance, googlerestaurant.wheelchair_accessible_seating, googlerestaurant.name, googlerestaurant.avg_rating,
    location.location_id, location.address, location.lat, location.lng, location.viewport_high_lat, location.viewport_high_lng, location.viewport_low_lat, location.viewport_low_lng, location.google_restaurant_id
from Restaurant
inner join GoogleRestaurant on GoogleRestaurant.google_restaurant_id = Restaurant.google_restaurant_id
inner join `Location` on `Location`.google_restaurant_id = Restaurant.google_restaurant_id
order by
    case when ? = "desc" then (
        case
            when ? = "name" then `GoogleRestaurant`.`name`
            when ? = "updated_at" then `GoogleRestaurant`.`updated_at`
            when ? = "created_at" then `created_at`
            when ? = "avg_rating" then `avg_rating`
            else `GoogleRestaurant`.`name`
        end
    ) end desc,
    case when ? != "desc" then (
        case
            when ? = "name" then `GoogleRestaurant`.`name`
            when ? = "updated_at" then `GoogleRestaurant`.`updated_at`
            when ? = "created_at" then `created_at`
            when ? = "avg_rating" then `avg_rating`
            else `GoogleRestaurant`.`name`
        end
    ) end asc,
    Restaurant.restaurant_id asc
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

prepare getTags from 'select tag.tag_id, tag.name, tag.restaurant_id from Tag where Tag.restaurant_id in (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)';

prepare getOpeningHours from 'select openinghours.opening_hours_id, openinghours.type, openinghours.google_restaurant_id, openingperiod.opening_period_id, openingperiod.close_day, openingperiod.opening_hours_id, openingperiod.open_day, openingperiod.open_time, openingperiod.close_time
from OpeningHours
inner join OpeningPeriod on OpeningPeriod.opening_hours_id = OpeningHours.opening_hours_id
where OpeningHours.google_restaurant_id in (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)';

execute getTags using @id1, @id2, @id3, @id4, @id5, @id6, @id7, @id8, @id9, @id10, @id11, @id12, @id13, @id14, @id15, @id16, @id17, @id18, @id19, @id20;
execute getOpeningHours using @id1, @id2, @id3, @id4, @id5, @id6, @id7, @id8, @id9, @id10, @id11, @id12, @id13, @id14, @id15, @id16, @id17, @id18, @id19, @id20;

set @reviewRestaurantId = 1;
set @orderCol2 = "created_at";

prepare getReviews from 'select review.review_id, review.comments, review.restaurant_id, review.user_id, review.created_at, review.updated_at, review.rating, review.is_anonymous
from Review
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
	googlerestaurant.google_restaurant_id, googlerestaurant.description, googlerestaurant.phone, googlerestaurant.website, googlerestaurant.google_url, googlerestaurant.business_status, googlerestaurant.price_level, googlerestaurant.supports_curbside_pickup, googlerestaurant.supports_delivery, googlerestaurant.supports_dine_in, googlerestaurant.supports_reservations, googlerestaurant.place_id, googlerestaurant.updated_at, googlerestaurant.supports_takeout, googlerestaurant.serves_breakfast, googlerestaurant.serves_brunch, googlerestaurant.serves_dinner, googlerestaurant.serves_lunch, googlerestaurant.serves_vegetarian_food, googlerestaurant.serves_wine, googlerestaurant.serves_beer, googlerestaurant.serves_cocktails, googlerestaurant.serves_coffee, googlerestaurant.serves_dessert, googlerestaurant.good_for_groups, googlerestaurant.good_for_watching_sports, googlerestaurant.has_outdoor_seating, googlerestaurant.has_restroom, googlerestaurant.accepts_credit_cards, googlerestaurant.accepts_debit_cards, googlerestaurant.accepts_cash_only, googlerestaurant.accepts_nfc, googlerestaurant.wheelchair_accessible_entrance, googlerestaurant.wheelchair_accessible_seating, googlerestaurant.name, googlerestaurant.avg_rating,
	location.location_id, location.address, location.lat, location.lng, location.viewport_high_lat, location.viewport_high_lng, location.viewport_low_lat, location.viewport_low_lng, location.google_restaurant_id,
	ST_Distance(
    	ST_SRID(Point(`Location`.lat, `Location`.lng), 4326),
    	ST_SRID(Point(?, ?), 4326),
    	"metre"
	) as distance
from Restaurant
inner join GoogleRestaurant on GoogleRestaurant.google_restaurant_id = Restaurant.google_restaurant_id
inner join `Location` on `Location`.google_restaurant_id = Restaurant.google_restaurant_id
where ST_Distance(
	ST_SRID(Point(`Location`.lat, `Location`.lng), 4326),
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
