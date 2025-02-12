package server

import (
	"CRUD/database"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
)

type user struct {
	ID    uint32 `json:"id"`
	Name  string `json: "name"`
	Email string `json: "email"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		w.Write([]byte("Error reading body"))
		return
	}

	var user user
	err = json.Unmarshal(body, &user)
	if err != nil {
		w.Write([]byte("Error unmarshalling body"))
		return
	}

	db, err := database.Connect()
	if err != nil {
		w.Write([]byte("Error connecting to database"))
		return
	}
	defer db.Close()

	statement, err := db.Prepare("INSERT INTO users(name, email) VALUES(?, ?)")
	if err != nil {
		_, err := w.Write([]byte(err.Error()))
		if err != nil {
			return
		}
	}
	defer statement.Close()

	insert, err := statement.Exec(user.Name, user.Email)
	if err != nil {
		w.Write([]byte("Error inserting user"))
		return
	}

	insertID, err := insert.LastInsertId()
	if err != nil {
		_, err := w.Write([]byte(err.Error()))
		if err != nil {
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("User Created! Id %d", insertID)))
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	db, err := database.Connect()
	if err != nil {
		w.Write([]byte("Error connecting to database"))
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		w.Write([]byte("Error getting users " + err.Error()))
		return
	}
	defer rows.Close()

	var users []user
	for rows.Next() {
		var user user
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			w.Write([]byte("Error scanning row: " + err.Error()))
			return
		}
		users = append(users, user)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(users); err != nil {
		w.Write([]byte("Error encoding response"))
		return
	}
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ID, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		w.Write([]byte("Error Parsing ID" + err.Error()))
		return
	}

	db, err := database.Connect()
	if err != nil {
		w.Write([]byte("Error connecting to database"))
		return
	}
	defer db.Close()

	row, err := db.Query("SELECT * FROM users WHERE id = ?", uint32(ID))
	if err != nil {
		w.Write([]byte("Error getting user " + err.Error()))
		return
	}
	defer row.Close()

	var user user
	if row.Next() {
		if err := row.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			w.Write([]byte("Error scanning row: " + err.Error()))
			return
		}
	}
	defer row.Close()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		w.Write([]byte("Error encoding response"))
		return
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("Error reading body"))
		return
	}
	var user user
	if err := json.Unmarshal(body, &user); err != nil {
		w.Write([]byte("Error unmarshalling body"))
		return
	}

	params := mux.Vars(r)
	ID, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		w.Write([]byte("Error Parsing ID" + err.Error()))
		return
	}

	db, err := database.Connect()
	if err != nil {
		w.Write([]byte("Error connecting to database"))
		return
	}
	defer db.Close()

	statement, err := db.Prepare("UPDATE users SET name = ?, email = ? WHERE id = ?")
	if err != nil {
		w.Write([]byte("Error updating user"))
		return
	}

	if _, err := statement.Exec(user.Name, user.Email, ID); err != nil {
		w.Write([]byte("Error updating user"))
		return
	}
	defer statement.Close()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ID, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		w.Write([]byte("Error Parsing ID" + err.Error()))
	}

	db, err := database.Connect()
	if err != nil {
		w.Write([]byte("Error connecting to database"))
		return
	}
	defer db.Close()

	statement, err := db.Prepare("DELETE FROM users WHERE id = ?")
	if err != nil {
		w.Write([]byte("Error deleting user"))
		return
	}

	if _, err := statement.Exec(ID); err != nil {
		w.Write([]byte("Error deleting user"))
		return
	}

	defer statement.Close()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}
