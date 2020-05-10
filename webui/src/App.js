import React from 'react';

import {connect} from 'react-redux';
import Navbar from 'react-bootstrap/Navbar';
import Container from "react-bootstrap/Container";
import {BrowserRouter as Router, Switch, Route, Redirect} from "react-router-dom";

import LoginForm from "./components/Login";
import SetupForm from "./components/Setup";
import NavDropdown from "react-bootstrap/NavDropdown";
import { logout, redirected } from './actions/auth';
import {IndexPage} from "./components/IndexPage";

// css imports
import 'bootstrap/dist/css/bootstrap.min.css';
import './App.css';


const NavUserInfo = connect(
    ({ auth }) => ({ user: auth.user }),
    ( dispatch ) => ({ logout: () => { dispatch(logout()); } }),
)(({ user, logout }) => {
    if (!user) {
        return null;
    }
    const userTitle = user.full_name || 'ID:'+ user.id;
    return (
        <Navbar.Collapse className="justify-content-end">
            <NavDropdown title={userTitle} id="basic-nav-dropdown" alignRight>
                <NavDropdown.Item onClick={(e) => { e.preventDefault(); logout(); }}>Logout</NavDropdown.Item>
            </NavDropdown>
        </Navbar.Collapse>
    );
})

const PrivateRoute = ({ children, user, ...rest }) => {
    return (
        <Route
            {...rest}
            render={({ location }) =>
                (!!user) ? (
                    children
                ) : (
                    <Redirect to={{pathname: "/login",  state: { from: location }}}/>
                )
            }
        />
    );
};


const App = ({ user, redirectTo, redirected }) => {

    return (
        <div className="App">
            <Navbar variant="dark" bg="dark" expand="md">
                <Navbar.Brand href="/">lakeFS</Navbar.Brand>
                <NavUserInfo/>
            </Navbar>
            <Container className={"main-app"}>
                <Router>
                    <Switch>

                        <Route path="/login" >
                            <LoginForm redirectTo={redirectTo} onRedirect={redirected}/>
                        </Route>

                        <Route path="/setup" >
                            <SetupForm />
                        </Route>

                        <PrivateRoute path="/" user={user}>
                            <IndexPage/>
                        </PrivateRoute>

                    </Switch>
                </Router>
            </Container>
        </div>
    );
};

export default connect(
    ({ auth }) => ({
        user: auth.user,
        redirectTo: auth.redirectTo,
    }), ({ redirected }))(App);
