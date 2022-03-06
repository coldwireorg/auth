#/bin/sh
export DSN=sqlite:///tmp/db.sqlite?_fk=true
export SECRETS_SYSTEM=ThisIsJustASuperToken!

hydra migrate sql -e --yes -c ./hydra-deva.yaml
hydra serve all -c ./hydra-dev.yaml --dangerous-force-http && \
hydra clients create \
    --id auth \
    --endpoint http://10.42.10.7:30200 \
    --grant-types authorization_code,refresh_token \
    --response-types code,id_token,profile \
    --scope openid,offline,profile \
    --callbacks https://auth.coldwire.org/api/callback \
    --token-endpoint-auth-method none