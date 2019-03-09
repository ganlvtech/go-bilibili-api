package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type RoomInitResponse struct {
	Code    int    `json:"code"`
	Msg     string `json:"msg"`
	Message string `json:"message"`
	Data    struct {
		RoomId      int  `json:"room_id"`
		ShortId     int  `json:"short_id"`
		Uid         int  `json:"uid"`
		NeedP2p     int  `json:"need_p2p"`
		IsHidden    bool `json:"is_hidden"`
		IsLocked    bool `json:"is_locked"`
		IsPortrait  bool `json:"is_portrait"`
		LiveStatus  int  `json:"live_status"`
		HiddenTill  int  `json:"hidden_till"`
		LockTill    int  `json:"lock_till"`
		Encrypted   bool `json:"encrypted"`
		PwdVerified bool `json:"pwd_verified"`
		LiveTime    int  `json:"live_time"`
		RoomShield  int  `json:"room_shield"`
		IsSp        int  `json:"is_sp"`
		SpecialType int  `json:"special_type"`
	} `json:"data"`
}

// 获取直播间基本信息
//
// GET https://api.live.bilibili.com/room/v1/Room/room_init?id=3
func RoomInit(shortId int) (*RoomInitResponse, error) {
	resp, err := http.Get(fmt.Sprintf("https://api.live.bilibili.com/room/v1/Room/room_init?id=%d", shortId))
	if err != nil {
		return nil, err
	}
	response := &RoomInitResponse{Code: -1}
	j := json.NewDecoder(resp.Body)
	err = j.Decode(response)
	if err != nil {
		return response, &ResponseJsonDecodeError{response.Message, err}
	}
	if response.Code != 0 {
		return response, &ResponseCodeNotZero{response.Message}
	}
	return response, nil
}

type RoomInfoResponse struct {
	Code    int    `json:"code"`
	Msg     string `json:"msg"`
	Message string `json:"message"`
	Data    struct {
		UID              int      `json:"uid"`
		RoomID           int      `json:"room_id"`
		ShortID          int      `json:"short_id"`
		Attention        int      `json:"attention"`
		Online           int      `json:"online"`
		IsPortrait       bool     `json:"is_portrait"`
		Description      string   `json:"description"`
		LiveStatus       int      `json:"live_status"`
		AreaID           int      `json:"area_id"`
		ParentAreaID     int      `json:"parent_area_id"`
		ParentAreaName   string   `json:"parent_area_name"`
		OldAreaID        int      `json:"old_area_id"`
		Background       string   `json:"background"`
		Title            string   `json:"title"`
		UserCover        string   `json:"user_cover"`
		Keyframe         string   `json:"keyframe"`
		IsStrictRoom     bool     `json:"is_strict_room"`
		LiveTime         string   `json:"live_time"`
		Tags             string   `json:"tags"`
		IsAnchor         int      `json:"is_anchor"`
		RoomSilentType   string   `json:"room_silent_type"`
		RoomSilentLevel  int      `json:"room_silent_level"`
		RoomSilentSecond int      `json:"room_silent_second"`
		AreaName         string   `json:"area_name"`
		Pendants         string   `json:"pendants"`
		AreaPendants     string   `json:"area_pendants"`
		HotWords         []string `json:"hot_words"`
		HotWordsStatus   int      `json:"hot_words_status"`
		Verify           string   `json:"verify"`
		NewPendants      struct {
			Frame interface{} `json:"frame"`
			Badge struct {
				Name     string `json:"name"`
				Position int    `json:"position"`
				Value    string `json:"value"`
				Desc     string `json:"desc"`
			} `json:"badge"`
			MobileFrame interface{} `json:"mobile_frame"`
			MobileBadge interface{} `json:"mobile_badge"`
		} `json:"new_pendants"`
		UpSession            string `json:"up_session"`
		PkStatus             int    `json:"pk_status"`
		PkID                 int    `json:"pk_id"`
		AllowChangeAreaTime  int    `json:"allow_change_area_time"`
		AllowUploadCoverTime int    `json:"allow_upload_cover_time"`
	} `json:"data"`
}

// 获取直播间详细信息
//
// GET https://api.live.bilibili.com/room/v1/Room/get_info?room_id=3
func RoomInfo(shortId int) (*RoomInfoResponse, error) {
	resp, err := http.Get(fmt.Sprintf("https://api.live.bilibili.com/room/v1/Room/get_info?room_id=%d", shortId))
	if err != nil {
		return nil, err
	}
	response := &RoomInfoResponse{Code: -1}
	j := json.NewDecoder(resp.Body)
	err = j.Decode(response)
	if err != nil {
		return response, &ResponseJsonDecodeError{response.Message, err}
	}
	if response.Code != 0 {
		return response, &ResponseCodeNotZero{response.Message}
	}
	return response, nil
}

type BagListResponse struct {
	Code    int    `json:"code"`
	Msg     string `json:"msg"`
	Message string `json:"message"`
	Data    struct {
		List []struct {
			BagId    int    `json:"bag_id"`
			GiftId   int    `json:"gift_id"`
			GiftName string `json:"gift_name"`
			GiftNum  int    `json:"gift_num"`
			GiftType int    `json:"gift_type"`
			ExpireAt int    `json:"expire_at"`
			CountMap []struct {
				Num  int    `json:"num"`
				Text string `json:"text"`
			} `json:"count_map"`
			CornerMark string `json:"corner_mark"`
		} `json:"list"`
		Time int `json:"time"`
	} `json:"data"`
}

// 获取免费的礼物包裹信息
//
// GET https://api.live.bilibili.com/gift/v2/gift/bag_list
func (b *BilibiliApiClient) GetBagList() (*BagListResponse, error) {
	resp, err := b.Client.Get("https://api.live.bilibili.com/gift/v2/gift/bag_list")
	if err != nil {
		return nil, err
	}
	response := &BagListResponse{Code: -1}
	j := json.NewDecoder(resp.Body)
	err = j.Decode(response)
	if err != nil {
		return response, &ResponseJsonDecodeError{response.Message, err}
	}
	if response.Code != 0 {
		return response, &ResponseCodeNotZero{response.Message}
	}
	return response, nil
}

type BagSendResponse struct {
	Code    int    `json:"code"`
	Msg     string `json:"msg"`
	Message string `json:"message"`
	Data    struct {
		Tid         string `json:"tid"`
		UID         int    `json:"uid"`
		Uname       string `json:"uname"`
		Face        string `json:"face"`
		GuardLevel  int    `json:"guard_level"`
		Ruid        int    `json:"ruid"`
		Rcost       int    `json:"rcost"`
		GiftID      int    `json:"gift_id"`
		GiftType    int    `json:"gift_type"`
		GiftName    string `json:"gift_name"`
		GiftNum     int    `json:"gift_num"`
		GiftAction  string `json:"gift_action"`
		GiftPrice   int    `json:"gift_price"`
		CoinType    string `json:"coin_type"`
		TotalCoin   int    `json:"total_coin"`
		PayCoin     int    `json:"pay_coin"`
		Metadata    string `json:"metadata"`
		Fulltext    string `json:"fulltext"`
		Rnd         string `json:"rnd"`
		TagImage    string `json:"tag_image"`
		EffectBlock int    `json:"effect_block"`
		Extra       struct {
			Wallet  interface{} `json:"wallet"`
			GiftBag struct {
				BagID   int `json:"bag_id"`
				GiftNum int `json:"gift_num"`
			} `json:"gift_bag"`
			TopList []interface{} `json:"top_list"`
			Follow  interface{}   `json:"follow"`
			Medal   interface{}   `json:"medal"`
			Title   interface{}   `json:"title"`
			Pk      struct {
				PkGiftTips string `json:"pk_gift_tips"`
			} `json:"pk"`
			Fulltext string `json:"fulltext"`
			Event    struct {
				EventScore     int `json:"event_score"`
				EventRedbagNum int `json:"event_redbag_num"`
			} `json:"event"`
			Capsule interface{} `json:"capsule"`
		} `json:"extra"`
		GiftEffect struct {
			Super            int           `json:"super"`
			SuperGiftNum     int           `json:"super_gift_num"`
			BroadcastMsgList []interface{} `json:"broadcast_msg_list"`
			SmallTvList      []interface{} `json:"small_tv_list"`
			BeatStorm        interface{}   `json:"beat_storm"`
		} `json:"gift_effect"`
	} `json:"data"`
}

// 送免费礼物
//
// POST https://api.live.bilibili.com/gift/v2/live/bag_send
func (b *BilibiliApiClient) BagSend(roomId int, receiverUid int, bagId int, giftId int, giftNum int) (*BagSendResponse, error) {
	v := url.Values{}
	v.Set("gift_id", strconv.Itoa(giftId))
	v.Set("ruid", strconv.Itoa(receiverUid))
	v.Set("gift_num", strconv.Itoa(giftNum))
	v.Set("bag_id", strconv.Itoa(bagId))
	v.Set("biz_id", strconv.Itoa(roomId))
	resp, err := b.Client.PostForm("https://api.live.bilibili.com/gift/v2/live/bag_send", v)
	if err != nil {
		return nil, err
	}
	response := &BagSendResponse{Code: -1}
	j := json.NewDecoder(resp.Body)
	err = j.Decode(response)
	if err != nil {
		return response, &ResponseJsonDecodeError{response.Message, err}
	}
	if response.Code != 0 {
		return response, &ResponseCodeNotZero{response.Message}
	}
	return response, nil
}

type DanmakuConfigResponse struct {
	Code    int    `json:"code"`
	Msg     string `json:"msg"`
	Message string `json:"message"`
	Data    struct {
		Color  int         `json:"color"`
		Mode   int         `json:"mode"`
		Length int         `json:"length"`
		RoomId interface{} `json:"roomid"`
	} `json:"data"`
}

// 获取在指定房间中可发送的弹幕长度、颜色、类型（滚动、顶部、底部）
//
// GET https://api.live.bilibili.com/userext/v1/danmuConf/getAll?roomid=3
func (b *BilibiliApiClient) GetDanmakuConfig(roomId int) (*DanmakuConfigResponse, error) {
	resp, err := b.Client.Get("https://api.live.bilibili.com/userext/v1/danmuConf/getAll?roomid=" + strconv.Itoa(roomId))
	if err != nil {
		return nil, err
	}
	response := &DanmakuConfigResponse{Code: -1}
	j := json.NewDecoder(resp.Body)
	err = j.Decode(response)
	if err != nil {
		return response, &ResponseJsonDecodeError{response.Message, err}
	}
	if response.Code != 0 {
		return response, &ResponseCodeNotZero{response.Message}
	}
	return response, nil
}

type ReceiveDailyBagResponse struct {
	Code    int    `json:"code"`
	Msg     string `json:"msg"`
	Message string `json:"message"`
	Data    struct {
		BagStatus       int `json:"bag_status"`
		BagExpireStatus int `json:"bag_expire_status"`
		BagToast        struct {
			ToastStatus  int    `json:"toast_status"`
			ToastMessage string `json:"toast_message"`
		} `json:"bag_toast"`
	} `json:"data"`
	BagList []interface{} `json:"bag_list"`
	Time    int           `json:"time"`
}

// 发送直播弹幕
//
// POST https://api.live.bilibili.com/msg/send
func (b *BilibiliApiClient) SendLiveMessage(roomId int, content string) (*SendLiveMessageResponse, error) {
	v := url.Values{}
	v.Set("color", "16777215")
	v.Set("fontsize", "25")
	v.Set("mode", "1")
	v.Set("msg", content)
	v.Set("rnd", Timestamp())
	v.Set("roomid", strconv.Itoa(roomId))
	biliJct, err := b.BiliJct()
	if err != nil {
		return nil, err
	}
	v.Set("csrf_token", biliJct)
	v.Set("csrf", biliJct)
	resp, err := b.Client.PostForm("https://api.live.bilibili.com/msg/send", v)
	if err != nil {
		return nil, err
	}
	response := &SendLiveMessageResponse{Code: -1}
	j := json.NewDecoder(resp.Body)
	err = j.Decode(response)
	if err != nil {
		return response, &ResponseJsonDecodeError{response.Message, err}
	}
	if response.Code != 0 {
		return response, &ResponseCodeNotZero{response.Message}
	}
	return response, nil
}

type SilverToCoinWebResponse struct {
	Code    int    `json:"code"`
	Msg     string `json:"msg"`
	Message string `json:"message"`
	Data    struct {
		Gold   string `json:"gold"`
		Silver string `json:"silver"`
		Tid    string `json:"tid"`
		Coin   int    `json:"coin"`
	} `json:"data"`
}

// 获取当天粉丝勋章赠送的辣条
//
// https://link.bilibili.com/p/help/index#/audience-fans-medal
//
// https://link.bilibili.com/p/help/index#/audience-level
func (b *BilibiliApiClient) ReceiveDailyBag() (*ReceiveDailyBagResponse, error) {
	payload := make(map[string]string)
	resp, err := b.Client.Get("https://api.live.bilibili.com/gift/v2/live/receive_daily_bag?" + b.SignPayload(payload).Encode())
	if err != nil {
		return nil, err
	}
	response := &ReceiveDailyBagResponse{Code: -1}
	j := json.NewDecoder(resp.Body)
	err = j.Decode(response)
	if err != nil {
		return response, &ResponseJsonDecodeError{response.Message, err}
	}
	if response.Code != 0 {
		return response, &ResponseCodeNotZero{response.Message}
	}
	return response, nil
}

type SendLiveMessageResponse struct {
	Code    int           `json:"code"`
	Msg     string        `json:"msg"`
	Message string        `json:"message"`
	Data    []interface{} `json:"data"`
}

// 用 700 银瓜子换 1 枚硬币
func (b *BilibiliApiClient) SilverToCoin(apiUrl string) (*SilverToCoinWebResponse, error) {
	resp, err := b.Client.Get(apiUrl)
	if err != nil {
		return nil, err
	}
	response := &SilverToCoinWebResponse{Code: -1}
	j := json.NewDecoder(resp.Body)
	err = j.Decode(response)
	if err != nil {
		return response, &ResponseJsonDecodeError{response.Message, err}
	}
	if response.Code != 0 {
		return response, &ResponseCodeNotZero{response.Message}
	}
	return response, nil
}

// 用 700 银瓜子换 1 枚硬币（电脑端）
// GET or POST https://api.live.bilibili.com/pay/v1/Exchange/silver2coin
func (b *BilibiliApiClient) SilverToCoinWeb() (*SilverToCoinWebResponse, error) {
	return b.SilverToCoin("https://api.live.bilibili.com/pay/v1/Exchange/silver2coin")
}

// 用 700 银瓜子换 1 枚硬币（移动端）
// GET or POST https://api.live.bilibili.com/AppExchange/silver2coin
func (b *BilibiliApiClient) SilverToCoinApp() (*SilverToCoinWebResponse, error) {
	return b.SilverToCoin("https://api.live.bilibili.com/AppExchange/silver2coin")
}

type TaskInfoResponse struct {
	Code    int    `json:"code"`
	Msg     string `json:"msg"`
	Message string `json:"message"`
	Data    struct {
		BoxInfo struct {
			FreeSilverFinish bool `json:"freeSilverFinish"`
			FreeSilverTimes  int  `json:"freeSilverTimes"`
			MaxTimes         int  `json:"max_times"`
			Minute           int  `json:"minute"`
			Silver           int  `json:"silver"`
			Status           int  `json:"status"`
			Times            int  `json:"times"`
			TimesMobile      int  `json:"times_mobile"`
			TimesWeb         int  `json:"times_web"`
			Type             int  `json:"type"`
		} `json:"box_info"`
		DoubleWatchInfo struct {
			Awards []struct {
				Name string `json:"name"`
				Num  int    `json:"num"`
				Type string `json:"type"`
			} `json:"awards"`
			MobileWatch int `json:"mobile_watch"`
			Progress    struct {
				Max int `json:"max"`
				Now int `json:"now"`
			} `json:"progress"`
			Status   int    `json:"status"`
			TaskID   string `json:"task_id"`
			WebWatch int    `json:"web_watch"`
		} `json:"double_watch_info"`
		LiveTimeInfo struct {
			Minute int  `json:"minute"`
			Status bool `json:"status"`
		} `json:"live_time_info"`
		LoginInfo struct {
			MobileLogin int `json:"mobile_login"`
			WebLogin    int `json:"web_login"`
		} `json:"login_info"`
		SignInfo struct {
			AllDays           int    `json:"allDays"`
			CurDate           string `json:"curDate"`
			CurDay            int    `json:"curDay"`
			CurMonth          int    `json:"curMonth"`
			CurYear           int    `json:"curYear"`
			HadSignDays       int    `json:"hadSignDays"`
			NewTask           int    `json:"newTask"`
			SignBonusDaysList []int  `json:"signBonusDaysList"`
			SignDaysList      []int  `json:"signDaysList"`
			SpecialText       string `json:"specialText"`
			Status            int    `json:"status"`
			Text              string `json:"text"`
		} `json:"sign_info"`
		WatchInfo struct {
			Awards []struct {
				Name string `json:"name"`
				Num  int    `json:"num"`
				Type string `json:"type"`
			} `json:"awards"`
			Progress struct {
				Max int `json:"max"`
				Now int `json:"now"`
			} `json:"progress"`
			Status int    `json:"status"`
			TaskID string `json:"task_id"`
		} `json:"watch_info"`
	} `json:"data"`
}

// 获取所有每日任务状态
//
// https://link.bilibili.com/p/center/index#/user-center/achievement/task
//
// GET https://api.live.bilibili.com/i/api/taskInfo
func (b *BilibiliApiClient) GetTaskInfo() (*TaskInfoResponse, error) {
	resp, err := b.Client.Get("https://api.live.bilibili.com/i/api/taskInfo")
	if err != nil {
		return nil, err
	}
	response := &TaskInfoResponse{Code: -1}
	j := json.NewDecoder(resp.Body)
	err = j.Decode(response)
	if err != nil {
		return response, &ResponseJsonDecodeError{response.Message, err}
	}
	if response.Code != 0 {
		return response, &ResponseCodeNotZero{response.Message}
	}
	return response, nil
}

type UserOnlineHeartbeatResponse struct {
	Code    int    `json:"code"`
	Msg     string `json:"msg"`
	Message string `json:"message"`
	Data    struct {
		GiftList []interface{} `json:"giftlist"`
	} `json:"data"`
}

// 双端直播间观看时长任务心跳
//
// 用户中心 > 任务成就 > 每日任务 > 双端观看直播
//
// https://link.bilibili.com/p/center/index#/user-center/achievement/task
//
// 每日使用移动端和网页端（双端）分别登录观看任意直播 5 分钟，即可获得以下奖励（需绑定手机）
//
// 奖励：
//
// 1. 700 银瓜子
//
// 2. 20 点粉丝勋章亲密度（发给当前佩戴勋章）
//
// 3. 1000 友爱金
//
// 备注：奖励请当天 24 点前领取，不然会失效哦
func (b *BilibiliApiClient) UserOnlineHeartbeat(postUrl string, roomId int) (*UserOnlineHeartbeatResponse, error) {
	v := url.Values{}
	v.Set("room_id", strconv.Itoa(roomId))
	resp, err := b.Client.PostForm(postUrl, v)
	if err != nil {
		return nil, err
	}
	response := &UserOnlineHeartbeatResponse{Code: -1}
	j := json.NewDecoder(resp.Body)
	err = j.Decode(response)
	if err != nil {
		return response, &ResponseJsonDecodeError{response.Message, err}
	}
	if response.Code != 0 {
		return response, &ResponseCodeNotZero{response.Message}
	}
	return response, nil
}

// 双端直播间观看时长任务心跳：网页端
//
// GET or POST http://api.live.bilibili.com/User/userOnlineHeart
func (b *BilibiliApiClient) UserOnlineHeartbeatWeb(roomId int) (*UserOnlineHeartbeatResponse, error) {
	return b.UserOnlineHeartbeat("https://api.live.bilibili.com/User/userOnlineHeart", roomId)
}

// 双端直播间观看时长任务心跳：移动端
//
// GET or POST https://api.live.bilibili.com/mobile/userOnlineHeart
func (b *BilibiliApiClient) UserOnlineHeartbeatMobile(roomId int) (*UserOnlineHeartbeatResponse, error) {
	return b.UserOnlineHeartbeat("https://api.live.bilibili.com/mobile/userOnlineHeart", roomId)
}

type SilverBoxGetCurrentTaskResponse struct {
	Code    int    `json:"code"`
	Msg     string `json:"msg"`
	Message string `json:"message"`
	Data    struct {
		Minute    int `json:"silver"`
		Silver    int `json:"minute"`
		TimeStart int `json:"time_start"`
		TimeEnd   int `json:"time_end"`
		Times     int `json:"times"`
		MaxTimes  int `json:"max_times"`
	} `json:"data"`
}

type SilverBoxTask struct {
	BilibiliApiClient *BilibiliApiClient
	Minute            int `json:"silver"`
	Silver            int `json:"minute"`
	TimeStart         int `json:"time_start"`
	TimeEnd           int `json:"time_end"`
	Times             int `json:"times"`
	MaxTimes          int `json:"max_times"`
}

// 获取当前银瓜子宝箱信息
//
// GET https://api.live.bilibili.com/lottery/v1/SilverBox/getCurrentTask
func (b *BilibiliApiClient) SilverBoxGetCurrentTask() (*SilverBoxTask, *SilverBoxGetCurrentTaskResponse, error) {
	resp, err := b.Client.Get("https://api.live.bilibili.com/lottery/v1/SilverBox/getCurrentTask")
	if err != nil {
		return nil, nil, err
	}
	response := &SilverBoxGetCurrentTaskResponse{Code: -1}
	j := json.NewDecoder(resp.Body)
	err = j.Decode(response)
	if err != nil {
		return nil, response, &ResponseJsonDecodeError{response.Message, err}
	}
	if response.Code != 0 {
		return nil, response, &ResponseCodeNotZero{response.Message}
	}
	return &SilverBoxTask{
		b,
		response.Data.Minute,
		response.Data.Silver,
		response.Data.TimeStart,
		response.Data.TimeEnd,
		response.Data.Times,
		response.Data.MaxTimes,
	}, response, nil
}

type SilverBoxFreeSilverAwardResponse struct {
	Code    int    `json:"code"`
	Msg     string `json:"msg"`
	Message string `json:"message"`
	Data    struct {
		Silver      string `json:"silver"`
		AwardSilver int    `json:"awardSilver"`
		IsEnd       int    `json:"isEnd"`
	} `json:"data"`
}

// 打开银瓜子宝箱
//
// GET or POST https://api.live.bilibili.com/mobile/freeSilverAward
func (s *SilverBoxTask) FreeSilverAward() (*SilverBoxFreeSilverAwardResponse, error) {
	v := url.Values{}
	v.Set("time_end", strconv.Itoa(s.TimeEnd))
	v.Set("time_start", strconv.Itoa(s.TimeStart))
	query := s.BilibiliApiClient.SignPayload2(v).Encode()
	resp, err := s.BilibiliApiClient.Client.Get("https://api.live.bilibili.com/mobile/freeSilverAward?" + query)
	if err != nil {
		return nil, err
	}
	response := &SilverBoxFreeSilverAwardResponse{Code: -1}
	j := json.NewDecoder(resp.Body)
	err = j.Decode(response)
	if err != nil {
		return response, &ResponseJsonDecodeError{response.Message, err}
	}
	if response.Code != 0 {
		return response, &ResponseCodeNotZero{response.Message}
	}
	return response, nil
}
