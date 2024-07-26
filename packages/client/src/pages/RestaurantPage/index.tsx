import React, {useState, useEffect} from "react"
import {ReviewCard} from "../../components/review"
import styles from "./index.module.scss"
import ContactInfoCard from "../../components/contactInfoCard"
import data from "./reviews.json"
import {useParams} from "react-router-dom"
import request from "../../utils/request"
import {apiURL, authURL} from "../../globals"

export const RestaurantPage: React.FC = () => {
    const {id: restaurantId} = useParams() as {id: string}
    const [isModalOpen, setIsModalOpen] = useState(false)
    const [name, setName] = useState<string>("")
    const [image, setImage] = useState<string | null>(null)

    useEffect(() => {
        ;(async () => {
            const currentRestaurant = await request(`${apiURL}/restaurants/${restaurantId}`, "GET", "json")
        })()
    }, [restaurantId])

    return (
        <div>
            <div className={styles.imageContainer}>
                {image ? <img src={image} alt={name} className={styles.image} /> : null}
                <div className={styles.overlay}></div>
                <h1 className={styles.textOverlay}>{name}</h1>
            </div>
            <div className={styles.margins}>
                <div className={styles.actions}>
                    <button className={styles.button}>Add review</button>
                </div>
                <div className={styles.content}>
                    <div className={styles.reviews}>
                        {data.map((review) => {
                            const date = new Date(review.updatedAt)
                            return (
                                <ReviewCard
                                    author={review.username}
                                    tags={["Comfort Food", "East Asian"]}
                                    comments={review.comments}
                                    order={"N/A"}
                                    date={date}
                                    rating={review.rating}
                                />
                            )
                        })}
                    </div>
                    <div className={styles.infoCard}>
                        <ContactInfoCard
                            phoneNumber={"416 1111 3333"}
                            location={"123 University Ave"}
                            website={"www.yunshang.ca"}
                            hours={[]}
                        />
                    </div>
                </div>
            </div>
        </div>
    )
}

// <div>
//     <ReviewCard
//         author={"Bob Dylan"}
//         tags={["Comfort Food", "East Asian"]}
//         text={
//             "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut"
//         }
//         order={"Tomato rice noodle"}
//     />
// </div>
