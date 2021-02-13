import React from 'react';
import personIcon from './resources/person.svg';
import styles from './index.module.css';

export default function UserInfo({ author, date }) {
    const formattedDate = new Date(date);
    return (
        <div className={styles.root}>
            <div><a><img src={personIcon} alt='person icon' /></a></div>
            <div className={styles.right}>
                <a className={styles.username}><span>{author.username}</span></a>
                <span className={styles.date}>{`${formattedDate.getDate()}-${formattedDate.getMonth() + 1}-${formattedDate.getFullYear()}`}</span>
            </div>
        </div>
    );
}
