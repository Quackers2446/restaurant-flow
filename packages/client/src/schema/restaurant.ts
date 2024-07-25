import {z} from "zod"

const nullableNumberBoolean = z
    .number()
    .nullable()
    .transform((value) => (value === null ? null : Boolean(value)))

const googleRestaurantSchema = z.object({
    acceptsCashOnly: nullableNumberBoolean,
    acceptsCreditCards: nullableNumberBoolean,
    acceptsDebitCards: nullableNumberBoolean,
    acceptsNfc: nullableNumberBoolean,
    avgRating: nullableNumberBoolean,
    businessStatus: z.object({
        googlerestaurantBusinessStatus: z
            .enum(["Operational", "ClosedTemporarily", "ClosedPermanently"])
            .optional(),
        valid: z.boolean(),
    }),
    description: z.string().nullable(),
    goodForGroups: nullableNumberBoolean,
    goodForWatchingSports: nullableNumberBoolean,
    googleRestaurantId: nullableNumberBoolean,
    googleUrl: z.string(),
    hasOutdoorSeating: nullableNumberBoolean,
    hasRestroom: nullableNumberBoolean,
    name: z.string(),
    phone: z.string(),
    placeId: z.string(),
    priceLevel: z.object({
        googlerestaurantPriceLevel: z
            .enum(["Free", "Inexpensive", "Moderate", "Expensive", "VeryExpensive"])
            .optional(),
        valid: z.boolean(),
    }),
    servesBeer: nullableNumberBoolean,
    servesBreakfast: nullableNumberBoolean,
    servesBrunch: nullableNumberBoolean,
    servesCocktails: nullableNumberBoolean,
    servesCoffee: nullableNumberBoolean,
    servesDessert: nullableNumberBoolean,
    servesDinner: nullableNumberBoolean,
    servesLunch: nullableNumberBoolean,
    servesVegetarianFood: nullableNumberBoolean,
    servesWine: nullableNumberBoolean,
    supportsCurbsidePickup: nullableNumberBoolean,
    supportsDelivery: nullableNumberBoolean,
    supportsDineIn: nullableNumberBoolean,
    supportsReservations: nullableNumberBoolean,
    supportsTakeout: nullableNumberBoolean,
    updatedAt: z.string(),
    website: z.string().nullable(),
    wheelchairAccessibleEntrance: nullableNumberBoolean,
    wheelchairAccessibleSeating: nullableNumberBoolean,
})

const locationSchema = z.object({
    address: z.string(),
    googleRestaurantId: z.number(),
    lat: z.number(),
    lng: z.number(),
    locationId: z.number(),
    viewportHighLat: z.number(),
    viewportHighLng: z.number(),
    viewportLowLat: z.number(),
    viewportLowLng: z.number(),
})

const openingPeriodSchema = z.object({
    openingPeriodId: z.number(),
    closeDay: z.number().optional(),
    openingHoursId: z.number(),
    openDay: z.number(),
    openTime: z.number(),
    closeTime: z.number().optional(),
})

const openingHoursSchema = z.record(z.string(), z.array(openingPeriodSchema))

const tagSchema = z.object({
    name: z.string(),
    restaurantId: nullableNumberBoolean,
    tagId: nullableNumberBoolean,
})

const getRestaurantsSchema = z.object({
    createdAt: z.string(),
    googleRestaurant: googleRestaurantSchema,
    googleRestaurantId: nullableNumberBoolean,
    location: locationSchema,
    openingHours: openingHoursSchema,
    restaurantId: z.number(),
    tags: z.array(tagSchema),
    updatedAt: z.string(),
})

export const getRestaurantsResponse = z.array(getRestaurantsSchema)

const getRestaurantsInAreaSchema = z.object({
    createdAt: z.string(),
    distance: z.number(),
    googleRestaurant: googleRestaurantSchema,
    googleRestaurantId: z.number().optional(),
    location: locationSchema,
    restaurantId: z.number().optional(),
    updatedAt: z.string(),
})

const getRestaurantReviewsSchema = z.object({
    comments: z.string(),
    createdAt: z.string(),
    isAnonymous: z.boolean(),
    rating: z.number(),
    restaurantId: z.number(),
    reviewId: z.number(),
    userId: z.string(),
})


export const getRestaurantReviewsResponse = z.array(getRestaurantReviewsSchema)

export const getRestaurantsInAreaResponse = z.array(getRestaurantsInAreaSchema)
