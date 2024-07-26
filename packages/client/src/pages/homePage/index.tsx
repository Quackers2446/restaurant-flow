import {Badge, Card, Container, Group, Image, SimpleGrid, Text, TextInput, Title} from "@mantine/core"
import React from "react"
import {CiSearch} from "react-icons/ci"
import styles from "./index.module.scss"
import {useNavigate} from "react-router-dom"

export const HomePage: React.FC = () => {
    const [searchQuery, setSearchQuery] = React.useState("")
    const navigate = useNavigate()

    const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault()
        if (searchQuery.trim()) {
            navigate("/search", {state: {query: searchQuery}})
        }
    }

    return (
        <Container>
            <Title className={styles.title}>Your UW Flow for good eats.</Title>
            <Text className={styles.subtitle}>Find the best meals near campus for your needs</Text>
            <Text className={styles.description}>
                Filter by categories tailored for students like you
                <br />
                Recommended by students like you
            </Text>
            <form onSubmit={handleSubmit}>
                <TextInput
                    placeholder={"Restaurant or food"}
                    className={styles.textInput}
                    radius="xl"
                    size="lg"
                    leftSection={<CiSearch size={24} />}
                    value={searchQuery}
                    onChange={(e) => setSearchQuery(e.currentTarget.value)}
                />
            </form>

            <SimpleGrid cols={3}>
                <Card shadow="sm" padding="lg" radius="20" className={styles.card}>
                    <Card.Section>
                        <Image
                            src="https://yunshang.ca/wp-content/uploads/%E6%96%B0-%E7%95%AA%E8%8C%84%E9%87%91%E9%92%88%E8%8F%87%E8%82%A5%E7%89%9B%E7%B1%B3%E7%BA%BF.jpg"
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
                        <Image
                            src="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSK05gWd8Vi0TR929n_R3NWoc_ALPGCmDhqdA&s"
                            alt="Wagyu"
                        />
                    </Card.Section>
                    <Text className={styles.cardTitle}>Wagyu | Daldongnae</Text>
                    <Group className={styles.cardBadge}>
                        <Badge className={styles.cardBadgeHealthy}>Healthy</Badge>
                        <Badge className={styles.cardBadgeEastAsian}>East Asian</Badge>
                    </Group>
                </Card>

                <Card shadow="sm" padding="lg" radius="20" className={styles.card}>
                    <Card.Section>
                        <Image
                            src="https://images.squarespace-cdn.com/content/v1/5c658114ab1a625acb417fc9/1575597714451-V4PZPRV96M4GJ0H1CB5X/IMG_3574.jpg"
                            alt="Loren Ipsum"
                        />
                    </Card.Section>
                    <Text className={styles.cardTitle}>Deerioca Fever | The Alley</Text>
                    <Group className={styles.cardBadge}>
                        <Badge className={styles.cardBadgeStudy}>Study snacks</Badge>
                        <Badge className={styles.cardBadgeCheap}>Cheap eats</Badge>
                    </Group>
                </Card>
            </SimpleGrid>
        </Container>
    )
}
