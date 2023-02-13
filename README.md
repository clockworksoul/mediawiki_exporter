# Mediawiki Exporter
The Mediawiki exporter is a [Prometheus exporter](https://prometheus.io/docs/instrumenting/exporters/) for [Mediawiki](https://www.mediawiki.org/wiki/MediaWiki). It reports on data such as the number or users, number of edits, and the number of [pages and articles](https://www.mediawiki.org/wiki/Manual:Article_count).

The container image for the exporter is available at [clockworksoul/mediawiki-exporter](https://hub.docker.com/repository/docker/clockworksoul/mediawiki-exporter/).

## Metrics

Emits the following metrics so far:

* `mediawiki_statistics_activeusers`: Current number of active users.
* `mediawiki_statistics_admins`: Current number of administrators.
* `mediawiki_statistics_articles`: Current number of articles.
* `mediawiki_statistics_edits`: Current number of edits.
* `mediawiki_statistics_images`: Current number of images.
* `mediawiki_statistics_jobs`: Current number of jobs.
* `mediawiki_statistics_pages`: Current number of pages.
* `mediawiki_statistics_users`: Current number of users.

Because all of these values can, in theory, decrease (such as when a resource is deleted), all of the above use gauges. In practice, however, significant downward movements should be rare, so Prometheus functions like `rate()` should still be useful.

## Execution

To execute this exporter, you must specify the URL of the Mediawiki instance that's being evaluated. This can be one of two ways.

1. As the only argument parameter.
2. Via the `MEDIAWIKI_API_URL`.

## Environment variables

This exporter supports three optional environment variables.

1. `MEDIAWIKI_API_URL` -- Specifies the URL of the evaluated Mediawiki instance.
3. `MEDIAWIKI_USERNAME`  -- The [bot username](https://www.mediawiki.org/wiki/Manual:Bot_passwords) (_not_ a user username), if authentication is required.
2. `MEDIAWIKI_PASSWORD` -- The [bot password](https://www.mediawiki.org/wiki/Manual:Bot_passwords) (_not_ a user password), if authentication is required.

## Exit status codes

Different failure modes return different status codes.

* `2`: General Mediawiki login failure
* `3`: API URL not specified
* `4`: Bot username/password are could not be authenticated
* `5`: Access denied by Mediawiki
