package tdameritrade

import (
	"chatgpt4/fixtures"
	tdameritradeTypes "chatgpt4/tdameritrade/types"
	"chatgpt4/utils"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/sashabaranov/go-openai"
)

func StartStreamingData(loc *time.Location) {
	// run in go routine
	go func() {
		userPrincipalsString := os.Getenv("TD_AMERITRADE_USER_PRINCIPALS")

		// connect to TD Ameritrade
		conn, requestId, userPrincipals, err := ConnectToTDAmeritradeWithExpressConnection(userPrincipalsString)
		if err != nil {
			log.Println("Error connecting to TD Ameritrade: ", err)
			return
		}
		defer conn.Close()

		// get candlestick data for ES
		*requestId += 1
		err = GetChartHistoryFutures("/ES", "1m", "1d", *requestId, conn, *userPrincipals)
		if err != nil {
			log.Println("Error streaming futures quotes: ", err)
			return
		}

		// loop forever listening for messages
		for {
			// logic on receiving a message
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Println("Error reading message:", err)
				return
			}

			// print message to console (useful for debugging)
			// log.Printf("recv: %s", message)

			messageStruct := tdameritradeTypes.StreamMessage{}

			// unmarshal into a struct
			err = json.Unmarshal(message, &messageStruct)
			if err != nil {
				log.Println("unmarshal failed: ", err)

				// write the message to a file for debugging
				utils.WriteByteArrayToFile(message, "message.txt")

				return
			}

			data := messageStruct.Data

			// loop at the entries
			for _, entry := range data {

				// stock quotes
				if entry.Service == "CHART_HISTORY_FUTURES" {

					// cast entry.Content to QuoteContent
					chartContent := entry.Content.(*[]tdameritradeTypes.ChartFuturesContent)

					// convert chart data to OHLCV
					ohlcv := ConvertChartToOHLCV(chartContent)

					// loop through the chart content and prompt gpt with the candlestick data
					for _, candlestick := range ohlcv {
						fmt.Println("Querying...")

						// convert candlestick to string
						content := fmt.Sprintf("%v", candlestick)

						client := openai.NewClient(os.Getenv("OPEN_AI_API_TOKEN"))
						resp, err := client.CreateChatCompletion(
							context.Background(),
							openai.ChatCompletionRequest{
								Model:    openai.GPT3Dot5Turbo0613,
								Messages: fixtures.ESOneMinuteTrader(content),
							},
						)

						if err != nil {
							fmt.Printf("ChatCompletion error: %v\n", err)
							return
						}

						fmt.Println(resp.Choices[0].Message.Content)
					}
				}

				// // futures quote
				// if entry.Service == "LEVELONE_FUTURES" {

				// 	// cast entry.Content to FuturesQuoteContent
				// 	quoteContent := entry.Content.(*[]tdameritradeTypes.QuoteContent)

				// 	// loop through the quote content
				// 	for _, content := range *quoteContent {
				// 		// print bid price (key 1 if it is not 0)
				// 		if content.One != 0 {
				// 			log.Println(content.Key+" bid: ", content.One)
				// 			socket.BroadcastMessage(models.Ticker{
				// 				Symbol: content.Key,
				// 				Price:  content.One,
				// 			})
				// 		}
				// 	}

				// }

			}
		}
	}()
}

func ConvertChartToOHLCV(chart *[]tdameritradeTypes.ChartFuturesContent) []tdameritradeTypes.OHLCV {
	ohlcv := make([]tdameritradeTypes.OHLCV, len(*chart))

	for i, item := range *chart {
		ohlcv[i] = tdameritradeTypes.OHLCV{
			O: item.Open,
			H: item.High,
			L: item.Low,
			C: item.Close,
			V: item.Volume,
		}
	}

	return ohlcv
}
