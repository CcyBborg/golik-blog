import axios from 'axios';
import {
    FETCH_POSTS_REQUEST,
    FETCH_POSTS_SUCCESS,
    FETCH_POSTS_FAILURE
} from './constants/action-types';

export function fetchPosts(offset, limit) {
    return async dispatch => {
        try {
            dispatch({
                type: FETCH_POSTS_REQUEST
            });
        
            const response = await axios.get('/posts', {
                params: {
                    offset,
                    limit
                }
            });

            dispatch({
                type: FETCH_POSTS_SUCCESS,
                payload: response.data
            });
        } catch (error) {
            dispatch({
                type: FETCH_POSTS_FAILURE
            });
        }
    };
}
