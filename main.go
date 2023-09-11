package main

import (
	"fmt"      /*Пакет вывода на экран*/
	"net/http" /*Пакет http*/
)

type User struct {
	name                  string
	age                   uint16 /* Целове число не может быть отрицательным */
	money                 int16
	avg_grades, happiness float64 /* Число с плавающей точкой */
}

func (u User) getAllInfo() string {
	return fmt.Sprintf("Ползователь %s имеет возраст %d и зарпралту в"+
		"размере %d.", u.name, u.age, u.money)
}

func home_page(w http.ResponseWriter, r *http.Request) {
	bob := User{"Bob", 25, -50, 4.4, 0.8}
	fmt.Fprintf(w, bob.getAllInfo())
}

func contacts_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Contacts easy!!")
}

func handleRequest() {
	http.HandleFunc("/", home_page)
	http.HandleFunc("/contacts/", contacts_page)
	http.ListenAndServe(":8080", nil)
}

func main() {
	//var bob User = ....
	//bob := User{ name : "Bob", age: 25, money: -50, avg_grades: 4.4, happiness: 0.8 }
	//bob := User{ "Bob", 25, -50, 4.4, 0.8 }
	handleRequest()
}
