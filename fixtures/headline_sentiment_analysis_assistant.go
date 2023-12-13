package fixtures

import "github.com/sashabaranov/go-openai"

func HeadLineSentimentAnalysis(input string) []openai.ChatCompletionMessage {
	return []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: "You are an advanced and accurate news headline sentiment analysis assistant. Analyze news articles and provide the answer in CSV format with the company or currency name in symbol format and then 'bullish', 'bearish', or 'neutral'. If a company or currency is not mentioned, provide the sector name, or commodity name, i.e. 'OIL' or 'CORN', or say 'MARKET WIDE'.",
		},
		{
			Role:    openai.ChatMessageRoleUser,
			Content: "Intel to sell $1.5 bln stake in Mobileye",
		},
		{
			Role:    "assistant",
			Content: "INTC,bearish",
		},
		{
			Role:    openai.ChatMessageRoleUser,
			Content: "London stocks muted ahead of British inflation data",
		},
		{
			Role:    "assistant",
			Content: "MARKET WIDE,bearish",
		},
		{
			Role:    openai.ChatMessageRoleUser,
			Content: "South African rand extends gains ahead of GDP data",
		},
		{
			Role:    "assistant",
			Content: "ZAR,bullish",
		},
		{
			Role:    openai.ChatMessageRoleUser,
			Content: "Saudi bourse inches higher on Fed rate pause hope; oil limits gains",
		},
		{
			Role:    "assistant",
			Content: "MARKET WIDE,bullish\nOIL,bearish",
		},
		{
			Role:    openai.ChatMessageRoleUser,
			Content: input,
		},
	}
}
