package main

import "fmt"

type IDataBase interface {
	Query()
}


type Knex struct {
	Port string
}

func (p Knex) Query() {
	fmt.Println("test")
}

type Sequelize struct {

}

func (p Sequelize) Query() {
	fmt.Println("test")
}

func main() {
	var a IDataBase
	a = Knex{
		Port: "123",
	}
	var b Knex = Knex{
		Port: "123"
	}
	//a = Sequelize{
	//}
	//_,  ok := a.(IDataBase)
	//_, ok := a.(Knex)
	//_, ok := a.(Sequelize)
	fmt.Println(a.Port)
	fmt.Println(b.Port)

}

