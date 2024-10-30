package todofunction

import "sync"


type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

type Todos struct {
	sync.Mutex 
	List []Todo
	IDcounter int 
}

func NewTodos() *Todos {
	return &Todos{
		List: []Todo{},
		IDcounter: 1,
	}
}

func (t *Todos) GetAll() []Todo {
	t.Lock()
	defer t.Unlock()
	return t.List

}

func (t * Todos) Add (todo Todo) Todo {
	t.Lock()
	defer t.Unlock()

	todo.ID = t.IDcounter
	t.IDcounter++
	t.List = append(t.List, todo)
	return todo
}

func (t *Todos) Update(id int, updatedTodo Todo) (Todo, bool) {
	t.Lock()
	defer t.Unlock()

	for i , todo := range t.List{
		if todo.ID == id{
			t.List[i].Title= updatedTodo.Title
			t.List[i].Done= updatedTodo.Done
			return t.List[i] , true 
		}
	}
	return Todo{}, false
}


func (t * Todos) Delete(id int) bool{
	t.Lock()
	defer t.Unlock()

	for i, todo :=range t.List {
		if todo.ID == id{
			t.List = append(t.List, t.List[i+1:]...)
			return true
		}
	}
	return false
}