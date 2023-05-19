connection "env0" {
  plugin = "env0"

  # For setting API key and secret see instructions at https://docs.env0.com/reference/authentication
  # `api_key`: The API key for the env0 account. (Required)
  # This can also be set via the `ENV0_API_KEY` environment variable. 
  # api_key = "asdpoblfth8acbd"

  # `api_secret`: The API secret of the env0 account. (Required)
  # This can also be set via the `ENV0_API_SECRET` environment variable. 
  # api_secret = "LjatOxDqNN9iKH1sLn14TojGkuH3GQAx"

  # If no credentials are specified, the plugin will use env0 CLI configuration, if it has been installed locally.
}
