# Bilibili API

Inspired by <https://github.com/lovelyyoshino/Bilibili-Live-API> and <https://github.com/lkeme/BiliHelper>

* `api.go` 无需登录的 API
* `auth.go` 登录相关的 API
* `client.go` 登录之后的 API

Don't use this repo as a package, it's still unstable now.
Methods may be changed at any time!

## Components

* GiftConfig
* RoomGiftList
* GetBagList
* BagSend
* GetDanmakuConfig
* SendLiveMessage
* ReceiveDailyBag
* SilverToCoinWeb
* SilverToCoinApp
* GetTaskInfo
* UserOnlineHeartbeatWeb
* UserOnlineHeartbeatMobile
* FreeSilverAward
* DailySignWeb
* DailySignApp
* GetSignInfoWeb
* ReceiveAward

| File                  | Description                |
|-----------------------|----------------------------|
| auth                  | 帐号登录组件               |
| receive_daily_bag     | 每日登录背包奖励辣条       |
| user_online_heartbeat | 双端直播间观看时长任务心跳 |
| silver_box            | 自动领宝箱                 |
| silver_to_coin        | 银瓜子换硬币               |
| Task                | 每日任务              |
| GiftSend            | 自动清空过期礼物      |

TODO

| File                | Description          |
|---------------------|-----------------------|
| MaterialObject      | 实物抽奖              |
| GroupSignIn         | 应援团签到            |
| Storm               | 节奏风暴              |
| RaffleHandler       | 小电视飞船            |
| RaffleHandler       | 摩天大樓              |
| RaffleHandler       | 小金人                |
| MasterSite          | 主站(观看、分享、投币)|
| Guard               | 舰长上船亲密度        |

