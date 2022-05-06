package main

import "strings"

// AnaliserQuery Função analisa querystring passada como parametro para retornar uma lista dos valores como lista
func AnaliserQuery(query string) map[string]string {
	var pares []string = strings.Split(query[1:], "&")
	var mapa = map[string]string{}
	for i := 0; i < len(pares); i++ {
		var chaveValor []string = strings.Split(pares[i], "=")
		var chave string = chaveValor[0]
		var valor string = chaveValor[1]

		mapa[chave] = valor
	}

	return mapa
}
