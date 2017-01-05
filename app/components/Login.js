import React from 'react';
import { Link } from 'react-router';

const Login = () =>
  <div>
    <div>
      <h1 className="text-primary text-center">Login</h1>
      <div className="box-login">
        <input type="text" placeholder="Enrollment Number" />
        <input type="password" placeholder="Password/PIN" />
        <button className="btn-accent">Login</button>
        <a href="#" className="text-center text-accent">Forgot Password</a>
        <Link to="/register" className="text-center text-accent">Register</Link>
      </div>
    </div>
  </div>

export default Login;
