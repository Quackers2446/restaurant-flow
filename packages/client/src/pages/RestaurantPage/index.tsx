import React, {useState, useEffect} from "react"
import {ReviewCard} from "../../components/review"
import styles from "./index.module.scss"
import ContactInfoCard from "../../components/contactInfoCard"
// import data from "./reviews.json"
import {useParams} from "react-router-dom"
import request from "../../utils/request"
import {apiURL, authURL} from "../../globals"
import ReviewModal from "../../components/ReviewModal"
import {getRestaurantResponse, getRestaurantReviewsResponse} from "../../schema/restaurant"

export const RestaurantPage: React.FC = () => {
    const {id: restaurantId} = useParams() as {id: string}
    const [isModalOpen, setIsModalOpen] = useState(false)
    // const [name, setName] = useState<string>("")
    const [image, setImage] = useState<string | null>(null)
    const [restaurant, setRestaurant] = useState<(typeof getRestaurantResponse)["_type"] | null>(null)
    const [reviews, setReviews] = useState<(typeof getRestaurantReviewsResponse)["_type"] | null>(null)

    useEffect(() => {
        ;(async () => {
            const currentRestaurant = await getRestaurantResponse.parseAsync(
                await request(`${apiURL}/restaurants/${restaurantId}`, "GET", "json"),
            )
            console.log(currentRestaurant)
            console.log(
                currentRestaurant
                    ? (currentRestaurant?.openingHours.main).map((day) => {
                          return (
                              String(day.openTime).substring(0, String(day.openTime).length - 2) +
                              ":" +
                              String(day.openTime).substring(
                                  String(day.openTime).length - 2,
                                  String(day.openTime).length,
                              ) +
                              " - " +
                              String(day.closeTime).substring(0, String(day.closeTime).length - 2) +
                              ":" +
                              String(day.closeTime).substring(
                                  String(day.closeTime).length - 2,
                                  String(day.closeTime).length,
                              )
                          )
                      })
                    : "",
            )
            setImage("https://" + currentRestaurant?.googleRestaurant.photos?.split("=")[0] + "=w2000-h200")
            setRestaurant(currentRestaurant)
        })()
    }, [restaurantId])

    useEffect(() => {
        ;(async () => {
            const currentReviews = await getRestaurantReviewsResponse.parseAsync(
                await request(`${apiURL}/restaurants/${restaurantId}/reviews`, "GET", "json"),
            )
            setReviews(currentReviews)
        })()
    }, [restaurantId])

    const handleOpenModal = () => {
        setIsModalOpen(true)
    }

    const handleCloseModal = () => {
        setIsModalOpen(false)
    }

    return (
        <div className={styles.page}>
            <div className={styles.imageContainer}>
                {image ? <img src={image} alt={restaurant?.googleRestaurant.name} className={styles.image} /> : null}
                <div className={styles.overlay}></div>
                <h1 className={styles.textOverlay}>{restaurant?.googleRestaurant.name}</h1>
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
                        {reviews
                            ? reviews.map((review) => {
                                  const date = new Date(review.createdAt)
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
                              })
                            : null}
                    </div>
                    <div className={styles.infoCard}>
                        <ContactInfoCard
                            phoneNumber={restaurant?.googleRestaurant.phone ? restaurant?.googleRestaurant.phone : ""}
                            location={restaurant?.location.address ? restaurant?.location.address : ""}
                            website={restaurant?.googleRestaurant.website ? restaurant?.googleRestaurant.website : ""}
                            hours={
                                restaurant
                                    ? (restaurant?.openingHours.main).map((day) => {
                                          return (
                                              [
                                                  "Monday",
                                                  "Tuesday",
                                                  "Wednesday",
                                                  "Thursday",
                                                  "Friday",
                                                  "Saturday",
                                                  "Sunday",
                                              ][day.closeDay ? day.closeDay : 0] +
                                              " " +
                                              String(day.openTime).substring(0, String(day.openTime).length - 2) +
                                              ":" +
                                              String(day.openTime).substring(
                                                  String(day.openTime).length - 2,
                                                  String(day.openTime).length,
                                              ) +
                                              " - " +
                                              String(day.closeTime).substring(0, String(day.closeTime).length - 2) +
                                              ":" +
                                              String(day.closeTime).substring(
                                                  String(day.closeTime).length - 2,
                                                  String(day.closeTime).length,
                                              )
                                          )
                                      })
                                    : []
                            }
                        />
                    </div>
                </div>
            </div>
        </div>
    )
}

function getRestaurantsSchema(arg0: {[key: string]: unknown}) {
    throw new Error("Function not implemented.")
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
