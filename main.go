// main.go
package main

import 	"github.com/SwanHtetAungPhyo/Scache/utils"

func main() {
	logfile := utils.InitLog()
	utils.CloseLogFile(logfile)
}
