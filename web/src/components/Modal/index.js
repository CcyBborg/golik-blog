import React from 'react';

export default function Modal({ children }) {
    return (
        <div className={styles.root}>
            <div className={styles.content}>
                {children}
            </div>
        </div>
    );
}