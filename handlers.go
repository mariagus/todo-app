package main

import (
	"net/http"
	"text/template"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	todos, err := getAllTodos()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) // Return an HTTP 500 error if the query fails
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/index.html")) // Parse the HTML template
	tmpl.Execute(w, todos)                                             // Render the template with the list of todos

}

func createHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		title := r.FormValue("title")
		urgency := r.FormValue("urgency")
		err := createTodo(title, urgency) // Get the title from the form data
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError) // Return an HTTP 500 error if insertion fails
			return
		}
		http.Redirect(w, r, "/", http.StatusSeeOther) // Redirect to the main page after successful creation
	}
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	err := deleteTodoByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) // Return an HTTP 500 error if deletion fails
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther) // Redirect to the main page after successful deletion
}
