import {createContext} from "react"
import {FullUser} from "./schema/user"

export interface Auth {
    accessToken: string
    user: FullUser
}

export interface AuthContextType {
    auth?: Auth | undefined | null

    setAccessToken: (accessToken: string) => void

    /** Set the current logged in user */
    setUser: (user: FullUser) => void

    /** Set the current logged in user from an unknown object that is validated */
    setUserFromUnknown: (user: unknown) => Promise<void>

    logout: () => void
}

export const AuthContext = createContext<AuthContextType>({
    auth: undefined,
    setAccessToken: () => {},
    setUser: () => {},
    setUserFromUnknown: () => new Promise((resolve) => resolve()),
    logout: () => {},
})
