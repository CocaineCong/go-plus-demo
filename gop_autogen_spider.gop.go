package main

import fmt "fmt"

func main() {
//line D:\Desktop\vscodeProject\go_spider\spider.gop:4
	arr := []int{1, 2, 3, 4, 5, 6}
//line D:\Desktop\vscodeProject\go_spider\spider.gop:5
	e := func() (_gop_ret [][]int) {
		for
//line D:\Desktop\vscodeProject\go_spider\spider.gop:5
		_, b := range arr {
//line D:\Desktop\vscodeProject\go_spider\spider.gop:5
			if b > 2 {
				for
//line D:\Desktop\vscodeProject\go_spider\spider.gop:5
				_, a := range arr {
//line D:\Desktop\vscodeProject\go_spider\spider.gop:5
					if a < b {
//line D:\Desktop\vscodeProject\go_spider\spider.gop:5
						_gop_ret = append(_gop_ret, []int{a, b})
					}
				}
			}
		}
//line D:\Desktop\vscodeProject\go_spider\spider.gop:5
		return
	}()
//line D:\Desktop\vscodeProject\go_spider\spider.gop:6
	fmt.Println(e)
}
