# UltraSlackBot

# UltraSlackBot ![Language Badge](https://img.shields.io/badge/Language-Go-blue.svg) ![Go Report](https://goreportcard.com/badge/github.com/DiSiqueira/UltraSlackBot) ![License Badge](https://img.shields.io/badge/License-MIT-blue.svg) ![Status Badge](https://img.shields.io/badge/Status-Development-brightgreen.svg)

The best Slack bot ever.

## Project Status

UltraSlackBot is on development. Pull Requests [are welcome](https://github.com/DiSiqueira/UltraSlackBot#social-coding)

## Features

## Installation


## Usage

### Basic usage

### Show help

## Program Help


## Contributing

### Bug Reports & Feature Requests

Please use the [issue tracker](https://github.com/DiSiqueira/UltraSlackBot/issues) to report any bugs or file feature requests.

### Developing

PRs are welcome. To begin developing, do this:

1. [Fork it] (https://github.com/DiSiqueira/UltraSlackBot/fork)
2. `$ mkdir -p $GOPATH/src/github.com/disiqueira/ultraslackbot`
3. `$ git clone https://github.com/#YOURGITHUBUSERNAMEHERE#/UltraSlackBot.git $GOPATH/src/github.com/disiqueira/ultraslackbot
4. `$ cd $GOPATH/src/github.com/disiqueira/ultraslackbot`
5. `$ cp .env.example .env`
6. Edit your new .env file and put your Slack token on it
7. `$ docker-compose up -d`
8. Profit! :white_check_mark:

Other info:

1. All the plugins should have it's own folder inside the `/plugins` folder
2. All the dependencies should me managed using [Glide](https://github.com/Masterminds/glide). You can use the [cnpj](https://github.com/DiSiqueira/UltraSlackBot/tree/master/plugins/cnpj) plugin as an example
3. You can see all the events that can be handled by plugins [here](https://godoc.org/github.com/nlopes/slack)
4. The [choose](https://github.com/DiSiqueira/UltraSlackBot/tree/master/plugins/choose) plugin is a good example of a complete plugin
5. The [catfact](https://github.com/DiSiqueira/UltraSlackBot/tree/master/plugins/catfact) plugin is a good example of a "Command" plugin

### Debugging

To see the logs you can use:

`docker-compose logs -f ultraslackbot`

## Social Coding

1. Create an issue to discuss about your idea
2. [Fork it] (https://github.com/DiSiqueira/UltraSlackBot/fork)
3. Create your feature branch (`git checkout -b my-new-feature`)
4. Commit your changes (`git commit -am 'Add some feature'`)
5. Push to the branch (`git push origin my-new-feature`)
6. Create a new Pull Request
7. Profit! :white_check_mark:

## License

The MIT License (MIT)

Copyright (c) 2013-2018 Diego Siqueira

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.  IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
