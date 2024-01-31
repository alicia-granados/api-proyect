package entities

type Champion struct {
	Id    int
	Name  string
	Title string
	Skins []Skins
	Lore  string
	Tags  []Tags
}

type Skins struct {
	Id          int
	Id_Num      string
	Num         int32
	Id_Champion int
	Name        string
}

type Tags struct {
	Id          int
	Id_Champion int
	Name        string
}
