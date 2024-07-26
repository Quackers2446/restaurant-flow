import {AppShell, MantineProvider} from "@mantine/core"
import "@mantine/core/styles.css"
import {QueryClient, QueryClientProvider} from "@tanstack/react-query"
import "leaflet/dist/leaflet.css"
import React, {useEffect, useState} from "react"
import ReactDOM from "react-dom/client"
import {Route, BrowserRouter as Router, Routes} from "react-router-dom"
import {AuthContext} from "./auth"
import {RestaurantMap} from "./components"
import {Header} from "./components/header"
import {apiURL, authURL} from "./globals"
import "./index.scss"
import "./leaflet"
import {HomePage, LoginPage, RestaurantPage, SearchPage} from "./pages"
import reportWebVitals from "./reportWebVitals"
import {FullUser, ownProfileResponse, refreshResponse} from "./schema/user"
import request from "./utils/request"

const queryClient = new QueryClient()

const App: React.FC = () => {
    const [currentUser, setCurrentUser] = useState<FullUser | undefined | null>(undefined)
    const [currentAccessToken, setCurrentAccessToken] = useState<string | undefined | null>(undefined)

    const setAccessToken = (accessToken: string): void => {
        setAccessToken(accessToken)
    }

    const setUser = (user: FullUser): void => {
        setCurrentUser(user)
    }

    const setUserFromUnknown = async (data: unknown): Promise<void> => {
        try {
            const user = await ownProfileResponse.parseAsync(data)

            if (user !== null && user !== undefined) {
                setCurrentUser(user)
            }
        } catch {}
    }

    const logout = async (): Promise<void> => {
        if (currentUser || currentAccessToken) {
            await fetch(`${authURL}/logout`, {
                method: "POST",
                credentials: "include",
                headers: {
                    "Content-Type": "application/json",
                },
            })
        }

        setCurrentAccessToken(null)
        setCurrentUser(null)
    }

    useEffect(() => {
        ;(async () => {
            try {
                const refreshResult = await request(`${authURL}/refresh`, {
                    method: "GET",
                    conversion: "json",
                    init: {credentials: "include"},
                })
                const {accessToken} = await refreshResponse.parseAsync(refreshResult)
                const getUserResult = await request(
                    `${apiURL}/users/own-profile`,
                    "GET",
                    "json",
                    undefined,
                    accessToken,
                )
                const userResult = await ownProfileResponse.parseAsync(getUserResult)

                setUser(userResult)
                setAccessToken(accessToken)
            } catch (err) {
                await logout()
            }

            setCurrentUser((_currentUser) => _currentUser)
        })()
    }, [])

    useEffect(() => {
        ;(async () => {
            try {
            } catch {
                console.error("Request failed")
            }
        })()
    })

    return (
        <div className="body">
            <AuthContext.Provider
                value={{
                    auth:
                        currentAccessToken === undefined || currentUser === undefined
                            ? undefined
                            : currentAccessToken === null || currentUser === null
                              ? null
                              : {
                                    accessToken: currentAccessToken,
                                    user: currentUser,
                                },
                    setUser,
                    setUserFromUnknown,
                    setAccessToken,
                    logout,
                }}
            >
                <QueryClientProvider client={queryClient}>
                    <MantineProvider>
                        <AppShell header={{height: 60}} padding="md">
                            <Header />
                            <AppShell.Main className="mantine-main">
                                <Router>
                                    <Routes>
                                        <Route path="/" element={<HomePage />} />
                                        <Route path="/map" element={<RestaurantMap />} />
                                        <Route path="/search" element={<SearchPage />} />
                                        <Route path="/login" element={<LoginPage type="login" />} />
                                        <Route path="/register" element={<LoginPage type="register" />} />
                                        <Route path="/restaurant/:id" element={<RestaurantPage />} />
                                    </Routes>
                                </Router>
                            </AppShell.Main>
                        </AppShell>
                    </MantineProvider>
                </QueryClientProvider>
            </AuthContext.Provider>
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
