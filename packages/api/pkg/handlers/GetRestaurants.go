package handlers

import (
	"net/http"
	"restaurant-flow/pkg/db"
	"restaurant-flow/pkg/httputil"

	"github.com/jackskj/carta"
	"github.com/labstack/echo/v4"
)

type getRestaurantsResult struct {
	db.Restaurant
	RestaurantId     *int                `json:"restaurantId" db:"restaurant_id" example:"0"` // Required for mapping
	GoogleRestaurant db.GoogleRestaurant `json:"googleRestaurant"`
}

// GetRestaurants
//
//	@Summary	get all restaurants
//	@Produce	json
//	@Success	200	{array}		getRestaurantsResult
//	@Failure	500	{object}	httputil.HTTPError
//	@Router		/restaurants [get]
func (handler Handler) GetRestaurants(context echo.Context) error {
	rows, err := handler.DB.Query(`
		select
			Res.restaurant_id,
			Res.created_at,
			Res.updated_at,

			GRes.google_restaurant_id,
			GRes.name,
			GRes.description,
			GRes.address,
			GRes.phone,
			GRes.website,
			GRes.google_url,
			GRes.avg_rating,
			GRes.business_status,
			GRes.price_level,

			GRes.supports_curbside_pickup,
			GRes.supports_delivery,
			GRes.supports_dine_in,
			GRes.supports_reservations,
			GRes.supports_takeout,

			GRes.serves_breakfast,
			GRes.serves_brunch,
			GRes.serves_dinner,
			GRes.serves_lunch,
			GRes.serves_vegetarian_food,
			GRes.serves_wine,
			GRes.serves_beer,
			GRes.serves_cocktails,
			GRes.serves_coffee,
			GRes.serves_dessert,

			GRes.good_for_groups,
			GRes.good_for_watching_sports,
			GRes.has_outdoor_seating,
			GRes.has_restroom,

			GRes.accepts_credit_cards,
			GRes.accepts_debit_cards,
			GRes.accepts_cash_only,
			GRes.accepts_nfc,

			GRes.wheelchair_accessible_entrance,
			GRes.wheelchair_accessible_seating,

			GRes.place_id,

			ST_X(GRes.coords) as coords_lat,
			ST_Y(GRes.coords) as coords_lng,
			ST_X(GRes.viewport_high) as viewport_high_lat,
			ST_Y(GRes.viewport_high) as viewport_high_lng,
			ST_X(GRes.viewport_low) as viewport_low_lat,
			ST_Y(GRes.viewport_low) as viewport_low_lng
		from Restaurant as Res
		inner join GoogleRestaurant GRes
		on GRes.google_restaurant_id = Res.google_restaurant_id;
		`)

	if err != nil {
		return httputil.NewError(context, http.StatusInternalServerError, err)
	}

	data := []getRestaurantsResult{}

	carta.Map(rows, &data)

	return context.JSON(http.StatusOK, data)
}
