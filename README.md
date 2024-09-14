# GitHub Inviter

GitHub Inviter is a web application that allows users to join a specific GitHub organization and team using an
invitation code. The application provides a simple web interface where users can enter their GitHub username and the
provided invitation code to get added to the organization's team.

## Features

- Simple web interface for user input
- Secure invitation process using a pre-defined invitation code
- Configurable for different GitHub organizations and teams
- Optional TLS support for secure connections, with automatic redirect

## Screenshot

![Screenshot](/assets/images/index.png)

## Configuration

The application is configured using environment variables. Here are the available configuration options:

| Environment Variable | Description                                             | Required | Default |
|----------------------|---------------------------------------------------------|----------|---------|
| `GITHUB_ORG_NAME`    | The name of your GitHub organization                    | Yes      | -       |
| `GITHUB_TOKEN`       | GitHub personal access token with necessary permissions | Yes      | -       |
| `GITHUB_GROUP_NAME`  | The name of the team in your organization               | Yes      | -       |
| `INVITE_CODE`        | The invitation code users need to provide               | Yes      | -       |
| `HTTP_PORT`          | The port on which the application will run              | No       | 80      |
| `HTTPS_PORT`         | The port on which the application will run (https)      | No       | 443     |
| `TLS_CERT`           | Path to the TLS certificate file                        | No       | -       |
| `TLS_KEY`            | Path to the TLS key file                                | No       | -       |

## Running with Docker Compose

Here's an example `docker-compose.yml` file to run the GitHub Inviter:

```yaml
services:
  github-inviter:
    container_name: github-inviter
    image: ghcr.io/frostwalk/github-inviter:latest
    environment:
      - GITHUB_ORG_NAME=your-org-name
      - GITHUB_TOKEN=your-github-token
      - GITHUB_GROUP_NAME=your-team-name
      - INVITE_CODE=your-invite-code
      - HTTP_PORT=80
      - HTTPS_PORT=443
      # Uncomment the following lines if you want to use TLS
      # - TLS_CERT=/path/to/your/cert.pem
      # - TLS_KEY=/path/to/your/key.pem
    ports:
      - "80:80"
    # Uncomment the following lines if you want to use TLS
    #  - "443:443"
    # volumes:
    #  - /path/to/your/cert.pem:/path/to/your/cert.pem:ro
    #  - /path/to/your/key.pem:/path/to/your/key.pem:ro
  ```

To run the application:

1. Save the above content in a file named `docker-compose.yml`
2. Replace the placeholder values with your actual configuration
3. Run the following command in the same directory as your `docker-compose.yml` file:

```
docker-compose up -d
```

The application will be available at `http://127.0.0.1:8080` (or `https://127.0.0.1:8080` if TLS is configured).

## About the token

in order for this app to work, it needs an
organization [Personal access token](https://docs.github.com/en/organizations/managing-programmatic-access-to-your-organization/managing-requests-for-personal-access-tokens-in-your-organization)
with **read and write permissions on members.**

### Credit

Thanks to
*[Code, Applied To Life](https://medium.com/code-applied-to-life/automated-github-organization-invites-3e940aa27040#.sikfvzyaj)*
this app is heavily inspired by their work.
