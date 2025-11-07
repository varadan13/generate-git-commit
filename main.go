package main

import (
	"bytes"
	"context"
	"fmt"
	"generate-git-commit/internal/gemini"
	"log"
	"os/exec"

	"google.golang.org/genai"
)

func gitStagedDiff() (string, error) {
	cmd := exec.Command("git", "diff", "--staged")

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("git diff --staged failed: %v (%s)", err, stderr.String())
	}

	return out.String(), nil
}

func main() {
	fmt.Println("||===================generate-git-command initialised===================||")

	ctx := context.Background()

	client, err := gemini.GetGenaiClient()

	if err != nil {
		log.Fatal(err)
	}

	result, err := client.Models.GenerateContent(
		ctx,
		"gemini-2.5-flash",
		genai.Text("What's the temperature in London?"),
		gemini.GetConfigWithTool(),
	)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result.Text())

	for _, cand := range result.Candidates {
		for _, part := range cand.Content.Parts {
			if part.FunctionCall != nil {
				fmt.Println("Function to call:", part.FunctionCall.Name)
				fmt.Println("Arguments:", part.FunctionCall.Args)
			}
		}
	}

	diff, err := gitStagedDiff()

	if err != nil {
		fmt.Println("❌ Error:", err)
		return
	}

	if diff == "" {
		fmt.Println("✅ No staged changes.")
	} else {
		fmt.Println("✅ Staged diff:\n", diff)
	}

}
