import {Button, TextInput} from "@mantine/core"
import {useForm} from "@mantine/form"
import {useMutation} from "@tanstack/react-query"
import React from "react"
import request from "../../utils/request"
import {authURL} from "../../globals"

export const Register: React.FC = () => {
    const {mutate} = useMutation({
        mutationFn: async (body: {name: string; username: string; email: string; password: string}) => {
            return await request(`${authURL}/register`, "POST", "json", body)
        },
        onSuccess(data, variables, context) {
            console.log({data, variables, context})
        },
    })
    const form = useForm({
        mode: "uncontrolled",
        initialValues: {
            name: "",
            username: "",
            email: "",
            password: "",
            password2: "",
        },
    })

    return (
        <>
            <form
                onSubmit={form.onSubmit(async (values) => {
                    if (values.password !== values.password2) {
                        throw new Error("passwords don't match") // TODO how to handle this??
                    }

                    mutate(values)
                })}
            >
                <TextInput label="Name" placeholder="Name" key={form.key("name")} {...form.getInputProps("name")} />
                <TextInput
                    label="Username"
                    placeholder="Username"
                    key={form.key("username")}
                    {...form.getInputProps("username")}
                />
                <TextInput label="Email" placeholder="Email" key={form.key("email")} {...form.getInputProps("email")} />
                <TextInput
                    label="Password"
                    placeholder="Password"
                    type="password"
                    key={form.key("password")}
                    {...form.getInputProps("password")}
                />
                <TextInput
                    label="Confirm Password"
                    placeholder="Confirm Password"
                    type="password"
                    key={form.key("password2")}
                    {...form.getInputProps("password2")}
                />
                <Button type="submit" mt="md" color="gray">
                    Submit
                </Button>
            </form>
        </>
    )
}
