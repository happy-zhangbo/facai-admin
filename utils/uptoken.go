package utils

import (
	"github.com/qiniu/api.v7/v7/auth"
	"github.com/qiniu/api.v7/v7/storage"
)

var (
	accessKey = "bbmhQ8GwYxN8Glb8cZ3l8aqXQUWyBKiITx0hxwQj"
	secretKey = "_0xLvao6ztOUvf5l_d9UTz3GLaseGKR3321E9bFh"
	bucket    = "facai-zhangbo"
)

func CreateUptoken() string {
	// 简单上传凭证
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := auth.New(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)
	return upToken
}
