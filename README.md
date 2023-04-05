<img alt="" src="https://codeberg.org/video-prize-ranch/rimgo/raw/branch/main/static/img/rimgo.svg" width="96" height="96" />

# rimgo
An alternative frontend for Imgur. Originally based on [rimgu](https://codeberg.org/3np/rimgu).

<a href="https://www.gnu.org/licenses/agpl-3.0.en.html">
  <img alt="License: AGPLv3" src="https://shields.io/badge/License-AGPL%20v3-blue.svg" height="20px">
</a>
<a href="https://matrix.to/#/#rimgo:nitro.chat">
  <img alt="Matrix" src="https://img.shields.io/badge/chat-matrix-blue" height="20px">
</a>

## Table of Contents
- [Features](#features)
- [Comparison](#comparison)
  - [Speed](#speed)
  - [Privacy](#privacy)
- [Usage](#usage)
- [Instances](#instances)
  - [Clearnet](#clearnet)
  - [Tor](#tor)
- [Automatically redirect links](#automatically-redirect-links)
  - [LibRedirect](#libredirect)
  - [GreaseMonkey script](#greasemonkey-script)
  - [Redirector](#redirector)
- [Install](#install)
  - [Docker (recommended)](#docker-recommended)
    - [Automatic updates](#automatic-updates)
  - [Build from source](#build-from-source)
    - [Requirements](#requirements)
- [Configuration](#configuration)
  - [Environment variables](#environment-variables)
- [Contributing](#contributing)
- [License](#license)

## Features
- Lightweight
- No JavaScript
- No ads or tracking
- No sign up or app install prompts
- Bandwidth efficient - automatically uses newer image formats (if enabled)

## Comparison
Comparing rimgo to Imgur.

### Speed
Tested using [Google PageSpeed Insights](https://pagespeed.web.dev/).

| | [rimgo](https://pagespeed.web.dev/report?url=https%3A%2F%2Fi.bcow.xyz%2Fgallery%2FgYiQLWy) | [Imgur](https://pagespeed.web.dev/report?url=https%3A%2F%2Fimgur.com%2Fgallery%2FgYiQLWy) |
| ------------------- | ------- | --------- |
| Performance         | 91      | 28        |
| Request count       | 29      | 340       |
| Resource Size       | 218 KiB | 2,542 KiB |
| Time to Interactive | 1.6s    | 23.8s     |

### Privacy
Imgur collects information about your device and uses tracking cookies for advertising, this is mentioned in their [privacy policy](https://imgur.com/privacy/). [Blacklight](https://themarkup.org/blacklight) found 31 trackers and 87 third-party cookies.

See what cookies and trackers Imgur uses and where your data gets sent: https://themarkup.org/blacklight?url=imgur.com

## Usage
Replace imgur.com or i.imgur.com with the instance domain. For i.stack.imgur.com, replace i.stack.imgur.com with the instance domain and add stack/ before the media ID. You can use a browser extension to do this [automatically](#automatically-redirect-links).

Imgur: `https://imgur.com/gallery/j2sOQkJ` -> `https://rimgo.bcow.xyz/gallery/j2sOQkJ`
Stack Overflow: `https://i.stack.imgur.com/KnO3v.jpg?s=64&g=1` -> `https://rimgo.bcow.xyz/stack/KnO3v.jpg?s=64&g=1`

## Instances
Open an issue to have your instance listed here! Instance privacy information is required for the instance list, see [Environment variables](#environment-variables).

> For more details on instance privacy, see https://librarian.codeberg.page/docs/usage/instance-privacy/

### Clearnet
To help distribute load, consider using instances other than the official one.

| URL                                                        	  | Country      | Provider                 | Privacy               | Notes |
| :------------------------------------------------------------ | :----------- | :----------------------- | :-------------------- | :---- |
| [rimgo.pussthecat.org](https://rimgo.pussthecat.org)       	  | 🇩🇪 DE        | Hetzner                  | ⚠️ Data collected     |       |
| [rimgo.totaldarkness.net](https://rimgo.totaldarkness.net) 	  | 🇨🇦 CA        | Vultr                    | ✅ Data not collected |       |
| [rimgo.bus-hit.me](https://rimgo.bus-hit.me)               	  | 🇨🇦 CA        | Oracle                   | ✅ Data not collected |       |
| [rimgo.esmailelbob.xyz](https://rimgo.esmailelbob.xyz)     	  | 🇨🇦 CA        | OVH                      | ⚠️ Data collected     |       |
| [imgur.artemislena.eu](https://imgur.artemislena.eu)       	  | 🇩🇪 DE        | Vodafone Deutschland     | ✅ Data not collected | Self-hosted, provider is ISP |
| [rimgo.vern.cc](https://rimgo.vern.cc)                        | 🇺🇸 US        | OVHCloud                 | ✅ Data not collected | [Edited theme](https://git.vern.cc/root/modifications/src/branch/master/rimgo) |
| [rim.odyssey346.dev](https://rim.odyssey346.dev/)             | 🇫🇷️ FR        | Trolling Solutions (OVH) | ✅ Data not collected |       |
| [rimgo.privacytools.io](https://rimgo.privacytools.io/)       | 🇸🇪 SE        | Cloudflare               | ✅ Data not collected |       |
| [i.habedieeh.re](https://i.habedieeh.re/)                     | 🇨🇦️ CA        | Oracle Cloud             | ✅ Data not collected |       |
| [rimgo.hostux.net](https://rimgo.hostux.net/)	                | 🇫🇷️ FR        | Gandi	                   | ⚠️ Data collected     |       |
| [ri.zzls.xyz](https://ri.zzls.xyz/)                           | 🇨🇱 CL        | TELEFÓNICA CHILE         | ✅ Data not collected | Self-hosted, provider is ISP |
| [rimgo.marcopisco.com](https://rimgo.marcopisco.com/)         | 🇵🇹 PT        | Cloudflare               | ⚠️ Data collected     |       |
| [rimgo.lunar.icu](https://rimgo.marcopisco.com/)              | 🇩🇪 DE        | Cloudflare               | ✅ Data not collected |       |
| [imgur.010032.xyz](https://imgur.010032.xyz/)                 | 🇰🇷 KR        | Cloudflare               | ✅ Data not collected |       |
| [rimgo.kling.gg](https://rimgo.kling.gg/)                     | 🇳🇱 NL        | RamNode                  | ✅ Data not collected |       |
| [i.01r.xyz](https://i.01r.xyz/)                               | 🇺🇸 US        | Cloudflare               | ✅ Data not collected |       |
| [rimgo.projectsegfau.lt](https://rimgo.projectsegfau.lt/)     | 🇱🇺 LU, 🇺🇸 US, 🇮🇳 IN | See below         | ✅ Data not collected |       |
| [rimgo.eu.projectsegfau.lt](https://rimgo.projectsegfau.lt/)  | 🇱🇺 LU | FranTech Solutions              | ✅ Data not collected |       |
| [rimgo.us.projectsegfau.lt](https://rimgo.projectsegfau.lt/)  | 🇺🇸 US | DigitalOcean                    | ✅ Data not collected |       |
| [rimgo.in.projectsegfau.lt](https://rimgo.projectsegfau.lt/)  | 🇮🇳 IN | Airtel                          | ✅ Data not collected |       |

### Tor

| URL | Privacy | Notes                    |
| :-- | :------ | :----------------------- |
| [rimgo.esmail5pdn24shtvieloeedh7ehz3nrwcdivnfhfcedl7gf4kwddhkqd.onion](http://rimgo.esmail5pdn24shtvieloeedh7ehz3nrwcdivnfhfcedl7gf4kwddhkqd.onion) | ✅ Data not collected | Onion of rimgo.esmailelbob.xyz |
| [rimgo.vernccvbvyi5qhfzyqengccj7lkove6bjot2xhh5kajhwvidqafczrad.onion](http://rimgo.vernccvbvyi5qhfzyqengccj7lkove6bjot2xhh5kajhwvidqafczrad.onion) | ✅ Data not collected | Onion of rimgo.vern.cc         |
| [imgur.lpoaj7z2zkajuhgnlltpeqh3zyq7wk2iyeggqaduhgxhyajtdt2j7wad.onion](http://imgur.lpoaj7z2zkajuhgnlltpeqh3zyq7wk2iyeggqaduhgxhyajtdt2j7wad.onion) | ✅ Data not collected | Onion of imgur.artemislena.eu  |
| [rim.odysfvr23q5wgt7i456o5t3trw2cw5dgn56vbjfbq2m7xsc5vqbqpcyd.onion](http://rim.odysfvr23q5wgt7i456o5t3trw2cw5dgn56vbjfbq2m7xsc5vqbqpcyd.onion)     | ⚠️ Data collected |  |
| [tdp6uqjtmok723suum5ms3jbquht6d7dssug4cgcxhfniatb25gcipad.onion](http://tdp6uqjtmok723suum5ms3jbquht6d7dssug4cgcxhfniatb25gcipad.onion)             | ✅ Data not collected | Onion of rimgo.privacytools.io |
| [i.habeehrhadazsw3izbrbilqajalfyqqln54mrja3iwpqxgcuxnus7eid.onion](http://i.habeehrhadazsw3izbrbilqajalfyqqln54mrja3iwpqxgcuxnus7eid.onion/)        | ✅ Data not collected | Onion of i.habedieeh.re |
| [rimgo.zzlsghu6mvvwyy75mvga6gaf4znbp3erk5xwfzedb4gg6qqh2j6rlvid.onion](http://rimgo.zzlsghu6mvvwyy75mvga6gaf4znbp3erk5xwfzedb4gg6qqh2j6rlvid.onion/) | ✅ Data not collected | Onion of ri.zzls.xyz |
| [tdn7zoxctmsopey77mp4eg2gazaudyhgbuyytf4zpk5u7lknlxlgbnid.onion/](http://tdn7zoxctmsopey77mp4eg2gazaudyhgbuyytf4zpk5u7lknlxlgbnid.onion/) | ✅ Data not collected | Onion of rimgo.kling.gg |

### I2P

| URL | Privacy | Notes                    |
| :-- | :------ | :----------------------- |
| [rimgo.i2p](http://rimgo.i2p) | ✅ Data not collected | i.habedieeh.re on I2P |
| [rimgov7l2tqyrm5txrtvhtnfyrzkc5d7ipafofavchbnnyog4r3q.b32.i2p](http://rimgov7l2tqyrm5txrtvhtnfyrzkc5d7ipafofavchbnnyog4r3q.b32.i2p) | ✅ Data not collected | Same as rimgo.i2p |
| [rimgo.zzls.i2p](http://rimgo.zzls.i2p) | ✅ Data not collected | ri.zzls.xyz on I2P |
| [p57356k2xwhxrg2lxrjajcftkrptv4zejeeblzfgkcvpzuetkz2a.b32.i2p](http://p57356k2xwhxrg2lxrjajcftkrptv4zejeeblzfgkcvpzuetkz2a.b32.i2p) | ✅ Data not collected | Same as rimgo.zzls.i2p |
| [ovzamsts5czfx3jasbbhbccyyl2z7qmdngtlqxdh4oi7abhdz3ia.b32.i2p](http://ovzamsts5czfx3jasbbhbccyyl2z7qmdngtlqxdh4oi7abhdz3ia.b32.i2p) | ✅ Data not collected | rimgo.kling.gg on I2P |

## Automatically redirect links

### LibRedirect
Use [LibRedirect](https://github.com/libredirect/libredirect) to automatically redirect Imgur links to rimgo!
- [Firefox](https://addons.mozilla.org/firefox/addon/libredirect/)
- [Chromium-based browsers (Brave, Google Chrome)](https://github.com/libredirect/libredirect#install-in-chromium-browsers)

### GreaseMonkey script
There is a script to redirect Imgur links to rimgo.
https://codeberg.org/zortazert/GreaseMonkey-Redirect/src/branch/main/imgur-to-rimgo.user.js

### Redirector
You can use the [Redirector](https://github.com/einaregilsson/Redirector) extension to redirect Imgur links to rimgo with the configuration below:

* Description: Imgur -> rimgo
* Example URL: https://imgur.com/a/H8M4rcp
* Include pattern: `^https?://i?.?imgur.com(/.*)?$`
* Redirect to: `https://rimgo.example.com$1`
* Pattern type: Regular Expression
* Advanced options:
  * Apply to:
    * [x] Main window (address bar)
    * [x] Images

For Stack Overflow images:
* Description: Stack Overflow Imgur -> rimgo
* Example URL: https://i.stack.imgur.com/BTKqD.png?s=128&g=1
* Include pattern: `^https?://i\.stack\.imgur\.com(/.*)?$`
* Redirect to: `https://rimgo.example.com/stack$1`
* Pattern type: Regular Expression
* Advanced options:
  * Apply to:
    * [x] Images

## Install
rimgo can run on any platform Go compiles on.

> It is strongly recommended to use [Caddy](https://caddyserver.com/) as your reverse proxy. Caddy is simple to configure, automatically manages your TLS certificates, and provides better performance with support for HTTP/2 and /3 (allow UDP port 443 in your firewall to use HTTP/3).

### Docker (recommended)
Install [Docker](https://docs.docker.com/engine/install/) and [docker-compose](https://docs.docker.com/compose/install/), then clone this repository.
```bash
git clone https://codeberg.org/video-prize-ranch/rimgo
cd rimgo
```

Edit the `docker-compose.yml` file using your favorite text editor.
```bash
nvim docker-compose.yml
```

You can now run rimgo.
```bash
sudo docker-compose up -d
```

#### Automatic updates
[Watchtower](https://containrrr.dev/watchtower/) can automatically update your Docker containers.

Create a new `docker-compose.yml` file or add the watchtower section to your existing `docker-compose.yml` file.
```yml
version: "3"
services:
  watchtower:
    image: containrrr/watchtower
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock
```

### Build from source

#### Requirements
* Go v1.16 or later

Clone the repository and `cd` into it.
```bash
git clone https://codeberg.org/video-prize-ranch/rimgo
cd rimgo
```

Build rimgo.
```bash
go build
```

You can now run rimgo.
```bash
./rimgo
```

To include version information use:
```bash
go build -ldflags "-X codeberg.org/video-prize-ranch/rimgo/pages.VersionInfo=$(date '+%Y-%m-%d')-$(git rev-list --abbrev-commit -1 HEAD)"
```

(optional) You can use a .env file to set environment variables for configuration.
```bash
cp .env.example .env
nvim .env
```

## Configuration

rimgo can be configured using environment variables. The path to the .env file can be changed the -c flag.

### Environment variables

> For more details on instance privacy, see https://librarian.codeberg.page/docs/usage/instance-privacy/

| Name                  | Default         | Note |
|-----------------------|-----------------|------|
| PORT                  | 3000            |      |
| ADDRESS               | 0.0.0.0         |      |
| IMGUR_CLIENT_ID       | 546c25a59c58ad7 |      |
| FORCE_WEBP            | 0               |      |
| PRIVACY_POLICY        |                 | Optional, URL to privacy policy |
| PRIVACY_MESSAGE       |                 | Optional, message to display on privacy page |
| PRIVACY_COUNTRY       |                 |      |
| PRIVACY_PROVIDER      |                 |      |
| PRIVACY_CLOUDFLARE    |                 |      |
| PRIVACY_NOT_COLLECTED |                 |      |
| PRIVACY_IP            |                 |      |
| PRIVACY_URL           |                 |      |
| PRIVACY_DEVICE        |                 |      |
| PRIVACY_DIAGNOSTICS   |                 |      |

## Contributing
Pull requests are welcome! If you have any questions or bug reports, open an [issue](https://codeberg.org/video-prize-ranch/rimgo/issues/new).

## License
This software is released under the AGPL-3.0 license. If you make any modifications to the code and distribute it (including use on a network server), you must publicly distribute your changes and release them under the AGPL-3.0.
