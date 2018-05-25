import React from 'react';
import axios from 'axios';

class FetchFromLocalHost extends React.Component {
  constructor(props) {
    super(props);
    console.log("fetching started");

    this.state = {  
        posts : []
    };
  }

  ComponentDidMount() {
    console.log("requesting");
    axios.get(`http://localhost:8080/${this.props.substring}`)
      .then(res => {
          
        console.log("fetched");
        console.log(res);
        const posts = res.data;
        this.setState({ posts });
      });
  }

  render() {
    this.ComponentDidMount();
    return (    
      <div>
        {this.state.posts}
      </div>
    );
  }
}

export default FetchFromLocalHost;