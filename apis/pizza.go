package apis

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"web/models"

	"github.com/gin-gonic/gin"
)

func AdminIndex(c *gin.Context) {
	c.String(http.StatusOK, "AdminIndex")
}

func AddPizza(c *gin.Context) {

	p := models.Pizza{}
	err := c.Bind(&p)
	if err != nil {
		log.Println("11", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}
	ra, err := p.Add()
	if err != nil {
		log.Println("22", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}
	msg := fmt.Sprintf("Add pizza %d successful %d", p.Id, ra)
	c.JSON(http.StatusOK, gin.H{
		"status": 1, "msg": msg,
	})
}

func GetPizzas(c *gin.Context) {
	var p models.Pizza
	pizzas, err := p.Gets()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"pizzas": pizzas,
	})
}
func GetPizza(c *gin.Context) {
	cid := c.Param("id")
	id, err := strconv.Atoi(cid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	p := models.User{Id: id}
	user, err := p.Get()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})

}

func ModPizza(c *gin.Context) {
	cid := c.Param("id")
	id, err := strconv.Atoi(cid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	p := models.User{Id: id}
	err = c.Bind(&p)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ra, err := p.Mod()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	msg := fmt.Sprintf("Update user %d successful %d", p.Id, ra)
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}

func DelPizza(c *gin.Context) {
	cid := c.Param("id")
	id, err := strconv.Atoi(cid)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	p := models.Pizza{Id: id}
	ra, err := p.Del()
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	msg := fmt.Sprintf("Delete pizza %d successful %d", id, ra)
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}
