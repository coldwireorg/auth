#/bin/sh
export DSN=sqlite:///tmp/db.sqlite?_fk=true
export SECRETS_SYSTEM=ThisIsJustASuperToken!

hydra migrate sql -e --yes -c ./hydra-dev.yaml
hydra serve all -c ./hydra-dev.yaml --dangerous-force-http