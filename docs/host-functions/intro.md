# 🛠️ Host functions

This section explains how to use the host functions with an Extism wasm plugin.

- [Display](display.md)
- [Environment](env.md)
- [Memory Cache](mem.md)
- [Redis Cache](redis.md)
- [Redis PubSub](redis-pubsub.md)
- [Nats Publish](nats.md)

!!! info "About Extism Host Functions"
    You can find more information into the **Extism** documentation: [https://extism.org/docs/concepts/host-functions](https://extism.org/docs/concepts/host-functions).
    Each [Plug-in Development Kits](https://extism.org/docs/concepts/pdk) provides the "bridge helpers" allowing using the host function callbacks of the host application from the guest wasm plug-in. Then you need to "assemble" the "bridge helpers" to call the host function. It's not always straightforward, and you need to do it for every language you use.

    Right now, we want to make sure that **SlingShot** is compliant with **Extism** PDKs and that's why we're only working on the "host" part.
    
    In the future, the **Slingshot** project will propose **layers of abstraction** to simplify the use of host functions.

    👀 you can follow these issues:

      - [Create slingshot-go-pdk](https://github.com/bots-garden/slingshot/issues/5)
      - [Create slingshot-rust-pdk](https://github.com/bots-garden/slingshot/issues/6)
