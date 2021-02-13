import {
    FETCH_POSTS_REQUEST,
    FETCH_POSTS_SUCCESS,
    FETCH_POSTS_FAILURE
} from '../constants/action-types';
import { POSTS_FETCH_LIMIT } from '../constants/limit';

export const initialState = {
    isLoaded: false,
    isLoading: false,
    isError: null,
    page: 0,
    ids: [],
    data: {},
};

/**
 * Редьюсер, обрабатывает данные о покупках
 *
 * @param {Object} [state=initialState] redux state
 * @param {Object} action redux action
 *
 * @returns {Object} updated state
 */
export default function posts(state = initialState, action) {
    switch (action.type) {
        case FETCH_POSTS_REQUEST: {
            return {
                ...state,
                isLoading: true,
                isError: false,
            };
        }

        case FETCH_POSTS_SUCCESS: {
            return {
                ...state,
                isLoading: false,
                ids: [...new Set(
                    [
                        ...state.ids,
                        ...action.payload.map(({ id }) => id),
                    ]
                )],
                data: {
                    ...state.data,
                    ...action.payload.reduce((result, post) => {
                        /* eslint-disable no-param-reassign */
                        result[post.id] = post;
                        /* eslint-enable no-param-reassign */

                        return result;
                    }, {}),
                },
                page: state.page + 1,
                isLoaded: action.payload.length < POSTS_FETCH_LIMIT,
            };
        }

        case FETCH_POSTS_FAILURE: {
            return {
                ...state,
                isLoading: false,
                isError: true
            };
        }

        default: {
            return state;
        }
    }
}
