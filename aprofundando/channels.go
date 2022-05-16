package main

import (
	"fmt"
	"time"
)

func ChannelSimples() {
	hello := make(chan string)

	go func() {
		hello <- "Hello mundo"
	}()

	result := <-hello
	fmt.Println(result)
}

//  Eu tentei almentar a complexidade do código nessa função
// e é simplismente incrivel como o Go é rápido, eu posso
// ter um código tão complexo quanto este a baixo e o tempo de
// processamento de 100000 repetições foi de 35 sec
func ChannelComplexo() {
	channel := make(chan int)

	go func() {
		i := 0
		repeticao := 0
		for {
			for j := 0; j < 5; j++ {
				if j == 2 {
					i += j
				}
				for l := 0; l < 5; l++ {
					channel <- l
					for d := 0; d < 5; d++ {
						fmt.Println("Sim, eu to tentando colocar fogo na minha máquina")
					}
				}
			}
			channel <- i
			i++
			repeticao++
			if repeticao == 100000 {
				fmt.Println(repeticao)
				break
			}
		}
	}()

	contadorForChannel := 0
	for x := range channel {
		fmt.Println("For do channel")
		fmt.Println(x)
		fmt.Println(contadorForChannel)
		contadorForChannel++
	}
}

func chamaChannel(pessoa string, mensagem chan int) {
	for x := range mensagem {
		fmt.Println("Pessoa:", pessoa, x)
		time.Sleep(time.Second)
	}
}

func VariosChannels() {
	mensagem := make(chan int)
	go chamaChannel("1", mensagem)
	go chamaChannel("2", mensagem)
	go chamaChannel("3", mensagem)
	go chamaChannel("4", mensagem)
	go chamaChannel("5", mensagem)
	go chamaChannel("6", mensagem)

	for i := 0; i < 10; i++ {
		mensagem <- i
	}
}

func main() {
	// ChannelSimples()
	// ChannelComplexo()
	VariosChannels()
}
