package apis

import (
	"log"
	"net/http"
	"fmt"
	. "test/rest/models"
	"github.com/gin-gonic/gin"
	"strconv"
)

func IndexApi(c *gin.Context) {
	c.String(http.StatusOK, "It works")
}

func AddPersonApi(c *gin.Context) {
	var p Person
	err := c.Bind(&p)
	if err != nil {
		log.Fatalln(err)
	}

	id, err := p.Add()
	if err != nil {
		log.Fatalln(err)
	}
	name := p.FirstName + " " + p.LastName
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf(" %d %s successfully created", id, name),
	})
}

func GetPersonsApi(c *gin.Context) {
	p := Person{}
	persons, err := p.GetAll()
	if err != nil {
		log.Fatalln(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"result": persons,
		"count":  len(persons),
	})
}

func GetPersonApi(c *gin.Context) {
	var result gin.H
	id := c.Param("id")

	Id, err := strconv.Atoi(id)
	if err != nil {
		log.Fatalln(err)
	}

	p := Person{
		Id: Id,
	}
	person, err := p.Get()
	if err != nil {
		result = gin.H{
			"result": nil,
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": person,
			"count":  1,
		}
	}
	c.JSON(http.StatusOK, result)
}

func ModPersonApi(c *gin.Context) {
	cid := c.Param("id")
	id, err := strconv.Atoi(cid)
	if err != nil {
		log.Fatalln(err)
	}
	p := Person{
		Id: id,
	}

	err = c.Bind(&p)
	if err != nil {
		log.Fatalln(err)
	}

	rows, err := p.Update()
	if err != nil {
		log.Fatalln(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Update person %d successful %d", id, rows),
	})
}

func DelPersonApi (c *gin.Context) {
	cid := c.Param("id")
	id, err := strconv.Atoi(cid)
	if err != nil {
		log.Fatalln(err)
	}
	p := Person{
		Id: id,
	}
	rows, err := p.Del()
	msg := fmt.Sprintf("Delete person %d successful %d", id, rows)
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}
