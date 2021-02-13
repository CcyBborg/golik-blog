import {
    FETCH_MYPOSTS_SUCCESS,
    FETCH_MYPOSTS_FAILURE
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
export default function myposts(state = initialState, action) {
    switch (action.type) {
        case FETCH_MYPOSTS_SUCCESS: {
            return {
                ...state,
                isLoading: false,
                isError: false,
                list: action.payload,
            };
        }

        case FETCH_MYPOSTS_FAILURE: {
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
