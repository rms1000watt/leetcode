#!/usr/bin/env bash

# 4.read innate file and parse the strings to count how many times an email address is found

grep "bbb@gmail.com" 04_emails.txt | uniq -c
