package api

import (
	"database/sql"
	"fmt"
)

type Recipe struct {
	RecipeID   int    `json:"id"`
	Name       string `json:"name"`
	PrepTime   string `json:"preptime"`
	Difficulty int    `json:"difficulty"`
	IsVeg      bool   `json:"isveg"`
}

type RecipeClient interface {
	GetAllRecipes() ([]Recipe, error)
	CreateRecipe(recipe Recipe) error
	GetRecipeById(id int) (Recipe, error)
	UpdateRecipe(recipe Recipe) error
	DeleteRecipeById(id int) error
	Search(search string) ([]Recipe, error)
}

type RecipeImpl struct {
	Db *sql.DB
}

func (rp *RecipeImpl) GetAllRecipes() ([]Recipe, error) {
	statement := fmt.Sprintf("SELECT * FROM recipes")
	rows, err := rp.Db.Query(statement)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	recipes := make([]Recipe, 0)

	for rows.Next() {
		var recipe Recipe
		err := rows.Scan(&recipe.RecipeID, &recipe.Name, &recipe.PrepTime, &recipe.Difficulty, &recipe.IsVeg)

		if err != nil {
			return nil, err
		}

		recipes = append(recipes, recipe)
	}

	return recipes, nil
}

func (rp *RecipeImpl) CreateRecipe(recipe Recipe) error {
	var lastId int
	statement := fmt.Sprintf("INSERT INTO recipes(Name, PrepTime, Difficulty, IsVeg) VALUES($1, $2, $3, $4) returning RecipeID")
	err := rp.Db.QueryRow(statement, recipe.Name, recipe.PrepTime, recipe.Difficulty, recipe.IsVeg).Scan(&lastId)
	if err != nil {
		return err
	}

	//_, err = rp.Db.Exec("INSERT INTO reciperatings VALUES($1)", lastId)

	return err
}

func (rp *RecipeImpl) GetRecipeById(id int) (Recipe, error) {
	statement := fmt.Sprintf("SELECT * FROM recipes WHERE RecipeID = $1")
	row := rp.Db.QueryRow(statement, id)

	var recipe Recipe
	err := row.Scan(&recipe.RecipeID, &recipe.Name, &recipe.PrepTime, &recipe.Difficulty, &recipe.IsVeg)

	if err != nil {
		return Recipe{}, err
	}

	return recipe, nil
}

func (rp *RecipeImpl) UpdateRecipe(recipe Recipe) error {
	statement := fmt.Sprintf("UPDATE recipes SET Name = $1, Difficulty= $2, PrepTime= $3, IsVeg=$4 WHERE RecipeID = $5")
	_, err := rp.Db.Exec(statement, recipe.Name, recipe.Difficulty, recipe.PrepTime, recipe.IsVeg, recipe.RecipeID)

	return err
}

func (rp *RecipeImpl) DeleteRecipeById(id int) error {
	statement := fmt.Sprintf("DELETE FROM recipes * WHERE RecipeID = $1")
	_, err := rp.Db.Exec(statement, id)

	return err
}

func (rp *RecipeImpl) Search(search string) ([]Recipe, error) {
	statement := "SELECT * FROM recipes WHERE Name LIKE '%' || $1 || '%'"
	rows, err := rp.Db.Query(statement, search)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	recipes := make([]Recipe, 0)

	for rows.Next() {
		var recipe Recipe
		err := rows.Scan(&recipe.RecipeID, &recipe.Name, &recipe.PrepTime, &recipe.Difficulty, &recipe.IsVeg)

		if err != nil {
			return nil, err
		}

		recipes = append(recipes, recipe)
	}

	return recipes, nil
}
