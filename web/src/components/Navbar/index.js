import React from 'react';
import barsIcon from './resources/bars.svg';
import styles from './index.module.css';

export default function Navbar() {
    return (
        <div className={styles.root}>
            <div>
                <img src={barsIcon} alt='menu icon' />
                <span className={styles.title}>Blog</span>
            </div>
            <div className={styles.right}>
                <a href='/login' className={styles.link}>Log In</a>
                <a href='/signup' className={styles.link}>Sign Up</a>
            </div>
        </div>
    );
}