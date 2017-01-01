import React from 'react';

const Footer = () =>
  <div className="footer bar">
    <ul>
      <li className="list-item"><a href="https://github.com/osdc/hermes/tags">Changelog</a></li>
      <li className="list-item"><a href="https://github.com/osdc/hermes">Github</a></li>
      <li className="list-item"><a href="https://raw.githubusercontent.com/osdc/hermes/master/AUTHORS">{ /* TODO: Add a modal-like display of contributors rather than redirecting to Github */ }Contributors</a></li>
    </ul>
    <div className="spacer" />
    <div className="vertical-center text-center">
      <span>Made with ❤ by <a href="https://github.com/osdc">OSDC</a></span>
    </div>
  </div>;


export default Footer;
