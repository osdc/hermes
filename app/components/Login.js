import React from 'react';

const Login = () =>
  <div>
    <div className="width-full">
      <h1 className="text-primary text-center">Login</h1>
      <div className="box-login">
        <input type="text" placeholder="Enrollment Number" />
        <input type="text" placeholder="Password/PIN" />
        <button className="btn-accent">Login</button>
      </div>
    </div>
  </div>

export default Login;
