<p align="center">
  <img src="docs/img/screenshot.png" alt="Snippetbox" width="500"/>
</p>

# Snippetbox

This application was developed by following the [Let's Go book by Alex Edwards](https://lets-go.alexedwards.net/).

It's a full-stack Go web application called Snippetbox that lets people paste and share snippets of text — a bit like GitHub's Gists or Pastebin.

## Features

- Save and view snippets.
- Middleware.
- RESTful routing.
- SSL/TLS web server using HTTP 2.0.
- User authentication. User can signup and login.
- Leveled logging.
- Data persistence using MySQL database.
- Dynamic HTML using Go templates
- Session management
- Web security

## Development

Software requirements:

- This project supports Go modules. Go 1.13+ is required.
- MySQL
- make

To start the local web server with HTTPS on port 4000:

```sh
$ git clone https://github.com/cedrickchee/snippetbox
$ cd snippetbox

# Generate TLS key and cert
$ make cert

$ make dev
```

To run the tests, run `make test`.
