package handlers

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetChampions(apiURL string) ([]byte, error) {
	// peticion get a la api de league of legends
	response, err := http.Get(apiURL)
	if err != nil {
		return nil, fmt.Errorf("error al hacer la solicitud: %v", err)
	}
	//se asegura que el body se cierre bien
	defer response.Body.Close()

	// Verificar el código de estado de la respuesta
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("la API devolvió un código de estado no válido: %v", response.Status)
	}

	// Leer todo el contenido del cuerpo de la respuesta
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {

		return nil, fmt.Errorf("error al leer el cuerpo de la response:%v", err)
	}

	return body, nil
}
