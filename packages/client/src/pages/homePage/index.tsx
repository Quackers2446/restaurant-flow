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
import { CiSearch } from "react-icons/ci"
import { apiURL } from "../../globals"
import { getRestaurantsResponse } from "../../schema/restaurant"
import { debounce } from "../../utils/debounce"
import styles from "./index.module.scss"
import { useNavigate } from 'react-router-dom';

export const HomePage: React.FC = () => {
    const [searchQuery, setSearchQuery] = React.useState('');
    const navigate = useNavigate();

    const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        if (searchQuery.trim()) {
            
            navigate('/search', { state: { query: searchQuery } });
        }
    };

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
