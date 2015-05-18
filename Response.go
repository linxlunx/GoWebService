package main

type NotFound struct {
	Status int `json:"status"`
	Message string `json:"message"`
	Result struct {} `json:"result,omitempty"`
}

type ResponseDetail struct {
	Status 	 int `json:"status"`
	Message string `json:"message"`
	Result User `json:"result"`
}

type ResponseList struct {
	Status 	 int `json:"status"`
	Message string `json:"message"`
	Result []*User `json:"result"`
}

type User struct {
	Rowid 	 int `json:"id"`
	Username string `json:"username"`
	Fullname string `json:"fullname"`
}