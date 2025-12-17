import api from "./api";

type User = {
    id: string;
    name: string;
    email: string;
}

export type registerRequest = {
    name : string;
    email : string;
    password : string;
}

export function register(req : registerRequest): Promise<{user: User}>{
    return api.post("/register", req)
}

export type loginRequest = {
    email : string;
    password : string;
}
export function login(req : loginRequest): Promise<{jwt : string}>{
    return api.post("/login", req)
}

export function getMe(): Promise<{user : User}>{
    return api.get(`/users/me`)
}

