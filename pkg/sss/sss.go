package sss

import (
	"bytes"
	"diplomaProject/pkg/globalVars"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"io"
	"mime/multipart"
	"strings"
	"time"
)

var sess = session.Must(session.NewSession(&aws.Config{
	Region: aws.String("eu-north-1"),
	Credentials: credentials.NewStaticCredentials(globalVars.TEAMUP_BUCKET_ID,
		globalVars.TEAMUP_BUCKET_SECRET, ""),
}))

var svc = s3.New(sess)

func UploadPic(form *multipart.Form, suffix string) (link string, err error) {
	fileHeaders := form.File["file"]
	if len(fileHeaders) == 0 {
		return "", errors.New("no file in multipart form")
	}

	file, err := fileHeaders[0].Open()
	if err != nil {
		return "", err
	}
	defer func() {
		err = file.Close()
	}()

	var buf bytes.Buffer
	_, err = io.Copy(&buf, file)
	splitName := strings.Split(fileHeaders[0].Filename, ".")
	ext := splitName[len(splitName)-1]

	t := time.Now()
	link = fmt.Sprintf("%d%d%d%d%d%d-%s", t.Year(),
		t.Month(), t.Day(), t.Hour(),
		t.Minute(), t.Second(), suffix) + "-pic." + ext

	_, err = svc.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(globalVars.TEAMUP_BUCKET_NAME),
		Key:    aws.String(link),
		Body:   strings.NewReader(buf.String()),
		ACL:    aws.String("public-read"), // make public
	})

	if err != nil {
		return "", err
	}

	link = "https://teamup-online.s3.eu-north-1.amazonaws.com/" + link

	return link, nil
}
