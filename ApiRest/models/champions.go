package models

type Data struct {
	Champion map[string]Champion `json:"data"`
}

type Champion struct {
	Id    int     `json:"id"`
	Name  string  `json:"name"`
	Title string  `json:"title"`
	Skins []Skins `json:"skins"`
	Lore  string  `json:"lore"`
	Tags  []Tags  `json:"tags"`
}

type Skins struct {
	Id          int    `json:"id"`
	Id_Num      string `json:"idNum"`
	Num         int    `json:"num"`
	Id_Champion int    `json:"idChampion"`
	Name        string `json:"name"`
}

type Tags struct {
	Id          int    `json:"id"`
	Id_Champion int    `json:"idChampion"`
	Name        string `json:"name"`
}