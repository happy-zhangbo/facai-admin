package datamodel

import (
	"time"
)

// Cart [...]
type Cart struct {
	CID     int     `gorm:"primary_key;column:c_id;type:int(10);not null"` // 主键
	CPsid   int     `gorm:"column:c_psid;type:int(10)"`                    // 规格id
	CCount  int     `gorm:"column:c_count;type:int(10)"`                   // 数量
	CUserid int     `gorm:"column:c_userid;type:int(10)"`                  // 用户id
	CTotal  float64 `gorm:"column:c_total;type:decimal(10,2)"`             // 总价
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
	OTotalAmount    float64   `gorm:"column:o_totalAmount;type:decimal(10,2)"` //总金额
	OUserid         int       //所属用户id
	ORemarks        string    //订单备注
	OAddress        string    //地址
	ODeliverytime   time.Time //送货时间
	OConfirmtime    time.Time //确认收货时间
	OLink           string    //收货人
	OLinkTel        string    //联系电话

	OrderDetail []OrderDetail `gorm:"ForeignKey:OdOid;AssociationForeignKey:OID"`
}

// OrderDetail [...]
type OrderDetail struct {
	OdID         int            `gorm:"primary_key;column:od_id;type:int(10);not null"` // 主键
	OdOid        int            `gorm:"column:od_oid;type:int(10)"`                     // 订单id
	OdPsid       int            `gorm:"column:od_psid;type:int(10)"`                    // 规格id
	OdCount      int            `gorm:"column:od_count;type:int(10)"`                   // 数量
	OdTotal      float64        `gorm:"column:od_total;type:decimal(10,2)"`             // 总价
	ProductSpecs []ProductSpecs `gorm:"ForeignKey:OdPsid;AssociationForeignKey:SID"`
}

// Product [...]
type Product struct {
	PID          int       `gorm:"primary_key;column:p_id;type:int(10);not null"` // 主键ID
	PTitle       string    `gorm:"column:p_title;type:varchar(155)"`              // 名称
	PState       int       `gorm:"column:p_state;type:int(1)"`                    // 状态
	PTypeid      int       `gorm:"column:p_typeid;type:int(2)"`                   // 分类id
	PCreatetime  time.Time `gorm:"column:p_createtime;type:datetime"`             // 创建时间
	PDetail      string    `gorm:"column:p_detail;type:varchar(3000)"`            // 介绍
	PBrief       string    `gorm:"column:p_brief;type:varchar(355)"`              // 简介
	PBrand       string    `gorm:"column:p_brand;type:varchar(20)"`               // 品牌
	PSource      string    `gorm:"column:p_source;type:varchar(20)"`              // 货源
	POrigin      string    `gorm:"column:p_origin;type:varchar(20)"`              // 产地
	PCover       string    `gorm:"column:p_cover;type:varchar(355)"`              // 封面图片
	PImgArray    string    `gorm:"column:p_imgArray;type:varchar(655)"`           // 轮播图片
	ProductSpecs []ProductSpecs
}

// ProductSpecs [...]
type ProductSpecs struct {
	SID        int     `gorm:"primary_key;column:s_id;type:int(10);not null"` // 主键
	SName      string  `gorm:"column:s_name;type:varchar(155)"`               // 规格名称
	SState     int     `gorm:"column:s_state;type:int(1)"`                    // 状态
	SPrice     float64 `gorm:"column:s_price;type:decimal(10,2)"`             // 价格
	SBrief     string  `gorm:"column:s_brief;type:varchar(355)"`              // 简介
	SStock     int     `gorm:"column:s_stock;type:int(10)"`                   // 库存
	SProductid int     `gorm:"column:s_productid;type:int(10)"`               // 所属产品id
	Product    Product
}

// ProductType [...]
type ProductType struct {
	PtID         int       `gorm:"primary_key;column:pt_id;type:int(10);not null"` // 主键
	PtName       string    `gorm:"column:pt_name;type:varchar(5)"`                 // 名称
	PtState      int       `gorm:"column:pt_state;type:int(1)"`                    // 状态
	PtCreatetime time.Time `gorm:"column:pt_createtime;type:datetime"`             // 创建时间
}

// Userinfo [...]
type Userinfo struct {
	UID         int       `gorm:"primary_key;column:u_id;type:int(10);not null"` // 主键
	UNickname   string    `gorm:"column:u_nickname;type:varchar(50)"`            // 昵称
	UPassword   string    `gorm:"column:u_password;type:varchar(30)"`            // 密码
	UState      int       `gorm:"column:u_state;type:int(1)"`                    // 状态
	UType       int       `gorm:"column:u_type;type:int(2)"`                     // 类型
	UOpenid     string    `gorm:"column:u_openid;type:varchar(155)"`             // 三方登录唯一标识
	URegType    int       `gorm:"column:u_regType;type:int(1)"`                  // 注册类型
	UCreatetime time.Time `gorm:"column:u_createtime;type:datetime"`             // 注册时间
	UPhone      string    `gorm:"column:u_phone;type:varchar(20)"`               // 手机号
	UEmail      string    `gorm:"column:u_email;type:varchar(155)"`              // 邮箱
	ULoginTime  time.Time `gorm:"column:u_loginTime;type:datetime"`              // 最近一次登录时间
	UAvatar     string    `gorm:"column:u_avatar;type:varchar(255)"`             // 头像
	Token       string
}

// Address [...]
type Address struct {
	AID       int    `gorm:"primary_key;column:a_id;type:int(10);not null"` // 地址id
	AUserID   int    `gorm:"column:a_userId;type:int(10)"`                  // 所属用户id
	ACity     string `gorm:"column:a_city;type:varchar(155)"`               // 所在城市
	AAddress  string `gorm:"column:a_address;type:varchar(255)"`            // 详细地址
	ADefatult int    `gorm:"column:a_defatult;type:int(1)"`                 // 是否默认
	ALink     string `gorm:"column:a_link;type:varchar(25)"`                // 联系人
	ATel      string `gorm:"column:a_tel;type:varchar(25)"`                 // 联系电话
}
