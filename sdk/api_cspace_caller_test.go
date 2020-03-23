package sdk

import (
	"github.com/flyingtime/dingtalk-sdk-golang/encrypt"
	"github.com/flyingtime/dingtalk-sdk-golang/file"
	"github.com/flyingtime/dingtalk-sdk-golang/json"
	"os"
	"testing"
	"time"
)

func TestDingTalkClient_FileUploadSingle(t *testing.T) {
	resp, _ := CreateClient().FileUploadSingle("C:\\Users\\admin\\Desktop\\dingding-test.jpg")
	t.Log(json.ToJson(resp))
}

func TestDingTalkClient_SendDingPanFileToSingleChat(t *testing.T) {
	client := CreateClient()

	mediaId := "#iAEHAqRmaWxlA6h5dW5kaXNrMATOCw_n8wXNBvQGzWSTB85dIrQTCM0iXw"
	filename := "test.jpg"

	resp, _ := client.SendDingPanFileToSingleChat("15001956402427783", encrypt.URLEncode(mediaId), encrypt.URLEncode(filename))
	t.Log(json.ToJson(resp))

}

func TestDingTalkClient_AddFileToUserCSpace(t *testing.T) {
	//TODO
}

func TestDingTalkClient_GetCustomSpace(t *testing.T) {
	resp, err := CreateClient().GetCustomSpace("test")
	t.Log(json.ToJson(resp))
	t.Log(err)
}

func TestDingTalkClient_GrantCustomSpace(t *testing.T) {
	//TODO
}

func TestDingTalkClient_UploadTransaction(t *testing.T) {
	//TODO

	//path := "D:\\download\\easyscheduler-EasyScheduler-dev.zip"
	path := "D:\\download\\20190528151901.png"
	count := 1
	fileInfo, err := os.Stat(path)
	if err != nil {
		t.Log(err)
	}

	client := CreateClient()
	fileSize := fileInfo.Size()

	beginTrans, err := client.BeginUploadTransaction(fileSize, count)
	if err != nil {
		t.Log(err)
	}
	t.Log(json.ToJson(beginTrans))

	uploadId := encrypt.URLEncode(beginTrans.UploadId)

	time.Sleep(time.Duration(2) * time.Second)

	buf, err := file.GetFileReader(path)
	if err != nil {
		t.Log(err)
	}
	for i := 1; i <= count; i++ {
		uploadChunkResp, err := client.BeginUploadChunk(uploadId, i, buf)
		if err != nil {
			t.Log(err)
		}
		t.Log(json.ToJson(uploadChunkResp))
		time.Sleep(time.Duration(2) * time.Second)
	}

	resp, _ := client.CommitUploadTransaction(fileSize, count, uploadId)
	t.Log(json.ToJson(resp))
}
