#!/bin/bash

# Set variables
DB_NAME="persons"
BACKUP_DIR="$HOME/www/playground/go-fun/dumps"

# Create backup directory if it doesn't exist
mkdir -p $BACKUP_DIR

# Backup command
while true; do
  BACKUP_CMD="sqlite3 $DB_NAME '.backup $BACKUP_DIR/backup_$(date +%Y%m%d_%H%M%S).db'"

  # Create cron job to run backup every 3 seconds
  # (crontab -l 2>/dev/null; echo "*/3 * * * * $BACKUP_CMD") | crontab -

  # Run the backup command immediately
  eval $BACKUP_CMD

  sleep 3;
done;
