---
Status: Published
Author: Steven Thurgood
Date: 2018-02-19
Reviewers: David Ferguson
Approvers: Betsy Beyer
Approval Date: 2018-02-20
Revisit Date: 2019-02-01
---

# Example SLO Document

> This document describes the SLOs for the Example Game Service.

## Service Overview

The Example Game Service allows Android and iPhone users to play a game with each other. The app runs on users’ phones, and moves are sent back to the API via a REST API. The data store contains the states of all current and previous games. A score pipeline reads this table and generates up-to-date league tables for today, this week, and all time. League table results are available in the app, via the API, and also on a public HTTP server.

The SLO uses a four-week rolling window.

## SLIs and SLOs

| Category           | SLI                                                                                                                                  | SLO                                                                       |
| ------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------- |
| **API**            |                                                                                                                                      |                                                                           |
| Availability       | count of "api" http_requests which do not have a 5XX status code divided by count of all "api" http_requests                         | 97% success                                                               |
| Latency            | count of "api" http_requests with a duration less than or equal to 400 milliseconds divided by count of all "api" http_requests      | 90% of requests < 400 ms                                                  |
|                    | count of "api" http_requests with a duration less than or equal to 850 milliseconds divided by count of all "api" http_requests      | 99% of requests < 850 ms                                                  |
| **HTTP server**    |                                                                                                                                      |                                                                           |
| Availability       | count of "web" http_requests which do not have a 5XX status code divided by count of all "web" http_requests                         | 99% success                                                               |
| Latency            | count of "web" http_requests with a duration less than or equal to 200 milliseconds divided by count of all "web" http_requests      | 90% of requests < 200 ms                                                  |
|                    | count of "web" http_requests with a duration less than or equal to 1,000 milliseconds divided by count of all "web" http_requests    | 99% of requests < 1,000 ms                                                |
| **Score pipeline** |                                                                                                                                      |                                                                           |
| Freshness          | count of all data_requests for "api" and "web" with freshness less than or equal to 1 minute divided by count of all data_requests   | 90% of reads use data written within the previous 1 minute.               |
|                    | count of all data_requests for "api" and "web" with freshness less than or equal to 10 minutes divided by count of all data_requests | 99% of reads use data written within the previous 10 minutes.             |
| Correctness        | count of all data_requests which were correct divided by count of all data_requests                                                  | 99.99999% of records injected by the prober result in the correct output. |
| Completeness       | count of all pipeline runs that processed 100% of the records divided by count of all pipeline runs                                  | 99% of pipeline runs cover 100% of the data.                              |

## Rationale

Availability and latency SLIs were based on measurement over the period 2018-01-01 to 2018-01-28. Availability SLOs were rounded down to the nearest 1% and latency SLO timings were rounded up to the nearest 50 ms. All other numbers were picked by the author and the services were verified to be running at or above those levels.

No attempt has yet been made to verify that these numbers correlate strongly with user experience.1

## Error Budget

Each objective has a separate error budget, defined as 100% minus (–) the goal for that objective. For example, if there have been 1,000,000 requests to the API server in the previous four weeks, the API availability error budget is 3% (100% – 97%) of 1,000,000: 30,000 errors.

We will enact the error budget policy when any of our objectives has exhausted its error budget.

## Clarifications and Caveats

- Request metrics are measured at the load balancer. This measurement may fail to accurately measure cases where user requests didn’t reach the load balancer.
- We only count HTTP 5XX status messages as error codes; everything else is counted as success.
- The test data used by the correctness prober contains approximately 200 tests, which are injected every 1s. Our error budget is 48 errors every four weeks.
