package main

import (
	"API/handlers"
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	//url para traer todos los ddatos de los campeones
	apiURL := `https://ddragon.leagueoflegends.com/cdn/13.24.1/data/es_MX/champion.json`

	body, err := handlers.GetChampions(apiURL)
	if err != nil {
		log.Fatalln("error al obtener los campeones:", err)
	}

	// Decodificar el JSON en una estructura de datos
	var data handlers.GenericData
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Fatalf("error al decodificar JSON: %v", err)
	}

	fmt.Println(data)

	for _, championHandler := range data.GenericChampions {

		info_campeom_URL := fmt.Sprintf("https://ddragon.leagueoflegends.com/cdn/13.24.1/data/es_MX/champion/%s.json", championHandler.Id)
		fmt.Println(info_campeom_URL)

		// Obtener la información específica del campeón
		infoBody, err := handlers.GetChampions(info_campeom_URL)

		if err != nil {
			fmt.Println("Error al obtener información del campeón:", err)
			return
		}
		var infoCampeon handlers.Data
		err = json.Unmarshal(infoBody, &infoCampeon)
		if err != nil {
			log.Fatalf("error al decodificar JSON: %v", err)
		}

		// Acceder a las skins del campeón específico
		//fmt.Println("Skins del campeón:", infoCampeon.Champion[championHandler.Id].Skins)
	}

}
