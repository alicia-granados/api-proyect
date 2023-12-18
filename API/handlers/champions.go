package handlers

type GenericData struct {
	GenericChampions map[string]GenericChampion `json:"data"`
}

type GenericChampion struct {
	Id string `json:"id"`
}

type Data struct {
	Champion map[string]Champion `json:"data"`
}

type Champion struct {
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
