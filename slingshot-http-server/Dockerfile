FROM scratch

ADD slingshot-v0.0.0-linux-arm64 ./
ADD k33g_go_handler_plugin.wasm ./

EXPOSE 8080

CMD ["./slingshot-v0.0.0-linux-arm64", "./k33g_go_handler_plugin.wasm", "handle", "8080"]
