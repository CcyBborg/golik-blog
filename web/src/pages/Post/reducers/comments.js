import {
    FETCH_COMMENTS_REQUEST,
    FETCH_COMMENTS_SUCCESS,
    FETCH_COMMENTS_FAILURE
} from '../constants/action-types';

export const initialState = {
    isLoading: true,
    isError: false,
    list: null
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
        case FETCH_COMMENTS_REQUEST: {
            return {
                ...state,
                isLoading: true,
                isError: false
            };
        }

        case FETCH_COMMENTS_SUCCESS: {
            return {
                ...state,
                isLoading: false,
                list: action.payload,
            };
        }

        case FETCH_COMMENTS_FAILURE: {
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
