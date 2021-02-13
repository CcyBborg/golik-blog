import Posts from './Posts';
import { Provider } from 'react-redux';
import store from './store';

const PostsPage = () => (
    <Provider store={store}>
        <Posts />
    </Provider>
);

export default PostsPage;
