How to Create a Private Discord Bot 
-----------------------------------


1. Log into https://discord.com/developers/applications
2. New Application
3. Installation -> Install Link -> None
4. Bot -> Public Bot -> Off  (This means only **YOU** can add the Bot to a server)
5. Bot -> Message Content Intent -> Enable (To allow the Bot to read and react to messages)
6. OAuth2 -> Reset Secret (To obtain your Discord Token)
7. OAuth2 -> OAuth2 URL Generator 
  - Scopes -> Enable 'bot'
  - Bot Permissions -> Enable 'Send Messages'
  - Integration Type -> Select 'Guild Install'
  - Generated URL -> Copy
8. Open the URL and add to selected Discord server
