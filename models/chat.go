package models

type ChatRequest struct {
	UserID  string `json:"user_id" binding:"required" example:"user123"`
	Message string `json:"message" binding:"required,min=1,max=1000" example:"What is the capital of France?"`
	Context struct {
		SystemPrompt string            `json:"system_prompt,omitempty"`
		Metadata     map[string]string `json:"metadata,omitempty"`
		Temperature  *float32          `json:"temperature,omitempty"`
	} `json:"context,omitempty"`
}

type ChatResponse struct {
	Success  bool   `json:"success"`
	Response string `json:"response,omitempty"`
	Error    string `json:"error,omitempty"`
}

type OpenAIRequest struct {
	Model       string    `json:"model"`
	Messages    []Message `json:"messages"`
	Temperature float32   `json:"temperature"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// New struct for maintaining conversation context
type Conversation struct {
	UserID      string    `json:"user_id"`
	Messages    []Message `json:"messages"`
	LastUpdated int64     `json:"last_updated"`
}

type OpenAIResponse struct {
	ID      string   `json:"id"`
	Object  string   `json:"object"`
	Created int      `json:"created"`
	Choices []Choice `json:"choices"`
}

type Choice struct {
	Index        int     `json:"index"`
	Message      Message `json:"message"`
	FinishReason string  `json:"finish_reason"`
}
