import React from "react"
import styles from "./index.module.scss"

type ReviewModalProps = {
    onClose: () => void
}

const ReviewModal: React.FC<ReviewModalProps> = ({onClose}) => {
    return (
        <div className={styles.modalOverlay}>
            <div className={styles.modalContent}>
                <h2>Write your review...</h2>
                <form className={styles.form}>
                    <div className={styles.formGroup}>
                        <label htmlFor="item">I got:</label>
                        <input type="text" id="item" className={styles.input} />
                    </div>
                    <div className={styles.formGroup}>
                        <label htmlFor="rating">Rating:</label>
                        <input type="text" id="rating" className={styles.input} />
                    </div>
                    <div className={styles.formGroup}>
                        <textarea
                            rows={4}
                            placeholder="Share your experience..."
                            className={styles.textarea}
                        ></textarea>
                    </div>
                    <div className={styles.modalActions}>
                        <button type="button" className={styles.button} onClick={onClose}>
                            Submit
                        </button>
                        <button type="button" className={styles.button} onClick={onClose}>
                            Cancel
                        </button>
                    </div>
                </form>
            </div>
        </div>
    )
}

export default ReviewModal
