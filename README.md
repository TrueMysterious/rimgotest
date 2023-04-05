<img src="https://codeberg.org/video-prize-ranch/rimgo/raw/branch/main/static/img/rimgo.svg" width="96" height="96" />
#Rimgo

<hr>
<hr>
An alternative frontend for Imgur. Based on [rimgu](https://codeberg.org/3np/rimgu) and rewritten in Go.
<a href="https://www.gnu.org/licenses/agpl-3.0.en.html">
  <img alt="License: AGPLv3" src="https://shields.io/badge/License-AGPL%20v3-blue.svg">
</a>
<a href="https://matrix.to/#/#rimgo:nitro.chat">
  <img alt="Matrix" src="https://img.shields.io/badge/chat-matrix-blue">
</a>
<a href="https://gitlab.com/overtime-zone-wildfowl/rimgo">
  <img alt="CI" src="https://gitlab.com/overtime-zone-wildfowl/rimgo/badges/main/pipeline.svg">
</a>

It's read-only and works without JavaScript. Images and albums can be viewed without wasting resources from downloading and running tracking scripts. No sign-up nags.

## Features

- URL-compatible with i.imgur.com - just replace the domain in the URL
- Images and videos (~~gifv~~, mp4)
- Galleries with comments
- Albums
- User page
- Tag page

Some things left to implement (contributions welcome!):

- [x] Streaming (currently media is downloaded in full in rimgu before it's returned)
- [ ] Localization/internationalization
- [x] Pretty CSS styling (responsive?)
- [ ] Support for other popular image sites
- [x] Filtering and exploration on user/tags pages
- [x] Responsive scaling of videos on user/tags pages
- [x] Logo

Things that are considered out of scope:

* Uploading, commenting, voting, etc - rimgo is read-only.
* Authentication, serving HTTPS, rate limiting, etc - Use a reverse proxy or load balancer like Caddy, Traefik, or NGINX.
* Anything requiring JavaScript or the client directly interacting with upstream servers.

## Instances

Open an issue to have your instance listed here!

### Clearnet

| URL                                                        | Country | Cloudflare |
| :--------------------------------------------------------- | :------ | :--------- |
| [i.bcow.xyz](https://i.bcow.xyz) (official)                | 🇳🇱️ NL   |            |
| [rimgo.pussthecat.org](https://rimgo.pussthecat.org)       | 🇩🇪 DE   |            |
| [img.riverside.rocks](https://img.riverside.rocks)         | 🇺🇸 US   |            |
| [rimgo.totaldarkness.net](https://rimgo.totaldarkness.net) | 🇨🇦 CA   |            |
| [rimgo.bus-hit.me](https://rimgo.bus-hit.me)               | 🇨🇦 CA   |            |
| [rimgo.esmailelbob.xyz](https://rimgo.esmailelbob.xyz)     | 🇨🇦 CA   |            |
| [rimgo.lunar.icu](https://rimgo.lunar.icu)                 | 🇩🇪 DE   | 😢         |
| [i.actionsack.com](https://i.actionsack.com)               | 🇺🇸 US   | 😢         |
| [rimgo.privacydev.net](https://irimgo.privacydev.net)      | 🇺🇸 US   |            |

### Tor

| URL | Country |
| :-- | :------ |
| [l4d4owboqr6xcmd6lf64gbegel62kbudu3x3jnldz2mx6mhn3bsv3zyd.onion](http://l4d4owboqr6xcmd6lf64gbegel62kbudu3x3jnldz2mx6mhn3bsv3zyd.onion/) | N/A |
| [jx3dpcwedpzu2mh54obk5gvl64i2ln7pt5mrzd75s4jnndkqwzaim7ad.onion](http://jx3dpcwedpzu2mh54obk5gvl64i2ln7pt5mrzd75s4jnndkqwzaim7ad.onion) | 🇺🇸 US |
| [rimgo.lqs5fjmajyp7rvp4qvyubwofzi6d4imua7vs237rkc4m5qogitqwrgyd.onion](http://rimgo.lqs5fjmajyp7rvp4qvyubwofzi6d4imua7vs237rkc4m5qogitqwrgyd.onion) | 🇨🇦 CA |
| [be7udfhmnzqyt7cxysg6c4pbawarvaofjjywp35nhd5qamewdfxl6sid.onion](http://be7udfhmnzqyt7cxysg6c4pbawarvaofjjywp35nhd5qamewdfxl6sid.onion) | 🇦🇺 AU |

### I2P

| URL | Country |
| :-- | :------ |
| [xazdnfgtzmcbcxhmcbbvr4uodd6jtn4fdiayasghywdn227xsmoa.b32.i2p](http://xazdnfgtzmcbcxhmcbbvr4uodd6jtn4fdiayasghywdn227xsmoa.b32.i2p) | 🇦🇺 AU |

## Automatically redirect links

### LibRedirect
Use [LibRedirect](https://github.com/libredirect/libredirect) to automatically redirect Imgur links to rimgo!
- [Firefox](https://addons.mozilla.org/firefox/addon/libredirect/)
- [Chromium-based browsers (Brave, Google Chrome)](https://github.com/libredirect/libredirect#install-in-chromium-brave-and-chrome)
- [Edge](https://microsoftedge.microsoft.com/addons/detail/libredirect/aodffkeankebfonljgbcfbbaljopcpdb)

### GreaseMonkey script
There is a script to redirect Imgur links to rimgo.
[https://codeberg.org/zortazert/GreaseMonkey-Redirect/src/branch/main/imgur-to-rimgo.user.js](https://codeberg.org/zortazert/GreaseMonkey-Redirect/src/branch/main/imgur-to-rimgo.user.js)

## Install
rimgo can run on any platform Go compiles on.

### Docker (recommended)
Install Docker and docker-compose, then clone this repository.
```
git clone https://codeberg.org/video-prize-ranch/rimgo
cd rimgo
```

Edit the `docker-compose.yml` file using your favorite text editor.
```
nvim docker-compose.yml
```

You can now run rimgo.
```
sudo docker-compose up -d
```

### Build from source

#### Requirements
* Go v1.16 or later

Clone the repository and `cd` into it.
```
git clone https://codeberg.org/video-prize-ranch/rimgo
cd rimgo
```

Build rimgo.
```
go build
```

You can now run rimgo.
```
./rimgo
```

## Configuration

rimgo can be configured using environment variables.

### Environment variables

| Name            | Default         |
|-----------------|-----------------|
| PORT            | 3000            |
| ADDRESS         | 0.0.0.0         |
| IMGUR_CLIENT_ID | 546c25a59c58ad7 |

## Contributing

PRs are welcome! You can also send patches to `cb.8a3w5@simplelogin.co` but pull requests are preferred.

This software is released under the AGPL 3.0 license. In short, this means that if you make any modifications to the code and then publish the result (e.g. by hosting the result on a web server), you must publicly distribute your changes and declare that they also use AGPL 3.0.
