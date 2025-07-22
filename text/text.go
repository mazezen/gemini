package text

import (
	"context"
	"fmt"
	"github.com/mazezen/gemini/client"
	"github.com/mazezen/gemini/model"
	"google.golang.org/genai"
	"os"
)

// TexGen 当个文本对话的示例
func TexGen() {

	ctx := context.Background()
	c := client.NewClient(ctx, os.Getenv("GEMINI_API_KEY"))

	thinkingBudgetVal := int32(0)

	res, err := c.Models.GenerateContent(
		ctx,
		model.Gemini25Flash,
		genai.Text("介绍一下AI的工作原理"),
		&genai.GenerateContentConfig{
			ThinkingConfig: &genai.ThinkingConfig{
				ThinkingBudget: &thinkingBudgetVal,
				// Turn off thinking:
				// ThinkingBudget: int32(0),
				// Turn on dynamic thinking:
				// ThinkingBudget: int32(-1),
			},
		},
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(res.Text())
}

// TexGenThinking 使用 Gemini 2.5 进行思考
// 2.5 Flash 和 Pro 模型默认启用了“思考”功能，以提升质量，这可能会导致运行时间延长并增加令牌用量。
// 使用 2.5 Flash 时，您可以通过将思考预算设置为零来停用思考功能.
func TexGenThinking() {
	ctx := context.Background()
	c := client.NewClient(ctx, os.Getenv("GEMINI_API_KEY"))

	thinkingBudgetVal := int32(1024)

	res, err := c.Models.GenerateContent(
		ctx,
		model.Gemini25Flash,
		genai.Text("AI是如何工作的?"),
		&genai.GenerateContentConfig{
			ThinkingConfig: &genai.ThinkingConfig{
				ThinkingBudget: &thinkingBudgetVal,
			},
		},
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(res.Text())
}

// TexGenCli 使用系统指令来引导 Gemini 模型的行为
func TexGenCli() {
	ctx := context.Background()
	c := client.NewClient(ctx, os.Getenv("GEMINI_API_KEY"))

	config := &genai.GenerateContentConfig{
		SystemInstruction: genai.NewContentFromText("我是一名代码搬运工,可以帮我写一个js的Hello world程序吗?", genai.RoleUser),
	}

	res, err := c.Models.GenerateContent(
		ctx,
		model.Gemini25Flash,
		genai.Text("你好"),
		config,
	)
	if err != nil {
		panic(err)
	}

	fmt.Println(res.Text())
}

// TexGenDefaultConfig 借助 GenerateContentConfig 对象, 替换默认生成参数
func TexGenDefaultConfig() {
	ctx := context.Background()
	c := client.NewClient(ctx, os.Getenv("GEMINI_API_KEY"))

	temp := float32(0.9)
	topP := float32(0.5)
	topK := float32(20.0)

	config := &genai.GenerateContentConfig{
		Temperature:      &temp,
		TopP:             &topP,
		TopK:             &topK,
		ResponseMIMEType: "application/json",
	}

	res, err := c.Models.GenerateContent(
		ctx,
		model.Gemini25Flash,
		genai.Text("What is the average size of a swallow?"),
		config,
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(res.Text())
}

// TexGenMultiModal 多模态输入，将文本与媒体文件组合使用
func TexGenMultiModal() {
	ctx := context.Background()
	c := client.NewClient(ctx, os.Getenv("GEMINI_API_KEY"))

	imagePath := "images/multimodal.png"
	imgData, err := os.ReadFile(imagePath)
	if err != nil {
		panic(err)
	}

	parts := []*genai.Part{
		genai.NewPartFromText("帮我分析图片中的代码"),
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

	res, err := c.Models.GenerateContent(
		ctx,
		model.Gemini25Flash,
		contents,
		nil,
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(res.Text())
}

// TexGenStream 模型仅在整个生成过程完成后才会返回回答
// 为了实现更流畅的互动，请使用流式传输在 GenerateContentResponse 实例生成时逐步接收这些实例
func TexGenStream() {
	ctx := context.Background()
	c := client.NewClient(ctx, os.Getenv("GEMINI_API_KEY"))

	stream := c.Models.GenerateContentStream(
		ctx,
		model.Gemini25Flash,
		genai.Text("帮我写一则小故事,讲计算机发展史"),
		nil,
	)

	for chunk, _ := range stream {
		part := chunk.Candidates[0].Content.Parts[0]
		fmt.Println(part.Text)
	}
}

// TexMultipleRoundsOfDialogue 多轮对话
// 注意 ：聊天功能仅在 SDK 中实现。在后台，它仍会使用 generateContent API。对于多轮对话，系统会在每次后续对话时将完整对话记录发送给模型。
func TexMultipleRoundsOfDialogue() {
	ctx := context.Background()
	c := client.NewClient(ctx, os.Getenv("GEMINI_API_KEY"))

	history := []*genai.Content{
		genai.NewContentFromText("我是一名代码搬运工", genai.RoleUser),
		genai.NewContentFromText("好的. 平常主要使用什么语言呢?", genai.RoleModel),
	}

	chat, err := c.Chats.Create(ctx, model.Gemini25Flash, nil, history)
	if err != nil {
		panic(err)
	}
	res, err := chat.SendMessage(ctx, genai.Part{Text: "我的职业是什么?"})
	if err != nil {
		panic(err)
	}

	if len(res.Candidates) > 0 {
		fmt.Println(res.Candidates[0].Content.Parts[0].Text)
	}
}

// TexGenStreamMultipleRoundsOfDialogue 流式响应多轮对话
func TexGenStreamMultipleRoundsOfDialogue() {
	ctx := context.Background()
	c := client.NewClient(ctx, os.Getenv("GEMINI_API_KEY"))

	history := []*genai.Content{
		genai.NewContentFromText("我是一名代码搬运工", genai.RoleUser),
		genai.NewContentFromText("好的. 平常主要使用什么语言呢?", genai.RoleModel),
	}

	chat, err := c.Chats.Create(ctx, model.Gemini25Flash, nil, history)
	if err != nil {
		panic(err)
	}

	stream := chat.SendMessageStream(ctx, genai.Part{Text: "我的职业是什么?"})
	for chunk, _ := range stream {
		part := chunk.Candidates[0].Content.Parts[0]
		fmt.Println(part.Text)
	}
}
