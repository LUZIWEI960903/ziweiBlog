package dao

import (
	"log"
	"ziweiBlog/models"
)

func GetAllCategory() ([]models.Category, error) {
	rows, err := DB.Query("select * from blog_category")
	if err != nil {
		log.Fatal("GetAllCategory [DB.Query] error:", err)
		return nil, err
	}

	var categorys []models.Category
	for rows.Next() {
		var category models.Category
		err = rows.Scan(&category.Cid, &category.Name, &category.CreateAt, &category.UpdateAt)
		if err != nil {
			log.Fatal("GetAllCategory [rows.Scan] error:", err)
			return nil, err
		}
		categorys = append(categorys, category)
	}
	return categorys, nil
}
