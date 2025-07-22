# <p style="text-align: center;">gemini</p>
> 调用 google gemini 示例

## gemini
> 谷歌推出的 AI 只能模型
* <a href="https://deepmind.google/models/gemini/" target="_blank">Gemini官网</a>
* <a href="https://gemini.google.com/app" target="_blank">Gemini Chat</a>
* <a href="https://ai.google.dev/gemini-api/docs?hl=zh-cn" target="_blank">Gemini开发者文档</a>
* <a href="https://github.com/googleapis/go-genai" target="_blank">Gemini SDK</a>

## 所有模型
| 模型变体                                                     | 输入                         | 输出                 | 优化目标                                             |
| :----------------------------------------------------------- | :--------------------------- | :------------------- | :--------------------------------------------------- |
| [Gemini 2.5 Pro](https://ai.google.dev/gemini-api/docs/models?hl=zh-cn#gemini-2.5-pro) `gemini-2.5-pro` | 音频、图片、视频、文本和 PDF | 文本                 | 增强的思考和推理能力、多模态理解能力、高级编码能力等 |
| [Gemini 2.5 Flash](https://ai.google.dev/gemini-api/docs/models?hl=zh-cn#gemini-2.5-flash) `gemini-2.5-flash` | 音频、图片、视频和文本       | 文本                 | 适应性思维，成本效益                                 |
| [Gemini 2.5 Flash-Lite 预览版](https://ai.google.dev/gemini-api/docs/models?hl=zh-cn#gemini-2.5-flash-lite) `gemini-2.5-flash-lite-preview-06-17` | 文本、图片、视频、音频       | 文本                 | 最具成本效益且支持高吞吐量的模型                     |
| [Gemini 2.5 Flash 原生音频](https://ai.google.dev/gemini-api/docs/models?hl=zh-cn#gemini-2.5-flash-native-audio) `gemini-2.5-flash-preview-native-audio-dialog` & `gemini-2.5-flash-exp-native-audio-thinking-dialog` | 音频、视频和文本             | 文本和音频，交错显示 | 高质量、自然的对话式音频输出，无论是否经过思考       |
| [Gemini 2.5 Flash 预览版 TTS](https://ai.google.dev/gemini-api/docs/models?hl=zh-cn#gemini-2.5-flash-preview-tts) `gemini-2.5-flash-preview-tts` | 文本                         | 音频                 | 低延迟、可控的单语音和多语音文字转语音音频生成       |
| [Gemini 2.5 Pro 预览版 TTS](https://ai.google.dev/gemini-api/docs/models?hl=zh-cn#gemini-2.5-pro-preview-tts) `gemini-2.5-pro-preview-tts` | 文本                         | 音频                 | 低延迟、可控的单语音和多语音文字转语音音频生成       |
| [Gemini 2.0 Flash](https://ai.google.dev/gemini-api/docs/models?hl=zh-cn#gemini-2.0-flash) `gemini-2.0-flash` | 音频、图片、视频和文本       | 文本                 | 新一代功能、速度和实时流式传输。                     |
| [Gemini 2.0 Flash 预览版图片生成](https://ai.google.dev/gemini-api/docs/models?hl=zh-cn#gemini-2.0-flash-preview-image-generation) `gemini-2.0-flash-preview-image-generation` | 音频、图片、视频和文本       | 文字、图片           | 对话式图片生成和编辑                                 |
| [Gemini 2.0 Flash-Lite](https://ai.google.dev/gemini-api/docs/models?hl=zh-cn#gemini-2.0-flash-lite) `gemini-2.0-flash-lite` | 音频、图片、视频和文本       | 文本                 | 成本效益和低延迟                                     |
| [Gemini 1.5 Flash](https://ai.google.dev/gemini-api/docs/models?hl=zh-cn#gemini-1.5-flash) `gemini-1.5-flash` | 音频、图片、视频和文本       | 文本                 | 在各种任务中提供快速而多样的性能 已弃用              |
| [Gemini 1.5 Flash-8B](https://ai.google.dev/gemini-api/docs/models?hl=zh-cn#gemini-1.5-flash-8b) `gemini-1.5-flash-8b` | 音频、图片、视频和文本       | 文本                 | 量大且智能程度较低的任务 已弃用                      |
| [Gemini 1.5 Pro](https://ai.google.dev/gemini-api/docs/models?hl=zh-cn#gemini-1.5-pro) `gemini-1.5-pro` | 音频、图片、视频和文本       | 文本                 | 需要更高智能的复杂推理任务 已弃用                    |
| [Gemini Embedding](https://ai.google.dev/gemini-api/docs/models?hl=zh-cn#gemini-embedding) `gemini-embedding-001` | 文本                         | 文本嵌入             | 衡量文本字符串的相关性                               |
| [Imagen 4](https://ai.google.dev/gemini-api/docs/models?hl=zh-cn#imagen-4) `imagen-4.0-generate-preview-06-06` `imagen-4.0-ultra-generate-preview-06-06` | 文本                         | 图片                 | 我们最新的图片生成模型                               |
| [Imagen 3](https://ai.google.dev/gemini-api/docs/models?hl=zh-cn#imagen-3) `imagen-3.0-generate-002` | 文本                         | 图片                 | 高质量图片生成模型                                   |
| [Veo 3 预览版](https://ai.google.dev/gemini-api/docs/models?hl=zh-cn#veo-3) `veo-3.0-generate-preview` | 文本                         | 带音频的视频         | 生成包含音效、环境噪音和对话的高品质视频             |
| [Veo 2](https://ai.google.dev/gemini-api/docs/models?hl=zh-cn#veo-2) `veo-2.0-generate-001` | 文字、图片                   | 视频                 | 高质量视频生成                                       |
| [Gemini 2.5 Flash Live](https://ai.google.dev/gemini-api/docs/models?hl=zh-cn#live-api) `gemini-live-2.5-flash-preview` | 音频、视频和文本             | 文字、音频           | 低延迟的双向语音和视频互动                           |
| [Gemini 2.0 Flash Live](https://ai.google.dev/gemini-api/docs/models?hl=zh-cn#live-api-2.0) `gemini-2.0-flash-live-001` | 音频、视频和文本             | 文字、音频           | 低延迟的双向语音和视频互动                           |

## 引入SDK
```shell
go get google.golang.org/genai
```

## 设置GEMINI API KEY 环境变量
```shell
export GEMINI_API_KEY="your-api-key"
```

## 示例

### 文本生成
1. 单个文本输入
2. 思考功能
3. 系统指令来引导 Gemini 模型的行为
4. 借助 GenerateContentConfig 对象, 替换默认生成参数
5. 多模态输入，将文本与媒体文件组合使用
6. 流式响应
7. 多轮对话
8. 流式响应多轮对话

### 图片生成
1. 文本生成图片
2. 图片编辑
3. 使用 Imagen 模型生成图片 (需要付费)

### 生成视频
1. 使用 Veo 3 生成视频 (要使用此模式，请确保您的帐户具有活动的GCP计费) 
2. 根据图片生成视频 (需要付费)

### 结构化输出
1. 结构化输出

### 文档理解
1. 传递内嵌 PDF 数据
2. 使用 File API 上传 PDF

### 图片理解
1. 传递内嵌图片数据
2. 使用 File API 上传图片

### 视频理解
1. YouTuBe

### 代码执行
1. 生成代码, 并执行代码
2. 在对话中使用代码执行

```go
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
```
