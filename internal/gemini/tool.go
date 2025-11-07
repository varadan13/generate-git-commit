package gemini

import (
	"google.golang.org/genai"
)

func GetConfigWithTool() *genai.GenerateContentConfig {
	tool := &genai.Tool{
		FunctionDeclarations: []*genai.FunctionDeclaration{
			{
				Name:        "get_current_temperature",
				Description: "Gets the current temperature for a given location.",
				Parameters: &genai.Schema{
					Type: "object",
					Properties: map[string]*genai.Schema{
						"location": {
							Type:        "string",
							Description: "The city name, e.g. San Francisco",
						},
					}, Required: []string{"location"}},
			}},
	}

	return &genai.GenerateContentConfig{Tools: []*genai.Tool{tool}}
}
