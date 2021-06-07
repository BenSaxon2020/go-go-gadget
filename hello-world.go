// +build ignore

package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("user connected")
	if r.URL.Path != "/favicon.ico" {
		var s = r.URL.Path
		s1 := strings.Trim(s, "/")
		// fmt.Println(" ")
		s2 := strings.Split(s1, "/")
		fmt.Println(s2)
		// fmt.Println(s1)
		// number, error := strconv.Atoi(s1)
		fmt.Println(r.URL.Path[3:])
		num1, error := strconv.Atoi(s2[0])
		var symb string = s2[1]
		num2, error := strconv.Atoi(s2[2])
		if error != nil {
			fmt.Fprintln(w, "\nplease enter just numbers after the backslash")
		} else {

			// fmt.Println("error:", error)
			// fmt.Println("number:", number)
			switch symb {
			case "*":
				var num_tot int = num1 * num2
				fmt.Fprintln(w, "\nYour sum is:", num1, symb, num2, "=", num_tot)
				break
			case "+":
				var num_tot int = num1 + num2
				fmt.Fprintln(w, "\nYour sum is:", num1, symb, num2, "=", num_tot)

				break
			case "D":
				var num_tot int = num1 / num2
				fmt.Fprintln(w, "\nYour sum is:", num1, "/", num2, "=", num_tot)
				break
			case "-":
				var num_tot int = num1 - num2
				fmt.Fprintln(w, "\nYour sum is:", num1, symb, num2, "=", num_tot)
				break
			default:
				fmt.Fprintln(w, "\nplease enter a number then on of the for symbles shown here: +, _, * or D for devide  followed by another number.")
			}

			sqlwrite(num1, num2, symb)
		}
	}

}

var num1 int
var num2 int
var symb string

func main() {
	fmt.Println("server online")
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func sqlwrite(num1 int, num2 int, symb string) {
	db, err := sql.Open("mysql", "[sql_user]:[sql_pass]@tcp([sql_ip]:[sql_port])/[database_name]")
	if err != nil {
		panic(err)
	}

	type Timeline struct {
		Id      int
		Content string
	}
	rows, err := db.Query(`INSERT INTO [table_name] (num1, num2, M_Equ) VALUES (?, ?, ?);`, num1, num2, symb)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	// for rows.Next() {
	// 	timeline := Timeline{}
	// 	err = rows.Scan(&timeline.Id, &timeline.Content)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	fmt.Println(timeline)
	// }
	err = rows.Err()
	if err != nil {
		panic(err)
	}

	defer db.Close()

}
