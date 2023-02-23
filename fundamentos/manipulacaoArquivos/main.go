package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// O pacote OS serve para manipulação de arquivos dentre outras coisas.
	// os.Create cria um novo arquivo.
	file, err := os.Create("teste.txt")
	if err != nil {
		panic(err)
	}

	// vai adiar o fechamento do arquivo para quando ele terminar de ser usado
	//defer file.Close()

	// file.Write nos permite passar/gravar dados, qualquer tipo de dados dentro do nosso arquvivo
	// desde que estejam convertidos para bytes
	tamanho, err := file.Write([]byte("Dados criados utilizando um slice de bytes"))
	if err != nil {
		panic(err)
	}

	file.Close()

	// escreve/grava strings no arquivo que a acabamos de criar acima
	// tamanho, err := file.WriteString("Hello, World")
	// if err != nil {
	// 	panic(err)
	// }

	fmt.Printf("Arquivo criado com sucesso! Tamanho: %d bytes\n\n", tamanho)

	//##LEITURA DE ARQUIVOS##
	arquivo, err := os.ReadFile("teste2.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(arquivo))

	/* quando o tamanho do arquivo a ser lido é maior do que a memória que temos disponível, nós utilizamos a leitura linha a linha.
	abrimos o arquivo de pouco em pouco, lemos linha a linha e manipulamos/gravamos da mesma forma. Isso nos possibilita resolver problemas de memória referente a leitura
	e gravação de arquivos de tamanhos muito grandes.

	#### LEITURA LINHA A LINHA ####
	*/

	arquivo2, err := os.Open("teste2.txt")
	if err != nil {
		panic(err)
	}

	defer arquivo2.Close()

	reader := bufio.NewReader(arquivo2)
	buffer := make([]byte, 15)
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			//log.Println("erro na leitura bufferizada:", err)
			break
		}
		fmt.Println(string(buffer[:n]))
	}

	//removendo arquivo
	err = os.Remove("teste.txt")
	if err != nil {
		panic(err)
	}

}
