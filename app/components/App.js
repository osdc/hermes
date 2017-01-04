import React from 'react';

import Navbar from './partials/Navbar';
import Footer from './partials/Footer';

const App = React.createClass({
  render() {
    return (
      <div className="hermes">
        <Navbar />
        <div className="content container">
          { this.props.children }
        </div>
        <Footer />
      </div>
    );
  }
});

export default App;
