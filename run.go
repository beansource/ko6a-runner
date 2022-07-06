package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func RunScript(test Ko6aTest) {
	replacer := strings.NewReplacer("/blob/", "/", "github.com", "raw.githubusercontent.com")
	fileContentUrl := replacer.Replace(test.scriptUrl)

	out, err := exec.Command("k6", "run", fileContentUrl).Output()

	if err != nil {
		fmt.Println(err)
	}

	output := string(out[:])
	fmt.Println(output)
}

// script := "https://github.com/beansource/stinky-mud-puddle/blob/main/using-k6/01-http-get.js"
// script = "https://github.com/beansource/stinky-mud-puddle/blob/main/using-k6/03-metrics.js"