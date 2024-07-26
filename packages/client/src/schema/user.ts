import {z} from "zod"

export const ownProfileResponse = z.object({
    createdAt: z.string(),
    email: z.string(),
    name: z.string(),
    userId: z.string(),
    username: z.string(),
})

export type FullUser = (typeof ownProfileResponse)["_type"]

export const refreshResponse = z.object({
    accessToken: z.string(),
})
