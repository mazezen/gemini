package main

import (
	"github.com/mazezen/gemini/coding"
	"github.com/mazezen/gemini/docs"
	"github.com/mazezen/gemini/format"
	"github.com/mazezen/gemini/photo"
	"github.com/mazezen/gemini/text"
	"github.com/mazezen/gemini/video"
	"github.com/mazezen/gemini/voe3"
)

/**
 * 开发文档: https://ai.google.dev/gemini-api/docs?hl=zh-cn
 */
func main() {

	// 文本生成
	text.TexGen()
	text.TexGenThinking()
	text.TexGenCli()
	text.TexGenDefaultConfig()
	text.TexGenMultiModal()
	text.TexGenStream()
	text.TexMultipleRoundsOfDialogue()
	text.TexGenStreamMultipleRoundsOfDialogue()

	// 图片生成
	photo.PhoGen()
	photo.PhoEdit()
	photo.ImaGen40GeneratePreview0606()

	// 生成视频
	voe3.VeoGen()
	voe3.ImageGenVideo()

	// 结构化输出
	format.FormatToJson()

	// 代码执行
	coding.EnableCodingExec()
	coding.DialogCodingExec()

	// 文档理解
	docs.InnerDoc()
	docs.UploadFileApi()

	// 图片理解
	photo.PhoInnerRead()
	photo.PhoFileApi()

	// 视频理解
	video.YouTuBe()
}
