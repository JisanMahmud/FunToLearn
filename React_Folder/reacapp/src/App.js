import React, { Component } from 'react';
import './App.css';
import FetchFromLocalHost from './components/FetchFromLocalHost.js'

class App extends Component {
  render() {
    return (
      <div className="App">
        <FetchFromLocalHost substring="asdfasdfa">
        </FetchFromLocalHost>
      </div>
    );
  }
}

export default App;
