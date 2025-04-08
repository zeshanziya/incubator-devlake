#!/bin/bash
export LD_LIBRARY_PATH=$LD_LIBRARY_PATH:/usr/local/lib
export DEVLAKE_PLUGINS=bamboo,bitbucket,circleci,customize,dora,gitextractor,github,github_graphql,gitlab,jenkins,jira,org,pagerduty,refdiff,slack,sonarqube,trello,webhook
exec "$SHELL"
