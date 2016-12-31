import React from 'react';
import DatePicker from 'react-datepicker';
import moment from 'moment';

import 'style!css!react-datepicker/dist/react-datepicker.css';

const Register = () =>
  <div>
    <div className="width-full">
      <h1 className="text-primary text-center">Register</h1>
      <div className="box-login">
        <input type="text" placeholder="Enrollment Number" />
        <DatePicker selected={moment()} />
        <input type="text" placeholder="Password/PIN" />
        <button className="btn-accent">Register</button>
        <a href="#" className="text-center text-accent">Login</a>
      </div>
    </div>
  </div>

export default Register;
