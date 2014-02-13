This is an experimental website/web server written in Go (golang). It employs the Tiedot (version 2.0) database system (which itself was written in Go) and the Martini web package.

# Dependencies

<b>go get github.com/codegangsta/martini</b>

<b>go get github.com/HouzuoGuo/tiedot</b>

#Configuration

Change the DATABASE_DIR setting in <b>/src/tiedotmartini2/config.go</b>, and make sure the new directory exists.

#Building and Running

Add this git's root directory to the GOPATH search path environment variable, then issue <b>go build main.go</b>. This will generate an executable (TiedotMartini2.exe in Windows). You will also need to issue <b>go build datainit.go</b> Run "datainit" before "main" in order for Tiedot 2 to work.

This website is pointed to <b>http://localhost:3000</b>. Set the PORT environment variable to change the port number. 
At the time this repostory was created, Tiedot was at version 2.0. Since 2.0 isn't compatible with 1.x-generated data, you might need to create a new data directory.      