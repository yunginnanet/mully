# mully

> [!NOTE]
> This project is in early development and is not yet ready for use.

interacts with a locally running `mullvad-daemon` instance with gRPC and golang, additionally provides a client for the https mullvad API.

gRPC and API definitions are generated automatically from official mullvad sources. protobuf defs are used for gRPC, and openapi for the API.

GNU Make orchestrates these tasks.

## quick start

```bash
_URL1=git.tcp.direct/kayos/mully.git
_URL2=github.com/yunginnanet/mully.git
(git clone "https://$_URL1" || git clone "https://$_URL2") || exit 1
cd mully && make
```
