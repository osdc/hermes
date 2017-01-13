import React from 'react';
import { Link } from 'react-router';

const IconClose = require('babel!svg-react!../../../public/ic_close_white_48px.svg?name=IconClose');
const IconMenu = require('babel!svg-react!../../../public/ic_menu_white_48px.svg?name=IconMenu');

class Navbar extends React.Component {
  constructor(props) {
    super(props);
    this.state = { isOpen: false };
  }

  toggleDropdown() {
    this.setState((prevState) => ({
      isOpen: !prevState.isOpen
    }));
    document.getElementsByClassName('dropdown-content')[0].classList.toggle('show');
  }

  render() {
    return (
      <div className="navbar bar">
        <div className="vertical-center logo text-center"><a href="#">Hermes</a></div>
        {
          this.state.isOpen ?
          <IconClose className="dropdown-trigger" onClick={this.toggleDropdown.bind(this)} />
          :
          <IconMenu className="dropdown-trigger" onClick={this.toggleDropdown.bind(this)} />
        }
        <div className="dropdown-content flex-row">
          <ul>
            <li className="list-item"><Link to="/">Dashboard</Link></li>
            <li className="list-item"><a href="#">Hubs</a></li>
            <li className="list-item"><a href="#">Events</a></li>
          </ul>
          <div className="spacer" />
          { /* Check if user is logged in */
            true ?
            <ul>
              <li className="list-item"><Link to="/login">Login</Link></li>
              <li className="list-item"><Link to="/register">Register</Link></li>
            </ul>
            :
            <ul>
              <li className="list-item"><a href="#"><span className="focus">{ /* TODO: Remove Hardcoded Value*/ }8</span></a></li>
              <li className="list-item"><a href="#">Preferences</a></li>
              <li className="list-item"><a href="#"> {/* TODO: Remove Harcoded Value */} Logout (Ujjwal Sharma)</a></li>
            </ul>
          }
        </div>
      </div>
    );
  }
}

export default Navbar;
