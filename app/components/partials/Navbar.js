import React from 'react';
import { Link } from 'react-router';

const ICON_CLOSE = 'https://cdnjs.cloudflare.com/ajax/libs/material-design-icons/3.0.1/navigation/1x_web/ic_close_white_48dp.png';
const ICON_MENU = 'https://cdnjs.cloudflare.com/ajax/libs/material-design-icons/3.0.1/navigation/1x_web/ic_menu_white_48dp.png';

function toggleDropdown() {
  let menuToggleIcon = document.getElementsByClassName('dropdown-trigger')[0];
  menuToggleIcon.src = (menuToggleIcon.src === ICON_CLOSE ? ICON_MENU : ICON_CLOSE);
  document.getElementsByClassName('dropdown-content')[0].classList.toggle('show');
}

const Navbar = () =>
  <div className="navbar bar">
    <div className="vertical-center logo text-center"><a href="#">Hermes</a></div>
    <img onClick={toggleDropdown}
         className="dropdown-trigger"
         src="https://cdnjs.cloudflare.com/ajax/libs/material-design-icons/3.0.1/navigation/1x_web/ic_menu_white_48dp.png" />
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
  </div>;

export default Navbar;
