import React, {useContext, useEffect} from 'react';
import {useForm} from 'react-hook-form';
import {useHistory} from 'react-router-dom'
import {toast} from 'react-toastify';
import {SocketContext} from "../context/socket";


export const CreateChannel = () => {
    const history = useHistory();
    const socket = useContext(SocketContext);
    const {register, handleSubmit} = useForm();

    useEffect(() => {
        socket.addEventListener('message', e => {
            const data = JSON.parse(e.data);
            console.log(data);
            if (data.message) {
                toast("Channel created", {type: 'success'});
                history.push(data.data);
            } else toast(data.error, {type: 'error'});
        });
    }, []);

    const onSubmit = data => socket.send(JSON.stringify({...data, type: 'create-channel'}));

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
            <input type="submit"/>
        </form>
    );
}
