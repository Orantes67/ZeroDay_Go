package polling


import (
	"encoding/json"
	"net/http"
)

func ShortPollingAlumno(w http.ResponseWriter, r *http.Request) {
	response := struct {
		NuevoAlumno bool `json:"nuevo_alumno"`
	}{
		NuevoAlumno: CheckNewAlumno(),
	}
	json.NewEncoder(w).Encode(response)
}
