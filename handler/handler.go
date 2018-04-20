package handler

import (
	"GoDocRestPostgres/api"
	"GoDocRestPostgres/config"
	"GoDocRestPostgres/utils"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var Logger *log.Logger

type RecipeHandler struct {
	Conf *config.Configuration
}

func InitializeHandler(router *mux.Router, conf *config.Configuration) {
	rh := RecipeHandler{
		Conf: conf,
	}
	Logger = conf.Logger
	router.HandleFunc("/recipes", rh.RecipesHandler).Methods("GET", "POST")
	router.HandleFunc("/recipes/search/", rh.SearchHandler).Methods("GET")
	router.HandleFunc("/recipes/{id}", rh.RecipeHandler).Methods("GET", "DELETE", "PUT")
}

func (rh *RecipeHandler) RecipesHandler(res http.ResponseWriter, req *http.Request) {
	conf := rh.Conf

	switch req.Method {

	case "GET":
		Logger.Printf("\n GET::Route : %s", "/recipes")
		recipes, err := conf.Client.GetAllRecipes()

		if err != nil {
			Logger.Println("error::", err)
			utils.ProcessError(res, err)
			return
		}

		utils.ProcessJson(res, http.StatusOK, recipes)

	case "POST":
		Logger.Printf("\n POST::Route : %s", "/recipes")
		recipeRequest, err := GetRecipeFromRequest(req)
		if err != nil {
			Logger.Println("error::", err)
			utils.ProcessError(res, err)
			return
		}

		err = conf.Client.CreateRecipe(recipeRequest)
		if err != nil {
			Logger.Println("error::", err)
			utils.ProcessError(res, err)
			return
		}

		utils.ProcessJson(res, http.StatusOK, map[string]string{"Result": "Successfully created"})
	}
}

func (rh *RecipeHandler) RecipeHandler(res http.ResponseWriter, req *http.Request) {
	conf := rh.Conf
	vars := mux.Vars(req)
	recipeId, err := strconv.Atoi(vars["id"])
	if err != nil {
		Logger.Println("error::", err)
		utils.ProcessError(res, err)
		return
	}

	switch req.Method {

	case "GET":
		Logger.Printf("\n GET::Route : %s :: recipeId %d", "/recipe", recipeId)
		recipe, err := conf.Client.GetRecipeById(recipeId)

		if err != nil {
			Logger.Println("error::", err)
			utils.ProcessError(res, err)
			return
		}

		utils.ProcessJson(res, http.StatusOK, recipe)

	case "DELETE":
		Logger.Printf("\n DELETE::Route : %s :: recipeId %d", "/recipe", recipeId)

		err = conf.Client.DeleteRecipeById(recipeId)
		if err != nil {
			Logger.Println("error::", err)
			utils.ProcessError(res, err)
			return
		}

		utils.ProcessJson(res, http.StatusOK, map[string]string{"Result": "Recipe successfully delete"})

	case "PUT":
		Logger.Printf("\n UPDATE::Router : %s :: recipeId %d", "/recipe", recipeId)

		recipeRequest, err := GetRecipeFromRequest(req)
		if err != nil {
			Logger.Println("error::", err)
			utils.ProcessError(res, err)
			return
		}

		recipeRequest.RecipeID = recipeId
		err = conf.Client.UpdateRecipe(recipeRequest)

		if err != nil {
			Logger.Println("error::", err)
			utils.ProcessError(res, err)
			return
		}

		utils.ProcessJson(res, http.StatusOK, map[string]string{"Result": "Recipe successfully updated"})
	}
}

func (rh *RecipeHandler) SearchHandler(res http.ResponseWriter, req *http.Request) {
	conf := rh.Conf
	//param := req.FormValue("search")
	param := req.URL.Query().Get("query")
	fmt.Println(param)

	Logger.Printf("\n GET::Route : %s :: Queryparam::%s", "/recipe/search", param)

	recipes, err := conf.Client.Search(param)

	if err != nil {
		Logger.Println("error::", err)
		utils.ProcessError(res, err)
		return
	}

	utils.ProcessJson(res, http.StatusOK, recipes)

}

func GetRecipeFromRequest(req *http.Request) (api.Recipe, error) {

	var recipe api.Recipe

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return recipe, errors.New("error in the request body")
	}

	err = json.Unmarshal(body, &recipe)
	if err != nil {
		return recipe, errors.New("Unable to unmarshal request body")
	}

	// you can also use decoder instead of unmarshalling
	// err := json.NewDecoder(req.Body).Decode(&recipe)
	// if err != nil {
	// 	return NewRecipe{}, errors.New("Invalid request body")
	// }
	return recipe, nil
}
