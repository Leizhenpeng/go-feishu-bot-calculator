package main

import (
	"context"
	"encoding/json"
	"fmt"
	"leizhenpeng/go-feishu-bot-calculator/calc"
	"net/http"
	"os"
	"regexp"

	larkcore "github.com/larksuite/oapi-sdk-go/v3/core"
	larkevent "github.com/larksuite/oapi-sdk-go/v3/event"

	"github.com/larksuite/oapi-sdk-go/v3/core/httpserverext"
	"github.com/larksuite/oapi-sdk-go/v3/event/dispatcher"
	larkim "github.com/larksuite/oapi-sdk-go/v3/service/im/v1"

	lark "github.com/larksuite/oapi-sdk-go/v3"
)

func setEnv() {
	os.Setenv("APP_ID", "xxx")
	os.Setenv("APP_SECRET", "xxx")
	os.Setenv("APP_VERIFICATION_TOKEN", "xxx")
	os.Setenv("APP_ENCRYPT_KEY", "xxx")
}

var client *lark.Client

func init() {
	setEnv()
}

func sendMsg(msg string, chatId *string) {
	content := larkim.NewTextMsgBuilder().
		Text(msg).
		Build()

	resp, err := client.Im.Message.Create(context.Background(), larkim.NewCreateMessageReqBuilder().
		ReceiveIdType(larkim.ReceiveIdTypeChatId).
		Body(larkim.NewCreateMessageReqBodyBuilder().
			MsgType(larkim.MsgTypeText).
			ReceiveId(*chatId).
			Content(content).
			Build()).
		Build())

	// 处理错误
	if err != nil {
		fmt.Println(err)
	}

	// 服务端错误处理
	if !resp.Success() {
		fmt.Println(resp.Code, resp.Msg, resp.RequestId())
	}
}
func msgFilter(msg string) string {
	//replace @到下一个非空的字段 为 ''
	regex := regexp.MustCompile(`@[^ ]*`)
	return regex.ReplaceAllString(msg, "")

}
func parseContent(content string) string {
	//"{\"text\":\"@_user_1  hahaha\"}",
	//only get text content hahaha
	var contentMap map[string]interface{}
	err := json.Unmarshal([]byte(content), &contentMap)
	if err != nil {
		fmt.Println(err)
	}
	text := contentMap["text"].(string)
	return msgFilter(text)
}
func main() {
	client = lark.NewClient(os.Getenv("APP_ID"), os.Getenv("APP_SECRET"))

	//// 注册消息处理器
	handler := dispatcher.NewEventDispatcher(os.Getenv("APP_VERIFICATION_TOKEN"), os.Getenv("APP_ENCRYPT_KEY")).
		OnP2MessageReceiveV1(func(ctx context.Context, event *larkim.P2MessageReceiveV1) error {
			fmt.Println(larkcore.Prettify(event))
			content := event.Event.Message.Content
			contentStr := parseContent(*content)
			out, err := calc.CalcStr(contentStr)
			if err != nil {
				fmt.Println(err)
			}
			sendMsg(calc.FormatMathOut(out), event.Event.Message.ChatId)
			return nil
		})

	// 注册 http 路由
	http.HandleFunc("/webhook/event", httpserverext.NewEventHandlerFunc(handler, larkevent.WithLogLevel(larkcore.LogLevelDebug)))

	// 启动 http 服务
	fmt.Println("http server started", "http://localhost:8080/webhook/event")

	err2 := http.ListenAndServe(":8080", nil)
	if err2 != nil {
		panic(err2)
	}

}
