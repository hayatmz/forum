package model

// Check if the category is existing in the table categories.
//
// If the category exist return its id.
//
// If the category doesn't exist, insert it in the table category,
// and return its id.
func getIdCategory(category string) (int64, error) {
	var err error
	var idCategory int64

	queryCategory := "SELECT id FROM categories WHERE category = ?"
	db.QueryRow(queryCategory, category).Scan(&idCategory)

	if idCategory == 0 {
		queryCategory := "INSERT INTO categories (category) VALUES (?)"
		idCategory, err = execQuery(queryCategory, category)
		if err != nil {
			return 0, err
		}
	}
	return idCategory, err
}

func getCategoriesPost(idPost string) []Category {
	queryCategories := `SELECT categories.id, categories.category FROM categories INNER JOIN post_categories 
	ON categories.id = post_categories.category_id WHERE post_categories.post_id = ?`
	rows, err := db.Query(queryCategories, idPost)
	if err != nil {
		return nil
	}
	defer rows.Close()

	var categories []Category
	for rows.Next() {
		var category Category
		err := rows.Scan(&category.IDCategory, &category.Category)
		if err == nil {
			categories = append(categories, category)
		}
	}
	return categories
}