package main
/*
import (
	"github.com/chenqinghe/baidu-ai-go-sdk/voice"
	"log"
	"encoding/base64"
	"fmt"
)

var apiKey = "5RijeBzVjQ82uPx8gxGGfeNXlfRt7yH6"
var	apiSecret = "keiyq3oKrkYsSPUcrf0gtRKneeTxjuqV"

//var speech = "T G I F是Thanks God It's Friday的意思，感天谢地今天周五了"
var speech = "你好"

func main() {
	client := voice.NewVoiceClient(apiKey, apiSecret)
	if err := client.Auth(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ClientId=%s, ClientSecret=%s, AccessToken=%s\n", client.ClientID, client.ClientSecret, client.AccessToken)
	file, err := client.TextToSpeech(speech)
    if err != nil {
		log.Fatal(err)
	}

	afterBase64Str := base64.StdEncoding.EncodeToString(file)
	fiLen := len(file)
	param := voice.ASRParams{
		Format: "wav",
		Rate: 16000,
		Channel: 1,
		Cuid: "adfadfaf12112", 
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
*/


