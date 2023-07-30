# slingshot

## Plugin creation

### Rust

```bash
cargo new --lib hello-rust-plugin --name hello
```

In the generated Cargo.toml, be sure to include:

```toml
[lib]
crate_type = ["cdylib"]
```
> ref: https://doc.rust-lang.org/reference/linkage.html

Add dependencies:
```bash
cd hello-rust-plugin
cargo add extism-pdk
cargo add serde
cargo add serde_json
```

#### Build 

```bash
# build
cargo clean
cargo build --release --target wasm32-wasi #--offline
# ls -lh *.wasm
ls -lh ./target/wasm32-wasi/release/*.wasm
```

#### Run

```bash
extism call ./target/wasm32-wasi/release/hello.wasm \
  hello --input "Bob Morane"  \
  --wasi
```