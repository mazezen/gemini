package docs

import (
	"context"
	"fmt"
	"github.com/mazezen/gemini/client"
	"github.com/mazezen/gemini/model"
	"google.golang.org/genai"
	"io"
	"net/http"
	"os"
)

// InnerDoc 传递内嵌 PDF 数据
func InnerDoc() {
	ctx := context.Background()
	c := client.NewClient(ctx, os.Getenv("GEMINI_API_KEY"))

	pdfResp, err := http.Get("https://discovery.ucl.ac.uk/id/eprint/10089234/1/343019_3_art_0_py4t4l_convrt.pdf")
	if err != nil {
		panic(err)
	}

	var pdfBytes []byte
	if pdfResp != nil && pdfResp.Body != nil {
		pdfBytes, _ = io.ReadAll(pdfResp.Body)
		pdfResp.Body.Close()
	}

	parts := []*genai.Part{
		&genai.Part{
			InlineData: &genai.Blob{
				MIMEType: "application/pdf",
				Data:     pdfBytes,
			},
		},
		genai.NewPartFromText("Summarize this document"),
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

// UploadFileApi 使用 File API 上传 PDF
// 您可以使用文件 API 上传更大的文档。当总请求大小（包括文件、文本提示、系统指令等）超过 20MB 时，请务必使用 File API。
func UploadFileApi() {
	ctx := context.Background()
	c := client.NewClient(ctx, os.Getenv("GEMINI_API_KEY"))

	pdfResp, err := http.Get("https://discovery.ucl.ac.uk/id/eprint/10089234/1/343019_3_art_0_py4t4l_convrt.pdf")
	if err != nil {
		panic(err)
	}
	defer pdfResp.Body.Close()

	localPdfPath := "images/convrt.pdf"
	outFile, err := os.Create(localPdfPath)
	if err != nil {
		panic(err)
	}
	defer outFile.Close()

	_, _ = io.Copy(outFile, pdfResp.Body)

	uploadConfig := &genai.UploadFileConfig{MIMEType: "application/pdf"}
	uploadedFile, _ := c.Files.UploadFromPath(ctx, localPdfPath, uploadConfig)

	promptParts := []*genai.Part{
		genai.NewPartFromURI(uploadedFile.URI, uploadedFile.MIMEType),
		genai.NewPartFromText("Summarize this document"),
	}
	contents := []*genai.Content{
		genai.NewContentFromParts(promptParts, genai.RoleUser), // Specify role
	}

	result, _ := c.Models.GenerateContent(
		ctx,
		model.Gemini25Flash,
		contents,
		nil,
	)

	fmt.Println(result.Text())
}
