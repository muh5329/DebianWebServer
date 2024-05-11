# DebianWebServer
Debian web server utilities and program


# Current Start Up Process ( Manual )

Navigate to: 
	ssh muh5329@192.XXX.XX.XXX
Navigate to 
	/var/www/demo
Pull repo
	git pull
build static files
	pnpm build
stop existing node server
	pm2 stop demo
start new process of node with pm2
	pm2 start pnpm --name demo -- start
Delete old process
	pm2 delete 0