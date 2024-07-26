import {Link} from "react-router-dom"
import styles from "./index.module.scss"
import {AppShell, Button} from "@mantine/core"
import {AuthContext} from "../../auth"
import {useContext} from "react"

export const Header: React.FC = () => {
    const authContext = useContext(AuthContext)

    return (
        <AppShell.Header className={styles.header}>
            <h1 className={styles.logo}>
                <Link to="/">UW Eats</Link>
            </h1>
            {authContext.auth ? (
                <Button variant="transparent" color="White" component={Link} to="/">
                    Hi, {authContext.auth.user.username}
                </Button>
            ) : (
                <Button variant="transparent" color="White" component={Link} to="/login">
                    Log In / Sign Up
                </Button>
            )}
        </AppShell.Header>
    )
}
