import React from "react";
import {SocketContext, socket} from './context/socket';
import {BrowserRouter as Router, Route, Switch,} from "react-router-dom";
import {Drawer} from "./components/Drawer";
import {ToastContainer} from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';
import {Home} from "./components/Home";

function App() {
    return (<SocketContext.Provider value={socket}>
        <Router>
            <Switch>
                <Route path="/:groupchat">
                    <Drawer/>
                </Route>
                <Route path="/">
                    <Home/>
                </Route>
            </Switch>
        </Router>
        <ToastContainer/>
    </SocketContext.Provider>);
}

export default App;
