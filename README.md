# transports

A HTTP proxy that aims to support different transports.

## Motivation

I started the [facebook tunnel](https://github.com/matiasinsaurralde/facebook-tunnel) project two years ago and I thought that it could be better to follow a modular approach for supporting other services (chat systems, platforms, *gram, *book?).

This repository includes some code to explore the idea, also I'm also planning to write a [pluggable transport](https://obfuscation.github.io/) for Tor in the future.

## Available transports

### Facebook Transport

This transports uses [surf](https://github.com/headzoo/surf), a stateful web browser built in Go.

Load your credentials by using ```export``` or the ```.env``` file:

```
FB_LOGIN=youraccount@facebook.com
FB_PASSWORD=supersecretpass
FB_FRIEND=yourtunnelfriend
```

### Whatsapp Transport

This transport uses a [HTTP wrapper](https://github.com/matiasinsaurralde/yowsup-http-wrapper) for [yowsup](https://github.com/tgalal/yowsup) to send/receive Whatsapp messages.

It would be good to have a "pure Golang" Whatsapp library but I think the current approach is fine for experimentation.

The following environment variables are used:

```
WA_CLIENT_LOGIN=123412341
WA_CLIENT_PASSWORD=whatsappgeneratedpassword123
WA_CLIENT_CONTACT=43214321

WA_SERVER_LOGIN=123412341
WA_SERVER_PASSWORD=whatsappgeneratedpassword123
```

**Requires Python 3**

## Contributors

* [Matias Insaurralde](https://github.com/matiasinsaurralde)

* [Carlos Carvallo](https://github.com/carloscarvallo)

## License

[MIT](https://github.com/matiasinsaurralde/transports/LICENSE)
