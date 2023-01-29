package services

import (
	"bytes"
	"douyin/model"
	"douyin/repositories"
	"douyin/setting"
	"douyin/util"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gin-gonic/gin"
	"github.com/u2takey/ffmpeg-go"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

var VideoService = newVideoService()

func newVideoService() *videoService {
	return &videoService{}
}

type videoService struct {
}

func (s *videoService) GetUserVideos(userId int64) (videos []model.Video) {
	videos = repositories.VideoRepository.Find(Db, userId)
	return
}

func (s *videoService) Publish(ctx *gin.Context, userId int64) error {
	title := ctx.PostForm("title")
	file, err := ctx.FormFile("data")
	if err != nil {
		return err
	}

	folderName := "video/" + strconv.Itoa(int(userId)) + "/" + time.Now().Format("2006-01-02-15-04-05")
	fileName := filepath.Base(file.Filename)

	// 使用阿里云
	videoUrl, err := pushFile(folderName + "/" + fileName, file, util.Buc)

	// 使用七牛云
	//c, _ := file.Open()
	//videoUrl, err := util.UpLoadFile(c, file.Size, folderName + "/" + fileName)

	if err != nil {
		log.Println(err)
		return err
	}

	// 使用阿里云
	coverUrl := VideoService.UploadCoverALi(videoUrl, folderName, 1, util.Buc)

	// 使用七牛云
	//bufCover.

	video := &model.Video{
		Model:      model.Model{
			Id: util.GetOnlyId(),
		},
		Title:      title,
		AuthorId:   userId,
		// must use this timestamp
		CreateTime: time.Now().Format("2006-1-2 15:04:05"),
		PlayUrl:    videoUrl,
		CoverUrl:   coverUrl,
	}
	err = repositories.VideoRepository.Create(Db, video)
	return err
}

func pushFile(filename string, data *multipart.FileHeader, bucket *oss.Bucket) (string, error) {
	file, _ := data.Open()

	defer file.Close()
	err := bucket.PutObject(filename, file)

	return setting.Conf.OSSAliConfig.SufferUrl + filename, err
}

func (s *videoService) GetLikes(id int64) []model.Video {
	ids := repositories.UserRepository.GetLikes(Db, id)
	videos := repositories.UserRepository.GetVideos(Db, ids)
	return videos
}

func (s *videoService) GetVideos(t string) (videos []model.Video) {
	videos = repositories.VideoRepository.FindByTime(Db, t)
	return
}

func (s *videoService) UploadCoverALi(videoPath, folderName string, frameNum int, bucket *oss.Bucket) string {
	buf := bytes.NewBuffer(nil)
	err := ffmpeg_go.Input(videoPath).
		Filter("select", ffmpeg_go.Args{fmt.Sprintf("gte(n, %d)", frameNum)}).
		Output("pipe:", ffmpeg_go.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()
	if err != nil {
		log.Println("generate cover failed：", err)
	}
	b := buf.Bytes()

	if err := bucket.PutObject(folderName + "/cover.jpeg", bytes.NewReader(b)); err != nil {
		log.Println("generate cover failed：", err)
	}

	return setting.Conf.OSSAliConfig.SufferUrl + folderName + "/cover.jpeg"
}

func (s *videoService) UploadCoverQiNiu(folderName string, bucket *oss.Bucket, buf *bytes.Buffer) string {
	var b []byte
	buf.Write(b)

	reader := bytes.NewReader(b)

	bc := make([]byte, len(b))
	if _, err := reader.Read(bc); err != nil {
		log.Fatal(err)
	}

	// 未更新上传七牛云
	if err := bucket.PutObject(folderName + "/cover.jpeg", bytes.NewReader(b)); err != nil {
		log.Println("generate cover failed：", err)
	}
	return setting.Conf.OSSAliConfig.SufferUrl + folderName + "/cover.jpeg"
}

func (s *videoService) GetVideoById(vid int64) model.Video {
	return repositories.VideoRepository.GetVideoById(Db, vid)
}