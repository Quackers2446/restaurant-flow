import React, {useState} from "react"
import {ReviewCard} from "../../components/review"
import styles from "./index.module.scss"
import ContactInfoCard from "../../components/contactInfoCard"
import data from "./reviews.json"
import ReviewModal from "../../components/ReviewModal"

type RestaurantPageProps = {
    name: string
    image: string
}

const RestaurantPage: React.FC<RestaurantPageProps> = ({name, image}) => {
    const [isModalOpen, setIsModalOpen] = useState(false)

    const handleOpenModal = () => {
        setIsModalOpen(true)
    }

    const handleCloseModal = () => {
        setIsModalOpen(false)
    }

    return (
        <div className={styles.page}>
            <div className={styles.imageContainer}>
                <img src={image} alt={name} className={styles.image} />
                <div className={styles.overlay}></div>
                <h1 className={styles.textOverlay}>{name}</h1>
            </div>
            <div className={styles.margins}>
                <div className={styles.actions}>
                    <button className={styles.button} onClick={handleOpenModal}>
                        Add review
                    </button>
                    {isModalOpen && <ReviewModal onClose={handleCloseModal} />}
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

export default RestaurantPage

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
