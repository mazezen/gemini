// Veo 3 是 Google 最先进的模型，可根据文字提示生成高保真 8 秒 720p 视频，
// 具有惊人的逼真效果和原生生成的音频。Veo 3 擅长各种视觉和电影风格。
// 选择以下示例，了解如何生成包含对话、电影级真实感或创意动画的视频

package voe3

import (
	"context"
	"github.com/mazezen/gemini/client"
	"github.com/mazezen/gemini/model"
	"log"
	"os"
	"time"
)

// VeoGen 使用 Veo 3 生成视频
func VeoGen() {
	ctx := context.Background()
	c := client.NewClient(ctx, os.Getenv("GEMINI_API_KEY"))

	prompt := `A close up of two people staring at a cryptic drawing on a wall, torchlight flickering.
    A man murmurs, 'This must be it. That's the secret code.' The woman looks at him and whispering excitedly, 'What did you find?'`

	operation, err := c.Models.GenerateVideos(
		ctx,
		model.Veo30GeneratePreview,
		prompt,
		nil,
		nil,
	)
	if err != nil {
		panic(err)
	}

	for !operation.Done {
		log.Println("Waiting for operation...")
		time.Sleep(10 * time.Second)
		operation, err = c.Operations.GetVideosOperation(ctx, operation, nil)
		if err != nil {
			panic(err)
		}
	}

	// download the generated video
	video := operation.Response.GeneratedVideos[0]
	c.Files.Download(ctx, video.Video, nil)
	fName := "images/dialogue_example.mp4"
	_ = os.WriteFile(fName, video.Video.VideoBytes, 0644)
	log.Printf("Generated video saved to %s\n", fName)
}

// ImageGenVideo 根据图片生成视频
// 使用 Imagen 生成图片，然后将该图片用作视频的起始帧
func ImageGenVideo() {
	ctx := context.Background()
	c := client.NewClient(ctx, os.Getenv("GEMINI_API_KEY"))

	prompt := "Panning wide shot of a calico kitten sleeping in the sunshine"

	// Step 1: Generate an image with Imagen
	imagenResponse, err := c.Models.GenerateImages(
		ctx,
		model.ImaGen30Generate002,
		prompt,
		nil,
	)
	if err != nil {
		panic(err)
	}

	// Step 2: Generate video with Veo 2 using the image
	operation, err := c.Models.GenerateVideos(
		ctx,
		model.Veo20Generate001,
		prompt,
		imagenResponse.GeneratedImages[0].Image,
		nil,
	)

	// Poll the operation status until the video is ready
	for !operation.Done {
		log.Println("Waiting for operation...")
		time.Sleep(10 * time.Second)
		operation, err = c.Operations.GetVideosOperation(ctx, operation, nil)
		if err != nil {
			panic(err)
		}
	}

	video := operation.Response.GeneratedVideos[0]
	c.Files.Download(ctx, video.Video, nil)
	fName := "images/veo2_with_image_input.mp4"
	_ = os.WriteFile(fName, video.Video.VideoBytes, 0644)
	log.Printf("Generated video saved to %s\n", fName)
}
