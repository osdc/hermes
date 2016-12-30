import React from 'react';
import { Route, IndexRoute } from 'react-router';
import Dashboard from './components/Dashboard';
import Login from './components/Login';
import Register from './components/Register'

export default (
	<div>
		<Route path="/" component={Dashboard} />
		<Route path="/login" component={Login} />
		<Route path="/register" component={Register} />
	</div>
);
