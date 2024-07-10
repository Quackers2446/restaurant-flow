import {AppShell, MantineProvider} from "@mantine/core"
import "@mantine/core/styles.css"
import {QueryClient, QueryClientProvider} from "@tanstack/react-query"
import Leaflet from "leaflet"
import iconShadow from "leaflet/dist/images/marker-shadow.png"
import "leaflet/dist/leaflet.css"
import React from "react"
import ReactDOM from "react-dom/client"
import {Route, BrowserRouter as Router, Routes} from "react-router-dom"
import {RestaurantMap} from "./components"
import {Header} from "./components/header"
import "./index.scss"
import icon from "./mapMarker.svg"
import {HomePage} from "./pages"
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
    return (
        <QueryClientProvider client={queryClient}>
            <MantineProvider defaultColorScheme="dark">
                <AppShell header={{height: 60}} padding="md">
                    <Header />
                    <AppShell.Main className="mantine-main">
                        <Router>
                            <Routes>
                                <Route path="/" element={<HomePage />} />
                                <Route path="/map" element={<RestaurantMap />} />
                            </Routes>
                        </Router>
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
