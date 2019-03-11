package api_test

import (
	"log"

	"github.com/ganlvtech/go-bilibili-api"
)

func ExampleBilibiliApiClient_Login() {
	b := api.NewBilibiliApiClient(false)
	err := b.Login("username", "password", "access_token", "refresh_token", []byte("cookie"))
	if err != nil {
		log.Println(err)
	}
}

func ExampleBilibiliApiClient_GetCookies() {
	b := BilibiliApiClient()
	cookie, err := b.SaveCookie()
	if err != nil {
		log.Println(err)
	}
	log.Println(cookie)
}

func ExampleBilibiliApiClient_LoadCookies() {
	b := BilibiliApiClient()
	cookie := "{\"bigfun.cn\":{\"bigfun.cn;/;DedeUserID\":{\"Name\":...}}}"
	err := b.LoadCookies([]byte(cookie))
	if err != nil {
		log.Println(err)
	}
}

func ExampleRoomInfo() {
	roomInfoResponse, err := api.RoomInfo(3)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(roomInfoResponse)
}

func ExampleRoomInit() {
	roomInitResponse, err := api.RoomInit(3)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(roomInitResponse)
}

// 返回一个已登录的 BilibiliApiClient
func BilibiliApiClient() *api.BilibiliApiClient {
	b := api.NewBilibiliApiClient(false)
	err := b.Login("username", "password", "access_token", "refresh_token", []byte("cookie"))
	if err != nil {
		log.Println(err)
	}
	return b
}

func ExampleGiftConfig() {
	giftConfigResponse, err := api.GiftConfig()
	if err != nil {
		log.Println(err)
	} else {
		for _, v := range giftConfigResponse.Data {
			draw := "不可抽奖"
			if v.Draw > 0 {
				draw = "可触发抽奖"
			}
			log.Println(v.Name, "id:", v.ID, "price:", v.Price, v.CoinType, v.Desc, v.Rights, draw)
		}
	}
}

func ExampleRoomGiftList() {
	giftConfigResponse, err := api.GiftConfig()
	if err != nil {
		log.Println(err)
		return
	}
	giftConfigMap := make(map[int]int)
	for k, v := range giftConfigResponse.Data {
		giftConfigMap[v.ID] = k
	}
	roomGiftListResponse, err := api.RoomGiftList(23058)
	if err != nil {
		log.Println(err)
		return
	}
	for _, v := range roomGiftListResponse.Data.List {
		v2 := giftConfigResponse.Data[giftConfigMap[v.ID]]
		draw := "不可抽奖"
		if v2.Draw > 0 {
			draw = "可触发抽奖"
		}
		log.Println(v2.Name, "id:", v2.ID, "price:", v2.Price, v2.CoinType, v2.Desc, v2.Rights, draw)
	}
}

func ExampleBilibiliApiClient_GetBagList() {
	b := BilibiliApiClient()
	bagListResponse, err := b.GetBagList()
	if err != nil {
		log.Println(err)
	} else {
		log.Println("包裹信息", bagListResponse)
	}
}

func ExampleBilibiliApiClient_BagSend() {
	roomInitResponse, err := api.RoomInit(3)
	if err != nil {
		log.Println(err)
		return
	}

	b := BilibiliApiClient()
	bagListResponse, err := b.GetBagList()
	if err != nil {
		log.Println(err)
		return
	}

	for _, v := range bagListResponse.Data.List {
		sendBagGiftResponse, err := b.BagSend(roomInitResponse.Data.RoomId, roomInitResponse.Data.Uid, v.BagId, v.GiftId, 1)
		if err != nil {
			log.Println("赠送失败", err.Error())
		} else {
			log.Println(sendBagGiftResponse)
		}
	}
}

func ExampleBilibiliApiClient_GetDanmakuConfig() {
	b := BilibiliApiClient()
	danmakuConfigResponse, err := b.GetDanmakuConfig(23058)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("Length", danmakuConfigResponse.Data.Length)
	log.Println("Color", danmakuConfigResponse.Data.Color)
	log.Println("Mode", danmakuConfigResponse.Data.Mode)
}

func ExampleBilibiliApiClient_SendLiveMessage() {
	b := BilibiliApiClient()
	_, err := b.SendLiveMessage(23058, "Hello, world!")
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("弹幕发送成功")
}

func ExampleBilibiliApiClient_ReceiveDailyBag() {
	b := BilibiliApiClient()
	_, err := b.ReceiveDailyBag()
	if err != nil {
		log.Println(err)
	} else {
		log.Println("领取每日礼包成功")
	}
}

func ExampleBilibiliApiClient_SilverToCoinWeb() {
	b := BilibiliApiClient()
	_, err := b.SilverToCoinWeb()
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("网页端银瓜子兑换硬币成功")
}

func ExampleBilibiliApiClient_SilverToCoinApp() {
	b := BilibiliApiClient()
	_, err := b.SilverToCoinApp()
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("移动端银瓜子兑换硬币成功")
}

func ExampleBilibiliApiClient_GetTaskInfo() {
	b := BilibiliApiClient()
	taskInfoResponse, err := b.GetTaskInfo()
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(taskInfoResponse.Data)
}

func ExampleBilibiliApiClient_UserOnlineHeartbeatWeb() {
	b := BilibiliApiClient()
	_, err := b.UserOnlineHeartbeatWeb(23058)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("网页端心跳")
	}
}

func ExampleBilibiliApiClient_UserOnlineHeartbeatMobile() {
	b := BilibiliApiClient()
	_, err := b.UserOnlineHeartbeatMobile(23058)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("移动端心跳")
	}
}

func ExampleSilverBoxTask_FreeSilverAward() {
	b := BilibiliApiClient()
	silverBoxTask, _, err := b.SilverBoxGetCurrentTask()
	if err != nil {
		log.Println(err)
	} else {
		award, err := silverBoxTask.FreeSilverAward()
		if err != nil {
			log.Println(err)
		} else {
			log.Println("领取银瓜子宝箱成功", award.Data.Silver)
		}
	}
}

func ExampleBilibiliApiClient_DailySignWeb() {
	b := BilibiliApiClient()
	_, err := b.DailySignWeb()
	if err != nil {
		log.Println(err)
		return
	}
}

func ExampleBilibiliApiClient_DailySignApp() {
	b := BilibiliApiClient()
	_, err := b.DailySignApp()
	if err != nil {
		log.Println(err)
		return
	}
}

func ExampleBilibiliApiClient_GetSignInfoWeb() {
	b := BilibiliApiClient()
	signInfoResponse, err := b.GetSignInfoWeb()
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(signInfoResponse)
}

func ExampleBilibiliApiClient_ReceiveAward() {
	b := BilibiliApiClient()
	_, err := b.ReceiveAward("double_watch_task")
	if err != nil {
		log.Println(err)
		return
	}
}
