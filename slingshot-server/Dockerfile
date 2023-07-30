FROM ubuntu:22.04 as extism-install

ARG DEBIAN_FRONTEND=noninteractive

RUN apt-get update
RUN apt-get -y install software-properties-common
RUN add-apt-repository universe
RUN apt-get update
RUN apt-get -y install build-essential curl python3.11 python3-pip git vim

# Install extism cli and sharedlib
RUN pip3 install poetry
RUN pip3 install git+https://github.com/extism/cli
RUN extism --prefix=/usr/local install latest
RUN pip3 install extism

FROM debian

ENV LD_LIBRARY_PATH="/usr/local/lib"

COPY --from=extism-install /usr/local/lib/libextism.so /usr/local/lib/libextism.so
COPY --from=extism-install /usr/local/include/extism.h /usr/local/include/extism.h

COPY slingshot-v0.0.0-linux-arm64 ./slingshot
COPY simple.wasm .

EXPOSE 8080

CMD ["./slingshot", "./simple.wasm", "handle", "8080"]

