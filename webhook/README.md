## Go webserver to handle GITHUB actions / Webhookup

Commands to run on webserver
 

  1. Navigate to cd /webhook/DebianWebServer/
  2. pull from latest, git pull
  3. Build
    go build
  4.  run process 
     nohup ./webhook &
  5. Verify that it is running 
     ps aux | grep webhook
