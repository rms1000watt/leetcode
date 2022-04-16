#!/usr/bin/env bash

hosts=(
  'vagrant@127.0.0.1 -p 2222'
  'vagrant@127.0.0.1 -p 2222'
  'vagrant@127.0.0.1 -p 2222'
  'vagrant@127.0.0.1 -p 2222'
)

file="01_process_report.txt"
echo "REPORT: $(date)" > "${file}"

for host in "${hosts[@]}"; do
  echo "handling $host" >> "${file}"
  ssh $host 'bash -c "ps aux | grep docker"'  >> "${file}"
done

echo "Now go send an email with this attachment"