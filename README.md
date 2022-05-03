# mariomenjr/go

A redirect server based on Go and MongoDB.

## Config

You'll need to provide the following variables in a .env file.

```env
PORT=1993

MONGO_USER=
MONGO_PASS=
MONGO_DB=
MONGO_HOST=
MONGO_PORT=
```

## Database

Once you are connected to a mongodb database, you can start creating a redirect endpoint using the following format:

```mongodb
{
    "_id" : ObjectId("626784992e827293a30b6614"),
    "created_at" : ISODate("2022-04-26T04:06:05.262Z"),
    "updated_at" : ISODate("2022-04-26T04:06:08.099Z"),
    "title" : "Mario's YouTube",
    "url" : "http://youtube.com/user/mariomenjr",
    "id" : "youtube"
}
```

## Run

In `development`:

```bash
sh ./setup.sh development
```

In `production`:

```bash
sh ./setup.sh production
```

## Example

[go.mariomenjr.com/github](https://go.mariomenjr.com/github)

## License 
The source code of this project is under [MIT License](https://opensource.org/licenses/MIT).
