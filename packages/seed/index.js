import "dotenv/config"
import mysql from "mysql2/promise"
import fetch from "node-fetch"
import {pick} from "@luke-zhang-04/utils"
import {semaphore} from "./semaphore.js"

const connection = await mysql.createConnection({
    host: "localhost",
    user: "user",
    password: "password",
    database: "restaurantFlow",
    port: 3306,
    namedPlaceholders: true,
    multipleStatements: true,
})

const businessStatusMap = {
    OPERATIONAL: "Operational",
    CLOSED_TEMPORARILY: "ClosedTemporarily",
    CLOSED_PERMANENTLY: "ClosedPermanently",
}

const priceLevelMap = {
    PRICE_LEVEL_FREE: "Free",
    PRICE_LEVEL_INEXPENSIVE: "Inexpensive",
    PRICE_LEVEL_MODERATE: "Moderate",
    PRICE_LEVEL_EXPENSIVE: "Expensive",
    PRICE_LEVEL_VERY_EXPENSIVE: "VeryExpensive",
}

const undefinedToNull = (obj) => {
    for (const key of Object.keys(obj)) {
        if (obj[key] === undefined) {
            obj[key] = null
        }
    }

    return obj
}

const insertPlace = async (place) => {
    // Holy crap
    // Bad design decision on my part: these should probably be tags instead of tables
    // https://stackoverflow.com/questions/9537710/is-there-a-way-to-use-on-duplicate-key-to-update-all-that-i-wanted-to-insert
    return await connection.query(
        /*sql*/ `
        start transaction;
        insert into GoogleRestaurant set
            name=:name,
            description=:description,
            phone=:phone,
            website=:website,
            google_url=:googleUrl,
            avg_rating=:avgRating,
            business_status=:businessStatus,
            price_level=:priceLevel,
            supports_curbside_pickup=:supportsCurbside_pickup,
            supports_delivery=:supportsDelivery,
            supports_dine_in=:supportsDine_in,
            supports_reservations=:supportsReservations,
            supports_takeout=:supportsTakeout,
            serves_breakfast=:servesBreakfast,
            serves_brunch=:servesBrunch,
            serves_dinner=:servesDinner,
            serves_lunch=:servesLunch,
            serves_vegetarian_food=:servesVegetarian_food,
            serves_wine=:servesWine,
            serves_beer=:servesBeer,
            serves_cocktails=:servesCocktails,
            serves_coffee=:servesCoffee,
            serves_dessert=:servesDessert,
            good_for_groups=:goodForGroups,
            good_for_watching_sports=:goodForWatchingSports,
            has_outdoor_seating=:hasOutdoorSeating,
            has_restroom=:hasRestroom,
            accepts_credit_cards=:acceptsCreditCards,
            accepts_debit_cards=:acceptsDebitCards,
            accepts_cash_only=:acceptsCashOnly,
            accepts_nfc=:acceptsNfc,
            wheelchair_accessible_entrance=:wheelchairAccessibleEntrance,
            wheelchair_accessible_seating=:wheelchairAccessibleSeating,
            place_id=:placeId
        as new
        on duplicate key update
            name=new.name,
            description=new.description,
            phone=new.phone,
            website=new.website,
            google_url=new.google_url,
            avg_rating=new.avg_rating,
            business_status=new.business_status,
            price_level=new.price_level,
            supports_curbside_pickup=new.supports_curbside_pickup,
            supports_delivery=new.supports_delivery,
            supports_dine_in=new.supports_dine_in,
            supports_reservations=new.supports_reservations,
            supports_takeout=new.supports_takeout,
            serves_breakfast=new.serves_breakfast,
            serves_brunch=new.serves_brunch,
            serves_dinner=new.serves_dinner,
            serves_lunch=new.serves_lunch,
            serves_vegetarian_food=new.serves_vegetarian_food,
            serves_wine=new.serves_wine,
            serves_beer=new.serves_beer,
            serves_cocktails=new.serves_cocktails,
            serves_coffee=new.serves_coffee,
            serves_dessert=new.serves_dessert,
            good_for_groups=new.good_for_groups,
            good_for_watching_sports=new.good_for_watching_sports,
            has_outdoor_seating=new.has_outdoor_seating,
            has_restroom=new.has_restroom,
            accepts_credit_cards=new.accepts_credit_cards,
            accepts_debit_cards=new.accepts_debit_cards,
            accepts_cash_only=new.accepts_cash_only,
            accepts_nfc=new.accepts_nfc,
            wheelchair_accessible_entrance=new.wheelchair_accessible_entrance,
            wheelchair_accessible_seating=new.wheelchair_accessible_seating;

        insert into Location set
            address=:address,
            lat=:coordLat,
            lng=:coordLng,
            viewport_high_lat=:highLat,
            viewport_high_lng=:highLng,
            viewport_low_lat=:lowLat,
            viewport_low_lng=:lowLng,
            google_restaurant_id=(select google_restaurant_id from GoogleRestaurant where place_id=:placeId)
        as new
        -- On duplicate key, no-op
        on duplicate key update google_restaurant_id=new.google_restaurant_id;

        insert into Restaurant set google_restaurant_id=(select google_restaurant_id from GoogleRestaurant where place_id=:placeId) as new
            -- On duplicate key, no-op
            on duplicate key update google_restaurant_id=new.google_restaurant_id;

        insert into OpeningHours set type="Main", google_restaurant_id=(select google_restaurant_id from GoogleRestaurant where place_id=:placeId) as new
            -- On duplicate key, no-op
            on duplicate key update google_restaurant_id=new.google_restaurant_id;

        ${
            place.regularOpeningHours?.periods
                .map(
                    (_, index) => /*sql*/ `
                    insert into OpeningPeriod set
                        open_day=:openDay${index},
                        open_time=:openTime${index},
                        close_day=:closeDay${index},
                        close_time=:closeTime${index},
                        opening_hours_id=(select opening_hours_id from OpeningHours where type="Main" and google_restaurant_id=(select google_restaurant_id from GoogleRestaurant where place_id=:placeId))
                    as new
                    on duplicate key update
                        open_day=new.open_day,
                        open_time=new.open_time,
                        close_day=new.close_day,
                        close_time=new.close_time;
                `,
                )
                ?.join("\n") ?? ""
        }
        commit;
    `,
        undefinedToNull({
            name: place.displayName.text,
            description: place.editorialSummary?.text,
            address: place.formattedAddress,
            phone: place.nationalPhoneNumber,
            website: place.websiteUri,
            googleUrl: place.googleMapsUri,
            avgRating: place.rating,
            businessStatus: businessStatusMap[place.businessStatus],
            priceLevel: priceLevelMap[place.priceLevel],

            supportsCurbsidePickup: place.curbsidePickup,
            supportsDelivery: place.delivery,
            supportsDineIn: place.dineIn,
            supportsReservations: place.reservable,
            supportsTakeout: place.takeout,

            ...pick(
                place,
                "servesBreakfast",
                "servesBrunch",
                "servesDinner",
                "servesLunch",
                "servesVegetarianFood",
                "servesWine",
                "servesBeer",
                "servesCocktails",
                "servesCoffee",
                "servesDessert",
                "goodForGroups",
                "goodForWatchingSports",
            ),

            hasOutdoorSeating: place.outdoorSeating,
            hasRestroom: place.restroom,
            acceptsCreditCards: place.paymentOptions.acceptsCreditCards,
            acceptsDebitCards: place.paymentOptions.acceptsDebitCards,
            acceptsCashOnly: place.paymentOptions.acceptsCashOnly,
            acceptsNfc: place.paymentOptions.acceptsNfc,

            wheelchairAccessibleEntrance: place.accessibilityOptions?.wheelchairAccessibleEntrance,
            wheelchairAccessibleSeating: place.accessibilityOptions?.wheelchairAccessibleSeating,
            placeId: place.id,
            coordLat: place.location.latitude,
            coordLng: place.location.longitude,
            highLat: place.viewport.high.latitude,
            highLng: place.viewport.high.longitude,
            lowLat: place.viewport.low.latitude,
            lowLng: place.viewport.low.longitude,

            ...Object.fromEntries(
                place.regularOpeningHours?.periods.flatMap((period, index) => [
                    [`openDay${index}`, period.open.day],
                    [`openTime${index}`, period.open.hour * 100 + period.open.minute],
                    [`closeDay${index}`, period.close === undefined ? null : period.close.day],
                    [
                        `closeTime${index}`,
                        period.close === undefined
                            ? null
                            : period.close.hour * 100 + period.close.minute,
                    ],
                ]) ?? [],
            ),
        }),
    )
}

let pageToken = undefined

for (let i = 0; i < 3; i++) {
    const response = await fetch("https://places.googleapis.com/v1/places:searchText", {
        method: "POST",
        body: JSON.stringify({
            textQuery: "Restaurants near the University of Waterloo",
            pageSize: 20,
            pageToken,
        }),
        headers: {
            "Content-Type": "application/json",
            "X-Goog-Api-Key": process.env.GOOGLE_PLACES_API_KEY,
            "X-Goog-FieldMask": "*",
        },
    })

    if (!response.ok) {
        console.log(`ERROR ${response.status}`, await response.text())
    } else {
        const data = await response.json()

        const {places, nextPageToken} = data

        try {
            console.log(`Inserting ${places.length} places...`)
            await semaphore(
                places.map((place) => async () => {
                    console.log(`Inserting ${place.displayName.text} ${place.id} ...`)

                    const val = await insertPlace(place)

                    console.log(`Inserted ${place.displayName.text} ${place.id}`)

                    return val
                }),
                10,
            )

            pageToken = nextPageToken
        } catch (err) {
            console.error(err)

            break
        }
    }
}

connection.destroy()
