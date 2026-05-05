package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type Taxes struct {
	Base  string             `json:"base"`
	Date  string             `json:"date"`
	Rates map[string]float64 `json:"rates"`
}

func main() {
	data, err := os.ReadFile("rates.json")

	if err != nil {
		fmt.Println("Erro ao tentar carregar as taxas.")
		return
	}

	var taxes Taxes

	err = json.Unmarshal(data, &taxes)

	if err != nil {
		fmt.Println("Erro ao converter os dados das taxas.")
		return
	}

	args := os.Args

	if len(args) != 3 {
		fmt.Println("Para converter uma moeda, utilize o padrão de conversão abaixo:\n./convert [valor_em_brl] [moeda_destino]")
		return
	}

	value, err := strconv.ParseFloat(args[1], 64)

	if err != nil {
		fmt.Println("Digite um valor válido.")
		return
	}

	rate, ok := taxes.Rates[args[2]]

	if !ok {
		fmt.Printf("A moeda %s não existe.", args[2])
		return
	}

	convertedValue := value * rate

	fmt.Printf("O resultado da conversão é %.2f\n", convertedValue)

}
