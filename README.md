# Mediawiki Exporter
A Prometheus exporter for Mediawiki.

Emits the following metrics so far:

* `mediawiki_statistics_activeusers`: Current number of active users.
* `mediawiki_statistics_admins`: Current number of administrators.
* `mediawiki_statistics_articles`: Current number of articles.
* `mediawiki_statistics_edits`: Current number of edits.
* `mediawiki_statistics_images`: Current number of images.
* `mediawiki_statistics_jobs`: Current number of jobs.
* `mediawiki_statistics_pages`: Current number of pages.
* `mediawiki_statistics_users`: Current number of users.

The container image for the exporter is available at `clockworksoul/mediawiki-exporter`.
