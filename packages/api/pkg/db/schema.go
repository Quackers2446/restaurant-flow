package db

import "time"

type GoogleRestaurant struct {
	GoogleRestaurantId *int       `json:"-" db:"google_restaurant_id" example:"0"`
	UpdatedAt          *time.Time `json:"updatedAt" db:"updated_at" example:"2000-01-01T00:00:00Z"`

	Name           *string  `json:"name" db:"name" example:"McDonald's"`
	Description    *string  `json:"description" db:"description" example:"Classic, long-running fast-food chain known for its burgers* & fries."`
	Address        *string  `json:"address" db:"address" example:"362 King St N, Waterloo, ON N2J 2Z2, Canada"`
	Phone          *string  `json:"phone" db:"phone" example:"(519) 772-0790"`
	Website        *string  `json:"website" db:"website" example:"https://www.mcdonalds.com/ca/en-ca/restaurant-locator.html?y_source=1*_MTQ1MTk5MzUtNzE1LWxvY2F0aW9uLndlYnNpdGU%3D"`
	GoogleUrl      *string  `json:"googleUrl" db:"google_url" example:"https://maps.google.com/?cid=217835057852927681"`
	AvgRating      *float32 `json:"avgRating" db:"avg_rating" example:"3.7"`
	BusinessStatus *string  `json:"businessStatus" db:"business_status" example:"Operational"`
	PriceLevel     *string  `json:"priceLevel" db:"price_level" example:"Moderate"`

	SupportsCurbsidePickup *bool `json:"supportsCurbsidePickup" db:"supports_curbside_pickup"`
	SupportsDelivery       *bool `json:"supportsDelivery" db:"supports_delivery"`
	SupportsDineIn         *bool `json:"supportsDineIn" db:"supports_dine_in"`
	SupportsReservations   *bool `json:"supportsReservations" db:"supports_reservations"`
	SupportsTakeout        *bool `json:"supportsTakeout" db:"supports_takeout"`

	ServesBreakfast      *bool `json:"servesBreakfast" db:"serves_breakfast"`
	ServesBrunch         *bool `json:"servesBrunch" db:"serves_brunch"`
	ServesDinner         *bool `json:"servesDinner" db:"serves_dinner"`
	ServesLunch          *bool `json:"servesLunch" db:"serves_lunch"`
	ServesVegetarianFood *bool `json:"servesVegetarianFood" db:"serves_vegetarian_food"`
	ServesWine           *bool `json:"servesWine" db:"serves_wine"`
	ServesBeer           *bool `json:"servesBeer" db:"serves_beer"`
	ServesCocktails      *bool `json:"servesCocktails" db:"serves_cocktails"`
	ServesCoffee         *bool `json:"servesCoffee" db:"serves_coffee"`
	ServesDessert        *bool `json:"servesDessert" db:"serves_dessert"`

	GoodForGroups         *bool `json:"goodForGroups" db:"good_for_groups"`
	GoodForWatchingSports *bool `json:"goodForWatchingSports" db:"good_for_watching_sports"`
	HasOutdoorSeating     *bool `json:"hasOutdoorSeating" db:"has_outdoor_seating"`
	HasRestroom           *bool `json:"hasRestroom" db:"has_restroom"`

	AcceptsCreditCards *bool `json:"acceptsCredit_cards" db:"accepts_credit_cards"`
	AcceptsDebitCards  *bool `json:"acceptsDebit_cards" db:"accepts_debit_cards"`
	AcceptsCashOnly    *bool `json:"acceptsCash_only" db:"accepts_cash_only"`
	AcceptsNfc         *bool `json:"acceptsNfc" db:"accepts_nfc"`

	WheelchairAccessibleEntrance *bool `json:"wheelchairAccessibleEntrance" db:"wheelchair_accessible_entrance"`
	WheelchairAccessibleSeating  *bool `json:"wheelchairAccessibleSeating" db:"wheelchair_accessible_seating"`

	PlaceId         *string `json:"placeId" db:"place_id" example:"ChIJp6htNPLzK4gRwU7zutTnBQM"`
	CoordsLat       float64 `json:"coordsLat" db:"coords_lat" example:"43.481748499999995"`
	CoordsLng       float64 `json:"coordsLng" db:"coords_lng" example:"-80.52557019999999"`
	ViewportHighLat float64 `json:"viewportHighLat" db:"viewport_high_lat" example:"-80.52418676970849"`
	ViewportHighLng float64 `json:"viewportHighLng" db:"viewport_high_lng" example:"43.483284130291494"`
	ViewportLowLat  float64 `json:"viewportLowLat" db:"viewport_low_lat" example:"43.480586169708495"`
	ViewportLowLng  float64 `json:"viewportLowLng" db:"viewport_low_lng" example:"-80.5268847302915"`
}

type Restaurant struct {
	RestaurantId *int       `json:"restaurantId" db:"restaurant_id" example:"0"`
	CreatedAt    *time.Time `json:"createdAt" db:"created_at" example:"2000-01-01T00:00:00Z"`
	UpdatedAt    *time.Time `json:"updatedAt" db:"updated_at" example:"2000-01-01T00:00:00Z"`
}
