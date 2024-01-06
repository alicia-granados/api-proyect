package entities

type Champion struct {
	Id    string
	Name  string
	Title string
	Skins []Skins
	Lore  string
	Tags  []Tags
}

type Skins struct {
	Id_Num string
	Num    int32
	Name   string
}

type Tags struct {
	Id_Champion int
	Name        string
}
