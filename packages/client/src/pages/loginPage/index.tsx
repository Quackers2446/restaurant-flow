import React from "react"

export interface LoginPageProps {
    type: "login" | "register"
}

export const LoginPage: React.FC<LoginPageProps> = ({type}) => {
    return <div></div>
}
