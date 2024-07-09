import {AppShell, Burger, MantineProvider} from "@mantine/core"
import "@mantine/core/styles.css"
import {useDisclosure} from "@mantine/hooks"
import {QueryClient, QueryClientProvider} from "@tanstack/react-query"
import Leaflet from "leaflet"
import icon from "./mapMarker.svg"
import iconShadow from "leaflet/dist/images/marker-shadow.png"
import "leaflet/dist/leaflet.css"
import React from "react"
import ReactDOM from "react-dom/client"
import {RestaurantMap} from "./components"
import "./index.scss"
import reportWebVitals from "./reportWebVitals"

const markerWidth = 25
const markerHeight = 37

Leaflet.Marker.prototype.options.icon = Leaflet.icon({
    iconUrl: icon,
    shadowUrl: iconShadow,
    iconAnchor: Leaflet.point(markerWidth / 2, markerHeight),
    popupAnchor: Leaflet.point(0, -markerHeight),
})

const queryClient = new QueryClient()

const App: React.FC = () => {
    const [opened, {toggle}] = useDisclosure()

    return (
        <QueryClientProvider client={queryClient}>
            <MantineProvider defaultColorScheme="dark">
                <AppShell header={{height: 60}} padding={0}>
                    <AppShell.Header>
                        <Burger opened={opened} onClick={toggle} hiddenFrom="sm" size="sm" />
                    </AppShell.Header>

                    <AppShell.Main className="mantine-main">
                        <RestaurantMap />
                    </AppShell.Main>
                </AppShell>
            </MantineProvider>
        </QueryClientProvider>
    )
}

const root = ReactDOM.createRoot(document.getElementById("root") as HTMLElement)
root.render(
    <React.StrictMode>
        <App />
    </React.StrictMode>,
)

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals()
