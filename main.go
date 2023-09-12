package main

import (
	"fmt"           /*Пакет вывода на экран*/
	"html/template" /* Пакет вывода html страниц */
	"net/http"      /*Пакет http*/
)

type User struct {
	Name                  string
	Age                   uint16 /* Целове число не может быть отрицательным */
	Money                 int16
	Avg_grades, Happiness float64 /* Число с плавающей точкой */
	Hobbies               []string
}

func (u User) getAllInfo() string {
	return fmt.Sprintf("Ползователь %s имеет возраст %d и зарпралту в"+
		"размере %d.", u.Name, u.Age, u.Money)
}

func (u *User) setNewName(newName string) {
	u.Name = newName
}

func home_page(w http.ResponseWriter, r *http.Request) {
	bob := User{"Bob", 25, -50, 4.4, 0.8, []string{"Footbal", "Basketball", "Skate", "Dance"}}
	//fmt.Fprintf(w, "<h1>Title</h1>")
	tmpl, _ := template.ParseFiles("templates/home_page.html")
	tmpl.Execute(w, bob)
}

func contacts_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Contacts easy!!")
}

func handleRequest() {
	http.HandleFunc("/", home_page)
	http.HandleFunc("/contacts/", contacts_page)
	// Подключение папки assets
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	http.ListenAndServe(":8080", nil)
}

func main() {
	//var bob User = ....
	//bob := User{ name : "Bob", age: 25, money: -50, avg_grades: 4.4, happiness: 0.8 }
	//bob := User{ "Bob", 25, -50, 4.4, 0.8 }
	handleRequest()
}
