import React from 'react';
import { Link } from 'react-router';
import DatePicker from 'react-datepicker';
import moment from 'moment';

import 'style!css!react-datepicker/dist/react-datepicker.css';

class Register extends React.Component {
  render() {
    return (
      <div>
        <div>
          <h1 className="text-primary text-center">Register</h1>
          <div className="box-login">
            <input type="text" placeholder="Enrollment Number" />
            <DatePicker selected={moment()} />
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
