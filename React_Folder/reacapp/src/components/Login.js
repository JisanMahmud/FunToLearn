import React, { Component } from 'react';
import { Button, FormGroup, FormControl, ControlLabel } from "react-bootstrap";
import VarifyLoginCred from './PostRequest.js'

class Login extends Component{
    constructor(props){
        super(props);
        console.log("Login page");

        this.state={
            loggedIn:false, 
            email : "",
            password : ""
        };
    }

    handleSubmit= (event)=> {
        event.preventDefault();
        console.log(event)
        console.log(this.state.email)
        console.log(this.state.password)
      
        var param = '/user/'+this.state.email+'/pass/'+this.state.password
        console.log(param)

        var LoggedIn = VarifyLoginCred(param)
        console.log("submited result" +LoggedIn)
    }

    handleChangeEmail = (event) => {
        console.log(event.target.value)
        this.setState({
          email: event.target.value
        });
    }

    handleChangePass = (event) => {

        console.log(event.target.value)
        this.setState({
          password: event.target.value
        });
    }


    render(){
        return(
      <div className="Login">
        <form onSubmit={this.handleSubmit}>
          <FormGroup controlId="email" bsSize="large">
            <ControlLabel>Email</ControlLabel>
            <FormControl
              autoFocus
              type="Input"
              value={this.state.email}
              onChange={this.handleChangeEmail}
            />
          </FormGroup>
          <FormGroup controlId="password" bsSize="large">
            <ControlLabel>Password</ControlLabel>
            <FormControl
              value={this.state.password}
              onChange={this.handleChangePass}
              type="password"
            />
          </FormGroup>
          <Button
            block
            bsSize="large"
            type="submit"
          >
            Login
          </Button>
        </form>
      </div>
        );
    }

}

export default Login