import type { AxiosResponse } from "axios";
import api from "./api";

type Todo = {
	id: string;
	title: string;
    description: string;
	done: boolean;
}

export function createTodo(todo : Omit<Todo, "id">): Promise<AxiosResponse<{todo: Todo}>>{
    return api.post("/todos", todo)
}

export function getTodo(id : string): Promise<AxiosResponse<{todo: Todo}>>{
    return api.get(`/todos/${id}`)
}

export function updateTodo(todo : Todo): Promise<AxiosResponse<{todo: Todo}>>{
    return api.patch(`/todos/${todo.id}`, todo)
}

export function deleteTodo(id : string): Promise<AxiosResponse<void>>{
    return api.delete(`/todos/${id}`)
}

export function getTodos(): Promise<AxiosResponse<{todos: Todo[]}>>{
    return api.get("/todos")
}
