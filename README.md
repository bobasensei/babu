# babu

Crawl wikipedia articles using the [Article Lookup API](https://enterprise.wikimedia.com/docs/on-demand/#article-lookup)
and the [Article Structured Contents API](https://enterprise.wikimedia.com/docs/on-demand/#article-structured-contents-beta)
and save the results in a PostgreSQL database.

## Setup 

First create a database by running this in your PostgreSQL instance, replacing USERNAME and PASSWORD with your desired values:
```
# create user USERNAME with encrypted password 'PASSWORD';
# create database babu with owner = 'USERNAME';
```

Set an environment variable to point to your database:
```
export BABU_DATABASE=postgres://USERNAME:PASSWORD@HOST:PORT/babu
```

Set an environment variable with your Wikimedia API token:
```
export BABU_WIKIMEDIA=...
```

Now use `babu` to build a local collection of structured articles.
```
babu init
babu fetch "Desi_hip_hop"
babu fetch "Genda_Phool"
babu fetch "Badshah_(rapper)"
babu fetch "Aastha_Gill"
babu fetch "Ek_Tha_Raja_(album)"
```

List the contents of your collection with `babu list` and get individual articles with `babu get`.

To use the Structured Contents API, add the `--structured-contents` flag to your `fetch`, `get`, and `list` commands.
