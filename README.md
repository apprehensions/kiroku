# kiroku

Discord user/bot program to retrieve all messages in a channel in plain text form.

Showcase:
```
[1088721562447917116] 2023-03-24 07:10:47 sewn (1062027356954439762): [started a call]
[1179759531316752454] 2023-11-30 12:23:11 sewn (1062027356954439762): hi
[1179764995484045422] 2023-11-30 12:44:53 sewn (1062027356954439762): [attachment: org.vinegarhq.Vinegar.png https://cdn.discordapp.com/attachments/1088100535623745588/1179764995203018843/org.vinegarhq.Vinegar.png ][reaction: 1 🐱]
[1179766266504941568] 2023-11-30 12:49:56 sewn (1062027356954439762): [sticker: Happy 796140620111544330]

[1179769629120401560] 2023-11-30 13:03:18 sewn (1062027356954439762): hi [edited at 2023-11-30 13:03:18]
```
The messages are retrieved in reverse order, meaning that the last message sent will be the first.
To make the last message the last in the output, you can pipe kiroku to `tac`.
