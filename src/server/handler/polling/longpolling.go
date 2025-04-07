package polling


import (
	"net/http"
	"time"

)

func LongPollingMaestro(w http.ResponseWriter, r *http.Request) {
	timeout := time.After(30 * time.Second)
	CheckAndResetNewMaestro() // Adjusted to call the function from the correct package

	for {
		select {
		case <-timeout:
			http.Error(w, "Timeout sin nuevos maestros", http.StatusRequestTimeout)
			return
		case <-time.After(1 * time.Second):
			if CheckAndResetNewMaestro() {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{"mensaje": "Nuevo maestro agregado"}`))
				return
			}
		}
	}
}
