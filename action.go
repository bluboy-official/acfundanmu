package acfundanmu

import (
	"fmt"
	"strconv"

	"github.com/valyala/fasthttp"
	"github.com/valyala/fastjson"
)

// 房管踢人
func (t *token) managerKick(kickedUID int64) (e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("managerKick() error: %w", err)
		}
	}()

	if len(t.cookies) == 0 {
		panic(fmt.Errorf("房管踢人需要登陆AcFun帐号"))
	}

	form := fasthttp.AcquireArgs()
	defer fasthttp.ReleaseArgs(form)
	form.Set("visitorId", strconv.FormatInt(t.userID, 10))
	form.Set("liveId", t.liveID)
	form.Set("kickedUserId", strconv.FormatInt(kickedUID, 10))
	resp, err := t.fetchKuaiShouAPI(managerKickURL, form)
	checkErr(err)
	defer fasthttp.ReleaseResponse(resp)
	body := resp.Body()

	var p fastjson.Parser
	v, err := p.ParseBytes(body)
	checkErr(err)
	if v.GetInt("result") != 1 || v.GetBool("data", "kickSucc") != true {
		panic(fmt.Errorf("房管踢人失败，响应为 %s", string(body)))
	}

	return nil
}

// 主播踢人
func (t *token) authorKick(kickedUID int64) (e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("authorKick() error: %w", err)
		}
	}()

	if len(t.cookies) == 0 {
		panic(fmt.Errorf("主播踢人需要登陆AcFun帐号"))
	}

	form := fasthttp.AcquireArgs()
	defer fasthttp.ReleaseArgs(form)
	form.Set("visitorId", strconv.FormatInt(t.userID, 10))
	form.Set("liveId", t.liveID)
	form.Set("kickedUserId", strconv.FormatInt(kickedUID, 10))
	resp, err := t.fetchKuaiShouAPI(authorKickURL, form)
	checkErr(err)
	defer fasthttp.ReleaseResponse(resp)
	body := resp.Body()

	var p fastjson.Parser
	v, err := p.ParseBytes(body)
	checkErr(err)
	if v.GetInt("result") != 1 || v.GetBool("data", "kickSucc") != true {
		panic(fmt.Errorf("主播踢人失败，响应为 %s", string(body)))
	}

	return nil
}

// 主播添加房管
func (t *token) addManager(managerUID int64) (e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("addManager() error: %w", err)
		}
	}()

	if len(t.cookies) == 0 {
		panic(fmt.Errorf("主播添加房管需要登陆AcFun帐号"))
	}

	form := fasthttp.AcquireArgs()
	defer fasthttp.ReleaseArgs(form)
	form.Set("visitorId", strconv.FormatInt(t.userID, 10))
	form.Set("managerUserId", strconv.FormatInt(managerUID, 10))
	resp, err := t.fetchKuaiShouAPI(addManagerURL, form)
	checkErr(err)
	defer fasthttp.ReleaseResponse(resp)
	body := resp.Body()

	var p fastjson.Parser
	v, err := p.ParseBytes(body)
	checkErr(err)
	if v.GetInt("result") != 1 {
		panic(fmt.Errorf("主播添加房管失败，响应为 %s", string(body)))
	}

	return nil
}

// 主播删除房管
func (t *token) deleteManager(managerUID int64) (e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("deleteManager() error: %w", err)
		}
	}()

	if len(t.cookies) == 0 {
		panic(fmt.Errorf("主播删除房管需要登陆AcFun帐号"))
	}

	form := fasthttp.AcquireArgs()
	defer fasthttp.ReleaseArgs(form)
	form.Set("visitorId", strconv.FormatInt(t.userID, 10))
	form.Set("managerUserId", strconv.FormatInt(managerUID, 10))
	resp, err := t.fetchKuaiShouAPI(deleteManagerURL, form)
	checkErr(err)
	defer fasthttp.ReleaseResponse(resp)
	body := resp.Body()

	var p fastjson.Parser
	v, err := p.ParseBytes(body)
	checkErr(err)
	if v.GetInt("result") != 1 {
		panic(fmt.Errorf("主播删除房管失败，响应为 %s", string(body)))
	}

	return nil
}

// ManagerKick 房管踢人，需要已经登陆AcFun帐号，不需要调用StartDanmu()
func (ac *AcFunLive) ManagerKick(kickedUID int64) error {
	return ac.t.managerKick(kickedUID)
}

// AuthorKick 主播踢人，需要已经登陆AcFun帐号，不需要调用StartDanmu()
func (ac *AcFunLive) AuthorKick(kickedUID int64) error {
	return ac.t.authorKick(kickedUID)
}

// AddManager 主播添加房管，需要已经登陆AcFun帐号，可以调用Init(0, cookies)，不需要调用StartDanmu()
func (ac *AcFunLive) AddManager(managerUID int64) error {
	return ac.t.addManager(managerUID)
}

// DeleteManager 主播删除房管，需要已经登陆AcFun帐号，可以调用Init(0, cookies)，不需要调用StartDanmu()
func (ac *AcFunLive) DeleteManager(managerUID int64) error {
	return ac.t.deleteManager(managerUID)
}