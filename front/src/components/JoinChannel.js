import React from 'react';
import {useForm} from 'react-hook-form';


export const JoinChannel = ({socket}) => {
    const {register, handleSubmit} = useForm();

    const onSubmit = data => socket.send(JSON.stringify({...data, type: 'join-channel'}));

    return (
        <form onSubmit={handleSubmit(onSubmit)}>
            <label>
                <input placeholder={'Nom de la salle'} {...register("room_name", {required: true, maxLength: 20, minLength: 3})}/>
            </label>
            <label>
                <input placeholder={'Pseudo'} {...register("username", {required: true, maxLength: 20, minLength: 3})}/>
            </label>
            <input type="submit" value={'Rejoindre'}/>
        </form>
    );
}
