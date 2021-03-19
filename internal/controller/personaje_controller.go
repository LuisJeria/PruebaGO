package controller

import (
	"encoding/json"
	"fmt"
	"hola/internal/domain"
	"hola/internal/service"
	"net/http"
)

type PersonajeController struct {
	service *service.PersonajeService
}

func NewPersonajeController(s *service.PersonajeService) *PersonajeController {
	personajeController := &PersonajeController{}
	personajeController.personajeRequestHandler()
	personajeController.service = s
	return personajeController

}

func (p *PersonajeController) personajeRequestHandler() {
	http.HandleFunc("/personaje",
		func(w http.ResponseWriter, r *http.Request) {

			switch m := r.Method; m {
			case "POST":
				w.WriteHeader(200)
				var personaje domain.Personaje
				err := json.NewDecoder(r.Body).Decode(&personaje)
				if err != nil {
					fmt.Fprintf(w, "400 bad request")
					return
				}
				err = p.service.Save(personaje)
				if err != nil {
					fmt.Fprint(w, "error de guardado: "+err.Error())
					return
				}
				fmt.Fprintf(w, "El dato ha sido guardado exitosamente")
			default:
				w.WriteHeader(405)
				fmt.Fprintf(w, "405 Method Not Allowed")
			}

		})
}
