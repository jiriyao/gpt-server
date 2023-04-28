package service

import (
	"fmt"
	"github.com/gogf/gf/util/gconv"
	"github.com/pkoukk/tiktoken-go"
	"github.com/sashabaranov/go-openai"
	"math"
)

const (
	GPT4_MODEL_TOKENS    = 32768
	GPT4_RESPONSE_TOKENS = 8192

	GPT3_MODEL_TOKENS    = 8192
	GPT3_RESPONSE_TOKENS = 2048
)

type ApiService struct {
	MaxModelTokens    int
	MaxResponseTokens int
	ApiMode           string //gpt3 or gpt4
	TimeoutMs         int
}

//api service for 切换管理每个api key的使用量和用法策略情况

func NewApiService(MaxModelTokens int, MaxResponseTokens int, ApiMode string, TimeoutMs int) *ApiService {
	return nil
}

func WithDefaultConfig(mode string) *ApiService {
	if mode == openai.GPT3Dot5Turbo {
		return &ApiService{
			MaxModelTokens:    GPT3_MODEL_TOKENS,
			MaxResponseTokens: GPT3_RESPONSE_TOKENS,
			ApiMode:           mode,
			TimeoutMs:         60000,
		}
	}

	return &ApiService{
		MaxModelTokens:    GPT4_MODEL_TOKENS,
		MaxResponseTokens: GPT4_RESPONSE_TOKENS,
		ApiMode:           mode,
		TimeoutMs:         60000,
	}

}

func (s *ApiService) DefaultSystemMessage() string {
	return "You are ChatGPT, a large language model trained by OpenAI. Follow the user's instructions carefully. Respond using markdown."
}

func (s *ApiService) QueryBalance(key string) float32 {
	return 0
}

func (s *ApiService) MaxTokens(numTokens int) int {
	var min = math.Min(gconv.Float64(s.MaxModelTokens-numTokens), gconv.Float64(s.MaxResponseTokens))
	var maxTokens = math.Max(1, min)
	return gconv.Int(maxTokens)
}

func (s *ApiService) isValidPrompt(numTokens int) bool {
	return numTokens < (s.MaxModelTokens - s.MaxResponseTokens)
}

//获取一个msg 里的token数量
func (s *ApiService) NumTokensFromMessages(messages []openai.ChatCompletionMessage, model string) (num_tokens int) {
	tkm, err := tiktoken.EncodingForModel(model)
	if err != nil {
		err = fmt.Errorf("EncodingForModel: %v", err)
		fmt.Println(err)
		return
	}

	var tokens_per_message int
	var tokens_per_name int
	if model == "gpt-3.5-turbo-0301" || model == "gpt-3.5-turbo" {
		tokens_per_message = 4
		tokens_per_name = -1
	} else if model == "gpt-4-0314" || model == "gpt-4" {
		tokens_per_message = 3
		tokens_per_name = 1
	} else {
		fmt.Println("Warning: model not found. Using cl100k_base encoding.")
		tokens_per_message = 3
		tokens_per_name = 1
	}

	for _, message := range messages {
		num_tokens += tokens_per_message
		num_tokens += len(tkm.Encode(message.Content, nil, nil))
		num_tokens += len(tkm.Encode(message.Role, nil, nil))
		if message.Name != "" {
			num_tokens += tokens_per_name
		}
	}
	num_tokens += 3
	return num_tokens
}
