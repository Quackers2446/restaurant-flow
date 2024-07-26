import React, { useEffect } from "react"
import { Button, Text, Flex, TextInput } from "@mantine/core"
import { CiSearch } from "react-icons/ci"
import styles from "./index.module.scss"
import { RestaurantMap, RestaurantCard } from "../../components"
import { getRestaurantsResponse, getRestaurantReviewsResponse } from "../../schema/restaurant"
import { debounce } from "../../utils/debounce"
import { apiURL } from "../../globals"
import { useLocation } from "react-router-dom"
import qs from "qs"

const Sidebar: React.FC = () => {
    return (
        <div className={styles.sidebar}>
            <div className={styles.filterSection}>
                <h2>Category</h2>
                <Flex direction="column" align="flex-start" gap="xs">
                    <Button className={styles.filterButton}>Comfort food</Button>
                    <Button className={styles.filterButton}>Treat yourself</Button>
                    <Button className={styles.filterButton}>Study snacks</Button>
                    <Button className={styles.filterButton}>Healthy</Button>
                    <Button className={styles.filterButton}>Indulgent</Button>
                    <Button className={styles.filterButton}>Grab and Go</Button>
                    <Button className={styles.filterButton}>Aesthetic</Button>
                </Flex>
            </div>
            <hr />
            <div className={styles.filterSection}>
                <h2>Price</h2>
                <Flex gap="sm">
                    <Button className={styles.filterButton}>$</Button>
                    <Button className={styles.filterButton}>$$</Button>
                    <Button className={styles.filterButton}>$$$</Button>
                </Flex>
            </div>
            <hr />
            <div className={styles.filterSection}>
                <h2>Cuisine</h2>
                <Flex direction="column" align="flex-start" gap="xs">
                    <Button className={styles.filterButton}>East Asian</Button>
                    <Button className={styles.filterButton}>Italian</Button>
                    <Button className={styles.filterButton}>Indian</Button>
                    <Button className={styles.filterButton}>Mexican</Button>
                </Flex>
            </div>
            <hr />
            <div className={styles.filterSection}>
                <h2>Location</h2>
                <Flex direction="column" align="flex-start" gap="xs">
                    <Button className={styles.filterButton}>University Ave. Plaza</Button>
                    <Button className={styles.filterButton}>On Campus</Button>
                    <Button className={styles.filterButton}>Uptown</Button>
                </Flex>
            </div>
        </div>
    )
}

export const SearchPage: React.FC = () => {
    const [data, setData] = React.useState<typeof getRestaurantsResponse._type>()
    const [allReviews, setAllReviews] = React.useState<(typeof getRestaurantReviewsResponse._type)[]>()
    const location = useLocation()
    const [query, setQuery] = React.useState<string>(location.state.query)

    useEffect(() => {
        const getRestauants = async () => {
            const res = await fetch(`${apiURL}/restaurants/search?${qs.stringify({ search: query })}`).then(
                async (res) => await getRestaurantsResponse.parseAsync(await res.json()),
            )
            setData(res)
        }

        getRestauants()
    }, [])

    useEffect(() => {
        const updateReviews = () => {
            const list: (typeof getRestaurantReviewsResponse._type)[] = []

            data?.map(async (restaurant) => {
                const res = await fetch(`${apiURL}/restaurants/${restaurant.restaurantId}/reviews`).then(
                    async (res) => await getRestaurantReviewsResponse.parseAsync(await res.json()),
                )
                list.push(res)
            })
            setAllReviews(list)
        }

        updateReviews()
    }, [data])

    return (
        <div className={styles.searchPage} style={{ height: "100dvh" }}>
            <Sidebar />
            <div className={styles.cardsContainer}>
                <TextInput
                    defaultValue={query}
                    placeholder={"Restaurant or Food"}
                    className={styles.textInput}
                    radius="xl"
                    size="lg"
                    leftSection={<CiSearch size={24} />}
                    onChange={debounce(async (event) => {
                        // TODO: use react query and clean this shit up
                        const res = await fetch(
                            `${apiURL}/restaurants/search?${qs.stringify({ search: event.target.value })}`,
                        ).then(async (res) => await getRestaurantsResponse.parseAsync(await res.json()))

                        setData(res)
                        console.log("event:", event.target.value);
                        setQuery(qs.stringify(event.target.value))
                    }, 500)}
                />
                <>
                    {data?.map((restaurant) => {
                        const review = allReviews?.find((review) => {
                            if (review && review[0]) return review[0].restaurantId == restaurant.restaurantId
                        })
                        return (
                            <RestaurantCard
                                id={restaurant.restaurantId}
                                imageUrl={"https://" + restaurant.googleRestaurant.photos?.split("=")[0] + "=w150-h150"} // replace with the actual path
                                title={restaurant.googleRestaurant.name}
                                rating={review ? review[0].rating : 5}
                                tag={["Comfort Food", "East Asian", "Affordable", "Healthy"][(restaurant.googleRestaurant.name.length % 4)]}
                                popularDish="XYZ"
                                review={review ? review[0].comments : ""}
                            />
                        )
                    })}
                </>
            </div>
            <div className={styles.map}>
                <RestaurantMap />
            </div>
        </div>
    )
}
