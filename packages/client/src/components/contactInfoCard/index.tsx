import React from "react"
import styles from "./index.module.scss"

type ContactInfoCardProps = {
    phoneNumber: string
    location: string
    website: string
    hours: string[]
}

const ContactInfoCard: React.FC<ContactInfoCardProps> = ({
    phoneNumber,
    location,
    website,
    hours,
}) => {
    return (
        <div className={styles.card}>
            <div className={styles.infoItem}>
                <span className={styles.label}>Phone Number:</span> {phoneNumber}
            </div>
            <div className={styles.infoItem}>
                <span className={styles.label}>Location:</span> {location}
            </div>
            <div className={styles.infoItem}>
                <span className={styles.label}>Website:</span>{" "}
                <a href={website} target="_blank" rel="noopener noreferrer">
                    {website}
                </a>
            </div>
            <div className={styles.infoItem}>
                <span className={styles.label}>Hours:</span>
                <div>
                    {hours.map((hour, index) => (
                        <div key={index}>{hour}</div>
                    ))}
                </div>
            </div>
        </div>
    )
}

export default ContactInfoCard
