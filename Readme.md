This is a simple Golang Chat-Room built in order to learn how WebSockets and real-time communication takes place in Golang.
This program uses gorilla/websocket library for WecSockets, net/http for routing and http server and JavaScript on the frontend to handle in-coming messages and users.

-> What this Program does (atm):
1. It takes a username from the user
2. It displays the username in the "Who's Online" section, so that one could know which users are currently in the room
3. It takes a message, from the user- and displays it with the username in the chatbox below
4. If a person refresh, leave or close the page- it removes them from the Online section, thus removing them from aaccessing further chat
5. It , at the moment, does not include any Auth, Database. It is just a simple example to learn real-time communication in Golang.


-> minor bug: The Send Message Button works perfectly fine, but when the message is sent with Pressing ENTER (i.e. on PC) it sends the message alongside an empty string, thus displaying both the message and the empty string inside the chatbox.
I'm figuring out the bug and will update the repo once I do.

website-link: https://go-chat-production.up.railway.app/
