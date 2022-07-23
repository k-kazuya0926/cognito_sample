package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"os"
)

func main() {
	username := "user1@example.com"
	password := "User1Pass"
	clientId := os.Getenv("CLIENT_ID")      // 「全般設定」画面で確認
	userPoolId := os.Getenv("USER_POOL_ID") // 「アプリクライアント」画面で確認

	svc := cognitoidentityprovider.New(session.New(), &aws.Config{Region: aws.String("ap-northeast-1")})

	// ログイン
	params := &cognitoidentityprovider.AdminInitiateAuthInput{
		AuthFlow: aws.String("ADMIN_NO_SRP_AUTH"),
		AuthParameters: map[string]*string{
			"USERNAME": aws.String(username),
			"PASSWORD": aws.String(password),
		},
		ClientId:   aws.String(clientId),
		UserPoolId: aws.String(userPoolId),
	}

	resp, err := svc.AdminInitiateAuth(params)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(resp)
}
