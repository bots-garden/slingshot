package main

import (
	slingshot "github.com/bots-garden/slingshot/go-pdk"
)

func Handler(argHandler []byte) ([]byte, error) {
	input := string(argHandler)
	slingshot.Print("👋 hello world 🌍 " + string(input))
	slingshot.Log("🙂 have a nice day 🏖️")
	//TODO: set header
	return []byte(`{msg:"hey!"}`), nil
}

func main() {
	slingshot.Print("👋 main function")
	slingshot.SetHandler(Handler)
}

/* sample with AWS
var gorillaLambda *gorillamux.GorillaMuxAdapter

func init() {
    r := mux.NewRouter()

    r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
        json.NewEncoder(w).Encode(Response{From: "gorilla", Message: time.Now().Format(time.UnixDate)})
    })

    gorillaLambda = gorillamux.New(r)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
    r, err := gorillaLambda.ProxyWithContext(ctx, *core.NewSwitchableAPIGatewayRequestV1(&req))
    return *r.Version1(), err
}

func main() {
    lambda.Start(Handler)
}

si pas de fonction par défaut
appeler main()
ou appeler main tout le temps ?
*/
