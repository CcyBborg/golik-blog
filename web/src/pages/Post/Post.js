import React, { useCallback, useEffect } from 'react';
import { useSelector, useDispatch } from 'react-redux';
import Box from '../../components/Box';
import PosHeader from '../../components/PostHeader';
import Spinner from '../../components/Spinner';
import Comments from '../../components/Comments';
import { fetchPost, fetchComments, postComment } from './actions';
import styles from './Post.module.css';

export default function Post({ match }) {
    const dispatch = useDispatch();
    const post = useSelector(state => state.post);
    const comments = useSelector(state => state.comments);
    const postId = Number(match.params.id);

    useEffect(() => dispatch(fetchPost(postId)), [postId, dispatch]);
    useEffect(() => dispatch(fetchComments(postId)), [postId, dispatch]);
    const handlePostComment = useCallback(content => dispatch(postComment(content, postId)), [dispatch, postId]);

    return (
        <Box>
            {post.isLoading && !post.isError ? (
                <Spinner />
            ) : (
                <>
                    <div className={styles.postContent}>
                        <PosHeader author={post.data.author} date={post.data.publishedAt} />
                        <h1>{post.data.title}</h1>
                        <p>{post.data.content}</p>
                    </div>
                    {comments.isLoading || comments.isError || <Comments list={comments.list} onPostComment={handlePostComment} />}
                    {comments.isLoading && <Spinner />}
                    {comments.isError && <p>Error comments!</p>}
                </>
            )}
            {post.isError && <p>Error while loading post!</p>}
        </Box>
    );
}
