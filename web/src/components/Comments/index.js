import React from 'react';
import PostHeader from '../../components/PostHeader';
import CommentInput from '../../components/CommentInput';
import styles from './index.module.css';

export default function Comments({ list, onPostComment }) {
    return (
        <section className={styles.root}>
            <h3 className={styles.title}>Comments <span>{list.length}</span></h3>
            <CommentInput onPostComment={onPostComment} />
            <ul className={styles.list}>
                {list.length == 0 ? (
                    <p>No comments yet, be the first one!</p>
                ) : (
                    list.map(comment => (
                        <li>
                            <PostHeader author={comment.author} date={comment.createdAt} />
                            <p>{comment.content}</p>
                        </li>
                )))}
            </ul>
        </section>
    );
}
