package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Food struct {
	Name  string `json:"name" form:"name"`
	Price int    `json:"price" form:"price"`
	Halal bool   `json:"halal" form:"halal"`
}

func main() {
	e := echo.New()

	e.GET("/foods", GetAllFoodsController)
	e.GET("/foods/:id", GetFoodController)
	e.POST("/foods", CreateFoodController)
	e.PUT("/foods/:id", UpdateFoodController)
	e.DELETE("/foods/:id", DeleteFoodController)

	e.Start(":8080")
}

var foods []Food

func GetAllFoodsController(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all foods data",
		"foods":   foods,
	})
}

func GetFoodController(c echo.Context) error {
	var food []Food

	fID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	for k, f := range foods {
		if fID == k+1 {
			food = append(food, f)
		}
	}
	if food == nil {
		return c.String(http.StatusNotFound, "invalid food id")
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "data found !",
		"food":    food[0],
	})
}

func CreateFoodController(c echo.Context) error {
	var inputFood Food

	if err := c.Bind(&inputFood); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]string{
			"message": "error fetching data",
		})
	}

	food := Food{
		Name:  inputFood.Name,
		Price: inputFood.Price,
		Halal: inputFood.Halal,
	}
	foods = append(foods, food)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success add data",
		"food":    food,
	})
}

func UpdateFoodController(c echo.Context) error {
	var food *Food

	fID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	if fID > len(foods) {
		return c.JSON(http.StatusNotFound, map[string]string{
			"message": "invalid index",
		})
	}
	for k, _ := range foods {
		if fID == k+1 {
			food = &foods[k]
		}
	}
	if err := c.Bind(food); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]string{
			"message": "error fetching data",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update data",
		"food":    food,
	})
}

func DeleteFoodController(c echo.Context) error {
	fID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	if fID > len(foods) {
		return c.JSON(http.StatusNotFound, map[string]string{
			"message": "invalid index",
		})
	}
	newFoods := foods[:fID-1]
	newFoods = append(newFoods, foods[fID:]...)
	foods = newFoods
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete",
		"foods":   foods,
	})

}
