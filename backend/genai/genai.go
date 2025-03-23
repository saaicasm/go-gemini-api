package genai

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

func ExampleGenerativeModel_GenerateContent_textOnly(animal string) string {
	godotenv.Load()
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("API_KEY")))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-1.5-flash")
	prompt := fmt.Sprintf("Write a story about %s doing random activity in a random city from random country. Keep the story only about 200 lines. Use simple language", animal)
	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		log.Fatal(err)
	}

	// Convert resp.Candidates[0].Content.Parts to a string
	var sb strings.Builder
	if len(resp.Candidates) > 0 && resp.Candidates[0].Content != nil {
		for _, part := range resp.Candidates[0].Content.Parts {
			if stringer, ok := part.(interface{ String() string }); ok {
				// It has a String() method, so use it
				sb.WriteString(stringer.String())
			} else {
				// Fallback: convert to string using fmt.Sprintf
				sb.WriteString(fmt.Sprintf("%v", part))
			}
		}
	}
	return sb.String()
}
