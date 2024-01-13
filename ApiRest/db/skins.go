package db

import (
	"ApiRest/models"
)

type Skins []models.Skins

func (r *RealDBRepo) GetSkins() (Skins, error) {
	sql := "SELECT Skins.Id, Skins.Id_Num, Skins.Num, Skins.Id_Champion, Skins.Name FROM Skins"

	Skinss := Skins{}
	rows, err := r.DB.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		Skins := models.Skins{}
		err := rows.Scan(&Skins.Id, &Skins.Id_Num, &Skins.Num, &Skins.Id_Champion, &Skins.Name)
		if err != nil {
			return nil, err
		}

		// AAdd the data object to the Skins slice
		Skinss = append(Skinss, Skins)
	}

	// Check for errors after exiting the loop
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return Skinss, nil
}

func (r *RealDBRepo) GetSkinId(SkinsId int) (Skins, error) {
	sql := "SELECT Skins.Id, Skins.Id_Num, Skins.Num, Skins.Id_Champion, Skins.Name FROM Skins WHERE Skins.Id = ? "

	Skinss := Skins{}
	rows, err := r.DB.Query(sql, SkinsId)
	if err != nil {
		return nil, err
	}
	defer rows.Close() // Make sure to close the rows at the end

	for rows.Next() {
		Skins := models.Skins{}

		err := rows.Scan(&Skins.Id, &Skins.Id_Num, &Skins.Num, &Skins.Id_Champion, &Skins.Name)

		if err != nil {
			return nil, err
		}

		Skinss = append(Skinss, Skins)

	}

	// Check for errors after exiting the loop
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return Skinss, nil
}
