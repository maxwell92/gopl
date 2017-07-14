package main
import (
	"github.com/chenqinghe/baidu-ai-go-sdk/voice"
	"log"
	"os"
	"io/ioutil"
	"encoding/base64"
	"fmt"

)

var apiKey = "5RijeBzVjQ82uPx8gxGGfeNXlfRt7yH6"
var	apiSecret = "keiyq3oKrkYsSPUcrf0gtRKneeTxjuqV"

//var speech = "T G I F是Thanks God It's Friday的意思，感天谢地今天周五了"
var speech = "你好"

func main() {
//	speech := os.Args[1]
	client := voice.NewVoiceClient(apiKey, apiSecret)
	if err := client.Auth(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ClientId=%s, ClientSecret=%s, AccessToken=%s\n", client.ClientID, client.ClientSecret, client.AccessToken)
/*
	file, err := client.TextToSpeech(speech)
    if err != nil {
		log.Fatal(err)
	}
	f, err := os.OpenFile("hello.mp3", os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	if _, err := f.Write(file); err != nil {
		log.Fatal(err)
	}
*/
	f, err := os.OpenFile("hello.wav", os.O_RDONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	fi, err1 := ioutil.ReadAll(f)
	if err1 != nil {
		log.Fatal(err1)
	}
	afterBase64Str := base64.StdEncoding.EncodeToString(fi)
	fiLen := len(fi)
	param := voice.ASRParams{
		Format: "wav",
		Rate: 16000,
		Channel: 1,
		Cuid: "12312312112", 
		Token: client.AccessToken,
		Lan: "zh",
		Speech: afterBase64Str,
		Len: fiLen,
	}
	// fmt.Printf("Token=%s\n", param.Token)
	fmt.Println(param.Token)
	rs, err2 := client.SpeechToText(param)
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Println(rs)
}


