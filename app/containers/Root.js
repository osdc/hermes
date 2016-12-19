import React, { Component, PropTypes } from 'react';
import { Provider } from 'react-redux';
import { Router } from 'react-router';
import routes from '../routes';

// Partials
import Navbar from '../components/partials/Navbar';
import Header from '../components/partials/Header';
import Footer from '../components/partials/Footer';

export default class Root extends Component {
  render() {
    const { store, history } = this.props;
    return (
        <Provider store={store}>
          <div>
            <Navbar />
            <Header />
            <div className="content container">
              <Router history={history} routes={routes} />
            </div>
            <Footer />
          </div>
        </Provider>
    );
  }
}

Root.propTypes = {
  store: PropTypes.object.isRequired,
  history: PropTypes.object.isRequired
};
