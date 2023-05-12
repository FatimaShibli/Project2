#!/bin/sh

# This is a fake shell script that demonstrates some basic shell commands

echo "Welcome to my fake shell script!"

# Create a new directory
mkdir mydir

# Change to the new directory
cd mydir

# Create some empty files
touch file1.txt file2.txt file3.txt

# Print the contents of the current directory
ls

# Wait for user input
read -p "Press Enter to continue..."

# Remove the files and directory
rm file1.txt file2.txt file3.txt
cd ..
rmdir mydir

echo "Thanks for using my fake shell script!"
