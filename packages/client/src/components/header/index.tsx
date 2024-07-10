import styles from "./index.module.scss"
import { AppShell, Button } from "@mantine/core"

export const Header: React.FC = () => {
    return (
        <AppShell.Header className={styles.header}>
            <h1 className={styles.logo}>UW Eats</h1>
            <Button variant="transparent" color="White"> Log In / Sign Up</Button>
        </AppShell.Header >
    );
}
