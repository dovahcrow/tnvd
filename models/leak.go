package models

import (
	"time"
)

type Leak struct {
	Id               int       `xorm:"pk autoincr"` //id
	TnvdId           string    `xorm:"pk notnull"`  //tnvd id
	BugtraqId        string    `xorm:""`            //bugtraq id
	Name             string    //漏洞名字
	LeakType         string    //漏洞类型
	DangerLevel      int       `xorm:"not null"` //危害等级
	InfluenceProduct string    `xorm:""`         //受影响的产品
	Desc             string    `xorm:"text"`     //漏洞描述
	Solution         string    `xorm:"text"`     //解决方案
	PatchInfo        string    `xorm:"text"`     //补丁信息
	Finder           string    //发现者
	Reporter         string    //报告者
	ReleaseTime      time.Time `xorm:"created"`         //漏洞公布时间
	ReportTime       time.Time `xorm:"created"`         //漏洞上报时间
	RecordTime       time.Time `xorm:"created"`         //漏洞收录时间
	UpdateTime       time.Time `xorm:"created updated"` //漏洞信息更新时间
	Attention        int       //关注
	Files            []string  `xorm:""` //附件
}

type leakCollection struct{}

var LeakCollection leakCollection

func (leakCollection) GetLeakNums() (int, error) {
	return 42, nil
}
