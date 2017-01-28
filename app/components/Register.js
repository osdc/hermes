import React from 'react';
import { Link } from 'react-router';

class Register extends React.Component {
  constructor(props) {
    super(props);
    this.state = {}
  }

  render() {
    return (
      <div>
        <div>
          <h1 className="text-primary text-center">Register</h1>
          <div className="box-login">
            <input type="text" placeholder="Enrollment Number" />
            <input type="password" placeholder="Password/PIN" />
            <button className="btn-accent">Register</button>
            <Link to="/login" className="text-center text-accent">Login</Link>
          </div>
        </div>
      </div>
    );
  }
}

export default Register;
