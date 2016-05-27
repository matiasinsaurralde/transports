# transports

A HTTP proxy that aims to support different transports.

## Motivation

I started the [facebook tunnel](https://github.com/matiasinsaurralde/facebook-tunnel) project two years ago and I thought that it could be better to follow a modular approach for supporting other services (chat systems, platforms?).

This repository includes some code to explore the idea, also I'm also planning to write a [pluggable transport](https://obfuscation.github.io/) for Tor in the future.

## Setup

### Facebook Transport

Load your credentials by using ```export``` or the ```.env``` file:

```
FB_LOGIN=youraccount@facebook.com
FB_PASSWORD=supersecretpass
FB_FRIEND=yourtunnelfriend
```

## Contributors

* [Matias Insaurralde](https://github.com/matiasinsaurralde)

* [Carlos Carvallo](https://github.com/carloscarvallo)

## License

[MIT](https://github.com/matiasinsaurralde/transports/LICENSE)
