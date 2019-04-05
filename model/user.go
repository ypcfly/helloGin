package model

type User struct {
	Username 	string	`json:"name" form:"name"`
	Age  		uint8	`json:"age" form:"age"`
	Mobile 		string	`json:"mobile" form:"mobile"`
	Sex			string	`json:"sex" form:"sex"`
	Address 	string	`json:"address" form:"address"`
	Id          uint16  `json:"id" form:"id"`
}


