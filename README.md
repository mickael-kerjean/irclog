
- search: keyword:
  SELECT created, author, message FROM Message WHERE 
  related_channel = '$1' AND ( author = '$2' OR message MATCH '$2' );
  
- pull data with constant paging before cetain timestamp
  SELECT created, author, message FROM Message WHERE 
     related_channel = '$1' AND created < '$2' 
     ORDER BY timestamp DESC


bot: invite
===========

1) someone invite the bot
2) bot sendout URL of the log in the public channel
