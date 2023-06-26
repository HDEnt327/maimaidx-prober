# maimai-prober API Documentation

This is the documentation for the maimai-prober API

All API routes link to https://www.diving-fish.com/api/*

> e.g. to access maimai/player/profile head to https://www.diving-fish.com/api/maimaidxprober/player/profile

---

## maimai

**Path:**

`/maimaidxprober`

---

### Query: Player

**Path:**

`/query/player`

**Parameters:**

`"username" OR "qq": str`

`"b50": bool`

**Usage:**

Returns a list of scores in the player's best40 or best50 in decending order of rating.

### Query: Plate

**Path:**

`/query/plate`

**Parameters:**

`"username" OR "qq": str`

`"version": str[]`

**Usage:**

Returns information related to a player's plate for a version. For list of versions, refer to https://github.com/Diving-Fish/maimaidx-prober/blob/main/database/models/maimai.py.

### Player: Profile
