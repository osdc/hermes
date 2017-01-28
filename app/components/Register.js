import React from 'react';
import { Link } from 'react-router';

import 'whatwg-fetch';

class Register extends React.Component {
  constructor(props) {
    super(props);
    this.state = { notice: '', name: '', enrollment_number: '', password: '' };
  }

  handleChangeName(event) {
    this.setState({ name: event.target.value });
  }

  handleChangeEnroll(event) {
    this.setState({ enrollment_number: event.target.value });
  }

  handleChangePass(event) {
    this.setState({ password: event.target.value });
  }

  handleSubmit() {
    fetch('http://localhost:4000/api/user/new', {
      method: 'POST',
      body: JSON.stringify(this.state),
      headers: {
        "Content-Type": "application/json"
      }
    }).then((response) => {
      if (response.statusText === 'Created') {
        this.setState({ notice: 'User created successfully. You can now log in.' });
      }
    }, (err) => {
      console.error(err);
    });
  }

  render() {
    return (
      <div>
        <div>
          <div>{ this.state.notice }</div>
          <h1 className="text-primary text-center">Register</h1>
          <div className="box-login">
            <input
              type="text"
              placeholder="Name"
              value={this.state.name}
              onChange={this.handleChangeName.bind(this)} />
            <input
              type="text"
              placeholder="Enrollment Number"
              value={this.state.enrollment_number}
              onChange={this.handleChangeEnroll.bind(this)} />
            <input
              type="password"
              placeholder="Password/PIN"
              value={this.state.pass}
              onChange={this.handleChangePass.bind(this)} />
            <button className="btn-accent" onClick={this.handleSubmit.bind(this)}>Register</button>
            <Link to="/login" className="text-center text-accent">Login</Link>
          </div>
        </div>
      </div>
    );
  }
}

export default Register;
