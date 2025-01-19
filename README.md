# Go URL Shortener Exercise

This repository contains a Go implementation of a URL shortener, inspired by the Gophercises exercise. The application provides an HTTP handler that redirects short URLs to their corresponding long URLs, similar to services like Bitly.

## Features

- **URL Redirection**: Maps short URLs to long URLs and redirects users accordingly.
- **Data Storage**: Supports storing URL mappings in YAML files.
- **Fallback Handling**: Provides a fallback handler for unmapped URLs.

## Setup

1. **Clone the Repository**:

   ```bash
   git clone https://github.com/saravanastar/go-url-shortner-exercise.git
   ```

2. **Configuring the url prefix and the path**
    Url prefix and respective redirect url can be configured in the ./src/config/url-shortner.ya,l

    ```yaml

        - path: /google
        url: https://www.google.com

        - path: /facebook
        url: https://www.facebook.com

        - path: /twitter
        url: https://www.twitter.com
    ```
3. **Running Local**
    To run this project locally either can run it as docker container or with the go command

    - Docker:
    ```bash
        make  run-docker
    ```
    - Go
    ```bash
        make run
    ```

## Contributing

Feel free to fork the repository, make improvements, and submit pull requests.