import React, {Dispatch, SetStateAction} from "react"
import styles from "./index.module.scss"

type ReviewModalProps = {
    onClose: () => void
    setDesc: Dispatch<SetStateAction<string | null>>
    setItem: Dispatch<SetStateAction<string | null>>
    setRating: Dispatch<SetStateAction<string | null>>
    onSubmit: React.FormEventHandler<HTMLFormElement>
}

const ReviewModal: React.FC<ReviewModalProps> = ({onClose, setRating, setItem, setDesc, onSubmit}) => {
    return (
        <div className={styles.modalOverlay}>
            <div className={styles.modalContent}>
                <h2>Write your review...</h2>
                <form className={styles.form} onSubmit={onSubmit}>
                    <div className={styles.formGroup}>
                        <label htmlFor="item">I got:</label>
                        <input
                            type="text"
                            id="item"
                            className={styles.input}
                            onChange={(event) => setItem(event.target.value)}
                        />
                    </div>
                    <div className={styles.formGroup}>
                        <label htmlFor="rating">Rating:</label>
                        <input
                            type="text"
                            id="rating"
                            className={styles.input}
                            onChange={(event) => setRating(event.target.value)}
                        />
                    </div>
                    <div className={styles.formGroup}>
                        <textarea
                            rows={4}
                            placeholder="Share your experience..."
                            className={styles.textarea}
                            onChange={(event) => setDesc(event.target.value)}
                        ></textarea>
                    </div>
                    <div className={styles.modalActions}>
                        <button type="submit" className={styles.button}>
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
