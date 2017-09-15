 #!/bin/bash
echo "Starting RethinkDB..."
rethinkdb --bind all --daemon&
echo "Starting backend..."
/backend/backend &
echo "Restoring demo data into RethinkDB..."
sleep 6
rethinkdb restore /rethinkdb/data.tar.gz
echo "Starting HTTP server..."
ruby -run -ehttpd /frontend -p8000
