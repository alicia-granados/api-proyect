package handlers

import (
	"API/db"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func APIRequest(apiURL string) ([]byte, error) {
	// Make a GET request to the League of Legends API
	response, err := http.Get(apiURL)
	if err != nil {
		return nil, fmt.Errorf("error making the request: %v", err)
	}
	// Ensure that the body is closed properly
	defer response.Body.Close()

	// Check the status code of the response
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("the API returned an invalid status code: %v", response.Status)
	}

	// Read the entire content of the response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading the response body: %v", err)
	}

	return body, nil
}

func ProcessChampions(infCampeon Champion) {

	championID, err := db.GetChampionID(infCampeon.Id)
	if err != nil {
		log.Fatalf("Error getting the champion ID:%s", err)
	}

	if championID == 0 {
		// Insert the champion and get its ID
		_, err := db.InsertChampion(infCampeon.Name, infCampeon.Title, infCampeon.Lore)
		if err != nil {
			log.Fatalf("Error inserting the champion: %s", err)
		}
	}

}

func ProcessTags(tags []string, championName string) {

	// Iterate over the champion's tags
	for _, tag := range tags {

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
