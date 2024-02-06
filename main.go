package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type CepInfo struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	for _, cep := range os.Args[1:] {

		req, err := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer requisição: %v\n", err)
		}
		defer req.Body.Close()

		res, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao ler body: %v\n", err)
		}

		var cepInfo CepInfo
		err = json.Unmarshal(res, &cepInfo)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer o parser: %v\n", err)
		}

		cepInfoFile, err := os.Create("cep_info.txt")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao criar arquivo .txt: %v\n", err)
		}
		_, err = cepInfoFile.WriteString("localidade: " + cepInfo.Localidade + "\n")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao escrever arquivo .txt: %v\n", err)
		}
	}
}
