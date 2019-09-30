package api

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"simple_rest/globals"
	"simple_rest/models"
)



func GetAll(w http.ResponseWriter, r *http.Request) error{
	var rows, err = globals.Pool().Query(context.Background(), "select main.id, inne.id from main inner join inne on main.inne_id = inne.id")
	if err != nil {
		return err
	}
	defer rows.Close()

	var all [] models.Main

	for rows.Next() {
		var main = models.Main{}
		var inner = models.Inner{}

		values, err := rows.Values()
		if err != nil {
			return err
		}
		inner.Id = values[1].(int32)
		main.Id = values[0].(int32)
		main.In = &inner
		all = append(all, main)
	}

	json.NewEncoder(w).Encode(all)
	return nil
}

func GetOne(w http.ResponseWriter, r *http.Request) error {
	params := mux.Vars(r)
	var rows, err = globals.Pool().Query(context.Background(), "select main.id, inne.id from main inner join inne on main.inne_id = inne.id where main.id = $1", params["id"])
	found := false
	if err != nil {
		return err
	}

	defer rows.Close()

	var main = models.Main{}
	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			return err
		}
		inner := models.Inner{}
		inner.Id = values[1].(int32)
		main.Id = values[0].(int32)
		main.In = &inner
		found = true
	}
	if found {
		json.NewEncoder(w).Encode(main)
	}
	json.NewEncoder(w).Encode(nil)
	return nil
}

func CreateOne(w http.ResponseWriter, r *http.Request) error {
	if r.FormValue("pass") == "pass" {
		inner := models.Inner{}
		main := models.Main{In: &inner}
		rows, err := globals.Pool().Query(context.Background(),"insert into inne default values returning id")
		if err != nil {
			return err
		}

		defer rows.Close()

		for rows.Next() {
			values, err := rows.Values()
			if err != nil {
				return err
			}

			inner.Id = values[0].(int32)
		}
		rows, err = globals.Pool().Query(context.Background(), "insert into main (inne_id) values($1) returning id", inner.Id)
		if err != nil {
			return err
		}

		defer rows.Close()

		for rows.Next() {
			values, err := rows.Values()
			if err != nil {
				return err
			}

			main.Id = values[0].(int32)
			main.In = &inner
		}
		json.NewEncoder(w).Encode(main)
		return nil
	}
	json.NewEncoder(w).Encode(models.Main{})
	return nil
}
