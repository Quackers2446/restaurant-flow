import { AppShell, MantineProvider } from "@mantine/core"
import "@mantine/core/styles.css"
import { QueryClient, QueryClientProvider } from "@tanstack/react-query"
import "leaflet/dist/leaflet.css"
import React from "react"
import ReactDOM from "react-dom/client"
import { Route, BrowserRouter as Router, Routes } from "react-router-dom"
import { RestaurantMap } from "./components"
import { Header } from "./components/header"
import "./index.scss"
import "./leaflet"
import { HomePage, SearchPage } from "./pages"
import reportWebVitals from "./reportWebVitals"
import Restaurant from "./pages/RestaurantPage"

const queryClient = new QueryClient()

const App: React.FC = () => {
    return (
        <div className="body" >
            <QueryClientProvider client={queryClient} >
                <MantineProvider >
                    <AppShell header={{ height: 60 }} padding="md">
                        <Header />
                        <AppShell.Main className="mantine-main">
                            <Router>
                                <Routes>
                                    <Route path="/" element={<HomePage />} />
                                    <Route path="/map" element={<RestaurantMap />} />
                                    <Route path="/search" element={<SearchPage />} />
                                    <Route
                                      path="/restaurant/" //:id
                                      element={
                                          <Restaurant
                                              name={"Yunshang Rice Noodle"}
                                              image={
                                                  "https://scontent.fyzd1-2.fna.fbcdn.net/v/t39.30808-6/299142732_494424549353128_4477296176078588222_n.jpg?_nc_cat=102&ccb=1-7&_nc_sid=cc71e4&_nc_ohc=3g4Hp-zV8hIQ7kNvgFe5ezi&_nc_ht=scontent.fyzd1-2.fna&gid=Avzg8Mj6U_L1qYXen_w0lGI&oh=00_AYDv3kFXTspuDgQEAhyp6Uw0MplbSWcGtW8JPSK5gM3Jdg&oe=66A8C043"
                                              }
                                          />
                                      }
                                    />
                                </Routes>
                            </Router>
                        </AppShell.Main>
                    </AppShell>
                </MantineProvider>
            </QueryClientProvider>
        </div>
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
