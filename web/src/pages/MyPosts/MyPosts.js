import React, { useEffect } from 'react';
import { useSelector, useDispatch } from 'react-redux';
import Box from '../../components/Box';
import MyPostsListItem from '../../components/MyPostsListItem';
import Spinner from '../../components/Spinner';
import { fetchMyPosts } from './actions';
import styles from './index.module.css';

export default function MyPosts() {
    const dispatch = useDispatch();
    const myposts = useSelector(state => state.myposts);

    useEffect(() => dispatch(fetchMyPosts()), [dispatch]);

    return (
        <Box>
            {myposts.isLoading || myposts.isError || (
                <>
                    <div className={styles.header}>
                        <h3>My Posts</h3>
                        <button className='btn-primary'>New Post</button>
                    </div>
                    <ul className={styles.list}>
                        {myposts.list.map(post => (
                            <MyPostsListItem {...post} />
                        ))}
                    </ul>
                </>
            )}
            {myposts.isLoading && <Spinner />}
            {myposts.isError && <p>Error fetching myposts!</p>}
        </Box>
    );
}
