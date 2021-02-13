import React, { useCallback } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import Box from '../../components/Box';
import PostList from '../../components/PostList';
import Spinner from '../../components/Spinner';
import { fetchPosts } from './actions';
import { POSTS_FETCH_LIMIT } from './constants/limit';

export default function Posts() {
    const dispatch = useDispatch();
    const posts = useSelector(state => state.posts);

    const handleFetchPosts = useCallback(
        () => dispatch(fetchPosts(posts.page, POSTS_FETCH_LIMIT)),
        [posts.page]
    );

    if (posts.isError) {
        return (<p>Произошла ошибка, попробуйте позже</p>);
    }

    return (
        <Box>
            <PostList
                posts={posts}
                onFetchPosts={handleFetchPosts} />
            {posts.isLoading && (
                <Spinner />
            )}
        </Box>
    );
}
