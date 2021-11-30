package main

import (
	fmt "fmt"
	json "encoding/json"
	ioutil "io/ioutil"
	log "log"
	http "net/http"
)

type BlinkResult struct {
	Code int64 "json:\"code\""
	Data []struct {
		Anonymous bool   "json:\"anonymous\""
		Avatar    string "json:\"avatar\""
		BizNo     string "json:\"bizNo\""
		Child     []struct {
			Anonymous  bool   "json:\"anonymous\""
			Avatar     string "json:\"avatar\""
			BizNo      string "json:\"bizNo\""
			ChildCount int64  "json:\"childCount\""
			Content    string "json:\"content\""
			CreateTime string "json:\"createTime\""
			Employee   struct {
				Bala            string "json:\"bala\""
				IsCertification bool   "json:\"isCertification\""
				Org             string "json:\"org\""
			} "json:\"employee\""
			Ext        string "json:\"ext\""
			FlowerName struct {
				FlowerName string "json:\"flowerName\""
				Level      string "json:\"level\""
			} "json:\"flowerName\""
			ID              int64  "json:\"id\""
			Level           int64  "json:\"level\""
			LikeCount       int64  "json:\"likeCount\""
			Nickname        string "json:\"nickname\""
			ParentID        int64  "json:\"parentId\""
			Platform        string "json:\"platform\""
			ReplyFlowerName struct {
				FlowerName string "json:\"flowerName\""
				Level      string "json:\"level\""
			} "json:\"replyFlowerName\""
			ReplyNickname string "json:\"replyNickname\""
			ReplyUsername string "json:\"replyUsername\""
			ResourceGroup string "json:\"resourceGroup\""
			ResourceID    string "json:\"resourceId\""
			ResourceOrder string "json:\"resourceOrder\""
			ResourceUser  string "json:\"resourceUser\""
			Score         int64  "json:\"score\""
			Status        int64  "json:\"status\""
			Student       struct {
				Bala            string "json:\"bala\""
				IsCertification bool   "json:\"isCertification\""
				Org             string "json:\"org\""
			} "json:\"student\""
			Title struct {
				ID       int64  "json:\"id\""
				TitleURL string "json:\"titleUrl\""
				Used     bool   "json:\"used\""
				Username string "json:\"username\""
			} "json:\"title\""
			Top      bool   "json:\"top\""
			UserLike bool   "json:\"userLike\""
			Username string "json:\"username\""
		} "json:\"child\""
		ChildCount int64  "json:\"childCount\""
		Content    string "json:\"content\""
		CreateTime string "json:\"createTime\""
		Employee   struct {
			Bala            string "json:\"bala\""
			IsCertification bool   "json:\"isCertification\""
			Org             string "json:\"org\""
		} "json:\"employee\""
		Ext        string "json:\"ext\""
		FlowerName struct {
			FlowerName string "json:\"flowerName\""
			Level      string "json:\"level\""
		} "json:\"flowerName\""
		FromType      string "json:\"fromType\""
		ID            int64  "json:\"id\""
		Level         int64  "json:\"level\""
		LikeCount     int64  "json:\"likeCount\""
		Nickname      string "json:\"nickname\""
		ParentID      int64  "json:\"parentId\""
		Platform      string "json:\"platform\""
		ReplyUsername string "json:\"replyUsername\""
		ResourceGroup string "json:\"resourceGroup\""
		ResourceID    string "json:\"resourceId\""
		ResourceOrder string "json:\"resourceOrder\""
		ResourceUser  string "json:\"resourceUser\""
		Score         int64  "json:\"score\""
		Status        int64  "json:\"status\""
		Student       struct {
			Bala            string "json:\"bala\""
			IsCertification bool   "json:\"isCertification\""
			Org             string "json:\"org\""
		} "json:\"student\""
		Title struct {
			ID       int64  "json:\"id\""
			TitleURL string "json:\"titleUrl\""
			Used     bool   "json:\"used\""
			Username string "json:\"username\""
		} "json:\"title\""
		Top      bool   "json:\"top\""
		UserLike bool   "json:\"userLike\""
		Username string "json:\"username\""
	} "json:\"data\""
	Msg string "json:\"msg\""
}
type LuckyBlinkPerson struct {
	UserName   string
	NickName   string
	CreateTime string
	Content    string
}

func main() {
//line D:\Desktop\vscodeProject\go_spider\main.gop:121
	client := &http.Client{}
//line D:\Desktop\vscodeProject\go_spider\main.gop:122
	reqSpider, err := http.NewRequest("GET", "https://blink-open-api.csdn.net/v1/pc/blink/allComment?pageNum=1&pageSize=1000&blinkId=1260435", nil)
//line D:\Desktop\vscodeProject\go_spider\main.gop:123
	if err != nil {
//line D:\Desktop\vscodeProject\go_spider\main.gop:124
		log.Fatal(err)
	}
//line D:\Desktop\vscodeProject\go_spider\main.gop:126
	reqSpider.Header.Set("content-length", "0")
//line D:\Desktop\vscodeProject\go_spider\main.gop:127
	reqSpider.Header.Set("accept", "*/*")
//line D:\Desktop\vscodeProject\go_spider\main.gop:128
	reqSpider.Header.Set("x-requested-with", "XMLHttpRequest")
//line D:\Desktop\vscodeProject\go_spider\main.gop:129
	respSpider, err := client.Do(reqSpider)
//line D:\Desktop\vscodeProject\go_spider\main.gop:130
	if err != nil {
//line D:\Desktop\vscodeProject\go_spider\main.gop:131
		log.Fatal(err)
	}
//line D:\Desktop\vscodeProject\go_spider\main.gop:133
	bodyText, _ := ioutil.ReadAll(respSpider.Body)
//line D:\Desktop\vscodeProject\go_spider\main.gop:134
	var result BlinkResult
//line D:\Desktop\vscodeProject\go_spider\main.gop:135
	_ = json.Unmarshal(bodyText, &result)
//line D:\Desktop\vscodeProject\go_spider\main.gop:136
	commentList := result.Data
//line D:\Desktop\vscodeProject\go_spider\main.gop:137
	var luckyBPList []LuckyBlinkPerson
	for
//line D:\Desktop\vscodeProject\go_spider\main.gop:138
	_, v := range commentList {
//line D:\Desktop\vscodeProject\go_spider\main.gop:139
		var luckBlinkPerson LuckyBlinkPerson
//line D:\Desktop\vscodeProject\go_spider\main.gop:140
		luckBlinkPerson.UserName = v.Username
//line D:\Desktop\vscodeProject\go_spider\main.gop:141
		luckBlinkPerson.NickName = v.Nickname
//line D:\Desktop\vscodeProject\go_spider\main.gop:142
		luckBlinkPerson.CreateTime = v.CreateTime
//line D:\Desktop\vscodeProject\go_spider\main.gop:143
		luckBlinkPerson.Content = v.Content
//line D:\Desktop\vscodeProject\go_spider\main.gop:144
		luckyBPList = append(luckyBPList, luckBlinkPerson)
	}
//line D:\Desktop\vscodeProject\go_spider\main.gop:162
	luckyBPList = removeBlinkRepByMap(luckyBPList)
	for
//line D:\Desktop\vscodeProject\go_spider\main.gop:147
	_, v := range luckyBPList {
//line D:\Desktop\vscodeProject\go_spider\main.gop:148
		fmt.Println(v.NickName, v.Content, v.CreateTime)
	}
}
func removeBlinkRepByMap(slc []LuckyBlinkPerson) []LuckyBlinkPerson {
//line D:\Desktop\vscodeProject\go_spider\main.gop:153
	var result []LuckyBlinkPerson
//line D:\Desktop\vscodeProject\go_spider\main.gop:154
	tempMap := map[LuckyBlinkPerson]byte{}
	for
//line D:\Desktop\vscodeProject\go_spider\main.gop:155
	_, e := range slc {
//line D:\Desktop\vscodeProject\go_spider\main.gop:156
		l := len(tempMap)
//line D:\Desktop\vscodeProject\go_spider\main.gop:157
		tempMap[e] = 0
//line D:\Desktop\vscodeProject\go_spider\main.gop:158
		if len(tempMap) != l {
//line D:\Desktop\vscodeProject\go_spider\main.gop:159
			result = append(result, e)
		}
	}
//line D:\Desktop\vscodeProject\go_spider\main.gop:162
	return result
}
