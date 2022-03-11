#/bin/sh

# Hydra
export DSN=sqlite:///tmp/db.sqlite?_fk=true
export SECRETS_SYSTEM=ThisIsJustASuperToken!
export LOG_LEAK_SENSITIVE_VALUES=true

hydra migrate sql -e --yes -c ./hydra-dev.yaml
hydra serve all -c ./hydra-dev.yaml --dangerous-force-http