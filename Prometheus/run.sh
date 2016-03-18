CURRENT=/tmp
rm -rf $CURRENT/src/github.com/luckykris/Cronus/Prometheus/db/
rm -rf $CURRENT/src/github.com/luckykris/Cronus/Prometheus/cfg/
cp -R db $CURRENT/src/github.com/luckykris/Cronus/Prometheus/
cp -R cfg $CURRENT/src/github.com/luckykris/Cronus/Prometheus/
go run prometheus.go -config prometheus.toml
