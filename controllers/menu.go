package controllers

import (
	"net/http"

	"../structs"
	"github.com/gin-gonic/gin"
)

func (idb *InDB) GetNavigation(c *gin.Context) {
	var (
		menu structs.Menu
		result gin.H
	)
	id := c.Param("id")
	err := idb.DB.Where("id = ?", id).First(&menu).Error
	if err != nil {
		result = gin.H{
			"result": err.Error(),
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": menu,
			"count":  1,
		}
	}

	c.JSON(http.StatusOK, result)
}

func (idb *InDB) GetNavigations(c *gin.Context) {
	var (
		menus []structs.Menu
		result  gin.H
	)

	idb.DB.Find(&menus)
	if len(menus) <= 0 {
		result = gin.H{
			"result": nil,
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": menus,
			"count":  len(menus),
		}
	}

	c.JSON(http.StatusOK, result)
}

func (idb *InDB) CreateNavigation(c *gin.Context) {
	var (
		menu structs.Menu
		result gin.H
	)
	name := c.PostForm("name")
	url := c.PostForm("url")
	menu.Name = name
	menu.Url = url
	idb.DB.Create(&menu)
	result = gin.H{
		"result": menu,
	}
	c.JSON(http.StatusOK, result)
}

func (idb *InDB) UpdateNavigation(c *gin.Context) {
	id := c.Query("id")
	name := c.PostForm("name")
	url := c.PostForm("url")
	var (
		menu    structs.Menu
		newMenu structs.Menu
		result    gin.H
	)

	err := idb.DB.First(&menu, id).Error
	if err != nil {
		result = gin.H{
			"result": "data not found",
		}
	}
	newMenu.Name = name
	newMenu.Url = url
	err = idb.DB.Model(&menu).Updates(newMenu).Error
	if err != nil {
		result = gin.H{
			"result": "update failed",
		}
	} else {
		result = gin.H{
			"result": "successfully updated data",
		}
	}
	c.JSON(http.StatusOK, result)
}

func (idb *InDB) DeleteNavigation(c *gin.Context) {
	var (
		menu structs.Menu
		result gin.H
	)
	id := c.Param("id")
	err := idb.DB.First(&menu, id).Error
	if err != nil {
		result = gin.H{
			"result": "data not found",
		}
	}
	err = idb.DB.Delete(&menu).Error
	if err != nil {
		result = gin.H{
			"result": "delete failed",
		}
	} else {
		result = gin.H{
			"result": "Data deleted successfully",
		}
	}

	c.JSON(http.StatusOK, result)
}