# Write and serve a plug-in as a nano-service

=== "Go"
    ```golang linenums="1"
    package main

    import (
        "strings"

        "github.com/extism/go-pdk"
        "github.com/valyala/fastjson"
    )

    var parser = fastjson.Parser{}

    //export handle
    func handle()  {
        // read function argument from the memory
        http_request_data := pdk.Input()

        var text string
        var code string

        JSONData, err := parser.ParseBytes(http_request_data)
        if err != nil {
            text = "😡 Error: " + err.Error()
            code = "500"
        } else {
            text = "🩵 Hello " + string(JSONData.GetStringBytes("body"))
            code = "200"
        }

        headers := []string{
            `"Content-Type": "application/json; charset=utf-8"`,
            `"X-Slingshot-version": "0.0.0"`,
        }

        headersStr := strings.Join(headers, ",")

        response := `{"headers":{` + headersStr + `}, "textBody": "` + text + `", "statusCode": `+ code +`}`

        mem := pdk.AllocateString(response)
        // copy output to host memory
        pdk.OutputMemory(mem)	
    }

    func main() {}
    ```

=== "Rust"
    ```rust linenums="1"
    use std::collections::HashMap;

    use extism_pdk::*;
    use json::Value;
    use serde::{Serialize, Deserialize};

    #[derive(Serialize)]
    #[serde(rename_all = "camelCase")]
    struct Response {
        pub text_body: String,
        pub json_body: Value,
        pub headers : HashMap<String,String>,
        pub status_code: i64
    }

    #[derive(Serialize, Deserialize)]
    #[serde(rename_all = "camelCase")]
    struct Request {
        pub body: String,
        pub base_url: String,
        pub headers : HashMap<String,String>,
        pub method: String
    }


    #[plugin_fn]
    pub fn handle(http_request_data: String) -> FnResult<Json<Response>> {

        let request : Request = serde_json::from_str(&http_request_data).unwrap();

        let message: String = "🩵 Hello ".to_string() + &request.body;

        let mut headers: HashMap<String, String> = HashMap::new();
        headers.insert("Content-Type".to_string(), "text/plain; charset=utf-8".to_string());
        headers.insert("X-Slingshot-version".to_string(), "0.0.0".to_string());

        let response : Response = Response { 
            text_body: message, 
            json_body: serde_json::from_str("{}")?, 
            headers , 
            status_code: 200 
        };
        
        Ok(Json(response))
    }
    ```

=== "JavaScript"
    ```javascript linenums="1"
    function handle() {

        // read function argument from the memory
        let http_request_data = Host.inputString()

        let JSONData = JSON.parse(http_request_data)

        let text = "💛 Hello " + JSONData.body

        let response = {
            headers: {
                "Content-Type": "application/json; charset=utf-8",
                "X-Slingshot-version": "0.0.0"
            },
            textBody: text,
            statusCode: 200
        }

        // copy output to host memory
        Host.outputString(JSON.stringify(response))
    }

    module.exports = {handle}
    ```

## Build

=== "Go"
    ```bash linenums="1"
    #!/bin/bash
    tinygo build -scheduler=none --no-debug \
    -o hello.wasm \
    -target wasi main.go
    ```
=== "Rust"
    ```bash linenums="1"
    #!/bin/bash
    cargo clean
    cargo build --release --target wasm32-wasi
    ls -lh ./target/wasm32-wasi/release/*.wasm
    cp ./target/wasm32-wasi/release/*.wasm .
    ```
=== "JavaScript"
    ```bash linenums="1"
    #!/bin/bash
    extism-js index.js -o hello.wasm
    ```
    
## Run

```bash linenums="1"
#!/bin/bash
./slingshot start --wasm=./hello.wasm --handler=handle --http-port=7070

🌍 http-port: 7070
🚀 handler  : handle
📦 wasm     : ./hello.wasm
🌍 slingshot server is listening on: 7070
```

### Query the service

```bash linenums="1"
#!/bin/bash
curl --verbose \
http://localhost:7070 \
-H 'content-type: text/plain; charset=utf-8' \
-d '😄 Bob Morane'
```

### Output

```bash linenums="1"
> POST / HTTP/1.1
> Host: localhost:7070
> User-Agent: curl/7.88.1
> Accept: */*
> content-type: text/plain; charset=utf-8
> Content-Length: 15
> 
< HTTP/1.1 200 OK
< Date: Tue, 15 Aug 2023 14:11:59 GMT
< Content-Type: application/json; charset=utf-8
< Content-Length: 26
< X-Slingshot-Version: 0.0.0
< 
* Connection #0 to host localhost left intact
🩵 Hello 😄 Bob Morane
```