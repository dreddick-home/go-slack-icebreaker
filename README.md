# go-slack-icebreaker

<p align="left">
<img width="100" height="100" src="https://github.com/dreddick-home/go-slack-icebreaker/blob/master/images/icebreaker.jpg">
</p>

Randomly select a channel user to pose an icebreaker question....and have fun!


<p align="left">
<img src="https://img.shields.io/github/go-mod/go-version/dreddick-home/go-slack-icebreaker">
<img src="https://img.shields.io/github/v/release/dreddick-home/go-slack-icebreaker">
<img src="https://github.com/dreddick-home/go-slack-icebreaker/workflows/CICD/badge.svg">
<img src="https://img.shields.io/badge/License-Apache%202.0-blue.svg">
<img src="https://goreportcard.com/badge/github.com/dreddick-home/go-slack-icebreaker">
</p>

## Overview

The command can randomly select a user from a list of active channel members (who are not bots). It will then send the icebreaker message to the channel mentioning them to go pick a question from a website.

The idea is to generate fun and interesting conversation to aid with team cohesion for remote workers.

Ideally the command should be scheduled using a cronjob to have a regular message sent to the channel.

## Usage

Use the following syntax to run the icebreaker commands from your terminal window:

```bash
icebreaker -token 'xxxxxxxxxxx' -channelid 'C01D63W3WKU'
```

The command rqeuires the following flags:
* token : the API token for accessing the Slack API
* channelid : the ID of the channel posting messages to

### Running with docker

```bash
docker run dreddick/go-slack-icebreaker:v0.1.1 -token 'xxxxxxxxxxx' -channelid 'C01D63W3WKU'
```


## Install

### Set up slack

Create an internal app:
* https://api.slack.com/apps
* Create App
* Add permissions "channels:read", "chat:write", "users:read" in the "OAuth & Permissions" section
* Copy "Bot User OAuth Access Token" - this is the token which will be used by the command
* Install app
* Invite app into channel in slack - for example: "/invite icebreaker"

### Build and Install the Binaries from Source

#### Prerequisite Tools

* Git
* Go 


#### Fetch from GitHub

```console
$ git clone https://github.com/dreddick-home/go-slack-icebreaker.git
$ cd go-slack-icebreaker
$ go build -o /usr/local/bin/icebreaker
```


### Releases

See https://github.com/dreddick-home/go-slack-icebreaker/releases

### TODO

* Scan message history to ensure same person doesnt get picked twice in a row

