package tickets

type Ticket struct {
	ID string
	Nombre string
	Email string
	PaisDestino string
	HoraVuelo string
	Precio float64
}

// ejemplo 1
func GetTotalTickets(destination string) (int, error) {}

// ejemplo 2
func GetMornings(time string) (int, error) {}

// ejemplo 3
func AverageDestination(destination string, total int) (int, error) {}
