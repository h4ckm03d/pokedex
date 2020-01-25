# pokedex

Demonstration simple aws lambda using golang and viewer using svelte

## Setup Guide

### Lambda

- Set `AWS_ROLE` env var, `export AWS_ROLE=<ur aws role>`
- Enter directory `lambda`, `cd lambda`
- `make build` - build and compress executable golang
- `make deploy` - deploy function for the first time
- `make redeploy` - update function if

### Front end

- create S3 bucket, on this example using `pokedextcg`. Please rename Makefile if using different name
- Enter dir `frontend`, `cd frontend`
- run `make deployfe`

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## Credits

All assets, css, and source API from https://pokemontcg.io