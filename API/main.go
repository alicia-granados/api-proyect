package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Data struct {
	Campeon map[string]Campeon `json:"data"`
}

type Campeon struct {
	Id    string   `json:"id"`
	Name  string   `json:"name"`
	Title string   `json:"title"`
	Skins []Skins  `json:"skins"`
	Lore  string   `json:"lore"`
	Tags  []string `json:"tags"`
}

type Skins struct {
	ID   string `json:"id"`
	Num  int32  `json:"num"`
	Name string `json:"name"`
}

// Definir la estructura de la url de la imagen
type imageURL struct {
	Url  string
	Name string
}

func get_Campions(apiURL string) ([]byte, error) {
	// peticion get a la api de league of legends
	response, err := http.Get(apiURL)
	if err != nil {
		return nil, fmt.Errorf("error al hacer la solicitud: %v", err)
	}
	//se asegura que el body se cierre bien
	defer response.Body.Close()

	// Verificar el código de estado de la respuesta
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("La API devolvió un código de estado no válido: %v", response.Status)
	}

	// Leer todo el contenido del cuerpo de la respuesta
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {

		return nil, fmt.Errorf("Error al leer el cuerpo de la response:%v", err)
	}
	return body, nil
}

func decoJSON(body []byte) (*Data, error) {

	// Decodificar el JSON en una estructura de datos
	var data Data
	err := json.Unmarshal(body, &data)
	if err != nil {
		return nil, fmt.Errorf("Error al decodificar JSON: %v", err)
	}
	return &data, nil

}
func main() {
	//url para traer todos los ddatos de los campeones
	apiURL := `https://ddragon.leagueoflegends.com/cdn/13.24.1/data/es_MX/champion.json`

	body, errAllCampions := get_Campions(apiURL)

	if errAllCampions != nil {
		fmt.Println("Error al obtener los campeones:", errAllCampions)
		return
	}

	data, errDecoJSON := decoJSON(body)
	if errDecoJSON != nil {
		fmt.Println("Error al decodificar json", errDecoJSON)
		return
	}

	// Obtener los primeros 7 campeones
	var primeros7Campeones []Campeon
	i := 0
	for _, campeon := range data.Campeon {
		if i < 1 {
			primeros7Campeones = append(primeros7Campeones, campeon)
			//id_campeon := primeros7Campeones[i].Id

			i++

		} else {
			break
		}

	}

	// Imprimir los primeros 7 campeones
	fmt.Println("Primeros 7 campeones:", primeros7Campeones[0].Id)

	// Realizar una segunda solicitud para obtener información específica del campeón
	info_campeom_URL := fmt.Sprintf("https://ddragon.leagueoflegends.com/cdn/13.24.1/data/es_MX/champion/%s.json", primeros7Campeones[0].Id)
	fmt.Println(info_campeom_URL)

	// Obtener la información específica del campeón
	infoBody, errCampeon := get_Campions(info_campeom_URL)

	if errCampeon != nil {
		fmt.Println("Error al obtener información del campeón:", errCampeon)
		return
	}

	campeonData, errDecoJSONCD := decoJSON(infoBody)
	if errDecoJSONCD != nil {
		fmt.Println("Error al decodificar json", errDecoJSONCD)
		return
	}
	fmt.Println(campeonData)

	campeonEspecifico := campeonData.Campeon[primeros7Campeones[0].Id]

	// Acceder a las skins del campeón específico
	fmt.Println("Skins del campeón:", campeonEspecifico.Skins)

	var URL_image []imageURL
	// Obtener url de los SKINS de los campeones
	for _, campeonSkin := range campeonEspecifico.Skins {
		URL_image = append(URL_image, imageURL{
			Url:  fmt.Sprintf("https://ddragon.leagueoflegends.com/cdn/13.24.1/data/es_MX/champion/%s_%d.json", campeonSkin.Name, campeonSkin.Num),
			Name: campeonSkin.Name,
		})
	}
	fmt.Println(URL_image)
}
