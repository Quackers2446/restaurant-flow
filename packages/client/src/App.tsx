import React from "react"
import logo from "./logo.svg"
import "./App.css"
import '@mantine/core/styles.css';
import { MantineProvider } from '@mantine/core';
import SearchBar from "./components/SearchBar";
import { AppShell, Burger } from '@mantine/core';
import { useDisclosure } from '@mantine/hooks'

function App() {
    const [opened, { toggle }] = useDisclosure();

    return (
        <MantineProvider defaultColorScheme='dark' >
            <AppShell
                header={{ height: 60 }}
                navbar={{
                    width: 300,
                    breakpoint: 'sm',
                    collapsed: { mobile: !opened },
                }}
                padding="md"
            >
                <AppShell.Header>
                    <Burger
                        opened={opened}
                        onClick={toggle}
                        hiddenFrom="sm"
                        size="sm"
                    />
                    <SearchBar />
                </AppShell.Header>

                <AppShell.Navbar p="lg">Navbar</AppShell.Navbar>

                <AppShell.Main>Main</AppShell.Main>
            </AppShell>
        </MantineProvider >
    );
}

export default App
