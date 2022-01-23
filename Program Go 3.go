/*Imagina que tienes que escribir unos cuantos correos electrónicos. Van a ser largos
y laboriosos, y la mejor forma de entretenerse es escuchar música mientras se escriben,
es decir, escuchar música “en paralelo” a la redacción de los emails.*/
//PARALELISMO
package main

import (
	"fmt"
	"sync"
	"time"
)

func imprimeHora(msg string) {
	fmt.Println(msg, time.Now().Format("15:04:05"))
}

// Tarea que se hará con el tiempo.
func escribeMail1(wg *sync.WaitGroup) {
	imprimeHora("Email #1 escrito.")
	wg.Done()
}
func escribeMail2(wg *sync.WaitGroup) {
	imprimeHora("Email #2 escrito.")
	wg.Done()
}
func escribeMail3(wg *sync.WaitGroup) {
	imprimeHora("Email #3 escrito.")
	wg.Done()
}

// Tarea realizada en paralelo
func escuchaPorSiempre() {
	for {
		imprimeHora("Escuchando...")
	}
}
func main() {
	var waitGroup sync.WaitGroup
	waitGroup.Add(3)

	go escuchaPorSiempre()

	// Darle algo de tiempo para que listenForever comience
	time.Sleep(time.Nanosecond * 10)

	// Escribiendo los emails
	go escribeMail1(&waitGroup)
	go escribeMail2(&waitGroup)
	go escribeMail3(&waitGroup)

	waitGroup.Wait()
}
