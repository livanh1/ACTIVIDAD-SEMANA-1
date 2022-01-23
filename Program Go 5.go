package main

import "fmt"
import "time"

func main() {
	t := time.Now()
	fecha := fmt.Sprintf("Fecha: %02d-%02d-%d Hora: %02d:%02d:%02d",
		t.Day(), t.Month(), t.Year(),
		t.Hour(), t.Minute(), t.Second())

	fmt.Println("LA FECHA Y HORA ACTUAL ES =>", fecha)
}
