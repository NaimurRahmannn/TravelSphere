package main

import (
	"TravelSphere/utils"

	_ "TravelSphere/routers"

	"github.com/beego/beego/v2/server/web"
)

func main() {
	// Expose formatting helpers to templates so views stay logic-free.
	web.AddFuncMap("population", utils.FormatPopulation)

	web.Run()
}