package db

import (
	"ApiRest/models"
)

type Tags []models.Tags

func (r *RealDBRepo) GetTags() (Tags, error) {
	sql := "SELECT Tags.Id, Tags.Id_Champion, Tags.Name FROM Tags"

	Tagss := Tags{}
	rows, err := r.DB.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		Tags := models.Tags{}
		err := rows.Scan(&Tags.Id, &Tags.IdChampion, &Tags.Name)
		if err != nil {
			return nil, err
		}

		// AAdd the data object to the Tags slice
		Tagss = append(Tagss, Tags)
	}

	// Check for errors after exiting the loop
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return Tagss, nil
}

func (r *RealDBRepo) GetTagId(TagsId int) (Tags, error) {
	sql := "SELECT Tags.Id, Tags.Id_Champion, Tags.Name FROM Tags WHERE Tags.Id = ? "

	Tagss := Tags{}
	rows, err := r.DB.Query(sql, TagsId)
	if err != nil {
		return nil, err
	}
	defer rows.Close() // Make sure to close the rows at the end

	for rows.Next() {
		Tags := models.Tags{}

		err := rows.Scan(&Tags.Id, &Tags.IdChampion, &Tags.Name)

		if err != nil {
			return nil, err
		}

		Tagss = append(Tagss, Tags)

	}

	// Check for errors after exiting the loop
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return Tagss, nil
}
