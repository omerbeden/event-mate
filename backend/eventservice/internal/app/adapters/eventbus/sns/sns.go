package sns

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

type SNSAdapter struct {
	Topic string
}

func (snsA *SNSAdapter) Subscribe() {

}

// TODO: generic olabilir , bir dusun
func (snsA *SNSAdapter) Publish() {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := sns.New(sess)

	result, err := svc.Publish(&sns.PublishInput{})
	if err != nil {
		fmt.Println("ERR publish event")
	}
	fmt.Println(result)
}
