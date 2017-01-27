import React from 'react';
import { Route, IndexRoute } from 'react-router';

import App from './components/App';
import Dashboard from './components/Dashboard';
import Login from './components/Login';
import Register from './components/Register';
import Hubs from './components/Hubs';

export default (
  <Route path="/" component={App}>
    <IndexRoute component={Dashboard} />
    <Route path="/login" component={Login} />
    <Route path="/register" component={Register} />
    <Route path="/hubs" component={Hubs} />
  </Route>
);
