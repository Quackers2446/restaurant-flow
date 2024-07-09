import { TextInput } from '@mantine/core';
import React from 'react';
import { Text, Title, Container, Card, Image, Badge, SimpleGrid, Group } from '@mantine/core';
import styles from './index.module.scss'
import { CiSearch } from "react-icons/ci";

export const HomePage: React.FC = () => {
    return (
        <Container>
            <Title className={styles.title}>
                Your UW Flow for good eats.
            </Title>
            <Text className={styles.subtitle}>
                Find the best meals near campus for your needs
            </Text>
            <Text className={styles.description}>
                Filter by categories tailored for students like you
                <br />
                Recommended by students like you
            </Text>
            <TextInput placeholder={"Restaurant or food"} className={styles.textInput}
                radius="xl" size="lg" leftSection={<CiSearch size={24} />} />

            <SimpleGrid cols={3}>
                <Card shadow="sm" padding="lg" radius="20" className={styles.card}>
                    <Card.Section>
                        <Image src="https://via.placeholder.com/150" alt="Tomato Rice Noodle Soup" />
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
        </Container >
    );
};
