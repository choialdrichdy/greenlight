package main

import (
	"fmt"
	"net/http"
	"time"

	"greenlight.aldrich.com/internal/data"
	"greenlight.aldrich.com/internal/validator"
)

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, "Create a new movie")
	var input struct {
		Title   string       `json:"title"`
		Year    int32        `json:"year"`
		Runtime data.Runtime `json:"runtime"`
		Genres  []string     `json:"genres"`
	}
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	movie := &data.Movie{
		Title:   input.Title,
		Year:    input.Year,
		Runtime: input.Runtime,
		Genres:  input.Genres,
	}
	v := validator.New()

	if data.ValidateMovie(v, movie); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}
	fmt.Fprintf(w, "%+v\n", input)
}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	data := data.Movie{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "Casa",
		Runtime:   102,
		Genres:    []string{"Action", "Romance", "Comedy"},
		Version:   1,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"movie": data}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	// fmt.Fprintf(w, "Show movie with the id %d\n", id)

}
