import {
    FETCH_POST_SUCCESS,
    FETCH_POST_FAILURE
} from '../constants/action-types';

export const initialState = {
    isLoading: true,
    isError: false,
    data: null
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
        case FETCH_POST_SUCCESS: {
            return {
                ...state,
                isLoading: false,
                isError: false,
                data: action.payload,
            };
        }

        case FETCH_POST_FAILURE: {
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
