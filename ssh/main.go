package main
import (
	"golang.org/x/crypto/ssh"
	"fmt"
	"time"
	"os"
)

func connect(user, password, host string ,port int ) ( *ssh.Session ,error) {
	var (
		auth		[]ssh.AuthMethod
		addr		string
		clientConfig	*ssh.ClientConfig
		client		*ssh.Client
		session		*ssh.Session
		err		error
	)
	auth = make([]ssh.AuthMethod,0)
	auth = append(auth, ssh.Password(password))
	clientConfig = &ssh.ClientConfig{
		User:	user,
		Auth:	auth,
		Timeout:	30 * time.Second,
	}

	addr = fmt.Sprintf("%s:%d",host,port)
	client, err = ssh.Dial("tcp", addr, clientConfig)
	if err != nil{
		fmt.Println(err)
		return nil,err
	}

	session,err = client.NewSession()
	if err != nil {
		fmt.Println(err)
		return nil,err
	}
	return session,nil
}


func main()  {
	// session, error := connect("app","yeepay.c0m","10.151.32.27", 22)
	session, error := connect("ubuntu","111111","123.206.58.28", 22)
	if error != nil{
		fmt.Println(error)
		os.Exit(1)
	}
	defer session.Close()
	session.Stdout = os.Stdout
	session.Stderr = os.Stderr
	session.Run("pwd")
}
