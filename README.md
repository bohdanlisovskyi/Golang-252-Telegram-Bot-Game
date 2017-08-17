# Golang-252-Telegram-Bot-Game

This repo have telegram bot code!
Each application stores in separate folder and has its own Makefile.

### Set Up Instructions
**1.Clone Repository**

    $   git clone https://github.com/bohdanlisovskyi/Golang-252-Telegram-Bot-Game.git

**2.Run to install glide and gometalinter**

    $   make install-helpers

**3.Run to install application dependencies**

    $   make dependencies

### Code quality check
    $   make code-quality

### How to import package

    $  glide get <some package>

Don't use `go get <some package>`

### Some useful links

* https://github.com/Masterminds/glide
* https://github.com/go-telegram-bot-api/telegram-bot-api
* https://github.com/alecthomas/gometalinter
* https://docs.travis-ci.com/user/languages/go/