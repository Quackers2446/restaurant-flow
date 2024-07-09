import Leaflet from "leaflet"
import iconShadow from "leaflet/dist/images/marker-shadow.png"
import mapMarker from "./mapMarker.svg"
import mapMarkerActive from "./mapMarkerActive.svg"

const markerWidth = 25
const markerHeight = 37

export const defaultIcon = Leaflet.icon({
    iconUrl: mapMarker,
    shadowUrl: iconShadow,
    iconAnchor: Leaflet.point(markerWidth / 2, markerHeight),
    popupAnchor: Leaflet.point(0, -markerHeight),
})

Leaflet.Marker.prototype.options.icon = defaultIcon

export const activeIcon = Leaflet.icon({
    iconUrl: mapMarkerActive,
    shadowUrl: iconShadow,
    iconAnchor: Leaflet.point(markerWidth / 2, markerHeight),
    popupAnchor: Leaflet.point(0, -markerHeight),
})
