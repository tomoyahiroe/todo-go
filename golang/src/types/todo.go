package types

type Todo struct {
	ID          int    `json:"id"`
	Task        string `json:"task"`
	Description string `json:"description"`
	Created_at  string `json:"created_at"`
	Updated_at  string `json:"updated_at"`
}

type CreateTodo struct {
	Task        string `json:"task"`
	Description string `json:"description"`
}
