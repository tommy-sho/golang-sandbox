package handler

import (
	// ...
	"net/http"
	"strconv"

	"github.com/tommy-sho/golang-sandbox/web/payjp-vue/backend-api/database"
)

// GetLists - get all items
func GetLists(c Context) {
	res, err := database.SelectAllItems()
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	c.JSON(http.StatusOK, res)
}

// GetItem - get item by id
func GetItem(c Context) {
	identifer, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	res, err := database.SelectItem(int64(identifer))
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	c.JSON(http.StatusOK, res)
}
