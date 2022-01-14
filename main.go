package main

import (
	"fmt"
	"log"

	"github.com/spro80/twitterGolang2021/bd"
	"github.com/spro80/twitterGolang2021/handlers"
)

//"log"
//"github.com/spro80/twitterGolang2021/bd"
//"github.com/spro80/twitterGolang2021/handlers"

func main() {
	fmt.Println("Hola")
	if bd.ChequeoConection() == 0 {
		log.Fatal("Sin conexion a la BD")
		return
	}
	handlers.Manejadores()

}
