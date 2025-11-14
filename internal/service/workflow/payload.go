package workflow

type Payload struct {
	WorkflowId string  `json:"workflow_id"`
	Inputs     Inputs  `json:"inputs"`
	Configs    Configs `json:"configs"`
}

type Inputs struct {
	TextProcessing TextProcessing `json:"text_processing"`
	LlmCalling     LlmCalling     `json:"llm_calling,omitempty"`
}

type TextProcessing struct {
	Query          any       `json:"query,omitempty"`
	TaskType       string    `json:"task_type,omitempty"`
	TargetLanguage string    `json:"target_language,omitempty"`
	FullText       string    `json:"full_text,omitempty"`
	LlmConfig      LlmConfig `json:"llm_config"`
	Temperature    float64   `json:"temperature"`
	MaxTokens      int       `json:"max_tokens"`
}

type LlmCalling struct {
	Messages  any       `json:"messages,omitempty"`
	FullText  string    `json:"full_text,omitempty"`
	LlmConfig LlmConfig `json:"llm_config"`
	Stream    bool      `json:"stream"`
}

type LlmConfig struct {
	Provider string `json:"provider"`
	Model    string `json:"model"`
	BaseUrl  string `json:"base_url"`
	ApiKey   string `json:"api_key"`
}

type Configs struct {
	EncryptionConfig EncryptionConfig `json:"encryption_config"`
}

type EncryptionConfig struct {
	Type   string `json:"type"`
	IsOpen bool   `json:"is_open"`
}

func NewPayload() *Payload {
	return &Payload{
		WorkflowId: "",
		Inputs: Inputs{
			TextProcessing: TextProcessing{
				Query:          "",
				TaskType:       "",
				TargetLanguage: "",
				FullText:       "",
				LlmConfig: LlmConfig{
					Provider: "openai_api_compatible",
					Model:    "Qwen2.5-VL-7B-Instruct",
					BaseUrl:  "https://studio.bd.kxsz.net:9443/v1",
					ApiKey:   "sk-7f794f453167c38fe362ea7a3ac77206",
				},
				Temperature: 0.3,
				MaxTokens:   2048,
			},
			LlmCalling: LlmCalling{
				Messages: "",
				FullText: "",
				LlmConfig: LlmConfig{
					Provider: "openai_api_compatible",
					Model:    "Qwen2.5-VL-7B-Instruct",
					BaseUrl:  "https://studio.bd.kxsz.net:9443/v1",
					ApiKey:   "sk-7f794f453167c38fe362ea7a3ac77206",
				},
				Stream: true,
			},
		},
		Configs: Configs{
			EncryptionConfig: EncryptionConfig{
				Type:   "PKCS1_OAEP",
				IsOpen: false,
			},
		},
	}
}

func (p *Payload) SetWorkflowID(workflowID string) *Payload {
	p.WorkflowId = workflowID
	return p
}

func (p *Payload) SetQuery(query any) *Payload {
	p.Inputs.TextProcessing.Query = query
	return p
}

func (p *Payload) SetTaskType(taskType string) *Payload {
	p.Inputs.TextProcessing.TaskType = taskType
	return p
}
func (p *Payload) SetTargetLanguage(targetLanguage string) *Payload {
	p.Inputs.TextProcessing.TargetLanguage = targetLanguage
	return p
}

func (p *Payload) SetFullText(fullText string) *Payload {
	p.Inputs.TextProcessing.FullText = fullText
	return p
}

func (p *Payload) SetLlmConfig(llmConfig LlmConfig) *Payload {
	p.Inputs.TextProcessing.LlmConfig = llmConfig
	return p
}

func (p *Payload) SetTemperature(temperature float64) *Payload {
	p.Inputs.TextProcessing.Temperature = temperature
	return p
}

func (p *Payload) SetMaxTokens(maxTokens int) *Payload {
	p.Inputs.TextProcessing.MaxTokens = maxTokens
	return p
}

func (p *Payload) SetEncryptionConfig(encryptionConfig EncryptionConfig) *Payload {
	p.Configs.EncryptionConfig = encryptionConfig
	return p
}

func (p *Payload) SetChatFullText(fullText string) *Payload {
	p.Inputs.LlmCalling.FullText = fullText
	return p
}

func (p *Payload) SetChatMessages(messages any) *Payload {
	p.Inputs.LlmCalling.Messages = messages
	return p
}
