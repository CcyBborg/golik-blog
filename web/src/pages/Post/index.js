import Post from './Post';
import { Provider } from 'react-redux';
import store from './store';

const PostPage = props => (
    <Provider store={store}>
        <Post {...props} />
    </Provider>
);

export default PostPage;
