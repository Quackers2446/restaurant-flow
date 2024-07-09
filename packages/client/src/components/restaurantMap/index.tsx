import styles from "./index.module.scss"
import { useQuery } from "@tanstack/react-query"
import qs from "qs"
import React from "react"
import { MapContainer, Marker, TileLayer } from "react-leaflet"
import { type LatLngLiteral } from "leaflet"
import { apiURL } from "../../globals"

// TODO: make these modifiable
const position: LatLngLiteral = { lat: 43.472587, lng: -80.537681 }

export const RestaurantMap: React.FC = () => {
    const { isPending, error, data } = useQuery({
        queryKey: ["restaurantList"],
        queryFn: () =>
            fetch(
                `${apiURL}/restaurants/in-area?${qs.stringify({ lat: position.lat, lng: position.lng, radius: 200 })}`,
            ).then((res) => res.json()),
    })

    return (
        <MapContainer
            center={position}
            zoom={50}
            scrollWheelZoom
            className={styles.leafletContainer}
        >
            <TileLayer
                attribution='&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
                url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
            />
            <Marker position={position}></Marker>
        </MapContainer>
    )
}
