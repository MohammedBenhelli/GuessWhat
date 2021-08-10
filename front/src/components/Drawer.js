import React, {useContext, useRef, useState} from 'react';
import useWebSocket, {ReadyState} from "react-use-websocket";
import {SOCKET_URL} from "../const";
import {toast} from 'react-toastify';
import CanvasDraw from "react-canvas-draw";
import {SliderPicker} from "react-color";
import {useParams} from 'react-router-dom'
import {SocketContext} from "../context/socket";


export const Drawer = () => {
    const {groupchat} = useParams();
    const socket = useContext(SocketContext);
    const [color, setColor] = useState('#426DEA');
    const drawRef = useRef(null);
    // const {
    //     sendJsonMessage,
    //     readyState,
    //     lastMessage
    // } = useWebSocket(SOCKET_URL, {
    //     onMessage: msg => {
    //         const data = JSON.parse(msg.data);
    //         console.log(data);
    //         if (data.message) {
    //             toast("Updated data", {type: 'success'});
    //         } else toast(data.error, {type: 'error'});
    //     }
    // });

    // console.log(lastMessage)

    const drawingEvent = e => {
        console.log(e);
        socket.send(JSON.stringify({canvas: e.getSaveData(), room_name: groupchat, type: 'update-canvas'}));
    }

    const props = {
        onChange: drawingEvent,
        lazyRadius: 10,
        brushRadius: 4,
        brushColor: color,
        hideGrid: true,
        canvasWidth: 700,
        canvasHeight: 400,
        disabled: false,
        immediateLoading: false,
        hideInterface: false
    };


    return (<>
            <SliderPicker color={color} onChangeComplete={c => setColor(c.hex)}/>
            <CanvasDraw ref={drawRef} {...props} />
            <button onClick={() => drawRef.current.undo()}>Previous</button>
        </>);
}
