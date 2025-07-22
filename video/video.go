package video

import (
	"context"
	"fmt"
	"github.com/mazezen/gemini/client"
	"google.golang.org/genai"
	"os"
)

func YouTuBe() {
	ctx := context.Background()
	c := client.NewClient(ctx, os.Getenv("GEMINI_API_KEY"))

	parts := []*genai.Part{
		genai.NewPartFromText("Please summarize the video in 3 sentences."),
		genai.NewPartFromURI("https://www.youtube.com/watch?v=9hE5-98ZeCg", "video/mp4"),
	}

	contents := []*genai.Content{
		genai.NewContentFromParts(parts, genai.RoleUser),
	}

	result, _ := c.Models.GenerateContent(
		ctx,
		"gemini-2.5-flash",
		contents,
		nil,
	)

	fmt.Println(result.Text())
}
