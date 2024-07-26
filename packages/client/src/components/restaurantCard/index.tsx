import React from "react"
import {Badge, Group, Text} from "@mantine/core"
import {IconStar, IconMessageCircle} from "@tabler/icons-react"
import styles from "./index.module.scss"
import {FaStar, FaStarHalfAlt, FaRegStar} from "react-icons/fa"
import {useNavigate} from "react-router-dom"

type RestaurantCardProps = {
    id: number
    imageUrl: string
    title: string
    rating: number
    tag: string
    popularDish: string
    review: string
}

export const RestaurantCard: React.FC<RestaurantCardProps> = ({
    id,
    imageUrl,
    title,
    rating,
    tag,
    popularDish,
    review,
}) => {
    const fullStars = Math.floor(rating / 2)
    const halfStars = rating % 2 >= 1 ? 1 : 0
    const emptyStars = 5 - fullStars - halfStars
    const navigate = useNavigate()

    return (
        <div className={styles.itemCard}>
            <div className={styles.itemImage}>
                <img src={imageUrl} alt={title} />
            </div>
            <div className={styles.itemDetails}>
                <Text className={styles.itemTitle}>{title}</Text>
                <Group className={styles.itemRating} gap="xs">
                    {Array.from({length: fullStars}, (_, i) => (
                        <FaStar key={`full-${i}`} />
                    ))}
                    {Array.from({length: halfStars}, (_, i) => (
                        <FaStarHalfAlt key={`half-${i}`} />
                    ))}
                    {Array.from({length: emptyStars}, (_, i) => (
                        <FaRegStar key={`empty-${i}`} />
                    ))}
                </Group>
                <Badge className={styles.itemTag}>{tag}</Badge>
                {/* <Text className={styles.itemPopularDish}>Popular dish: {popularDish}</Text> */}
                <div className={styles.itemReview}>
                    <IconMessageCircle size={16} />
                    <Text className={styles.reviewText}>"{review}..."</Text>
                    <Text className={styles.moreLink} onClick={() => navigate(`/restaurant/${id}`)}>
                        more
                    </Text>
                </div>
            </div>
        </div>
    )
}
