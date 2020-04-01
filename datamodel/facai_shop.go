package datamodel

import (
	"time"
)

// ProductType [...]
type ProductType struct {
	PtID         int       `gorm:"primary_key"` // 主键
	PtName       string    // 名称
	PtState      int       // 状态
	PtCreatetime time.Time // 创建时间
}

// Userinfo [...]
type Userinfo struct {
	UID         int       `gorm:"primary_key"` // 主键
	UNickname   string    // 昵称
	UPassword   string    // 密码
	UState      int       // 状态
	UType       int       // 类型
	UOpenid     string    // 三方登录唯一标识
	URegType    int       // 注册类型
	UCreatetime time.Time // 注册时间
	UPhone      string    // 手机号
	UEmail      string    // 邮箱
	ULoginTime  time.Time // 最近一次登录时间
	UAvatar     string    // 头像
	Token       string
}

// Address [...]
type Address struct {
	AID       int    `gorm:"primary_key"` // 地址id
	AUserID   int    // 所属用户id
	ACity     string // 所在城市
	AAddress  string // 详细地址
	ADefatult int    // 是否默认
	ALink     string // 联系人
	ATel      string // 联系电话
}

// Cart [...]
type Cart struct {
	CID     int     `gorm:"primary_key"` // 主键
	CPsid   int     // 规格id
	CCount  int     // 数量
	CUserid int     // 用户id
	CTotal  float64 // 总价
}

// Order [...]
type Order struct {
	OID             int       `gorm:"primary_key"` // 主键
	OSerialNum      string    //交易号
	OPayMethod      int       //支付方式(0在线)
	OCreatetime     time.Time //创建时间
	OState          int       //交易状态(-1订单取消、-2支付失败、0待支付、1、备货中、2、已发货、3订单已完成)
	OType           int       //交易类型(1微信)
	OTransactionNum string    //第三方交易号
	OTotalAmount    float64   //总金额
	OUserid         int       //所属用户id
	ORemarks        string    //订单备注
	OAddress        string    //地址
	ODeliverytime   time.Time //送货时间
	OConfirmtime    time.Time //确认收货时间
	OLink           string    //收货人
	OLinkTel        string    //联系电话
	OrderDetails    []OrderDetail
}

// OrderDetail [...]
type OrderDetail struct {
	OdID    int     `gorm:"primary_key"` // 主键
	OdOid   int     // 订单id
	OdPsid  int     // 规格id
	OdCount int     // 数量
	OdTotal float64 // 总价
}

// Product [...]
type Product struct {
	PID         int       `gorm:"primary_key"` // 主键ID
	PTitle      string    // 名称
	PState      int       // 状态
	PTypeid     int       // 分类id
	PCreatetime time.Time // 创建时间
	PDetail     string    // 介绍
	PBrief      string    // 简介
	PBrand      string    // 品牌
	PSource     string    // 货源
	POrigin     string    // 产地
	PCover      string    // 封面图片
	PImgArray   string    // 轮播图片
}

// ProductSpecs [...]
type ProductSpecs struct {
	SID        int     `gorm:"primary_key"` // 主键
	SName      string  // 规格名称
	SState     int     // 状态
	SPrice     float64 // 价格
	SBrief     string  // 简介
	SStock     int     // 库存
	SProductid int     // 所属产品id
}
