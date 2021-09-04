package config

import "time"

const NumberOfTopEntriesToConsider = 10
const NumberOfLastEntriesToConsiderWhenSearchingByTag = 50
const Salt = "eoriu69045trhgeo4t5780ejr9t78340gjtu9eu9034oij490et74564564asdas8dgyas8d68743r584frth437ry49r58478r5984r984rewhf8943hf349yfg34921eu23ru3efr4eftu43t784tf980eufg8u4f4uft"
const JwtExpirationPeriod = time.Hour * 6
const CachePeriod = 1800 // 30 minutes
const CommonCacheTag = "default"
const User = ""
const DbName = ""
const PostgresConfig = "postgres://username@localhost:5432/postgres?sslmode=disable" //update creds as needed
