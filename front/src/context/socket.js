import React from "react";
import { SOCKET_URL } from "../const";

export const socket = new WebSocket(SOCKET_URL);
export const SocketContext = React.createContext(undefined, undefined);
