import React from 'react';

const Navbar = () =>
  <div className="navbar bar">
    <div className="vertical-center"><a href="#">Hermes</a></div>
    <ul>
      <li className="list-item"><a href="#">Dashboard</a></li>
      <li className="list-item"><a href="#">Hubs</a></li>
      <li className="list-item"><a href="#">Events</a></li>
    </ul>
    <div className="spacer" />
    <ul>
      <li className="list-item"><a href="#"><span className="focus">8</span></a></li>
      <li className="list-item"><a href="#">Preferences</a></li>
      <li className="list-item"><a href="#"> {/* TODO: Remove Harcoded Value */} Logout (Ujjwal Sharma)</a></li>
    </ul>
  </div>;


export default Navbar;
