## Stable.World Client

Stable.world provides easy to use wrapper scripts around common applications

### Installation

#### Mac OSX

```sh
curl https://install.stable.world/osx | bash -s -- <wrapper-name> [... <wrapper-name>]
```

eg to install `scurl`

```sh
curl https://install.stable.world/osx | bash -s -- scurl
```

### Setup

1. You need to go to https://console.stable.world and set up an account.
2. Get your `STABLE_WORLD_BUCKET` token from the site.
3. export your `STABLE_WORLD_BUCKET` as an environment variable.

#### Example

```
# Export the variable
export STABLE_WORLD_BUCKET="xoxox-339339-21-x"

# Run stable.world commands
scurl http://httpbin.org/get
```

### Commands

#### scurl

Wrapper around `curl`

#### spip

Wrapper around `pip`
