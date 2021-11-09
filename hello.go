package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoramentos = 5
const delay = 5

func main() {

	exibeIntroducao()
	for {
		exibeMenu()

		cmd := leComando()

		switch cmd {
		case 1:
			iniciarMonitoramento()
		case 2:
			imprimeLogs()
		case 0:
			fmt.Println("Saindo do programa ...")
			os.Exit(0)
		default:
			fmt.Println("Não conheco esse programa ...")
			os.Exit(-1)
		}
	}

}

func exibeIntroducao() {
	nome := "João"
	versao := 1.1
	fmt.Println("Olá, Sr. ", nome)
	fmt.Println("Este programa está na  versão ", versao)
}

func exibeMenu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0- Sair do programa")
}

func leComando() int {
	var cmd int
	fmt.Scan(&cmd)
	fmt.Println("Comando escolhido foi", cmd)
	return cmd
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando ...")
	sites := leSiteDoArquivo()

	for i := 0; i < monitoramentos; i++ {
		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			testaSite(site)
		}
		fmt.Println("")
		time.Sleep(delay * time.Second)
	}
	fmt.Println("")
}

func testaSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("O site", site, "Foi carregado com sucesso")
		registraLog(site, true)
	} else {
		fmt.Println("Site", site, "Esta com problemas.", resp.StatusCode)
		registraLog(site, false)
	}

}

func leSiteDoArquivo() []string {

	var sites []string

	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		sites = append(sites, linha)

		if err == io.EOF {
			fmt.Println("Abacou o arquivo:", err)
			break
		}

	}

	arquivo.Close()

	return sites

}

func registraLog(site string, status bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + "- online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}

func imprimeLogs() {
	fmt.Println("Exibindo logs ...")

	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(arquivo))

}
