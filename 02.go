package main

import "fmt"

func mergeMaps(map1, map2 map[string]string) map[string]string {
	mergedMap := make(map[string]string)

	for key, value := range map1 {
		mergedMap[key] = value
	}

	for key, value := range map2 {
		mergedMap[key] = value
	}

	return mergedMap
}

func main() {
	map1 := map[string]string{
		"chave1": "valor1",
		"chave2": "valor2",
		"chave3": "valor3",
	}

	map2 := map[string]string{
		"chave2": "novo_valor2",
		"chave4": "valor4",
	}

	mergedMap := mergeMaps(map1, map2)

	for key, value := range mergedMap {
		fmt.Printf("%s: %s \n", key, value)
	}
}
