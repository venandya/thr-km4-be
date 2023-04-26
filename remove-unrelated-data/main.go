package main

import "fmt"

func removeUnrelated(dataMap map[string]any, key string) map[string]any {
	delete(dataMap, key)
	return dataMap
}

func main() {
	data1 := map[string]any{
		"name":    "Edo",
		"age":     20,
		"address": "Jakarta",
	}
	fmt.Println(removeUnrelated(data1, "address")) // output: map[name:Edo age:20]

	data2 := map[string]any{
		"run":  "lari",
		"jump": "loncat",
		"swim": "berenang",
	}
	fmt.Println(removeUnrelated(data2, "jump")) // output: map[run:lari swim:berenang]

	data3 := map[string]any{
		"satu":  "ichi",
		"dua":   "ni",
		"tiga":  "san",
		"empat": "yon",
		"lima":  "go",
	}
	fmt.Println(removeUnrelated(data3, "tiga")) // output: map[satu:ichi dua:ni empat:yon lima:go]
}
