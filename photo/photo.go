// Gemini 可以通过对话方式生成和处理图片。您可以使用文本、图片或两者结合来提示 Gemini，以执行各种与图片相关的任务，例如图片生成和编辑。
//您必须在配置中添加 responseModalities: ["TEXT", "IMAGE"]。这些模型不支持仅输出图片。

package photo

import (
	"context"
	"fmt"
	"github.com/mazezen/gemini/client"
	"github.com/mazezen/gemini/model"
	"google.golang.org/genai"
	"os"
)

// PhoGen 生成图片
func PhoGen() {
	ctx := context.Background()
	c := client.NewClient(ctx, os.Getenv("GEMINI_API_KEY"))

	config := &genai.GenerateContentConfig{
		ResponseModalities: []string{"TEXT", "IMAGE"},
	}

	result, _ := c.Models.GenerateContent(
		ctx,
		model.Gemini20FlashPreviewImageGeneration,
		genai.Text("帮我生成一张大熊猫吃竹子🎋的照片"),
		config,
	)

	for _, part := range result.Candidates[0].Content.Parts {
		if part.Text != "" {
			fmt.Println(part.Text)
		} else if part.InlineData != nil {
			imageBytes := part.InlineData.Data
			outputFilename := "panda.png"
			_ = os.WriteFile(outputFilename, imageBytes, 0644)
		}
	}
}

// PhoEdit 图片编辑（文字和图片转图片）
func PhoEdit() {
	ctx := context.Background()
	c := client.NewClient(ctx, os.Getenv("GEMINI_API_KEY"))

	imagePath := "images/panda.png"
	imgData, err := os.ReadFile(imagePath)
	if err != nil {
		panic(err)
	}

	parts := []*genai.Part{
		genai.NewPartFromText("这是一张大熊猫图片,可以帮我在大熊猫旁边加一个饲养员吗?"),
		&genai.Part{
			InlineData: &genai.Blob{
				MIMEType: "image/png",
				Data:     imgData,
			},
		},
	}

	contents := []*genai.Content{
		genai.NewContentFromParts(parts, genai.RoleUser),
	}

	config := &genai.GenerateContentConfig{
		ResponseModalities: []string{"TEXT", "IMAGE"},
	}

	res, err := c.Models.GenerateContent(
		ctx,
		model.Gemini20FlashPreviewImageGeneration,
		contents,
		config,
	)
	if err != nil {
		panic(err)
	}

	for _, part := range res.Candidates[0].Content.Parts {
		if part.Text != "" {
			fmt.Println(part.Text)
		} else if part.InlineData != nil {
			imageBytes := part.InlineData.Data
			outputFilename := "images/panda_breeder.png"
			_ = os.WriteFile(outputFilename, imageBytes, 0644)
		}
	}
}

// ImaGen40GeneratePreview0606 使用 Imagen 模型生成图片
func ImaGen40GeneratePreview0606() {
	ctx := context.Background()
	c := client.NewClient(ctx, os.Getenv("GEMINI_API_KEY"))

	config := &genai.GenerateImagesConfig{
		NumberOfImages: 4,
	}

	res, err := c.Models.GenerateImages(
		ctx,
		model.ImaGen40GeneratePreview0606,
		"Robot holding a red stakeboard",
		config,
	)
	if err != nil {
		panic(err)
	}

	for n, image := range res.GeneratedImages {
		fName := fmt.Sprintf("images/imgen-%d.png", n)
		_ = os.WriteFile(fName, image.Image.ImageBytes, 0644)
	}
}
