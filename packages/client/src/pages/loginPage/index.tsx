import React from "react"
import {Register} from "./register"
import {Center, Container} from "@mantine/core"
import {Login} from "./login"

export interface LoginPageProps {
    type: "login" | "register"
}

export const LoginPage: React.FC<LoginPageProps> = ({type}) => {
    return (
        <Center h="100%">
            <Container>{type === "login" ? <Login /> : <Register />}</Container>
        </Center>
    )
}
