import React from 'react';
import {useForm} from 'react-hook-form';


export const CreateChannel = ({socket}) => {
    const {register, handleSubmit} = useForm();

    const onSubmit = data => socket.send(JSON.stringify({...data, type: 'create-channel'}));

    return (
        <form onSubmit={handleSubmit(onSubmit)}>
            <label>
                <input placeholder={'Nom de la salle'} {...register("room_name", {required: true, maxLength: 20, minLength: 3})}/>
                <input placeholder={'Pseudo'} {...register("username", {required: true, maxLength: 20, minLength: 3})}/>
            </label>
            <input type="submit" value={'Creer'}/>
        </form>
    );
}
