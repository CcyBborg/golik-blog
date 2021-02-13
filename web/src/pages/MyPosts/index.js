import MyPosts from './MyPosts';
import { Provider } from 'react-redux';
import store from './store';

const MyPostsPage = props => (
    <Provider store={store}>
        <MyPosts {...props} />
    </Provider>
);

export default MyPostsPage;
