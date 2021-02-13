import React from 'react';
import PostListItem from '../PostListItem';
import InfiniteScroll from '../InfiniteScroll';
import styles from './index.module.css';

export default function PostList({ posts, onFetchPosts }) {
    return (
        <ul className={styles.list}>
            <InfiniteScroll isLoaded={posts.isLoaded} isLoading={posts.isLoading} onFetch={onFetchPosts}>
                {posts.ids.map(id => (
                    <PostListItem key={id} {...posts.data[id]} />
                ))}
            </InfiniteScroll>
        </ul>
    );
}
