package model

type GreetingMessage struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	created_at string `json:"created_at"`
	updated_at string `json:"updated_at"`
}