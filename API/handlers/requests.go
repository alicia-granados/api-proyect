package handlers

import (
	"API/db"
	"fmt"
	"io/ioutil"
	"log"
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

func ProcessTags(tags []string, championName string) {

	// Iterate over the champion's tags
	for _, tag := range tags {
		fmt.Println("TAGSSS-------------------------", tag)

		championID, err := db.GetChampionID(championName)
		if err != nil {
			log.Fatalf("Error getting the champion ID:%s", err)
			continue
		}
		err = db.InsertTag(championID, tag)
		if err != nil {
			log.Fatalf("Error inserting the tag: %s", err)
		}

	}

}

func ProcesSkins(skins []Skins, championName string) {

	for _, skin := range skins {

		// Check if the skin already exists in the Skins table
		skinID, err := db.GetSkinID(skin.Id_Num)
		if err != nil {
			log.Fatalf("Error getting the skin ID:%s", err)
			continue
		}

		// If the skin doesn't exist, insert it and get its ID
		if skinID == 0 {
			championID, err := db.GetChampionID(championName)
			if err != nil {
				log.Fatalf("Error getting the champion ID:%s", err)
				continue
			}
			err = db.InsertSkins(skin.Id_Num, skin.Num, championID, skin.Name)
			if err != nil {
				log.Fatalf("Error inserting the skin: %s", err)
				continue
			}
		}
	}
}
