import React, {useContext, useEffect} from "react";
import {CreateChannel} from "./CreateChannel";
import {JoinChannel} from "./JoinChannel";
import {toast} from "react-toastify";
import {SocketContext} from "../context/socket";
import {useHistory} from "react-router-dom";


export const Home = () => {

    const history = useHistory();
    const socket = useContext(SocketContext);

    useEffect(() => {
        socket.onmessage = e => {
            const data = JSON.parse(e.data);
            console.log(data);
            if (data.message === 'Channel created') {
                toast("Channel created", {type: 'success'});
                history.push(data.data);
            }else if (data.message === 'Added to channel') {
                toast("Added to channel", {type: 'success'});
                history.push(data.data);
            } else toast(data.error, {type: 'error'});
        };
    }, []);
    return (<div id={'home'}>
        <h1 id={'title'}>Guess What ?</h1>
            <CreateChannel socket={socket}/>
            <JoinChannel socket={socket}/>
        </div>)
}
