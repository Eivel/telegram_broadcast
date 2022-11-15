## Telegram Broadcast
Telegram Broadcast is a simple webserver for publishing messages over REST API. It was created to work together with a Lua script inside CC: Tweaked Minecraft mod to provide bidirectional communication between a Telegram group and a Minecraft chat, without any external dependencies.

## Setup
1. Run the webserver on your VPS and expose the endpoints.
2. Run both `tele_reader.lua` and `tele_writer.lua` inside a ComputerCraft computer. Ensure that there's a chunk loader nearby. To handle server restarts, set the programs to run on the computer startup.
