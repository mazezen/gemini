package coding

import (
	"context"
	"fmt"
	"github.com/mazezen/gemini/client"
	"github.com/mazezen/gemini/model"
	"google.golang.org/genai"
	"os"
)

// EnableCodingExec 启用代码执行
func EnableCodingExec() {
	ctx := context.Background()
	c := client.NewClient(ctx, os.Getenv("GEMINI_API_KEY"))

	config := &genai.GenerateContentConfig{
		Tools: []*genai.Tool{
			{CodeExecution: &genai.ToolCodeExecution{}},
		},
	}

	res, err := c.Models.GenerateContent(
		ctx,
		model.Gemini25Flash,
		genai.Text("What is the sum of the first 50 prime numbers? "+
			"Generate and run code for the calculation, and make sure you get all 50."),
		config,
	)
	if err != nil {
		panic(err)
	}

	fmt.Println(res.Text())
	fmt.Println(res.ExecutableCode())
	fmt.Println(res.CodeExecutionResult())
}

// DialogCodingExec 在对话中使用代码执行
func DialogCodingExec() {
	ctx := context.Background()
	c := client.NewClient(ctx, os.Getenv("GEMINI_API_KEY"))

	config := &genai.GenerateContentConfig{
		Tools: []*genai.Tool{
			{CodeExecution: &genai.ToolCodeExecution{}},
		},
	}

	chat, err := c.Chats.Create(
		ctx,
		model.Gemini25Flash,
		config,
		nil,
	)
	if err != nil {
		panic(err)
	}

	result, err := chat.SendMessage(
		ctx,
		genai.Part{Text: "What is the sum of the first 50 prime numbers? " +
			"Generate and run code for the calculation, and " +
			"make sure you get all 50.",
		},
	)
	if err != nil {
		panic(err)
	}

	fmt.Println(result.Text())
	fmt.Println(result.ExecutableCode())
	fmt.Println(result.CodeExecutionResult())
}
