# ga-dtfy

A Github Action helping you scan whatever you ship with Detectify, by triggering a [Deep Scan](https://detectify.com/product/deep-scan). Currently works for web applications.

# Instructions

1. Sign up to Detectify, add your asset and validate the ownership of your application domain.
2. Generate a Detectify API key 
3. Create a scan profile for your application and fetch its token.
4. Trigger a Deep Scan with your GA workflow. 

An example follows - take it as a very simple example which can be adapted as needed to fit your scanning needs:

```yaml
name: "Security scan"
on: [ push ]

jobs:
  run_scan:
    runs-on: ubuntu-latest
    name: A job to start a Detectify scan
    steps:
      - name: Start scan
        uses: koslib/ga-dtfy@master
        id: start_scan
        with:
          api_key: ${{ secrets.DETECTIFY_API_KEY }}
          scan_profile_token: ${{ secrets.DETECTIFY_SCAN_PROFILE_TOKEN }}
      - name: Start scan result
        run: echo "Detectify responded with ${{ steps.start_scan.outputs.result }}"

```

The example workflow file can be found also [here](.github/workflows/example.yml) which you can add into your GA config.

> Note: an empty (blank-string) response but no error is a good sign! The scan has been triggered!

# Config

This Github Action requires two secrets set in your repository:

1. `DETECTIFY_API_KEY`: your Detectify API key, the one you generated at step #2 of the instructions above.
2. `DETECTIFY_SCAN_PROFILE_TOKEN`: your scan profile token, which you fetched at step #3 of the instructions above.
