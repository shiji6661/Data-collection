package pkg

import (
	"common/appconfig"
	"context"
	"fmt"
	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss"
	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss/credentials"
	"log"
)

const BUCKET_NAME = "yooz1"

func AliYunUpload(localFile, objectName string) {
	// 加载默认配置并设置凭证提供者和区域
	cfg := oss.LoadDefaultConfig().
		WithCredentialsProvider(credentials.NewStaticCredentialsProvider(appconfig.NaCos.ALiYun.AccessKeyId, appconfig.NaCos.ALiYun.AccessKeySecret)).
		WithRegion(appconfig.NaCos.ALiYun.Address)

	// 创建OSS客户端
	client := oss.NewClient(cfg)

	// 创建上传对象的请求
	putRequest := &oss.PutObjectRequest{
		Bucket:       oss.Ptr(BUCKET_NAME),     // 存储空间名称
		Key:          oss.Ptr(objectName),      // 对象名称
		StorageClass: oss.StorageClassStandard, // 指定对象的存储类型为标准存储
		Acl:          oss.ObjectACLPublicRead,  // 指定对象的访问权限为公有访问
		Metadata: map[string]string{
			"yourMetadataKey1": "yourMetadataValue1", // 设置对象的元数据
		},
	}

	// 执行上传对象的请求
	result, err := client.PutObjectFromFile(context.TODO(), putRequest, localFile)
	if err != nil {
		log.Fatalf("failed to put object from file %v", err)
	}

	// 打印上传对象的结果
	log.Printf("put object from file result:%#v\n", result)
}

func GetUrl(filename string) string {
	return fmt.Sprintf("https://%s.oss-cn-shanghai.aliyuncs.com/%s", BUCKET_NAME, filename)
}
