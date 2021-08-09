import React, {useCallback, useMemo, useRef} from "react";
import {
    BrowserRouter as Router,
    Switch,
    Route,
} from "react-router-dom";
import {CreateChannel} from "./components/CreateChannel";

function App() {



    return (
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
    );
}

export default App;
