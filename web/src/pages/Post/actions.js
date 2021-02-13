import axios from 'axios';
import url from 'url';
import {
    FETCH_POST_SUCCESS,
    FETCH_POST_FAILURE,

    FETCH_COMMENTS_REQUEST,
    FETCH_COMMENTS_SUCCESS,
    FETCH_COMMENTS_FAILURE
} from './constants/action-types';

export function fetchPost(postId) {
    return async dispatch => {
        try {
            const response = await axios.get(`/posts/${postId}`);

            dispatch({
                type: FETCH_POST_SUCCESS,
                payload: response.data
            });
        } catch (error) {
            dispatch({
                type: FETCH_POST_FAILURE
            });
        }
    };
}

export function fetchComments(postId) {
    return async dispatch => {
        dispatch({
            type: FETCH_COMMENTS_REQUEST
        });
        try {
            const response = await axios.get(`/posts/${postId}/comments`);

            dispatch({
                type: FETCH_COMMENTS_SUCCESS,
                payload: response.data
            });
        } catch (error) {
            dispatch({
                type: FETCH_COMMENTS_FAILURE
            });
        }
    };
}

export function postComment(content, postId) {
    return async dispatch => {
        try {
            const params = new URLSearchParams();
            params.append('content', content);
            await axios.post(`/posts/${postId}/comments`, params);

            dispatch(fetchComments(postId));
        } catch (error) {
            dispatch({
                type: FETCH_COMMENTS_FAILURE
            });
        }
    };
}
