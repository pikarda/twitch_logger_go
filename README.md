# twitch_logger_go

App that logs message sent in twitch chat and send it to specific discord channel by webhook.

It uses the Notion database as a database with the information provided to run the application.

This is convenient for me, because with Notion I can manage the list of monitored streamers, change links to webhooks, add users that the program will ignore.
And all this without restarting the program.

---

There is 2 secrets inside .env file:

- **NOTION_TOKEN** - token of your integration app that has connection to your notion_database.

You can add new integration here https://www.notion.so/my-integrations

- **DB_ID** - ID of notion_database with all information for working logger.

There is database with 4 columns: _name_, _value_, _tag_, _color_.

"**tag**" column must be with type "_select_"

This names is case sensitive.

---

Example of table (Inline database)

| **name**          | **value**                                                              | **tag**   | **color** |
| ----------------- | ---------------------------------------------------------------------- | --------- | --------- |
| TOKEN             | qwertyuiopasdg123456                                                   | secret    |           |
| CLIENT_ID         | zxcvbnmlkjhg98765                                                      | secret    |           |
|                   |                                                                        |           |           |
| streamer1         | https://discord.com/api/v10/webhooks/1234567890/qwerty123456?wait=true | user      | #26FACF   |
| streamer2         | https://discord.com/api/v10/webhooks/1234567890/qwerty123456?wait=true | user      | #FA4000   |
| streamer3         | https://discord.com/api/v10/webhooks/1234567890/qwerty123456?wait=true | user      | #2C7DFA   |
|                   |                                                                        |           |           |
| blacklisted users | nightbot, moobot, streamelements                                       | blacklist |           |
|                   |                                                                        |           |           |
| reload            | 0                                                                      | reload    |           |

Colors may not be specified.

- **TOKEN** - oauth2 token that will be using for get profile picture of chat user.

- **CLIENT_ID** - client id that will be using for get profile picture of chat user.

You can use your own user secrets or register twitch developer app here https://dev.twitch.tv/console

---

Periodicaly app wil check value of "**reload**" row. If it's equals 1 then app will refetch data from database and update streamer list.

So if you want change something - do it and then put 1 in reload value.

Or add notion button that will do it for you.

---

![logger](https://github.com/pikarda/twitch_logger_go/assets/25252682/1fa76b0d-8c5c-4222-9c22-5c3f2bb6fc79)

