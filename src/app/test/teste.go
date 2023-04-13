package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

const monitoring = 3
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
			fmt.Println("Iniciando Logs")
			printLogs()
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
		logsSite(site, true)

	} else {
		fmt.Println("Site", site, "esta com problemas")
		fmt.Println("___")
		logsSite(site, false)

	}
}

func verifySite() {
	fmt.Println("Iniciando Monitoramento")
	fmt.Println("___")

	sitesURLs := readSitesFile()

	fmt.Println("sites lidos", sitesURLs)

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

func logsSite(site string, status bool) {

	file, err := os.OpenFile("logs.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("erro", err.Error())
	}

	file.WriteString(time.Now().Format("02/01/2006-15-04-05") + "-" + site + "- Online " + strconv.FormatBool(status) + "\n")

	file.Close()

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

func readSitesFile() []string {

	var sites []string
	file, err := os.Open("sites.txt")
	fileRead, err := ioutil.ReadFile("sites.txt")

	if err != nil {
		fmt.Println("error", err.Error())
	}

	fmt.Println("file", file)
	fmt.Println("file", string(fileRead))

	leitor := bufio.NewReader(file)
	for {
		row, err := leitor.ReadString('\n')
		row = strings.TrimSpace(row)
		sites = append(sites, row)

		if err == io.EOF {
			break
		}

	}
	file.Close()
	return sites
}

func printLogs() {
	file, err := ioutil.ReadFile("logs.txt")

	if err != nil {
		fmt.Println("erro", err.Error())
	}

	fmt.Println(string(file))

}
