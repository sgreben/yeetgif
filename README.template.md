# yeetgif

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


## Licensing

- [Modified copy](pkg/imaging) of `github.com/disintegration/imaging`: [MIT License](pkg/imaging/LICENSE)
- `yeetgif` itself: [MIT License](LICENSE)
