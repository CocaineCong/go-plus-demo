package main

import (
	fmt "fmt"
	strconv "strconv"
	json "encoding/json"
	ioutil "io/ioutil"
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
	NickName   string
	CreateTime string
	Content    string
	FromType   string
}

func main() {
//line D:\Desktop\vscodeProject\go_spider\main.gop:121
	client := &http.Client{}
//line D:\Desktop\vscodeProject\go_spider\main.gop:122
	reqSpider, err := http.NewRequest("GET", "https://blink-open-api.csdn.net/v1/pc/blink/allComment?pageNum=1&pageSize=1000&blinkId=1260435", nil)
//line D:\Desktop\vscodeProject\go_spider\main.gop:123
	if err != nil {
//line D:\Desktop\vscodeProject\go_spider\main.gop:124
		fmt.Println(err)
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
		fmt.Println(err)
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
		luckBlinkPerson.FromType = v.FromType
//line D:\Desktop\vscodeProject\go_spider\main.gop:141
		luckBlinkPerson.NickName = v.Nickname
//line D:\Desktop\vscodeProject\go_spider\main.gop:142
		luckBlinkPerson.CreateTime = v.CreateTime
//line D:\Desktop\vscodeProject\go_spider\main.gop:143
		luckBlinkPerson.Content = v.Content
//line D:\Desktop\vscodeProject\go_spider\main.gop:144
		luckyBPList = append(luckyBPList, luckBlinkPerson)
	}
//line D:\Desktop\vscodeProject\go_spider\main.gop:164
	luckyBPList = removeBlinkRepByMap(luckyBPList)
	dataHandler(luckyBPList)
}
func removeBlinkRepByMap(slc []LuckyBlinkPerson) []LuckyBlinkPerson {
//line D:\Desktop\vscodeProject\go_spider\main.gop:155
	var result []LuckyBlinkPerson
//line D:\Desktop\vscodeProject\go_spider\main.gop:156
	tempMap := map[LuckyBlinkPerson]byte{}
	for
//line D:\Desktop\vscodeProject\go_spider\main.gop:157
	_, e := range slc {
//line D:\Desktop\vscodeProject\go_spider\main.gop:158
		l := len(tempMap)
//line D:\Desktop\vscodeProject\go_spider\main.gop:159
		tempMap[e] = 0
//line D:\Desktop\vscodeProject\go_spider\main.gop:160
		if len(tempMap) != l {
//line D:\Desktop\vscodeProject\go_spider\main.gop:161
			result = append(result, e)
		}
	}
//line D:\Desktop\vscodeProject\go_spider\main.gop:164
	return result
}
func dataHandler(items []LuckyBlinkPerson) {
//line D:\Desktop\vscodeProject\go_spider\main.gop:178
	countAndroid, countIOS, countPC, countOrder, all := 0, 0, 0, 0, len(items)
//line D:\Desktop\vscodeProject\go_spider\main.gop:179
	meanDayIOS, meanDayAndroid, meanDayPC := 0, 0, 0
//line D:\Desktop\vscodeProject\go_spider\main.gop:180
	tmp := 0
//line D:\Desktop\vscodeProject\go_spider\main.gop:181
	var userAndroid []LuckyBlinkPerson
//line D:\Desktop\vscodeProject\go_spider\main.gop:182
	var userIOS []LuckyBlinkPerson
//line D:\Desktop\vscodeProject\go_spider\main.gop:183
	var userPC []LuckyBlinkPerson
	for
//line D:\Desktop\vscodeProject\go_spider\main.gop:185
	_, item := range items {
//line D:\Desktop\vscodeProject\go_spider\main.gop:186
		switch item.FromType {
//line D:\Desktop\vscodeProject\go_spider\main.gop:187
		case "CSDN-APP:Android:":
//line D:\Desktop\vscodeProject\go_spider\main.gop:188
			countAndroid++
//line D:\Desktop\vscodeProject\go_spider\main.gop:189
			tmp, _ = strconv.Atoi(item.CreateTime[:2])
//line D:\Desktop\vscodeProject\go_spider\main.gop:190
			userAndroid = append(userAndroid, item)
//line D:\Desktop\vscodeProject\go_spider\main.gop:191
			meanDayAndroid += tmp
//line D:\Desktop\vscodeProject\go_spider\main.gop:192
		case "CSDN-APP:iOS:":
//line D:\Desktop\vscodeProject\go_spider\main.gop:193
			countIOS++
//line D:\Desktop\vscodeProject\go_spider\main.gop:194
			userIOS = append(userIOS, item)
//line D:\Desktop\vscodeProject\go_spider\main.gop:195
			tmp, _ = strconv.Atoi(item.CreateTime[:2])
//line D:\Desktop\vscodeProject\go_spider\main.gop:196
			meanDayIOS += tmp
//line D:\Desktop\vscodeProject\go_spider\main.gop:197
		case "pc":
//line D:\Desktop\vscodeProject\go_spider\main.gop:198
			countPC++
//line D:\Desktop\vscodeProject\go_spider\main.gop:199
			userPC = append(userPC, item)
//line D:\Desktop\vscodeProject\go_spider\main.gop:200
			tmp, _ = strconv.Atoi(item.CreateTime[:2])
//line D:\Desktop\vscodeProject\go_spider\main.gop:201
			meanDayPC += tmp
//line D:\Desktop\vscodeProject\go_spider\main.gop:202
		default:
//line D:\Desktop\vscodeProject\go_spider\main.gop:203
			countOrder++
		}
	}
//line D:\Desktop\vscodeProject\go_spider\main.gop:206
	fmt.Printf("来自Android有:%d人 平均评论天:%d天前\n", countAndroid, meanDayAndroid/len(userAndroid))
//line D:\Desktop\vscodeProject\go_spider\main.gop:207
	fmt.Printf("来自iOS有:%d人 平均评论天:%d天前\n", countIOS, meanDayIOS/len(userIOS))
//line D:\Desktop\vscodeProject\go_spider\main.gop:208
	fmt.Printf("来自pc有:%d人 平均评论天:%d天前\n", countPC, meanDayPC/len(userPC))
//line D:\Desktop\vscodeProject\go_spider\main.gop:209
	fmt.Printf("其他的有:%d人\n", countOrder)
//line D:\Desktop\vscodeProject\go_spider\main.gop:210
	fmt.Printf("一共有%d人\n", all)
//line D:\Desktop\vscodeProject\go_spider\main.gop:221
	mapDay := map[string]int{}
//line D:\Desktop\vscodeProject\go_spider\main.gop:222
	meanDayIOS, meanDayAndroid, meanDayPC = 0, 0, 0
	for
//line D:\Desktop\vscodeProject\go_spider\main.gop:223
	_, user := range items {
//line D:\Desktop\vscodeProject\go_spider\main.gop:224
		mapDay[user.CreateTime]++
	}
	for
//line D:\Desktop\vscodeProject\go_spider\main.gop:226
	k, v := range mapDay {
//line D:\Desktop\vscodeProject\go_spider\main.gop:227
		fmt.Printf("%s 评论了 %d 条信息\n", k, v)
	}
}
