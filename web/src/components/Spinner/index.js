import React from 'react';
import styles from './index.module.css';

export default function Spinner() {
    return (
        <div className={styles.root}>
            <div className={styles.loader}>Loading...</div>
        </div>
    );
}
