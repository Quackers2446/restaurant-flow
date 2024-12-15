// UserReviews.js
import React from "react"
import styles from "./index.module.scss"
import {ReviewCard} from "../../components/review"
import data from "./reviews.json"

const UserReviews = () => {
    return (
        <div className={styles.userReviews}>
            <h1 className={styles.title}>Hi, Bob!</h1>
            <h2 className={styles.subtitle}>Your previous reviews</h2>
            <div className={styles.reviewsList}>
                {data.map((review) => {
                    const date = new Date(review.updatedAt)
                    return (
                        <ReviewCard
                            id={review.reviewId}
                            key={review.userId}
                            author={review.username} // better to have some sort of getRestaurant function using id
                            tags={["Comfort food", "East Asian"]}
                            comments={review.comments}
                            order={"N/A"}
                            date={date}
                            rating={review.rating}
                            onDelete={() => {}}
                        />
                    )
                })}
            </div>
        </div>
    )
}

export default UserReviews
