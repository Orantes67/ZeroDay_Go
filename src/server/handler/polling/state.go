package polling


import "sync"

var (
	newMaestroAdded bool
	newAlumnoAdded  bool
	mu              sync.Mutex
)

func SetNewMaestroAdded() {
	mu.Lock()
	newMaestroAdded = true
	mu.Unlock()
}

func CheckAndResetNewMaestro() bool {
	mu.Lock()
	defer mu.Unlock()
	if newMaestroAdded {
		newMaestroAdded = false
		return true
	}
	return false
}

func SetNewAlumnoAdded() {
	mu.Lock()
	newAlumnoAdded = true
	mu.Unlock()
}

func CheckNewAlumno() bool {
	mu.Lock()
	defer mu.Unlock()
	return newAlumnoAdded
}
