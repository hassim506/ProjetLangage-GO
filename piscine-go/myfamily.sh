curl https://learn.zone01dakar.sn/assets/superhero/all.json | jq ' . [] | select ( .id == '${HERO_ID}' )' | jq '.connections ' | jq  '.relatives' | sed -e 's/"//g'
