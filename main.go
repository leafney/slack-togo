/**
 * @Author:      leafney
 * @GitHub:      https://github.com/leafney
 * @Project:     slack-togo
 * @Date:        2023-09-16 23:25
 * @Description:
 */

package main

import (
	"github.com/leafney/rose"
	"github.com/slack-go/slack"
	"log"
)

const (
	// bot token
	TOKEN = "xoxb-"
	// user token
	//TOKEN     = "xoxp-"
	CHANNELID = ""
)

func main() {
	api := slack.New(TOKEN)

	ConversationHistory(api)

	//PostMessage(api)

}

// 历史消息记录
func ConversationHistory(api *slack.Client) {
	// 获取历史消息记录
	result, err := api.GetConversationHistory(&slack.GetConversationHistoryParameters{
		IncludeAllMetadata: true,
		Limit:              1000,
		ChannelID:          CHANNELID,
	})
	if err != nil {
		log.Println(err)
		return
	}

	data := rose.JsonMarshalStr(result)
	//log.Println(data)
	rose.FWriteFile("111.json", data, false)
	//for _, msg := range result.Messages {
	//	msg.
	//}

}

// 二级内容列表
func ConversationReplies(api *slack.Client) {
	//	获取消息项的回复列表
	//- [conversations.replies method | Slack](https://api.slack.com/methods/conversations.replies)
	result, _, _, _ := api.GetConversationReplies(&slack.GetConversationRepliesParameters{
		IncludeAllMetadata: true,
		Timestamp:          "1687698556.243869",
		ChannelID:          CHANNELID,
	})

	data := rose.JsonMarshalStr(result)
	rose.FWriteFile("456.json", data, false)
}

// 发送消息
func PostMessage(api *slack.Client) {
	//	发送消息
	resCid, resTs, _ := api.PostMessage(
		CHANNELID,
		slack.MsgOptionText("Hello world as User 444", false),
		//slack.MsgOptionAsUser(true),
		//slack.MsgOptionUser("U0563221SBY"),
	)

	log.Printf("result channelid %s timestamp %s", resCid, resTs)
}