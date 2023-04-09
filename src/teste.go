package main

import (
	"fmt"
	"net/http"
	"os"
	"reflect"
	"time"
)

const monitoring = 10
const delay = 5

func main() {

	getName()

	for {
		opcao := getOpcao()

		// if opcao == 1 {
		// 	fmt.Println("Iniciando Programa")
		// } else if opcao == 2 {
		// 	fmt.Println("Iniciando logs")
		// } else if opcao == 0 {
		// 	fmt.Println("Saindo do Programa")
		// } else {
		// 	fmt.Println("Nao conheco esta opcao")
		// }

		switch opcao {
		case 1:
			verifySite()
		case 2:
			logsSite()
		case 0:
			fmt.Println("Saindo do Programa")
			os.Exit(0)
		default:
			fmt.Println("Nao conheco este comando")
			os.Exit(-1)
		}
	}

}

func getName() {
	var name string
	fmt.Println("Bem vindo, Digite o seu nome")
	fmt.Scanf("%s", &name)

	fmt.Println("Holla que tal Sr.", name, "Saudacoes")
	fmt.Println("o tipo da variavel e:", reflect.TypeOf(name))

}

func getOpcao() int {
	var opcao int
	println("Escolha uma opcao")
	println("1- Iniciar programa")
	println("2- Ixibir Logs")
	println("0- Sair")

	fmt.Scanf("%d", &opcao)
	fmt.Println("O comando escolhido foi", opcao)
	return opcao
}

func statusSite(site string) {

	response, _ := http.Get(site)
	if response.StatusCode == 200 {
		fmt.Println("Site", site, "foi carregado com sucesso")
		fmt.Println("___")

	} else {
		fmt.Println("Site", site, "esta com problemas")
		fmt.Println("___")

	}
}

func verifySite() {
	fmt.Println("Iniciando Monitoramento")
	fmt.Println("___")

	sitesURLs := []string{"https://www.google.com/oi",
		"https://www.google.com/", "https://www.google.com/oi"}

	// for i := 0; i < len(sitesURLs); i++ {
	// 	fmt.Println(sitesURLs[i])
	// }

	for i := 0; i < monitoring; i += 1 {
		for i, site := range sitesURLs {
			fmt.Println("Monitorando", i+1, site)
			statusSite(site)
		}
		time.Sleep(delay * time.Second)
	}
}

func logsSite() {
	fmt.Println("Iniciando Logs")
	sitesURLs := []string{"https://www.google.com/oi",
		"https://www.google.com/", "https://www.google.com/oi"}

	// for i := 0; i < len(sitesURLs); i++ {
	// 	fmt.Println(sitesURLs[i])
	// }

	for i, site := range sitesURLs {
		fmt.Println("Monitorando", i+1, site)
	}
}

func tested() {
	var sites [3]string
	sites[0] = "www1"
	sites[1] = "www2"
	sites[2] = "www3"

	fmt.Println(sites)
	fmt.Println("o tamanho do array e", len(sites))

	nomes := []string{"Flavio", "EricK", "Junior", "Renata", "Rogerio", "Flavio", "EricK", "Junior", "Renata", "Rogerio"}
	fmt.Println(nomes)
	fmt.Println("A quantidade do array e", len(nomes))
	fmt.Println("A capacidade do array e", cap(nomes))

	nomes = append(nomes, "Lidia", "Celio")

	fmt.Println("o nova quantidade do array e", len(nomes))
	fmt.Println("A nova capacidade do array e", cap(nomes))

	fmt.Println(nomes)
}
