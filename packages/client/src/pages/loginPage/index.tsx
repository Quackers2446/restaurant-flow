import React from "react"
import {Register} from "./register"
import {Center, Container} from "@mantine/core"

export interface LoginPageProps {
    type: "login" | "register"
}

export const LoginPage: React.FC<LoginPageProps> = ({type}) => {
    return (
        <Center h="100%">
            <Container>{type === "login" ? <></> : <Register />}</Container>
        </Center>
    )
}
