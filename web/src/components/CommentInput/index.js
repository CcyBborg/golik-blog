import React, { useState } from 'react';
import styles from './index.module.css';

export default function CommentInput({ onPostComment }) {
    const [comment, setComment] = useState('');

    return (
        <form onSubmit={event => { event.preventDefault(); onPostComment(comment); }}>
            <textarea className={styles.input} value={comment} rows={5} onChange={({ target }) => setComment(target.value)}></textarea>
            <button className='btn-primary' type='submit'>Post</button>
        </form>
    );
}