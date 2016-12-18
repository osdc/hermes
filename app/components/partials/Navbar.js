import React from 'react';

const Navbar = () =>
    <nav>
        <span className="logo">Hermes</span>
        <ul className="nav-left visible-lg">
            <li><a href="#">Dashboard</a></li>
            <li><a href="#">Hubs</a></li>
            <li><a href="#">Events</a></li>
        </ul>
        <ul className="nav-right">
            <li><a href="#"><span className="focus">8</span></a></li>
            <li><a href="#">Preferences</a></li>
            <li><a href="#">Logout (Ujjwal Sharma)</a></li>
        </ul>
    </nav>;


export default Navbar;
