package main

import (
	"log"
	"os"
	_ "study/day02/mathers"
	"study/day02/search"
)

func init(){
	log.SetOutput(os.Stdout)
}

func main()  {
	search.Run("president")
}