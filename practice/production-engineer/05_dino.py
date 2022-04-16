#!/usr/bin/env python3

import petl as etl
import math

table1 = etl.fromcsv("05_dino1.csv")
table2 = etl.fromcsv("05_dino2.csv")

table = etl.join(table1, table2, key="NAME")
out = etl.select(table, lambda rec: rec["STANCE"] == "bipedal")

mappings = {}
mappings["NAME"] = "NAME"
mappings["DIET"] = "DIET"
mappings["STANCE"] = "STANCE"
mappings["LEG_LENGTH"] = "LEG_LENGTH"
mappings["STRIDE_LENGTH"] = "STRIDE_LENGTH"

# Given the following formula, speed = ((STRIDE_LENGTH / LEG_LENGTH) - 1) * SQRT(LEG_LENGTH * g)
# Where g = 9.8 m/s^2 (gravitational constant)
mappings["SPEED"] = lambda rec: ((float(rec["STRIDE_LENGTH"]) / float(rec["LEG_LENGTH"])) - 1.0) * math.sqrt(float(rec["LEG_LENGTH"]) * 9.8)

out = etl.fieldmap(out, mappings)
print(out)
