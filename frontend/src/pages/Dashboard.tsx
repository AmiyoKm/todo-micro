import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import { CheckCircle, Circle, Plus, Trash2 } from "lucide-react";
import { useState } from "react";
import {
    AlertDialog,
    AlertDialogAction,
    AlertDialogCancel,
    AlertDialogContent,
    AlertDialogDescription,
    AlertDialogFooter,
    AlertDialogHeader,
    AlertDialogTitle,
    AlertDialogTrigger,
} from "../components/ui/alert-dialog";
import { Badge } from "../components/ui/badge";
import { Button } from "../components/ui/button";
import { Card, CardContent, CardHeader, CardTitle } from "../components/ui/card";
import { Input } from "../components/ui/input";
import { createTodo, deleteTodo, getTodos, updateTodo } from "../services/todo";

export function Dashboard() {

  const queryClient = useQueryClient();
  const [newTodoTitle, setNewTodoTitle] = useState("");
  const [newTodoDesc, setNewTodoDesc] = useState("");


  const { data, isLoading, error } = useQuery({
    queryKey: ["todos"],
    queryFn: getTodos,
  });

  const createMutation = useMutation({
    mutationFn: createTodo,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["todos"] });
      setNewTodoTitle("");
      setNewTodoDesc("");

    },
  });

  const updateMutation = useMutation({
    mutationFn: updateTodo,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["todos"] });
    },
  });

  const deleteMutation = useMutation({
    mutationFn: deleteTodo,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["todos"] });
    },
  });

  const handleCreate = (e: React.FormEvent) => {
    e.preventDefault();
    if (!newTodoTitle.trim()) return;
    createMutation.mutate({
      title: newTodoTitle,
      description: newTodoDesc,
      done: false,
    });
  };

  const toggleTodo = (todo: any) => {
    updateMutation.mutate({ ...todo, done: !todo.done });
  };



  if (isLoading) return <div className="p-8 text-center">Loading todos...</div>;
  if (error) return <div className="p-8 text-center text-red-500">Error loading todos</div>;

  const todos = data?.data.todos || [];

  return (
    <div className="min-h-screen bg-gray-50 dark:bg-gray-950 p-4 md:p-8">
      <div className="max-w-4xl mx-auto space-y-6">
        <div className="flex items-center justify-between mb-8">
            <h1 className="text-3xl font-bold tracking-tight">My Todos</h1>
        </div>

        {/* Create Todo Section */}
        <Card>
          <CardHeader>
            <CardTitle className="text-lg">Add New Task</CardTitle>
          </CardHeader>
          <CardContent>
            <form onSubmit={handleCreate} className="space-y-4">
              <div className="flex flex-col gap-4 md:flex-row">
                <Input
                  placeholder="What needs to be done?"
                  value={newTodoTitle}
                  onChange={(e) => setNewTodoTitle(e.target.value)}
                  className="flex-1"
                />
                <Input
                  placeholder="Description (optional)"
                  value={newTodoDesc}
                  onChange={(e) => setNewTodoDesc(e.target.value)}
                  className="flex-1"
                />
                <Button type="submit" disabled={createMutation.isPending}>
                  <Plus className="mr-2 h-4 w-4" />
                  Add
                </Button>
              </div>
            </form>
          </CardContent>
        </Card>

        {/* Todo List */}
        <div className="space-y-4">
          {todos.length === 0 ? (
            <div className="text-center py-12 text-muted-foreground">
              No todos yet. Add one above!
            </div>
          ) : (
            todos.map((todo: any) => (
              <Card key={todo.id} className="transition-all hover:shadow-md">
                <CardContent className="p-4 flex items-center gap-4">
                  <Button
                    variant="ghost"
                    size="icon"
                    className={`shrink-0 ${todo.done ? "text-green-500 hover:text-green-600" : "text-gray-400 hover:text-gray-500"}`}
                    onClick={() => toggleTodo(todo)}
                  >
                    {todo.done ? (
                      <CheckCircle className="h-6 w-6" />
                    ) : (
                      <Circle className="h-6 w-6" />
                    )}
                  </Button>

                  <div className="flex-1 min-w-0">
                    <h3 className={`font-semibold ${todo.done ? "line-through text-muted-foreground" : ""}`}>
                      {todo.title}
                    </h3>
                    {todo.description && (
                      <p className={`text-sm text-muted-foreground truncate ${todo.done ? "line-through" : ""}`}>
                        {todo.description}
                      </p>
                    )}
                  </div>

                  <div className="flex items-center gap-2">
                    <Badge variant={todo.done ? "secondary" : "default"}>
                      {todo.done ? "Done" : "Active"}
                    </Badge>
                    <AlertDialog>
                      <AlertDialogTrigger asChild>
                        <Button
                          variant="ghost"
                          size="icon"
                          className="text-red-500 hover:text-red-600 hover:bg-red-50"
                        >
                          <Trash2 className="h-4 w-4" />
                        </Button>
                      </AlertDialogTrigger>
                      <AlertDialogContent>
                        <AlertDialogHeader>
                          <AlertDialogTitle>Are you absolutely sure?</AlertDialogTitle>
                          <AlertDialogDescription>
                            This action cannot be undone. This will permanently delete your todo.
                          </AlertDialogDescription>
                        </AlertDialogHeader>
                        <AlertDialogFooter>
                          <AlertDialogCancel>Cancel</AlertDialogCancel>
                          <AlertDialogAction
                            className="bg-red-500 hover:bg-red-600"
                            onClick={() => deleteMutation.mutate(todo.id)}
                          >
                            Delete
                          </AlertDialogAction>
                        </AlertDialogFooter>
                      </AlertDialogContent>
                    </AlertDialog>
                  </div>
                </CardContent>
              </Card>
            ))
          )}
        </div>
      </div>
    </div>
  );
}
