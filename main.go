// By XiaoZhiSoC
// 2023.07.17
// Happy middle school graduation!

package main

import (
	"fmt"
	"github.com/valyala/fastjson"
	_ "github.com/valyala/fastjson"
	"io"
	"net/http"
	"time"
)

func main() {
	req, err := http.NewRequest("GET", "https://www.nnzkzs.com/api/services/app/student/GetMatriculateInfo", nil)
	if err != nil {
		fmt.Println("Request Error.")
		fmt.Println(err)
		return
	}

	set_headers(req)

	client := http.Client{}
	order_old := 0
	for {
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Do Error.")
			fmt.Println(err)
			return
		}
		bytes, _ := io.ReadAll(resp.Body)
		value, err := fastjson.Parse(string(bytes))
		name := string(value.GetStringBytes("result", "signupSchoolName"))
		code := string(value.GetStringBytes("result", "signupSchoolCode"))
		order := value.GetInt("result", "calculateResult", "order")
		guide_reason := string(value.GetStringBytes("result", "calculateResult", "guideReason"))
		inst_reason := string(value.GetStringBytes("result", "calculateResult", "instReason"))
		if order != 0 && order_old != order {
			order_old = order
			fmt.Printf("学校: %s, 排名: %d, 指导: %s, 非定向: %s\n", name, order, guide_reason, inst_reason)

			req2, err := http.NewRequest("GET", "https://www.nnzkzs.com/api/services/app/student/GetStat?highSchoolCodes="+code, nil)
			if err != nil {
				fmt.Println("Request Error.")
				fmt.Println(err)
				return
			}
			set_headers(req2)
			resp2, err := client.Do(req2)
			if err != nil {
				fmt.Println("Do Error.")
				fmt.Println(err)
				return
			}
			bytes, _ := io.ReadAll(resp2.Body)
			value2, _ := fastjson.Parse(string(bytes))
			alter_plan := value2.GetInt("result", "alterPlan")
			dir_plan := value2.GetInt("result", "dirPlan")
			inst_plan := value2.GetInt("result", "instPlan")
			guide_plan := value2.GetInt("result", "guidePlan")
			plan := alter_plan + dir_plan + inst_plan + guide_plan
			fmt.Printf("总计划: %d, 差距: %d\n\n", plan, order-plan)
			time.Sleep(time.Duration(1) * time.Second)
		} else {
			time.Sleep(time.Duration(15) * time.Second)
		}
	}
}

func set_headers(req *http.Request) {

	// After Login, Replace This !
	req.Header.Set("Cookie", "")
	req.Header.Set("X-Xsrf-Token", "")
	
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/116.0.0.0 Safari/537.36 Edg/116.0.0.0")
}
