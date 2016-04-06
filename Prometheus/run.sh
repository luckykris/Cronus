source /etc/profile
CURRENT=/data/gopath
rm -rf $CURRENT/src/github.com/luckykris/Cronus/Prometheus/db/
rm -rf $CURRENT/src/github.com/luckykris/Cronus/Prometheus/cfg/
rm -rf $CURRENT/src/github.com/luckykris/Cronus/Prometheus/prometheus/
rm -rf $CURRENT/src/github.com/luckykris/Cronus/Prometheus/global/
rm -rf $CURRENT/src/github.com/luckykris/Cronus/Prometheus/http/
cp -R db $CURRENT/src/github.com/luckykris/Cronus/Prometheus/
cp -R cfg $CURRENT/src/github.com/luckykris/Cronus/Prometheus/
cp -R prometheus $CURRENT/src/github.com/luckykris/Cronus/Prometheus/
cp -R global $CURRENT/src/github.com/luckykris/Cronus/Prometheus/
cp -R http $CURRENT/src/github.com/luckykris/Cronus/Prometheus/
go run prometheus.go -config /etc/prometheus.toml
