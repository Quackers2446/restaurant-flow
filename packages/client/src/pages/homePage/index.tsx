import {
    Badge,
    Card,
    Container,
    Group,
    Image,
    SimpleGrid,
    Text,
    TextInput,
    Title,
} from "@mantine/core"
import qs from "qs"
import React from "react"
import {CiSearch} from "react-icons/ci"
import {apiURL} from "../../globals"
import {getRestaurantsResponse} from "../../schema/restaurant"
import {debounce} from "../../utils/debounce"
import styles from "./index.module.scss"

export const HomePage: React.FC = () => {
    const [data, setData] = React.useState<typeof getRestaurantsResponse._type>()

    return (
        <Container>
            <Title className={styles.title}>Your UW Flow for good eats.</Title>
            <Text className={styles.subtitle}>Find the best meals near campus for your needs</Text>
            <Text className={styles.description}>
                Filter by categories tailored for students like you
                <br />
                Recommended by students like you
            </Text>
            <TextInput
                placeholder={"Restaurant or food"}
                className={styles.textInput}
                radius="xl"
                size="lg"
                leftSection={<CiSearch size={24} />}
                onChange={debounce(async (event) => {
                    // TODO: use react query and clean this shit up
                    const res = await fetch(
                        `${apiURL}/restaurants/search?${qs.stringify({search: event.target.value})}`,
                    ).then(
                        async (res) => await getRestaurantsResponse.parseAsync(await res.json()),
                    )

                    setData(res)
                }, 500)}
            />

            {/* TODO: this was temporary */}
            <Text>
                {data?.map((restaurant) => (
                    <React.Fragment key={restaurant.restaurantId}>
                        {restaurant.googleRestaurant.name}
                        <br />
                    </React.Fragment>
                ))}
            </Text>

            <SimpleGrid cols={3}>
                <Card shadow="sm" padding="lg" radius="20" className={styles.card}>
                    <Card.Section>
                        <Image
                            src="https://via.placeholder.com/150"
                            alt="Tomato Rice Noodle Soup"
                        />
                    </Card.Section>
                    <Text className={styles.cardTitle}>Tomato Rice Noodle Soup | Yunshang</Text>
                    <Group className={styles.cardBadge}>
                        <Badge className={styles.cardBadgeComfort}>Comfort food</Badge>
                        <Badge className={styles.cardBadgeEastAsian}>East Asian</Badge>
                    </Group>
                </Card>

                <Card shadow="sm" padding="lg" radius="20" className={styles.card}>
                    <Card.Section>
                        <Image src="https://via.placeholder.com/150" alt="Loren Ipsum" />
                    </Card.Section>
                    <Text className={styles.cardTitle}>Loren ipsum | Loren ipsum</Text>
                    <Group className={styles.cardBadge}>
                        <Badge className={styles.cardBadgeHealthy}>Healthy</Badge>
                        <Badge className={styles.cardBadgeCheap}>Cheap eats</Badge>
                    </Group>
                </Card>

                <Card shadow="sm" padding="lg" radius="20" className={styles.card}>
                    <Card.Section>
                        <Image src="https://via.placeholder.com/150" alt="Loren Ipsum" />
                    </Card.Section>
                    <Text className={styles.cardTitle}>Loren ipsum | Loren ipsum</Text>
                    <Group className={styles.cardBadge}>
                        <Badge className={styles.cardBadgeStudy}>Study snacks</Badge>
                    </Group>
                </Card>
            </SimpleGrid>
        </Container>
    )
}
