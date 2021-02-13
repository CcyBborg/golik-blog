import React from 'react';
import PostHeader from '../PostHeader';
import CommentsCount from '../CommentsCount';
import styles from './index.module.css';

export default function PostListItem({ id, title, summary, author, commentsCount, publishedAt }) {
    return (
        <li className={styles.root}>
            <PostHeader author={author} date={publishedAt} />
            <a href={`posts/${id}/`}><h2>{title}</h2></a>
            <a href={`posts/${id}/`}><p>{summary}</p></a>
            <CommentsCount count={commentsCount} />
        </li>
    );
}