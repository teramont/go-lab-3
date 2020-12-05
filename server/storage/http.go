package machines

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/teramont/go-lab-3/server/tools"
)

type HttpHandlerFunc http.HandlerFunc

// HttpHandler creates a new instance of channels HTTP handler.
func HttpHandler(storage *Storage) HttpHandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			handleListMachines(storage, rw)
		} else if r.Method == "PATCH" {
			handleDiskConnection(r, rw, storage)
		} else {
			rw.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

func handleDiskConnection(r *http.Request, rw http.ResponseWriter, store *Storage) {
	var conn Connect
	if err := json.NewDecoder(r.Body).Decode(&conn); err != nil {
		log.Printf("Error decoding channel input: %s", err)
		tools.WriteJsonBadRequest(rw, "bad JSON payload")
		return
	}
	err := store.Connect(conn)
	if err == nil {
		tools.WriteJsonOk(rw, &conn)
	} else {
		log.Printf("Error connecting disk: %s", err)
		tools.WriteJsonInternalError(rw)
	}
}

func handleListMachines(storage *Storage, rw http.ResponseWriter) {
	res, err := storage.ListMachines()
	if err != nil {
		log.Printf("Error making query to the db: %s", err)
		tools.WriteJsonInternalError(rw)
		return
	}
	tools.WriteJsonOk(rw, res)
}
