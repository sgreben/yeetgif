# yeetgif

![<doc/terminal.png gif -n 0 resize -x 500 | gif compose -s 0.8 -p right doc/wobble.gif | gif fried -t 0 -a 0 -j 0 -n 0.2 -u 1.0 -o 1 -j 1 | gif optimize -x 0 -y 0 --kb=500 > doc/terminal.gif](doc/terminal.gif)

<!-- TOC -->

- [Get it](#get-it)
- [Use it](#use-it)
- [Usage](#usage)
    - [roll](#roll)
    - [wobble](#wobble)
    - [pulse](#pulse)
    - [zoom](#zoom)
    - [shake](#shake)
    - [woke](#woke)
    - [fried](#fried)
    - [hue](#hue)
    - [tint](#tint)
    - [resize](#resize)
    - [crop](#crop)
    - [optimize](#optimize)
    - [compose](#compose)
    - [crowd](#crowd)
    - [nop](#nop)
    - [meta](#meta)
- [Licensing](#licensing)

<!-- /TOC -->

## Get it

```sh
go get -u github.com/sgreben/yeetgif/cmd/gif
```

Or [download the binary](https://github.com/sgreben/yeetgif/releases/latest) from the releases page.

```sh
# Linux
curl -LO https://github.com/sgreben/yeetgif/releases/download/${VERSION}/gif_${VERSION}_linux_x86_64.tar.gz | tar xz

# OS X
curl -LO https://github.com/sgreben/yeetgif/releases/download/${VERSION}/gif_${VERSION}_osx_x86_64.tar.gz | tar xz

# Windows
curl -LO https://github.com/sgreben/yeetgif/releases/download/${VERSION}/gif_${VERSION}_windows_x86_64.zip
unzip gif_${VERSION}_windows_x86_64.zip
```

**NOTE**: To use the `optimize` command, you'll also need the [`giflossy`](https://github.com/kornelski/giflossy) fork of `gifsicle` installed:

```sh
brew install giflossy
```

## Use it

```sh
<doc/yeet.png | gif fried | gif wobble | gif crop >doc/yeet.gif
```
![before](doc/yeet.png)
![after](doc/yeet.gif)


## Usage

```text
${USAGE}
```

### roll

![before](doc/eggplant.png)![after](doc/roll.gif)

```text
${USAGE_roll}
```

### wobble

![before](doc/eggplant.png)![after](doc/wobble.gif)

```text
${USAGE_wobble}
```

### pulse

![before](doc/eggplant.png)![after](doc/pulse.gif)

```text
${USAGE_pulse}
```

### zoom

![before](doc/eggplant.png)![after](doc/zoom.gif)

```text
${USAGE_zoom}
```

### shake

![before](doc/eggplant.png)![after](doc/shake.gif)

```text
${USAGE_shake}
```

### woke

![before](doc/yeet.png)![after](doc/woke.gif)

```text
${USAGE_woke}
```

### fried

![before](doc/yeet.png)![after](doc/fried.gif)

```text
${USAGE_fried}
```

### hue

![before](doc/eggplant.png)![after](doc/hue.gif)

```text
${USAGE_hue}
```

### tint

![before](doc/eggplant.png)![after](doc/tint.gif)

```text
${USAGE_tint}
```

### resize

```text
${USAGE_resize}
```

### crop

```text
${USAGE_crop}
```

### optimize

```text
${USAGE_optimize}
```

### compose

![before](doc/yeet.png)![before](doc/eggplant.png)![after](doc/compose.gif)

```text
${USAGE_compose}
```

### crowd

![before](doc/wobble.gif)![after](doc/crowd.gif)

```text
${USAGE_crowd}
```

### nop

```text
${USAGE_nop}
```

### meta

```text
${USAGE_meta}
```

## Licensing

- [Modified copy](pkg/imaging) of `github.com/disintegration/imaging`: [MIT License](pkg/imaging/LICENSE)
- `yeetgif` itself: [MIT License](LICENSE)
