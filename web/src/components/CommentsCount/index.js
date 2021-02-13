import React from 'react';
import commentIcon from './resources/comment.svg';
import styles from './index.module.css';

export default function CommentsCount({ count }) {
    return (
        <div className={styles.root}>
            <img src={commentIcon} alt='comments icon' />
            <span className={styles.count}>{count}</span>
        </div>
    );
}
