package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type Item struct {
	Title string
	Body  string
}

type API int

var database []Item

func (a *API) GetByName(title string, reply *Item) error {
	for _, item := range database {
		if item.Title == title {
			*reply = item
			return nil
		}
	}
	return nil
}

func (a *API) AddItem(item Item, reply *Item) error {
	database = append(database, item)
	*reply = item
	return nil
}

func (a *API) EditItem(edit Item, reply *Item) error {
	for i, val := range database {
		if val.Title == edit.Title {
			database[i] = Item{edit.Title, edit.Body}
			*reply = database[i]
			return nil
		}
	}
	return nil
}

func (a *API) DeleteItem(item Item, reply *Item) error {
	for i, val := range database {
		if val.Title == item.Title {
			database = append(database[:i], database[i+1:]...)
			*reply = item
			return nil
		}
	}
	return nil
}

func (a *API) GetDB(title string, reply *[]Item) error {
	*reply = database
	return nil
}

func main() {
	var api = new(API)
	err := rpc.Register(api)

	if err != nil {
		log.Fatal("Error registering API: ", err)
	}

	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", ":4040")

	if err != nil {
		log.Fatal("Error listening: ", err)
	}

	log.Printf("Serving on port %d", 4040)
	err = http.Serve(listener, nil)
	if err != nil {
		log.Fatal("Error serving: ", err)
	}
}
