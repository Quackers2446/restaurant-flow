import React from "react"
import styles from "./index.module.scss"
import {Badge, Group} from "@mantine/core"
import {ReactComponent as Star} from "./star.svg"
import {ReactComponent as GreyStar} from "./grey-star.svg"
import {ReactComponent as HalfStar} from "./half-star.svg"

interface ReviewCardProps {
    author: string
    tags: string[]
    comments: string
    date: Date
    order: string
    rating: number
}

export const ReviewCard: React.FC<ReviewCardProps> = ({
    author,
    tags,
    comments,
    order,
    date,
    rating,
}) => {
    return (
        <div className={styles.reviewCard}>
            <div className={styles.reviewHeader}>
                <div className={styles.authorDetails}>
                    <div className={styles.authorName}>{author}</div>
                </div>
            </div>
            <div className={styles.star}>
                {Array.from({length: Math.floor(rating / 2)}, (_, index) => (
                    <Star width="20px" height="20px" padding-right="10px" />
                ))}
                {!Number.isInteger(rating / 2) ? (
                    <HalfStar width="20px" height="20px" padding-right="10px" />
                ) : (
                    <></>
                )}
                {Array.from(
                    {length: 5 - Math.floor(rating / 2) - (Number.isInteger(rating / 2) ? 0 : 1)},
                    (_, index) => (
                        <GreyStar width="20px" height="20px" padding-right="10px" />
                    ),
                )}
            </div>
            <Group className={styles.cardBadge}>
                <Badge className={styles.cardBadgeComfort}>Comfort food</Badge>
                <Badge className={styles.cardBadgeEastAsian}>East Asian</Badge>
            </Group>
            <p className={styles.reviewText}>{comments}</p>
            <div>
                <strong className={styles.order}>What they got: </strong>
                {order}
            </div>
        </div>
    )
}
