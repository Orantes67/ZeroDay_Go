package server


import (
	"ZeroDay_Go/src/server/handler/polling"

	"net/http"
)



func Run(mux *http.ServeMux) {
	mux.HandleFunc("/Longpolling/maestro", polling.LongPollingMaestro)
	mux.HandleFunc("/shortPolling/alumno", polling.ShortPollingAlumno)
}
