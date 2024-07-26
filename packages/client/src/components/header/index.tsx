import {Link} from "react-router-dom"
import styles from "./index.module.scss"
import {AppShell, Button} from "@mantine/core"

export const Header: React.FC = () => {
    return (
        <AppShell.Header className={styles.header}>
            <h1 className={styles.logo}>
                <Link to="/">UW Eats</Link>
            </h1>
            <Button variant="transparent" color="White" component={Link} to="/login">
                Log In / Sign Up
            </Button>
        </AppShell.Header>
    )
}
