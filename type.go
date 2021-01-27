package main

type a struct {
	X  int    `json:"x"`
	Y  string `json:"y"`
	Id int    `json:"id"`
}

type Person struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Birthday string `json:"birthday"`
}
