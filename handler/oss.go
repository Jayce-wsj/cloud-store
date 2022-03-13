package handler

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"os"
	"time"
	"workspace/cloud-pan/config"
)

//简单上传
func SimpleUpload(objectKey string, filePath string) (err error) {
	// 创建OSSClient实例。
	// yourEndpoint填写Bucket对应的Endpoint，以华东1（杭州）为例，填写为https://oss-cn-hangzhou.aliyuncs.com。其它Region请按实际情况填写。
	// 阿里云账号AccessKey拥有所有API的访问权限，风险很高。强烈建议您创建并使用RAM用户进行API访问或日常运维，请登录RAM控制台创建RAM用户。
	client, _ := oss.New(config.OSSEndpoint, config.OSSAccesskeyID, config.OSSAccessKeySecret)

	// 填写存储空间名称，例如examplebucket。
	bucket, err := client.Bucket(config.OSSBucket)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// 依次填写Object的完整路径（例如exampledir/exampleobject.txt）和本地文件的完整路径（例如D:\\localpath\\examplefile.txt）。
	err = bucket.PutObjectFromFile(objectKey, filePath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	return
}

//分片上传
func MultipartUpload(objectName string, locaFilename string) (err error) {
	// 创建OSSClient实例。
	// yourEndpoint填写Bucket所在地域对应的Endpoint。以华东1（杭州）为例，Endpoint填写为https://oss-cn-hangzhou.aliyuncs.com。
	// 阿里云账号AccessKey拥有所有API的访问权限，风险很高。强烈建议您创建并使用RAM用户进行API访问或日常运维，请登录RAM控制台创建RAM用户。
	client, _ := oss.New(config.OSSEndpoint, config.OSSAccesskeyID, config.OSSAccessKeySecret)
	// 填写存储空间名称。
	bucketName := config.OSSBucket
	// 填写Object完整路径。Object完整路径中不能包含Bucket名称。
	//objectName := "exampleobject.txt"
	// 填写本地文件的完整路径。如果未指定本地路径，则默认从示例程序所属项目对应本地路径中上传文件。
	//locaFilename := "D:\\localpath\\examplefile.txt"

	// 获取存储空间。
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	// 将本地文件分片，且分片数量指定为3。
	chunks, err := oss.SplitFileByPartNum(locaFilename, 3)
	fd, err := os.Open(locaFilename)
	defer fd.Close()

	// 指定过期时间。
	expires := time.Date(2049, time.January, 10, 23, 0, 0, 0, time.UTC)
	// 如果需要在初始化分片时设置请求头，请参考以下示例代码。
	options := []oss.Option{
		oss.MetadataDirective(oss.MetaReplace),
		oss.Expires(expires),
		// 指定该Object被下载时的网页缓存行为。
		// oss.CacheControl("no-cache"),
		// 指定该Object被下载时的名称。
		// oss.ContentDisposition("attachment;filename=FileName.txt"),
		// 指定该Object的内容编码格式。
		// oss.ContentEncoding("gzip"),
		// 指定对返回的Key进行编码，目前支持URL编码。
		// oss.EncodingType("url"),
		// 指定Object的存储类型。
		// oss.ObjectStorageClass(oss.StorageStandard),
	}

	// 步骤1：初始化一个分片上传事件，并指定存储类型为标准存储。
	imur, err := bucket.InitiateMultipartUpload(objectName, options...)
	// 步骤2：上传分片。
	var parts []oss.UploadPart
	for _, chunk := range chunks {
		fd.Seek(chunk.Offset, os.SEEK_SET)
		// 调用UploadPart方法上传每个分片。
		part, err := bucket.UploadPart(imur, fd, chunk.Size, chunk.Number)
		if err != nil {
			fmt.Println("Error:", err)
			return err
		}
		parts = append(parts, part)
	}

	// 指定Object的读写权限为公共读，默认为继承Bucket的读写权限。
	objectAcl := oss.ObjectACL(oss.ACLPublicRead)

	// 步骤3：完成分片上传，指定文件读写权限为公共读。
	cmur, err := bucket.CompleteMultipartUpload(imur, parts, objectAcl)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("cmur:", cmur)

	return
}

//授权下载
func SignUrl(objectName string) (signedURl string, err error) {
	client, _ := oss.New(config.OSSEndpoint, config.OSSAccesskeyID, config.OSSAccessKeySecret)
	bucketName := config.OSSBucket
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	signedURl, err = bucket.SignURL(objectName, oss.HTTPGet, 3600)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	return
}
