package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"raw-http/todoFunctions"
)

var todos = todofunction.NewTodos()

func main() {

	http.HandleFunc("/todos" , todosHandler)
	http.HandleFunc("/todos/" , todoHandler)

	fmt.Println("Server is listening")
	http.ListenAndServe(":8080" , nil)
}

func todosHandler(w http.ResponseWriter , r * http.Request){
	switch r.Method{
	case http.MethodGet:getTodos(w)
	case http.MethodPost:createTodo(w,r)
	default:
		http.Error(w,"Method not allowed" , http.StatusMethodNotAllowed)
	}
}


func todoHandler (w http.ResponseWriter , r * http.Request ){
	id  , err := parseID(r.URL.Path)
	if err != nil{
		http.Error(w , "Invalid Id" , http.StatusBadRequest)
		return 
	}
	switch r.Method{
	case http.MethodGet: getTodoById(w , id)
	case http.MethodPut: updatedTodo(w ,r, id)
	case http.MethodDelete: deleteTodo(w, id)
	default:
		http.Error(w,"Method not allowed" , http.StatusMethodNotAllowed)
	}
	

}


func parseID(path string)(int , error){

	var id int
	_, err := fmt.Sscanf(path , "/todos/%id", &id)
	return id , err

}




func getTodos(w http.ResponseWriter){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos.GetAll())
}


func createTodo(w http.ResponseWriter , r * http.Request){
	var todo todofunction.Todo
	if err :=json.NewDecoder(r.Body).Decode(&todo); err !=nil{
		http.Error(w , err.Error(), http.StatusBadRequest)
		return 
	}

	createTodo := todos.Add(todo)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createTodo)

}


func getTodoById(w http.ResponseWriter,  id int ){

	for _, todo := range todos.GetAll(){
		if todo.ID ==id{
			w.Header().Set("Content-Type" , "application/json")
			json.NewEncoder(w).Encode(todo)
			return 
		}
	}
	http.NotFound(w , nil)
}

func updatedTodo( w http.ResponseWriter, r *http.Request, id int){
	var updatedTodo todofunction.Todo

	if err := json.NewDecoder(r.Body).Decode(&updatedTodo); err != nil{
		http.Error(w, err.Error() , http.StatusBadRequest)
		return 
	}

	if todo, ok:=todos.Update(id , updatedTodo); ok{
		w.Header().Set("Content-Type" , "application/json")
		json.NewEncoder(w).Encode(todo)
	}else {
		http.NotFound(w,nil)	
	}

}


func deleteTodo ( w http.ResponseWriter, id int){
	if todos.Delete(id){
		w.WriteHeader(http.StatusNoContent)
	}else {
		http.NotFound(w , nil)
	}
}