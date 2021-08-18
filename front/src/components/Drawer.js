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

    useEffect(() => {
        socket.onmessage = e => {
            const {data, error, message} = JSON.parse(e.data);
            console.log(JSON.parse(e.data));
            if (message === 'update-canvas' && !isAdmin) {
                console.log(data)
                drawRef.current.loadSaveData(data, true);
            } else if (message === 'is-admin') {
                setIsAdmin(data === 'true');
            }else if (message === 'is-drawer') {
                setIsDrawer(data === 'true');
            } else if (message === 'start-game'){
                setState(JSON.parse(data));
                console.log(state)
                setIsStarted(true);
                socket.send(JSON.stringify({type: 'is-drawer', room_name: groupchat}));
            } else toast(error, {type: 'error'});
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


    return (<>
        <SliderPicker color={color} onChangeComplete={c => setColor(c.hex)}/>
        <CanvasDraw ref={drawRef} {...props} />
        {isDrawer && <button onClick={() => drawRef.current.undo()}>Previous</button>}
        {isAdmin && !started && <button onClick={startGame}>Start</button>}
        {started && <h1>{state.word}</h1>}
    </>);
}
