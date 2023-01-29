package services

import (
	"douyin/model"
	"douyin/repositories"
	"errors"
	"gorm.io/gorm"
	"log"
)

var CommentService = newCommentService()

func newCommentService() *commentService {
	return &commentService{}
}

type commentService struct {
}

func (s *commentService) CommentVideo(uid int64, vid int64, actionType int, cid int64, t string) error{
	var err error
	switch actionType {
	case 1:
		err = Db.Transaction(func(tx *gorm.DB) error {
			if err = repositories.CommentRepository.Publish(Db, uid, vid, t); err != nil {
				log.Println(err)
				return err
			}
			if err = repositories.CommentRepository.IncComment(Db, vid); err != nil {
				log.Println(err)
				return err
			}
			return nil
		})
		if err != nil {
			log.Println(err)
			return err
		}
	case 2:
		// check rights, only video's and comment's author have authorization
		c := CommentService.GetCommentById(cid)
		v := VideoService.GetVideoById(vid)

		if uid != c.UserId || uid != v.AuthorId{
			return errors.New("no permission")
		}

		err = Db.Transaction(func(tx *gorm.DB) error {
			if err = repositories.CommentRepository.Delete(Db, cid); err != nil {
				log.Println(err)
				return err
			}
			if err = repositories.CommentRepository.DecComment(Db, vid); err != nil {
				log.Println(err)
				return err
			}
			return nil
		})
		if err != nil {
			log.Println(err)
			return err
		}
	}
	return nil
}

func (s *commentService) GetComments(vid int64) (comments []model.Comment) {
	comments = repositories.CommentRepository.FindByTimeDesc(Db, vid)
	return
}

func (s *commentService) GetCommentById(id int64) model.Comment {
	return repositories.CommentRepository.GetCommentById(Db, id)
}
