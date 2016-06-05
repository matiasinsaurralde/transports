# transports

A HTTP proxy that aims to support different transports.

## Motivation

I started the [facebook tunnel](https://github.com/matiasinsaurralde/facebook-tunnel) project two years ago and I thought that it could be better to follow a modular approach for supporting other services (chat systems, platforms, *gram, *book?).

This repository includes some code to explore the idea.

### Why not a TCP/UDP tunnel?

At this time I'm not planning tuntap support (like I did in the previous project). I would like to focus on the transports. Also, I think that a HTTP proxy is easier to port and run, especially when considering that the project is built on Golang, where the output is a static binary. For example, it'll be very easy to build a binary for ARM.

## Available transports

I've been working on these transports during the past week:

### Facebook Transport (early stage, sorry!)

This transport uses [surf](https://github.com/headzoo/surf), a stateful web browser built in Go.

Load your credentials by using ```export``` or the ```.env``` file:

```
FB_LOGIN=youraccount@facebook.com
FB_PASSWORD=supersecretpass
FB_FRIEND=yourtunnelfriend
```

I'm looking for collaborators from countries where the [Internet.org](https://info.internet.org/en/) campaigns like "Free Basics" are active, they could benefit from it :)

### Whatsapp Transport (status: you can perform some GETs)

This transport uses a [HTTP wrapper](https://github.com/matiasinsaurralde/yowsup-http-wrapper) for [yowsup](https://github.com/tgalal/yowsup) to send/receive Whatsapp messages.

I recorded this small video, showing some interactions with this transport. For the demonstration I point my browser to the proxy and perform a test request to Akamai, the communication happens between two Whatsapp clients running on the same computer:

[![Whatsapp Transport](https://img.youtube.com/vi/5KhS7fueK9k/0.jpg)](http://bit.ly/1TTu9wo)

It would be good to have a "pure Golang" Whatsapp library but I think the current approach is fine for experimentation (anyone considering writing this?).

The following environment variables are used:

```
WA_CLIENT_LOGIN=123412341
WA_CLIENT_PASSWORD=whatsappgeneratedpassword123
WA_CLIENT_CONTACT=43214321

WA_SERVER_LOGIN=123412341
WA_SERVER_PASSWORD=whatsappgeneratedpassword123
```

**Requires Python 3**

## Marshalers

I'm working on providing a set of "marshalers" and a simple API to combine them, this could be useful for conducting network/system usage benchmark experiments & performing a good choice.

[Protocol buffers](https://github.com/google/protobuf) sound like a good option, instead of JSON (which is what I'm actually using for the Whatsapp transport). Also [brotli](https://github.com/google/brotli) looks promising. A combination of these two is a very interesting thing to consider.


## Contributors

* [Matias Insaurralde](https://github.com/matiasinsaurralde)

* [Carlos Carvallo](https://github.com/carloscarvallo)

## License

[MIT](https://github.com/matiasinsaurralde/transports/LICENSE)
