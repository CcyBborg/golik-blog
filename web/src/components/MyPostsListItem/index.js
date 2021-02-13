import React from 'react';
import CommentsCount from '../CommentsCount';
import styles from './index.module.css'

const ZERO_DATE = '0001-01-01T00:00:00Z';

export default function MyPostsListItem({ id, title, publishedAt, commentsCount }) {
    const formattedDate = new Date(publishedAt);
    const isPublished = publishedAt !== ZERO_DATE;

    return (
        <li>
            <h3>{title}</h3>
            {isPublished && (
                <span className={styles.date}>Published at {`${formattedDate.getDate()}-${formattedDate.getMonth() + 1}-${formattedDate.getFullYear()}`}</span>
            )}
            {isPublished && (
                <CommentsCount count={commentsCount} />
            )}
            <div className={styles.actions}>
                {isPublished ? (
                    <button className={styles.editBtn}>Unpublish</button>
                ) : (
                    <button className={styles.editBtn}>Publish</button>
                )}
                <button className={styles.editBtn}>Edit</button>
                <button className={styles.deleteBtn}>Delete</button>
            </div>
        </li>
    );
}
