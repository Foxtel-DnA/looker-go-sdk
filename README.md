# looker-go-sdk

```go
package main

import (
	apiclient "looker-go-sdk/client"
	"looker-go-sdk/client/api_auth"
	"looker-go-sdk/client/group"

	"github.com/go-openapi/strfmt"

	httptransport "github.com/go-openapi/runtime/client"
)

func main() {
	transport := httptransport.New("[YOUR LOOKER HOST].looker.com:19999", "/api/3.0/", nil)
	client := apiclient.New(transport, strfmt.Default)

	clientID := "[YOUR CLIENT_ID]"
	clientSecret := "[YOUR_CLIENT_SECRET]"

	pd := api_auth.NewLoginParams()
	pd.ClientID = &clientID
	pd.ClientSecret = &clientSecret

	resp, err := client.APIAuth.Login(pd)

	if err != nil {
		println("error: " + err.Error())
		return
	}

	token := resp.Payload.AccessToken
	println("token: " + token)

	authInfoWriter := httptransport.APIKeyAuth("Authorization", "header", "token "+token)
	transport.DefaultAuthentication = authInfoWriter

	authClient := apiclient.New(transport, strfmt.Default)

	groupParams := group.NewGroupParams()
	groupParams.GroupID = 45

	group, err := authClient.Group.Group(groupParams)
	if err != nil {
		println("error: " + err.Error())
		return
	}

	println("group: " + group.Payload.Name + ", group id: " + string(group.Payload.ID))
}
```