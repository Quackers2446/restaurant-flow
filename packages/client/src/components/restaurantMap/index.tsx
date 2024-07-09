import styles from "./index.module.scss"
import {useQuery} from "@tanstack/react-query"
import qs from "qs"
import React from "react"
import {MapContainer, Marker, Popup, TileLayer} from "react-leaflet"
import {type LatLngLiteral} from "leaflet"
import {apiURL} from "../../globals"
import {getRestaurantsInAreaResponse} from "../../schema/restaurant"

// TODO: make these modifiable
const position: LatLngLiteral = {lat: 43.472587, lng: -80.537681}

export const RestaurantMap: React.FC = () => {
    const {error, data} = useQuery({
        queryKey: ["restaurantList"],
        queryFn: () =>
            fetch(
                `${apiURL}/restaurants/in-area?${qs.stringify({lat: position.lat, lng: position.lng, radius: 200})}`,
            ).then(async (res) => await getRestaurantsInAreaResponse.parseAsync(await res.json())),
    })

    console.log({data, error})

    return (
        <MapContainer
            center={position}
            maxZoom={20}
            zoom={17}
            scrollWheelZoom
            className={styles.leafletContainer}
        >
            <TileLayer
                attribution='&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
                url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
                maxZoom={20}
            />
            {data?.map((restaurant) => (
                <Marker position={[restaurant.location.lat, restaurant.location.lng]}>
                    <Popup>{restaurant.googleRestaurant.name}</Popup>
                </Marker>
            ))}
        </MapContainer>
    )
}