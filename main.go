package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"log"
	"os"
)

func main() {
	username := "user1@example.com"
	password := "User1Pass"
	clientId := os.Getenv("CLIENT_ID")      // 「全般設定」画面で確認
	userPoolId := os.Getenv("USER_POOL_ID") // 「アプリクライアント」画面で確認

	session, err := session.NewSession()
	if err != nil {
		log.Fatalln(err.Error())
	}
	svc := cognitoidentityprovider.New(session, &aws.Config{Region: aws.String("ap-northeast-1")})

	// TODO ユーザー作成

	// TODO メールアドレス認証

	// ログイン
	adminInitiateAuthInput := &cognitoidentityprovider.AdminInitiateAuthInput{
		AuthFlow: aws.String("ADMIN_NO_SRP_AUTH"),
		AuthParameters: map[string]*string{
			"USERNAME": aws.String(username),
			"PASSWORD": aws.String(password),
		},
		ClientId:   aws.String(clientId),
		UserPoolId: aws.String(userPoolId),
	}

	adminInitiateAuthOutput, err := svc.AdminInitiateAuth(adminInitiateAuthInput)
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println(adminInitiateAuthOutput)

	// TODO パスワードリセット

	// TODO パスワード変更

	// TODO メールアドレス変更

	// TODO 表示名変更

	// ログアウト
	globalSignOutInput := &cognitoidentityprovider.GlobalSignOutInput{
		AccessToken: aws.String(*adminInitiateAuthOutput.AuthenticationResult.AccessToken),
	}
	globalSignOutOutput, err := svc.GlobalSignOut(globalSignOutInput)
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Printf("globalSignOutOutput: %+v\n", globalSignOutOutput)

	// TODO ユーザー削除
}
