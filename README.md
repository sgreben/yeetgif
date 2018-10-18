# yeetgif

Composable GIF effects CLI, with reasonable defaults. Made for custom Slack/Discord emoji :)

![terminal](doc/terminal.gif)

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
    - [erase](#erase)
    - [chop](#chop)
    - [text](#text)
    - [nop](#nop)
    - [meta](#meta)
- [Hall of Fame](#hall-of-fame)
- [Licensing](#licensing)

<!-- /TOC -->

## Get it

```sh
go get -u github.com/sgreben/yeetgif/cmd/gif
```

Or [download the binary](https://github.com/sgreben/yeetgif/releases/latest) from the releases page.

```sh
# Linux
curl -LO https://github.com/sgreben/yeetgif/releases/download/1.15.0/gif_1.15.0_linux_x86_64.tar.gz | tar xz

# OS X
curl -LO https://github.com/sgreben/yeetgif/releases/download/1.15.0/gif_1.15.0_osx_x86_64.tar.gz | tar xz

# Windows
curl -LO https://github.com/sgreben/yeetgif/releases/download/1.15.0/gif_1.15.0_windows_x86_64.zip
unzip gif_1.15.0_windows_x86_64.zip
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

Usage: gif [OPTIONS] COMMAND [arg...]

                     
Options:             
  -n                 Duplicate a single input image this many times (default 20)
  -q, --quiet        Disable all log output (stderr)
  -d, --delay-ms     Frame delay in milliseconds (default 20)
  -p, --pad          Pad images (default true)
      --write-meta   Write command line options into output GIF metadata (default true)
                     
Commands:            
  roll               (☭ ͜ʖ ☭)
  wobble             🍆( ͡° ͜ʖ ͡°)🍆
  pulse              ( ͡◉ ͜ʖ ͡◉)
  zoom               (⌐▀͡ ̯ʖ▀)
  shake              ˵(˵ ͡⚆ ͜ʖ ͡⚆˵)˵
  woke               💯  W O K E F L A R E S ( ͡ 🅱️ ͜ʖ ͡ 🅱️ ) 💯
  fried              fr͍͈i̗̟̲̻e͕̗d̬ m̷͔͊e̶̪̿m̷̙̈́é̵̤s̷̺͒
  hue                ( ͡☆ ͜ʖ ͡☆)
  tint               🎨༼ຈل͜ຈ༽
  resize             (° ͜ʖ°)¯\_( ͡☉ ͜ʖ ͡☉)_/¯
  crop               ┬┴┬┴┤ ͜ʖ ͡°)
  optimize           👌( ͡ᵔ ͜ʖ ͡ᵔ )👌
  compose            (ﾉ ͡° ͜ʖ ͡°)ﾉ*:･ﾟ✧
  crowd              (⟃ ͜ʖ ⟄) ͜ʖ ͡°)( ° ͜ʖ( ° ͜ʖ °)
  erase              ( ͡° ͜ʖ ͡°)=ε/̵͇̿̿/'̿̿ ̿ ̿ ̿ ̿ ̿
  chop               ✂️( ͡°Ĺ̯ ͡° )🔪
  text               🅰️乁(˵ ͡☉ ͜ʖ ͡☉˵)┌🅱️
  nop                乁(ᴗ ͜ʖ ᴗ)ㄏ
  meta               (🧠 ͡ಠ ʖ̯ ͡ಠ)┌
                     
Run 'gif COMMAND --help' for more information on a command.
```

### roll

![before](doc/eggplant.png)![after](doc/roll.gif)

```text

Usage: gif roll [OPTIONS]

(☭ ͜ʖ ☭)
                      
Options:              
  -r, --revolutions   (default 1)
  -s, --scale         (default 1)
  -p, --phase         (default 0)
```

### wobble

![before](doc/eggplant.png)![after](doc/wobble.gif)

```text

Usage: gif wobble [OPTIONS]

🍆( ͡° ͜ʖ ͡°)🍆
                    
Options:            
  -f, --frequency   (default 1)
  -a, --amplitude   (default 20)
  -p, --phase       (default 0)
  -t, --type        one of [sine snap saw sticky] (default sine)
      --custom      comma-separated angles (°), e.g. 0,10,0,60,0 (default [])
```

### pulse

![before](doc/eggplant.png)![after](doc/pulse.gif)

```text

Usage: gif pulse [OPTIONS]

( ͡◉ ͜ʖ ͡◉)
                    
Options:            
  -0, --from        (default 1)
  -1, --to          (default 1.5)
  -f, --frequency   (default 1)
  -p, --phase       (default 0)
```

### zoom

![before](doc/eggplant.png)![after](doc/zoom.gif)

```text

Usage: gif zoom [OPTIONS]

(⌐▀͡ ̯ʖ▀)
                 
Options:         
  -0, --from     (default 1)
  -1, --to       (default 1.5)
  -c, --custom   (default [])
```

### shake

![before](doc/eggplant.png)![after](doc/shake.gif)

```text

Usage: gif shake [OPTIONS]

˵(˵ ͡⚆ ͜ʖ ͡⚆˵)˵
                    
Options:            
  -f, --frequency   (default [1])
  -a, --amplitude   (default [7])
  -r, --random      🌀 (default 0.75)
```

### woke

![before](doc/yeet.png)![after](doc/woke.gif)

```text

Usage: gif woke [OPTIONS] POINTS

💯  W O K E F L A R E S ( ͡ 🅱️ ͜ʖ ͡ 🅱️ ) 💯
                          
Arguments:                
  POINTS                  flare locations, JSON, e.g. "[[123,456],[-100,23]]" (default &[])
                          
Options:                  
  -c, --clip              clip flares to image alpha (default true)
  -t, --type              (default full)
  -s, --scale             (default 0.9)
  -u, --hue               (default 0.8)
  -l, --lightness         (default 1)
  -a, --alpha             (default 0.8)
  -p, --alpha-pow         (default 2)
      --alpha-threshold   (default 0.15)
  -r, --random            🌀 (default 0.5)
```

### fried

![before](doc/yeet.png)![after](doc/fried.gif)

```text

Usage: gif fried [OPTIONS]

fr͍͈i̗̟̲̻e͕̗d̬ m̷͔͊e̶̪̿m̷̙̈́é̵̤s̷̺͒
                     
Options:             
      --clip         (default true)
  -j, --jpeg         [0,100] (default 84)
  -w, --walk         🌀 (default 10)
  -i, --iterations   (default 1)
  -a                 🅰️ (default 0.33)
  -b                 🅱️ (default 0.2)
  -c                 🆑 (default 0.9)
  -n, --noise        🌀️ (default 1)
      --noise1       🌀️ (default 0.02)
      --noise2       🌀️ (default 0.5)
      --noise3       🌀 (default 0.1)
  -u, --saturation   (default 3)
  -o, --contrast     (default 6)
  -t, --tint         tint (default 0.4)
```

### hue

![before](doc/eggplant.png)![after](doc/hue.gif)

```text

Usage: gif hue [OPTIONS]

( ͡☆ ͜ʖ ͡☆)
                    
Options:            
  -f, --frequency   (default 1)
  -a, --amplitude   (default 0.1)
```

### tint

![before](doc/eggplant.png)![after](doc/tint.gif)

```text

Usage: gif tint [OPTIONS]

🎨༼ຈل͜ຈ༽
                    
Options:            
  -f, --frequency   (default 1)
  -0, --from        (default 0.7)
  -1, --to          (default 0.9)
  -i, --intensity   (default 0.95)
```

### resize

```text

Usage: gif resize [OPTIONS]

(° ͜ʖ°)¯\_( ͡☉ ͜ʖ ͡☉)_/¯
                 
Options:         
  -s, --scale    (default 1)
  -x, --width    width (pixels) (default 0)
  -y, --height   height (pixels) (default 0)
```

### crop

```text

Usage: gif crop [OPTIONS]

┬┴┬┴┤ ͜ʖ ͡°)
                    
Options:            
  -t, --threshold   (default 0)
```

### optimize

```text

Usage: gif optimize [OPTIONS]

👌( ͡ᵔ ͜ʖ ͡ᵔ )👌
                 
Options:         
      --kb       target file size (KB) (default 128)
  -x, --width    target width (pixels) (default 128)
  -y, --height   target height (pixels) (default 128)
```

### compose

![before](doc/yeet.png)![before](doc/eggplant.png)![after](doc/compose.gif)

```text

Usage: gif compose [OPTIONS] INPUT

(ﾉ ͡° ͜ʖ ͡°)ﾉ*:･ﾟ✧
                   
Arguments:         
  INPUT            
                   
Options:           
  -x               (default 0)
  -y               (default 0)
  -z, --z-order    one of [under over] (default over)
  -p, --position   one of [center left right top bottom abs] (default center)
  -s, --scale      (default 1)
```

### crowd

![before](doc/wobble.gif)![after](doc/crowd.gif)

```text

Usage: gif crowd [OPTIONS]

(⟃ ͜ʖ ⟄) ͜ʖ ͡°)( ° ͜ʖ( ° ͜ʖ °)
                 
Options:         
  -n             crowd size (default 3)
      --flip     🌀 flip (default true)
  -x             🌀 x (default 0.5)
  -y             🌀 y (default 0.25)
  -s, --scale    🌀 [0.0,1.0] (default 0.25)
  -r, --rotate   🌀 [0.0,1.0] (default 0.1)
  -a, --alpha    🌀 [0.0,1.0] (default 0)
  -o, --offset   🌀 [0.0,1.0] (default 1)
```

### erase

![before](doc/skeledance.gif)![after](doc/erase.gif)

```text

Usage: gif erase [OPTIONS]

( ͡° ͜ʖ ͡°)=ε/̵͇̿̿/'̿̿ ̿ ̿ ̿ ̿ ̿
                    
Options:            
  -x, --sample-x    (default 3)
  -y, --sample-y    (default 3)
  -t, --tolerance   (default 0.2)
  -u                (default 1)
  -s                (default 0.5)
  -l                (default 1)
```

### chop

```text

Usage: gif chop COMMAND [arg...]

✂️( ͡°Ĺ̯ ͡° )🔪
               
Commands:      
  shuffle      
  duplicate    
  drop-every   
  drop-first   
  drop-last    
  reverse      
               
Run 'gif chop COMMAND --help' for more information on a command.
```

### text

![before](doc/gunther.jpg)![after](doc/gunther.gif)
> woke | text | fried

```text

Usage: gif text [OPTIONS] [TEXT]

🅰️乁(˵ ͡☉ ͜ʖ ͡☉˵)┌🅱️
                             
Arguments:                   
  TEXT                       (default "#yeetgif")
                             
Options:                     
  -a, --background-alpha     (default 0.7)
  -s, --font-size            (default 18.5)
  -y, --text-y               (default 0.3)
  -p, --background-padding   (default 3)
```

### nop

```text

Usage: gif nop

乁(ᴗ ͜ʖ ᴗ)ㄏ
```

### meta


![input](doc/yeet.gif)
```sh
$ <doc/yeet.gif gif meta show

[2018-10-05T13:08:57+02:00] gif fried
[2018-10-05T13:08:58+02:00] gif wobble
[2018-10-05T13:08:58+02:00] gif crop
[2018-10-05T13:08:58+02:00] gif optimize -x 0
```

```text

Usage: gif meta COMMAND [arg...]

(🧠 ͡ಠ ʖ̯ ͡ಠ)┌
               
Commands:      
  show         show 🧠
  add          add 🧠
  clear        remove 🧠
               
Run 'gif meta COMMAND --help' for more information on a command.
```

## Hall of Fame

Tweet a GIF made with yeetgif using the [`#yeetgif`](https://twitter.com/hashtag/yeetgif) hashtag. Best ones end up below :)

> No entries yet

## Licensing

- [Modified copy](pkg/imaging) of `github.com/disintegration/imaging`: [MIT License](pkg/imaging/LICENSE)
- `yeetgif` itself: [MIT License](LICENSE)
- [Roboto Regular TrueType Font](pkg/gifstatic/roboto.go): [Apache License 2.0](pkg/gifstatic/roboto.go-LICENSE)
