import React from 'react';
import {useForm} from 'react-hook-form';
import useWebSocket, {ReadyState} from "react-use-websocket";
import {SOCKET_URL} from "../const";
import {useHistory} from 'react-router-dom'
import {toast} from 'react-toastify';


export const CreateChannel = () => {
    const history = useHistory();
    const {register, handleSubmit} = useForm();
    const {
        sendJsonMessage,
        readyState,
    } = useWebSocket(SOCKET_URL, {
        onMessage: msg => {
            const data = JSON.parse(msg.data);
            console.log(data);
            if (data.message) {
                toast("Channel created", {type: 'success'});
                history.push(data.data);
            } else toast(data.error, {type: 'error'});
        }
    });

    const onSubmit = data => sendJsonMessage({...data, type: 'create-channel'});

    return (
        <form onSubmit={handleSubmit(onSubmit)}>
            <label>
                Room name
                <input {...register("room_name", {required: true, maxLength: 20, minLength: 3})}/>
            </label>
            <label>
                Username
                <input {...register("username", {required: true, maxLength: 20, minLength: 3})}/>
            </label>
            <input disabled={readyState !== ReadyState.OPEN} type="submit"/>
        </form>
    );
}
