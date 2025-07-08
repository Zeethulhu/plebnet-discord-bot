

```
docker run \
  -v $PWD/vector.yaml:/etc/vector/vector.yaml:ro \
  -p 8686:8686 \
  --name vector-enshrouded-log-emulator \
  timberio/vector:0.48.0-debian
```
