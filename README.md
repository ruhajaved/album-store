# Vinyl Album Store

## Description
This is a go project implementing a vinyl album store. You can add albums and also lookup albums.

## Setup Database
Create the DB by:
```
# Copy SQL script to temp since using default user to access DB.
cp AlbumStore.sql /tmp
# Create db and populate w/ data. This uses the default user "postgres".
sudo -u postgres psql -f /tmp/AlbumStore.sql
```

To look at the database:
```
sudo -u postgres psql album_store
```