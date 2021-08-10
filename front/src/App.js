import React from "react";
import {SocketContext, socket} from './context/socket';
import {BrowserRouter as Router, Route, Switch,} from "react-router-dom";
import {CreateChannel} from "./components/CreateChannel";
import {Drawer} from "./components/Drawer";
import {ToastContainer} from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';

function App() {
    return (<SocketContext.Provider value={socket}>
        <Router>
            <Switch>
                <Route path="/create-channel">
                    <CreateChannel/>
                </Route>
                <Route path="/:groupchat">
                    <Drawer/>
                </Route>
                <Route path="/">
                    <CreateChannel/>
                </Route>
            </Switch>
        </Router>
        <ToastContainer/>
    </SocketContext.Provider>);
}

export default App;
