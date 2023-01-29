package util

import (
	"context"
	"douyin/setting"
	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
	"mime/multipart"
)

func UpLoadFile(file multipart.File, fileSize int64, key string) (string, error) {
	putPolicy := storage.PutPolicy{
		Scope: setting.Conf.OSSQiNiuConfig.Bucket,
	}
	mac := qbox.NewMac(setting.Conf.OSSQiNiuConfig.AccessKey, setting.Conf.OSSQiNiuConfig.SecretKey)
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{
		Zone:          &storage.ZoneHuanan,
		UseCdnDomains: false,
		UseHTTPS:      false,
	}

	putExtra := storage.PutExtra{}

	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{
		Key: key,
	}

	err := formUploader.PutFile(context.Background(), &ret, upToken, "", "", &putExtra)
	if err != nil {
		return "", err
	}
	url := setting.Conf.OSSQiNiuConfig.QiNiuServer + ret.Key
	return url, err
}
