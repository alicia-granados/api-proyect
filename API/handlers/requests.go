package handlers

import (
	"API/db"
	"fmt"
	"io/ioutil"
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

func ProcessChampions(dbRepo db.DatabaseRepo, infCampeon Champion) error {
	championID, err := dbRepo.GetChampionID(infCampeon.Name)
	if err != nil {
		return fmt.Errorf("error getting the champion ID:%v", err)
	}
	if championID == 0 {
		// Insert the champion and get its ID
		_, err := dbRepo.InsertChampion(infCampeon.Name, infCampeon.Title, infCampeon.Lore)
		if err != nil {
			return fmt.Errorf("error inserting the champion: %v", err)
		}
	}
	return nil

}

func ProcessTags(dbRepo db.DatabaseRepo, tags []string, championName string) error {

	// Iterate over the champion's tags
	for _, tag := range tags {

		championID, err := dbRepo.GetChampionID(championName)
		if err != nil {
			return fmt.Errorf("error getting the champion ID:%v", err)
		}
		err = dbRepo.InsertTag(championID, tag)
		if err != nil {
			return fmt.Errorf("error inserting the tag:%v", err)
		}

	}
	return nil
}

func ProcesSkins(dbRepo db.DatabaseRepo, skins []Skins, championName string) error {

	for _, skin := range skins {
		// Check if the skin already exists in the Skins table
		skinID, err := dbRepo.GetSkinID(skin.Id_Num)
		if err != nil {
			return fmt.Errorf("error getting the skin ID:%v", err)
		}
		// If the skin doesn't exist, insert it and get its ID
		if skinID == 0 {
			championID, err := dbRepo.GetChampionID(championName)
			if err != nil {
				return fmt.Errorf("error getting the champion ID:%v", err)
			}
			err = dbRepo.InsertSkins(skin.Id_Num, skin.Num, championID, skin.Name)
			if err != nil {
				return fmt.Errorf("error inserting the skin: %v", err)
			}
		}
	}
	return nil
}
