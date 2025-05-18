create private game
```js
POST /games/create?private=true
{
    "Title": ...,
    "MaxPlayers": ...,
    "GameMode": ...,
}
```

Join private game
```js
WS /game?code=12346
```

<!-- ```js
POST /games/join?private=true&code=12346
``` -->

JOIN QUEUE
```js
WS /queue
```

JOIN GAME
```js
WS /game?id=andfz156fez4iz
```