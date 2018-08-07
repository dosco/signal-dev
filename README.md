# Signal-DEV, A development server for Signal apps

While playing around with the signal iOS/Android and desktop apps I felt the need for
a simpler signal backend that could be used to debug the protocol as well as quickly
experiment with new features.

This is not a real signal backend and only implements a subset of the API at this moment.
The orginal signal server is in Java while this is entirely in GoLang. We used boltDB as
a lite persistance layer.

#### Quickstart

To start the service. It runs on 8080 (HTTP)

```console
go run *.go
```

#### Extra

If you want to use TLS then you would need self-signed certificates for localhost. Only thing that matters here is common name so enter 'localhost' when prompted for that all other values can be blank

```console
openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout key.pem -out cert.pem
```

#### Signal iOS

You need to patch the iOS app to allow for unsecured localhost connections. The patch is included just apply it.

```console
git clone https://github.com/signalapp/Signal-iOS.git
cd Signal-iOS
git am ~/signal_dev_checkout_folder/signal-ios.patch
```

