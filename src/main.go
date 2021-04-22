package main

import (
	"log"
	"github.com/gin-gonic/gin"
)

type Passive struct{
	Name string `json:"passivename" binding:"required"`
	Scaling string `json:"scaling" binding: "required"`
}

type Champion struct{
	Name string `json:"name" binding:"required"`
	DamageSource string `json:"damagesource" binding:"required"`
	PassiveAbility Passive `json:"passiveability" binding:"required"`
}

var Champions = make([]Champion, 0, 100)

func respondToGetChampion(c *gin.Context){
	c.JSON(200, Champions)
}

func respondToPostChampion(c *gin.Context){
	var champion Champion
	
	err := c.ShouldBindJSON(&champion)

	log.Println(champion)
	log.Println(err)

	if err != nil {
		c.JSON(400, gin.H{"message": "Bad request, failed to create champion"})
		return
	}
	Champions = append(Champions, champion)
	c.JSON(200, gin.H{"message": "Success, created new champion"})
}


var LeageChamps = make([]Champion, 0, 100)

func main(){
	r := gin.Default()

	r.GET("/champion", respondToGetChampion)
	r.POST("/champion", respondToPostChampion)
	r.PATCH("/champion/:id", respondToGetChampion)
	r.DELETE("/champion/:id", respondToGetChampion)

	r.Run()
}