## How to run

## NextJS server
```
    #Start docker daemon 
    dockerd
    #or
    sudo systemctl start docker

    # Build the image
    docker build -t nextjs-docker .


    # Run docker image
    docker run -p 3001:3000 nextjs-docker

```



docker compose up -d --no--deps  --scale server=2 --no-recreate server