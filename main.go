package main

import (
	"database/sql"
	"fmt"           /*Пакет вывода на экран*/
	"html/template" /* Пакет вывода html страниц */
	"net/http"      /*Пакет http*/

	/* Пакет работы с SQL */

	_ "github.com/go-sql-driver/mysql" /* Пакет работы с SQL */
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

func index(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/index.html")
	tmpl.Execute(w, r)
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

func pages_sign_up(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/pages-sign-up.html")
	tmpl.Execute(w, r)
}
func pages_sign_in(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/pages-sign-in.html")
	tmpl.Execute(w, r)
}
func pages_profile(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/pages-profile.html")
	tmpl.Execute(w, r)
}

func handleRequest() {
	http.HandleFunc("/", index)
	http.HandleFunc("/home_page/", home_page)
	http.HandleFunc("/contacts/", contacts_page)
	http.HandleFunc("/pages-sign-up/", pages_sign_up)
	http.HandleFunc("/pages-sign-in/", pages_sign_in)
	http.HandleFunc("/pages-profile/", pages_profile)
	// Подключение папки assets
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	http.ListenAndServe(":8080", nil)
}

func main() {
	//var bob User = ....
	//bob := User{ name : "Bob", age: 25, money: -50, avg_grades: 4.4, happiness: 0.8 }
	//bob := User{ "Bob", 25, -50, 4.4, 0.8 }

	db, err := sql.Open("mysql", "cowboy_kpda:22031986Te!@tcp(cowboy.beget.tech)/cowboy_kpda")

	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Установка данных

	insert, err := db.Query("INSERT INTO `users` (`name`,`age`) VALUES ('Bob', 35)")
	if err != nil {
		panic(err)
	}
	defer insert.Close()

	fmt.Println("Подключено в БД")

	handleRequest()
}
