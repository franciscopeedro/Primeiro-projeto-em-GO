package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

const monitoramentos = 3
const delay = 3

func main() {

	for {
		exibeIntroducao()
		LeSitesDoArquivo()
		exibeMenu()

		comando := leComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo Logs...")
		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando")
			os.Exit(-1)
		}

	}

}

func exibeIntroducao() {
	nome := "Pedro"
	versao := 1.1
	fmt.Println("Olá, sr(a).", nome)
	fmt.Println("Este programa está na versão", versao)
}

func exibeMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi:", comandoLido)
	fmt.Println()

	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")

	sites := LeSitesDoArquivo()

	fmt.Println(sites)
	for i := 0; i < monitoramentos; i++ {
		for i, sites := range sites {
			fmt.Println("Testando site", i, ":", sites)
			testasite(sites)
		}
		time.Sleep(delay * time.Second)
		fmt.Println()
	}
	fmt.Println()
}

func testasite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro!!", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
	}
	fmt.Println()
}

func LeSitesDoArquivo() []string {

	var sites []string

	arquivo, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("Ocorreu um erro!!", err)
	}

	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')

		fmt.Println(linha)
		if err != io.EOF {
			break
		}
	}

	return sites
}
