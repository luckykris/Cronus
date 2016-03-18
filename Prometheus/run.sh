rm -rf src/github.com/luckykris/Cronus/Prometheus/db/
rm -rf src/github.com/luckykris/Cronus/Prometheus/cfg/
cp -R db src/github.com/luckykris/Cronus/Prometheus/
cp -R cfg src/github.com/luckykris/Cronus/Prometheus/
go run prometheus.go -config prometheus.toml
