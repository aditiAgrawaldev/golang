package main

import "fmt"

func main() {
	fmt.Println("Maps in go")
	languages := make(map[string]string)

	languages["JS"] = "JavaScript"
	languages["Py"] = "Python"
	languages["J"] = "Java"
	languages["C++"] = "C++"
	languages["Rb"] = "Ruby"

	fmt.Println("List of all languages", languages)
	fmt.Println("JS shorts for", languages["JS"])

	delete(languages, "Rb")
	fmt.Println("List of all languages after deleting Ruby", languages)

	for key, value := range languages {
		fmt.Println(key, "shorts for", value)
	}
}
