package dto

type Meta struct {
	Total      int         `json:"total"`
	Page       int         `json:"page"`
	NextCursor interface{} `json:"next_cursor"`
}
