package gemini

import (
	"context"
	"fmt"
	"generate-git-commit/internal/git"
	"log"

	"google.golang.org/genai"
)

func getGenaiClient() (*genai.Client, error) {
	ctx := context.Background()

	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  GetApiKey(),
		Backend: genai.BackendGeminiAPI,
	})

	return client, err
}

func StreamContent() {
	ctx := context.Background()

	client, err := getGenaiClient()

	if err != nil {
		log.Fatal(err)
	}

	diff, err := git.GetStagedDiff()

	if err != nil {
		log.Fatal(err)
	}

	systemprompt := `
						Use the below style guide to generate a commit message:

						Git Commit Style Guide

						Format of the Commit Message
						----------------------------

						{type}({scope}): {subject}
						<BLANK LINE>
						{body}
						<BLANK LINE>
						{footer}

						Rules for Commit Message
						-----------------------

						#### Length

						- Keep lines under 80 characters in width.
						- Subject line must not be longer than 60 characters (one line in Github PR description).

						#### Subject - {subject}

						Summary of the changes made.

						> The subject line contains a succinct description of the change to the logic.

						- Must be present tense
						- Written in the imperative
						- First letter is not capitalized
						- Does not end with a '.'

						#### Allowed Types - {types}

						- feat -> feature
						- fix -> bug fix
						- docs -> documentation
						- style -> formatting, lint stuff
						- refactor -> code restructure without changing external behavior
						- test -> adding missing tests
						- chore -> maintenance
						- init -> initial commit
						- rearrange -> files moved, added, deleted etc
						- update -> update code (versions, library compatibility)

						#### Scope - {scope}

						Where the change was (i.e. the file, the component, the package).
						
						#### Message Body - {body}

						This gives details about the commit, including:

						- motivation for the change (broken code, new feature, etc)
						- contrast with previous behavior

						Some rules for the body:

						- Must be in present tense.
						- Should be imperative.
						- Lines must be less than 80 characters long.


						#### Message Footer - {footer}

						These are notes that someone should be aware of. Format footer in category blocks.

						- TESTING -> how to test the change
						- BREAKING CHANGE -> what is different now, additional things now needed, etc
						`

	prompt := fmt.Sprintf("**Prompt:** You are an AI specialized in generating clear and concise commit messages based on git diffs.Your task is to analyze the provided git diff and summarize the changes in a structured commit message.Follow these guidelines: 1.**Identify the Purpose**: Determine the main purpose of the changes (e.g., bug fix, feature addition, refactoring, documentation update).2.**Summarize Changes**: List the key modifications made, focusing on what files were changed and the type of changes (additions, deletions, modifications).3.**Use Imperative Mood**: Write the commit message in the imperative mood, starting with a verb (e.g., Add, Fix, Update).4.**Limit Length**: Keep the summary line to 50 characters or less, followed by a more detailed explanation if necessary.5.**Include Context**: If there are any related issues or tickets, mention them at the end of the message.6.**Format**: Ensure that the commit message follows conventional commit standards if applicable.**Input Git Diff**: ``` %s ```", *diff)

	contents := []*genai.Content{
		{
			Role: "user",
			Parts: []*genai.Part{
				{Text: systemprompt},
			},
		},
		{
			Role: "user",
			Parts: []*genai.Part{
				{Text: prompt},
			},
		},
	}

	for result, err := range client.Models.GenerateContentStream(
		ctx,
		"gemini-2.5-flash",
		contents,
		nil,
	) {
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print(result.Candidates[0].Content.Parts[0].Text)
	}

}
