package main

import (
	"fmt"
	"net/http"
	"os"
	"reflect"
)

func main() {

	getName()
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
		statusSite()
	case 2:
		logsSite()
	case 0:
		fmt.Println("Saindo do Programa")
		os.Exit(0)
	default:
		fmt.Println("Nao conheco este comando")
		os.Exit(-1)
	}

	fmt.Println("o tipo da variavel e:", reflect.TypeOf(opcao))
	fmt.Println("o endereco da variavel e:", &opcao)
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

func statusSite() {
	fmt.Println("Iniciando Monitoramento")
	site := "https://www.google.com/oi"
	response, _ := http.Get(site)
	fmt.Println(response)
	if response.StatusCode == 200 {
		fmt.Println("Site", site, "foi carregado com sucesso")
	} else {
		fmt.Println("Site", site, "esta com problemas")
	}
}

func logsSite() {
	fmt.Println("Iniciando Logs")
	sitesURLs := []string{"https://www.google.com/oi",
		"https://www.google.com/", "https://www.google.com/oi"}

	// for i := 0; i < len(sitesURLs); i++ {
	// 	fmt.Println(sitesURLs[i])
	// }

	for _, site := range sitesURLs {
		fmt.Println(site)
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
