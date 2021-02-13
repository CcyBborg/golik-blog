import axios from 'axios';
import {
    FETCH_MYPOSTS_SUCCESS,
    FETCH_MYPOSTS_FAILURE
} from './constants/action-types';

export function fetchMyPosts() {
    return async dispatch => {
        try {
            const response = await axios.get('/myposts');

            dispatch({
                type: FETCH_MYPOSTS_SUCCESS,
                payload: response.data
            });
        } catch (error) {
            dispatch({
                type: FETCH_MYPOSTS_FAILURE
            });
        }
    };
}
