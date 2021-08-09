import React from 'react';
import {useForm} from 'react-hook-form';
import useWebSocket, {ReadyState} from "react-use-websocket";
import {SOCKET_URL} from "../const";


export const CreateChannel = () => {
    const {register, handleSubmit} = useForm();
    const {
        sendJsonMessage,
        lastMessage,
        readyState,
    } = useWebSocket(SOCKET_URL);

    const onSubmit = data => sendJsonMessage({data, type: 'create-channel'});

    return (
        <form onSubmit={handleSubmit(onSubmit)}>
            <label>
                Room name
                <input {...register("room_name", { required: true, maxLength: 20, minLength: 3 })}/>
            </label>
            <label>
                Username
                <input {...register("username", { required: true, maxLength: 20, minLength: 3 })}/>
            </label>
            <input disabled={readyState !== ReadyState.OPEN} type="submit"/>
        </form>
    );
}
