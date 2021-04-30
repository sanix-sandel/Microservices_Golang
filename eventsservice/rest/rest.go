package rest

import (
	"net/http"

	"github.com/Microservices/lib/msgqueue"
	"github.com/Microservices/lib/persistence"
	"github.com/gorilla/mux"
)

func ServeAPI(endpoint string, dbHandler persistence.DatabaseHandler, eventEmitter msgqueue.EventEmitter) error {
	handler := NewEventHandler(dbHandler, eventEmitter)

	r := mux.NewRouter()
	eventsrouter := r.PathPrefix("/events").Subrouter()
	eventsrouter.Methods("GET").Path("/{SearchCriteria}/{search}").HandlerFunc(handler.findEventHandler)
	eventsrouter.Methods("GET").Path("").HandlerFunc(handler.AllEventHandler)
	eventsrouter.Methods("GET").Path("/{eventID}").HandlerFunc(handler.oneEventHandler)
	eventsrouter.Methods("POST").Path("").HandlerFunc(handler.NewEventHandler)

	locationRouter := r.PathPrefix("/locations").Subrouter()
	locationRouter.Methods("GET").Path("").HandlerFunc(handler.allLocationsHandler)
	locationRouter.Methods("POST").Path("").HandlerFunc(handler.newLocationHandler)

	return http.ListenAndServe(endpoint, r)

}
