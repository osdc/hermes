import 'babel-polyfill';
import React from 'react';
import { render } from 'react-dom';
import { browserHistory } from 'react-router';
import { syncHistoryWithStore } from 'react-router-redux';
import { AppContainer } from 'react-hot-loader';
import configureStore from './store/configureStore';
import Root from './containers/Root';
import './styles/main.scss';

const store = configureStore();
const history = syncHistoryWithStore(browserHistory, store);

const target = document.getElementById('root');
const node = (
    <AppContainer>
      <Root store={store} history={history} />
    </AppContainer>);
render(
  node,
  target
);

if (module.hot) {
  module.hot.accept('./containers/Root', () => {
    const NewRoot = require('./containers/Root').default;
    const newNode = (
        <AppContainer>
          <NewRoot store={store} history={history} />
        </AppContainer>);
    render(
      newNode,
      target
    );
  });
}

