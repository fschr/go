package models

type (
	User struct {
		Name string `json:"name"`
		Id 	 int	`json:"id"`
		Rank int	`json:"rank"`
	}
)