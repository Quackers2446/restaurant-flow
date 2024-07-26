import {Button, TextInput} from "@mantine/core"
import {useForm} from "@mantine/form"
import {useMutation} from "@tanstack/react-query"
import React, {useContext} from "react"
import request from "../../utils/request"
import {apiURL, authURL} from "../../globals"
import {AuthContext} from "../../auth"
import {refreshResponse, ownProfileResponse} from "../../schema/user"
import {useNavigate} from "react-router-dom"

export const Login: React.FC = () => {
    const authContext = useContext(AuthContext)
    const nav = useNavigate()

    const {mutate} = useMutation({
        mutationFn: async (body: {email: string; password: string}) => {
            return await request(`${authURL}/login`, "POST", "json", body, undefined, undefined, {
                credentials: "include",
            })
        },
        onSuccess: async (data) => {
            try {
                const {accessToken} = await refreshResponse.parseAsync(data)

                authContext.setAccessToken(accessToken)

                const getUserResult = await request(
                    `${apiURL}/users/own-profile`,
                    "GET",
                    "json",
                    undefined,
                    accessToken,
                )
                const userResult = await ownProfileResponse.parseAsync(getUserResult)

                authContext.setUser(userResult)

                nav("/")
            } catch (err) {
                console.error(err, data)
                authContext.logout()
            }
        },
    })
    const form = useForm({
        mode: "uncontrolled",
        initialValues: {
            email: "",
            password: "",
        },
    })

    return (
        <>
            <form
                onSubmit={form.onSubmit(async (values) => {
                    mutate(values)
                })}
            >
                <TextInput label="Email" placeholder="Email" key={form.key("email")} {...form.getInputProps("email")} />
                <TextInput
                    label="Password"
                    placeholder="Password"
                    type="password"
                    key={form.key("password")}
                    {...form.getInputProps("password")}
                />
                <Button type="submit" mt="md" color="gray">
                    Submit
                </Button>
            </form>
        </>
    )
}
