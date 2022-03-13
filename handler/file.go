package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mingzhehao/goutils/filetool"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
	db "workspace/cloud-pan/db/mysql"
	"workspace/cloud-pan/util/logs"
)

type FileQueryReq struct {
	Type     string `form:"type" json:"type"`
	Username string `form:"username" json:"username"`
}

//文件列表
func FileQueryHandler(c *gin.Context) {
	rsp := &Result{}
	defer func() {
		c.JSON(http.StatusOK, rsp)
	}()

	req := &FileQueryReq{}
	if err := c.ShouldBindQuery(&req); err != nil {
		logs.Errorf("err:%v", err)
		return
	}
	logs.Debugf("req:%+v", req)

	loginUser, _ := c.Cookie("username")
	where := fmt.Sprintf("username='%s'",loginUser)
	if req.Type != "all" {
		where = fmt.Sprintf("%s and type='%s'",where,req.Type)
	}
	logs.Debugf("where:%v",where)
	rsp.Data, _ = db.Select().From("files").Where(where).Order("uploadtime desc").Query()
}

type Sizer interface {
	Size() int64
}

const (
	LOCAL_FILE_DIR    = "static/upload/"
	MAX_FILE_SIZE     = 200 //文件最大限制20M
	IMAGE_TYPES       = "(jpg|gif|png)"
	MOV_TYPES         = "(mp4)"
	DOC_TYPES         = "(doc|docx|xls|xlsx)"
	OTHER_TYPES       = "(txt)"
	ACCEPT_FILE_TYPES = "(jpg|gif|png|mp4|doc|docx|xls|xlsx|txt)"
)

var (
	imageTypes      = regexp.MustCompile(IMAGE_TYPES)
	movTypes        = regexp.MustCompile(MOV_TYPES)
	docTypes        = regexp.MustCompile(DOC_TYPES)
	otherTypes      = regexp.MustCompile(OTHER_TYPES)
	acceptFileTypes = regexp.MustCompile(ACCEPT_FILE_TYPES)
)

type FileInfo struct {
	Url          string `json:"url,omitempty"`
	ThumbnailUrl string `json:"thumbnailUrl,omitempty"`
	Name         string `json:"name"`
	Ext          string `json:"ext"`
	Type         string `json:"type"`
	Size         int64  `json:"size"`
	Error        string `json:"error,omitempty"`
	DeleteUrl    string `json:"deleteUrl,omitempty"`
	DeleteType   string `json:"deleteType,omitempty"`
}

func (fi *FileInfo) ValidateType() (valid bool) {
	if acceptFileTypes.MatchString(fi.Ext) {
		return true
	}
	fi.Error = "Filetype not allowed"
	return false
}

//文件上传
func FileOssHandler(c *gin.Context) {
	rsp := &Result{}
	defer func() {
		c.JSON(http.StatusOK, rsp)
	}()

	path := LOCAL_FILE_DIR

	h, err := c.FormFile("file")
	if err != nil {
		rsp.Status = 500
		rsp.Msg = "上传失败"
	} else {
		ext := filetool.Ext(h.Filename)
		fi := &FileInfo{
			Name: h.Filename,
			Ext:  ext,
			Size:h.Size,
		}

		logs.Debugf("ext:%v", fi.Ext)
		if !fi.ValidateType() {
			rsp.Status = 500
			rsp.Msg = "非法的文件格式"
			return
		}
		filesize := h.Size / 1024 /1024
		if filesize > MAX_FILE_SIZE {
			rsp.Status = 500
			rsp.Msg = fmt.Sprintf("超过文件最大限制 %vM", MAX_FILE_SIZE)
			return
		}

		if imageTypes.MatchString(fi.Ext) {
			fi.Type = "pic"
		} else if movTypes.MatchString(fi.Ext) {
			fi.Type = "mov"
		} else if docTypes.MatchString(fi.Ext) {
			fi.Type = "doc"
		} else if otherTypes.MatchString(fi.Ext) {
			fi.Type = "other"
		}
		logs.Debugf("type:%v", fi.Type)

		//先上传到本地临时文件夹
		fileExt := strings.TrimLeft(ext, ".")
		fileSaveName := fmt.Sprintf("%s_%d%s", fileExt, time.Now().UnixNano(), ext)
		dir, _ := os.Getwd()
		imgPath := fmt.Sprintf("%s\\%s%s", dir, strings.Replace(path, "/", "\\", -1), fileSaveName)

		filetool.InsureDir(path)
		if err := c.SaveUploadedFile(h, path+fileSaveName); err != nil {
			c.String(http.StatusBadRequest, "保存失败 Error:%s", err.Error())
			return
		}

		logs.Debugf("imgpath:%v", imgPath)

		logs.Debugf("dir:%v", dir)


		//上传到oss
		if filesize > 10 {
			//大于10M用分片上传
			err = MultipartUpload(fileSaveName, imgPath)
			if err != nil {
				rsp.Status = 500
				rsp.Msg = "上传失败"
			}
		} else {
			//小于10M用简单上传
			err = SimpleUpload(fileSaveName, imgPath)
			if err != nil {
				rsp.Status = 500
				rsp.Msg = "上传失败"
			}
		}

		username, _ := c.Cookie("username")

		logs.Debugf("filesize:%v",fi.Size)
		//记录存入数据库
		db.Insert().Table("files").Value(db.GetValues().Adds(
			"username", username,
			"name", fileSaveName,
			"type", fi.Type,
			"size", fi.Size/1024,
			"uploadtime", time.Now().Format("2006-01-02 15:04:05"),
			"updatetime", time.Now().Format("2006-01-02 15:04:05"),
			)).Exec()
	}
}

type DownloadFileReq struct {
	ObjectName string `form:"objectname" json:"objectname"`
}

//oss授权文件下载
func DownloadURLHandler(c *gin.Context) {
	rsp := &Result{}
	defer func() {
		c.JSON(http.StatusOK, rsp)
	}()

	req := &DownloadFileReq{}
	if err := c.ShouldBindQuery(&req); err != nil {
		logs.Errorf("err:%v", err)
		return
	}
	logs.Debugf("req:%+v", req)
	url, err := SignUrl(req.ObjectName)
	if err != nil {
		rsp.Status = 500
		rsp.Msg = "下载地址获取失败"
	}
	rsp.Data = url
}
