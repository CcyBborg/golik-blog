import React from 'react';
import {
  BrowserRouter as Router,
  Switch,
  Route
} from 'react-router-dom';
import Navbar from './components/Navbar';
import Posts from './pages/Posts';
import Post from './pages/Post';
import MyPosts from './pages/MyPosts';


function App() {
  return (
    <Router>
      <div>
        <Navbar />
        <Switch>
          <Route path='/posts/:id(\d+)' component={({ match }) => (<Post match={match} />)}>
          </Route>
          <Route path='/myposts'>
            <MyPosts />
          </Route>
          <Route path='/' exact>
            <Posts />
          </Route>
        </Switch>
      </div>
    </Router>
  );
}

export default App;
