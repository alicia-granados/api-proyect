package main

import (
	"API/db"
	"API/handlers"
	"encoding/json"
	"fmt"
	"log"
)

func main() {

	// URL to fetch all champion data
	apiURL := `https://ddragon.leagueoflegends.com/cdn/13.24.1/data/es_MX/champion.json`
	body, err := handlers.APIRequest(apiURL)
	if err != nil {
		log.Fatalln("error fetching champions:", err)
	}

	// Decode JSON into a data structure
	var data handlers.GenericData
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Fatalf("error decoding JSON: %v", err)
	}

	db.Connect()
	defer db.Close()

	for _, championHandler := range data.GenericChampions {

		infoCampeonURL := fmt.Sprintf("https://ddragon.leagueoflegends.com/cdn/13.24.1/data/es_MX/champion/%s.json", championHandler.Id)
		fmt.Println(infoCampeonURL)

		// Get specific champion information
		infoBody, err := handlers.APIRequest(infoCampeonURL)

		if err != nil {
			fmt.Println("Error fetching champion information:", err)
			return
		}
		var infoCampeon handlers.Data
		err = json.Unmarshal(infoBody, &infoCampeon)
		if err != nil {
			log.Fatalf("error decoding JSON: %v", err)
		}

		// Access the champion's
		infCampeon := infoCampeon.Champion[championHandler.Id]
		handlers.ProcessChampions(infCampeon)

		// Access the champion's tags
		tags := infoCampeon.Champion[championHandler.Id].Tags
		handlers.ProcessTags(tags, championHandler.Id)

		// Iterate over the champion's skins
		// Access the champion's skins
		skins := infoCampeon.Champion[championHandler.Id].Skins
		handlers.ProcesSkins(skins, championHandler.Id)

	}

}
