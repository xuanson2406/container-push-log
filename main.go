package main

import (
	"bytes"
	"context"
	"log"
	"os"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"k8s.io/klog/v2"
)

func main() {
	ctx := context.Background()
	endpoint := "s3-hfx03.fptcloud.com"
	accessKeyID := "VQNDG4HVO3GGV6P85YL2"
	secretAccessKey := "uMhpUKfhC3SmhpN0tTJwLoE2YgjGcJ4yDaheVDvS"
	useSSL := true
	bucketName := "1c7896b7-15ea-424d-9bda-89624019ded5"
	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
		Region: "",
	})
	if err != nil {
		log.Fatalln(err)
	}
	time.Sleep(5 * time.Second)
	envValue := os.Getenv("CLUSTER_NAME")
	fileName := envValue + ".log"
	file, err := os.Open("/etc/kubernetes/audit-log/" + fileName)
	if err != nil {
		log.Fatalln("Can not open file: ", err)
	}
	defer file.Close()

	// Get file info and size
	fileInfo, err := file.Stat()
	if err != nil {
		log.Fatalln(err)
	}
	// // Đọc dữ liệu từ file vào một byte array
	fileSize := fileInfo.Size()
	buffer := make([]byte, fileSize)
	_, err = file.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}
	reader := bytes.NewReader(buffer)

	// minioClient.IN
	objectName := os.Getenv("OBJECT_NAME")
	_, err = minioClient.PutObject(ctx, bucketName, objectName, reader, fileSize, minio.PutObjectOptions{})
	if err != nil {
		log.Fatal(err)
	}
	klog.Infof("Log file of cluster %s uploaded successfully", envValue)
	// fileA := "/tmp/logA.log"
	// fileB := "/tmp/logB.log"
	// err = minioClient.FGetObject(ctx, bucketName, "fke-k8s-1l0kct8f.log", fileA, minio.GetObjectOptions{})
	// if err != nil {
	// 	fmt.Printf("get A")
	// 	log.Fatal(err)
	// }
	// err = minioClient.FGetObject(ctx, bucketName, "fke-k8s-1l0kct8f-1.log", fileB, minio.GetObjectOptions{})
	// if err != nil {
	// 	fmt.Printf("get B")
	// 	log.Fatal(err)
	// }
	// fileLogA, err := os.Open(fileA)
	// if err != nil {
	// 	log.Fatalln("Can not open file: ", err)
	// }
	// defer fileLogA.Close()
	// fileLogB, err := os.Open(fileB)
	// if err != nil {
	// 	log.Fatalln("Can not open file: ", err)
	// }
	// defer fileLogB.Close()
	// fileAinfo, err := fileLogA.Stat()
	// if err != nil {
	// 	fmt.Printf("file A")
	// 	log.Fatalln(err)
	// }
	// fileBinfo, err := fileLogB.Stat()
	// if err != nil {
	// 	fmt.Printf("file B")
	// 	log.Fatalln(err)
	// }
	// // // Đọc dữ liệu từ file vào một byte array
	// fileSizeA := fileAinfo.Size()
	// bufferA := make([]byte, fileSizeA)
	// _, err = fileLogA.Read(bufferA)
	// if err != nil {
	// 	fmt.Printf("read A")
	// 	log.Fatal(err)
	// }
	// fileSizeB := fileBinfo.Size()
	// bufferB := make([]byte, fileSizeB)
	// _, err = fileLogB.Read(bufferB)
	// if err != nil {
	// 	fmt.Printf("read B")
	// 	log.Fatal(err)
	// }
	// buffer := make([]byte, fileSizeA+fileSizeB)
	// buffer = append(bufferA, bufferB...)
	// reader := bytes.NewReader(buffer)
	// _, err = minioClient.PutObject(ctx, bucketName, "fileName-test", reader, fileSizeA+fileSizeB, minio.PutObjectOptions{})
	// if err != nil {
	// 	fmt.Printf("PUT")
	// 	log.Fatal(err)
	// }
	// klog.Infof("Log file of cluster  uploaded successfully")
	return
}
