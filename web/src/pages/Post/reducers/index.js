import { combineReducers } from 'redux';
import post from './post';
import comments from './comments';

const reducer = combineReducers({ post, comments });

export default reducer;
