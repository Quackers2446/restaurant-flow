import {AppShell, Burger, MantineProvider} from "@mantine/core"
import "@mantine/core/styles.css"
import {useDisclosure} from "@mantine/hooks"
import React from "react"
import ReactDOM from "react-dom/client"
import reportWebVitals from "./reportWebVitals"

const App: React.FC = () => {
    const [opened, {toggle}] = useDisclosure()

    return (
        <MantineProvider defaultColorScheme="dark">
            <AppShell
                header={{height: 60}}
                navbar={{
                    width: 300,
                    breakpoint: "sm",
                    collapsed: {mobile: !opened},
                }}
                padding="md"
            >
                <AppShell.Header>
                    <Burger opened={opened} onClick={toggle} hiddenFrom="sm" size="sm" />
                </AppShell.Header>

                <AppShell.Navbar p="lg">Navbar</AppShell.Navbar>

                <AppShell.Main>Main</AppShell.Main>
            </AppShell>
        </MantineProvider>
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
