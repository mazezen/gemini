package format

import (
	"context"
	"fmt"
	"github.com/mazezen/gemini/client"
	"github.com/mazezen/gemini/model"
	"google.golang.org/genai"
	"os"
)

// FormatToJson 结构化输出
func FormatToJson() {
	ctx := context.Background()
	c := client.NewClient(ctx, os.Getenv("GEMINI_API_KEY"))

	config := &genai.GenerateContentConfig{
		ResponseMIMEType: "application/json",
		ResponseSchema: &genai.Schema{
			Type: genai.TypeArray,
			Items: &genai.Schema{
				Type: genai.TypeObject,
				Properties: map[string]*genai.Schema{
					"recipeName": {Type: genai.TypeString},
					"ingredients": {
						Type:  genai.TypeArray,
						Items: &genai.Schema{Type: genai.TypeString},
					},
				},
				PropertyOrdering: []string{"recipeName", "ingredients"},
			},
		},
	}

	res, err := c.Models.GenerateContent(
		ctx,
		model.Gemini25Flash,
		genai.Text("List a few popular cookie recipes, and include the amounts of ingredients."),
		config,
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(res.Text())

}
