import React, {useContext, useEffect, useRef, useState} from 'react';
import {toast} from 'react-toastify';
import CanvasDraw from "react-canvas-draw";
import {SliderPicker} from "react-color";
import {useParams} from 'react-router-dom'
import {SocketContext} from "../context/socket";


export const Drawer = () => {
    const {groupchat} = useParams();
    const socket = useContext(SocketContext);
    const [color, setColor] = useState('#426DEA');
    const [state, setState] = useState(null)
    const [isAdmin, setIsAdmin] = useState(false)
    const [isDrawer, setIsDrawer] = useState(false)
    const [started, setIsStarted] = useState(false)
    const drawRef = useRef(null);

    const sleep = (milliseconds) => {
        return new Promise(resolve => setTimeout(resolve, milliseconds))
    }

    useEffect(() => {
        socket.onmessage = async e => {
            const {data, error, message} = JSON.parse(e.data);
            console.log(JSON.parse(e.data));
            if (message === 'update-canvas' && !isAdmin) {
                console.log(data)
                drawRef.current.loadSaveData(data, true);
            } else if (message === 'is-admin') {
                setIsAdmin(data === 'true');
            } else if (message === 'is-drawer') {
                setIsDrawer(data === 'true');
            } else if (message === 'start-game') {
                setState(JSON.parse(data));
                setIsStarted(true);
                socket.send(JSON.stringify({type: 'is-drawer', room_name: groupchat}));
            } else if (message === 'new-message') {
                setState(JSON.parse(data));
                console.log(JSON.parse(data))
                await sleep(500)
                socket.send(JSON.stringify({type: 'is-drawer', room_name: groupchat}));
            } else {
                toast(error, {type: 'error'});
            }
        };
        socket.send(JSON.stringify({type: 'is-admin', room_name: groupchat}));
    }, [groupchat]);

    const drawingEvent = (e) => {
        console.log(e, JSON.parse(e.getSaveData()));
        if (isDrawer)
            socket.send(JSON.stringify({canvas: e.getSaveData(), room_name: groupchat, type: 'update-canvas'}));
    }

    const startGame = () => socket.send(JSON.stringify({room_name: groupchat, type: 'start-game'}));

    const props = {
        onChange: drawingEvent,
        lazyRadius: 10,
        brushRadius: 4,
        brushColor: color,
        hideGrid: true,
        canvasWidth: 700,
        canvasHeight: 400,
        disabled: !isDrawer,
        immediateLoading: false,
        hideInterface: false
    };


    return (<div id={'channel'}>
        <PersonsList users={state?.teams}/>
        <div id="drawer">
            {started && isDrawer && <h1 id={'word'}>{state.word}</h1>}
            <SliderPicker color={color} onChangeComplete={c => setColor(c.hex)}/>
            <CanvasDraw ref={drawRef} {...props} />
            {isDrawer && <button onClick={() => drawRef.current.undo()}>Previous</button>}
            {isAdmin && !started && <button onClick={startGame}>Start</button>}
        </div>
        <Chat messages={state?.chat?.messages || []} socket={socket} groupchat={groupchat}/>
    </div>);
}


const PersonsList = ({users}) => {
    return (<div id="users-list">
        <div id="team-1">
            Team #1
            {users && users[0]?.users.map((user, i) => <div key={i}>
                <p>{user.username} score: {user.score}</p>
            </div>)}
        </div>
        <br/><br/>
        <div id="team-2">
            Team #2
            {users && users[1]?.users.map((user, i) => <div key={i}>
                <p>{user.username} score: {user.score}</p>
            </div>)}
        </div>
    </div>);
}

const Chat = ({messages, socket, groupchat}) => {
    return (<div id={'chat-box'}>
        <div id="messages">
            {messages.map((message, i) => <p key={i} title={message?.timestamp} className={'message'}>{message?.user?.username}: {message?.text}</p>)}
        </div>
        <input type="text" placeholder={'Votre message'}
               onKeyPress={(e) => e.key === 'Enter' && socket.send(JSON.stringify({
                   type: 'new-message',
                   room_name: groupchat,
                   message: e.target.value
               }))}/>
    </div>)
}
