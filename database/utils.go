package database

// import "context"

type Todo struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

// func createTable(dbInstance service) error {
// 	_, err := dbInstance.DB.Exec(context.Background(), "create table if not exists todos (id integer primary key, title text, description text)")
// 	if err != nil {
// 		return err
// 	}
//
// 	return nil
// }

// TODO: Add these:
// getTodos() ([]Todo, error)
// createTodo(todo Todo) (bool, error)
// getTodoByID(id int) (Todo, error)
// updateTodoByID(id int) (bool, error)
// deleteTodoByID(id int) (bool, error)
