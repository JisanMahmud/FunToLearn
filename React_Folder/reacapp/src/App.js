import React, { Component } from 'react';
import './App.css';
import FetchFromLocalHost from './components/FetchFromLocalHost.js'
import { BrowserRouter,Switch, Route, Link } from 'react-router-dom'
import Login from './components/Login'

class App extends Component {
  render() {

    console.log(this.props)
    return (
            <BrowserRouter>
              <div>
                <Switch>
                  <Route
                    exact path="/"
                    render={()=><FetchFromLocalHost substring="home"/>} />
                  <Route
                    path="/login"
                    render={()=><Login/>} />
                  <Route path="/blog" children={({match}) => (
                    <li className={match ? 'active' : ''}>
                      <Link to="/blog">Blog</Link>
                    </li>)} />
                  <Route render={() => <h1>Page not found</h1>} />
                </Switch>
              </div>
            </BrowserRouter>
    );
  }
}

export default App;
