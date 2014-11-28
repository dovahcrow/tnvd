package models

import (
	"time"
)

type Patch struct {
	Id       int64     //id
	LeakId   int       // 漏洞id
	Reporter string    // 报告者
	Link     string    //补丁链接
	Desc     string    `xorm:"text"` //补丁描述
	File     []string  //附件
	Opinion  string    `xorm:"string"`  //审核意见
	Time     time.Time `xorm:"created"` //时间
}
