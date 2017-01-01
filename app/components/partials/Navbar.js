import React from 'react';
import { Link } from 'react-router';

const iconClose = 'https://cdnjs.cloudflare.com/ajax/libs/material-design-icons/3.0.1/navigation/1x_web/ic_close_white_48dp.png';
const iconMenu = 'https://cdnjs.cloudflare.com/ajax/libs/material-design-icons/3.0.1/navigation/1x_web/ic_menu_white_48dp.png';

function toggleDropdown() {
  var img = document.getElementById('dropdown-trigger');
  img.src = (img.src == iconClose ? iconMenu : iconClose);
  var x = document.getElementsByClassName('dropdown-content');
  for (var i = 0; i < x.length; i++) {
    x[i].classList.toggle('show');
  }
}

const Navbar = () =>
  <div className="navbar bar">
    <div className="vertical-center logo text-center"><a href="#">Hermes</a></div>
    <img onClick={toggleDropdown} id="dropdown-trigger" src="https://cdnjs.cloudflare.com/ajax/libs/material-design-icons/3.0.1/navigation/1x_web/ic_menu_white_48dp.png" />
    <div className="dropdown-content flex-row">
      <ul>
        <li className="list-item"><Link to="/">Dashboard</Link></li>
        <li className="list-item"><a href="#">Hubs</a></li>
        <li className="list-item"><a href="#">Events</a></li>
      </ul>
      <div className="spacer" />
      <ul>
        <li className="list-item"><a href="#"><span className="focus">8</span></a></li>
        <li className="list-item"><a href="#">Preferences</a></li>
        <li className="list-item"><a href="#"> {/* TODO: Remove Harcoded Value */} Logout (Ujjwal Sharma)</a></li>
      </ul>
    </div>
  </div>;


export default Navbar;
