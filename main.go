package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

func main() {
	email := "user1@example.com"
	password := "User1Pass"
	clientID := os.Getenv("CLIENT_ID")      // 「全般設定」画面で確認
	userPoolID := os.Getenv("USER_POOL_ID") // 「アプリクライアント」画面で確認

	sess, err := session.NewSession()
	// 次のようにしている記事もある
	//sess := session.Must(session.NewSessionWithOptions(session.Options{
	//	SharedConfigState: session.SharedConfigEnable,
	//}))
	if err != nil {
		log.Fatalln(err.Error())
	}
	cognitoIdentityProvider := cognitoidentityprovider.New(sess, &aws.Config{Region: aws.String("ap-northeast-1")})

	//// ユーザー作成
	//email2 := "user2@example.com"
	//adminCreateUserInput := &cognitoidentityprovider.AdminCreateUserInput{
	//	//ClientMetadata: nil,
	//	DesiredDeliveryMediums: []*string{
	//		aws.String("EMAIL"),
	//	},
	//	//ForceAliasCreation: nil,
	//	//MessageAction:      nil,
	//	//TemporaryPassword:  nil,
	//	UserAttributes: []*cognitoidentityprovider.AttributeType{
	//		{
	//			Name:  aws.String("email"),
	//			Value: aws.String(email2),
	//		},
	//	},
	//	UserPoolId: &userPoolID,
	//	Username:   &email2,
	//	//ValidationData: nil,
	//}
	//fmt.Printf("adminCreateUserInput: %+v\n", adminCreateUserInput)
	//adminCreateUserOutput, err := cognitoIdentityProvider.AdminCreateUser(adminCreateUserInput)
	//if err != nil {
	//	log.Fatalln(err.Error())
	//}
	//fmt.Printf("adminCreateUserOutput: %+v\n", adminCreateUserOutput)

	// TODO メールアドレス認証

	// TODO パスワードリセット

	// ログイン
	adminInitiateAuthInput := &cognitoidentityprovider.AdminInitiateAuthInput{
		//AnalyticsMetadata: &cognitoidentityprovider.AnalyticsMetadataType{
		//	AnalyticsEndpointId: nil,
		//},
		AuthFlow: aws.String(cognitoidentityprovider.AuthFlowTypeAdminNoSrpAuth),
		AuthParameters: map[string]*string{
			"USERNAME": aws.String(email),
			"PASSWORD": aws.String(password),
		},
		ClientId: aws.String(clientID),
		//ClientMetadata: nil,
		//ContextData: &cognitoidentityprovider.ContextDataType{
		//	EncodedData: nil,
		//	HttpHeaders: nil,
		//	IpAddress:   nil,
		//	ServerName:  nil,
		//	ServerPath:  nil,
		//},
		UserPoolId: aws.String(userPoolID),
	}
	adminInitiateAuthOutput, err := cognitoIdentityProvider.AdminInitiateAuth(adminInitiateAuthInput)
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Printf("adminInitiateAuthOutput: %+v\n", adminInitiateAuthOutput)

	// TODO 初期パスワードの変更
	//adminRespondToAuthChallengeInput := &cognitoidentityprovider.AdminRespondToAuthChallengeInput{
	//	//AnalyticsMetadata: &cognitoidentityprovider.AnalyticsMetadataType{
	//	//	AnalyticsEndpointId: nil,
	//	//},
	//	ChallengeName: aws.String("NEW_PASSWORD_REQUIRED"),
	//	ChallengeResponses: map[string]*string{
	//		"NEW_PASSWORD": aws.String(password),
	//		"USERNAME":     aws.String(email),
	//	},
	//	ClientId: aws.String(clientID),
	//	//ClientMetadata: nil,
	//	//ContextData: &cognitoidentityprovider.ContextDataType{
	//	//	EncodedData: nil,
	//	//	HttpHeaders: nil,
	//	//	IpAddress:   nil,
	//	//	ServerName:  nil,
	//	//	ServerPath:  nil,
	//	//},
	//	Session:    adminInitiateAuthOutput.Session,
	//	UserPoolId: aws.String(userPoolID),
	//}
	//
	//adminRespondToAuthChallengeOutput, err := cognitoIdentityProvider.AdminRespondToAuthChallenge(adminRespondToAuthChallengeInput)
	//if err != nil {
	//	log.Fatalln(err.Error())
	//}
	//fmt.Printf("adminRespondToAuthChallengeOutput: %+v\n", adminRespondToAuthChallengeOutput)

	// TODO メールアドレス変更

	// TODO 表示名変更

	// ログアウト
	globalSignOutInput := &cognitoidentityprovider.GlobalSignOutInput{
		AccessToken: aws.String(*adminInitiateAuthOutput.AuthenticationResult.AccessToken),
	}
	globalSignOutOutput, err := cognitoIdentityProvider.GlobalSignOut(globalSignOutInput)
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Printf("globalSignOutOutput: %+v\n", globalSignOutOutput)

	// TODO ユーザー削除
}
