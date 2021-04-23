package main

import (
	"log"
	"github.com/gin-gonic/gin"
)

var IdCounter string = "1"

type Passive struct{
	Name string `json:"passivename" binding:"required"`
	Scaling string `json:"scaling" binding: "required"`
}

type Champion struct{
	Id string
	Name string `json:"name" binding:"required"`
	DamageSource string `json:"damagesource" binding:"required"`
	PassiveAbility Passive `json:"passiveability" binding:"required"`
}

var LeagueChampions map[string]Champion = make(map[string]Champion, 0)

func respondToGetChampion(c *gin.Context){
	allChampions := make([]Champion, 0)
	for _, leagueChampion:= range LeagueChampions {
		allChampions = append(allChampions, leagueChampion)
	}
	c.JSON(200, allChampions)
}

func respondToPostChampion(c *gin.Context){
	var champion Champion
	
	err := c.ShouldBindJSON(&champion)
	champion.Id = IdCounter

	log.Println("Creating this chamption: ", champion)
	log.Println(err)

	if err != nil {
		c.JSON(400, gin.H{"message": "Bad request, failed to create champion"})
		return
	}
	LeagueChampions[string(IdCounter)] = champion
	c.JSON(201, gin.H{"message": "Success, created new champion"})
	IdCounter += "1"

	log.Println("This is the id counter after POST: ", IdCounter)
}

func checkChampionAvailability(c *gin.Context, id string) bool{
	log.Println("This is the deleted ID:", id)

	_, ok := LeagueChampions[id]
	if ok == false {
		c.JSON(404, gin.H{"message": "This champion wasn't found"})
		return false
	}

	return true
}


func respondToDeleteChampion(c *gin.Context){
	id := c.Query("id")

	if !checkChampionAvailability(c, id){
		return
	}

	delete(LeagueChampions, id)

	c.JSON(200, gin.H{"message": `The champion with has been deleted`})
}

func respondToPatchChampion(c *gin.Context){
	var champion Champion
	id := c.Query("id")


	if !checkChampionAvailability(c, id){
		return
	}
	
	err := c.ShouldBindJSON(&champion)

	log.Println("Creating this chamption: ", champion)
	log.Println(err)

	if err != nil {
		c.JSON(400, gin.H{"message": "Bad request, failed to create champion"})
		return
	}

	LeagueChampions[id] = champion
	c.JSON(200, gin.H{"message": "Modified champion"})
}

func main(){
	r := gin.Default()

	r.GET("/champion", respondToGetChampion)
	r.POST("/champion", respondToPostChampion)
	r.PATCH("/champion", respondToPatchChampion)
	r.DELETE("/champion", respondToDeleteChampion)

	r.Run()
}