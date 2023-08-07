FROM scratch

ADD slingshot-v0.0.0-linux-arm64 ./
ADD simple.wasm ./

EXPOSE 8080

CMD ["./slingshot-v0.0.0-linux-arm64", "./simple.wasm", "handle", "8080"]
