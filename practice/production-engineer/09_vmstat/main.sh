#!/usr/bin/env bash

echo "Mach Virtual Memory Statistics: (page size of 4096 bytes)
free	active	inac	wire	faults	copy	zerofill	reactive	pageins	pageout
49314	97619	154001	26746	4236	41	3208	24148787	145786	50308	3459
49341	97814	153974	26551	2902	21	1429	0	0	0" \
  | go run .
