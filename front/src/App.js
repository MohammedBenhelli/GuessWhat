import React from "react";
import {BrowserRouter as Router, Route, Switch,} from "react-router-dom";
import {CreateChannel} from "./components/CreateChannel";
import {ToastContainer} from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';

function App() {
    return (<>
        <Router>
            <Switch>
                <Route path="/create-channel">
                    <CreateChannel/>
                </Route>
                <Route path="/">
                    <CreateChannel/>
                </Route>
            </Switch>
        </Router>
        <ToastContainer/>
    </>);
}

export default App;
