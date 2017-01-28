import React, { Component, PropTypes } from 'react';

import Navbar from './partials/Navbar';
import Footer from './partials/Footer';

class App extends Component {
  constructor(props){
      super(props);
  }
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
}
App.PropTypes = {
    children : PropTypes.object.isRequired
};

export default App;
