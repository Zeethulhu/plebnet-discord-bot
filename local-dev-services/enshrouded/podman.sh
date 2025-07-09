podman pod create \
  --name emulated-enshrouded-logs \
  --replace \
  -p 8687:8686 \
  -p 4223:4222 \
  -p 8222:8222 \
  -p 6222:6222 \
  -v $PWD/vector.yaml:/etc/vector/vector.yaml:ro,Z

podman run \
  --name emulated-enshrouded-nats \
  --pod emulated-enshrouded-logs \
  --replace \
  -d \
  nats:latest

podman run \
  --name emulated-enshrouded-vector \
  --pod emulated-enshrouded-logs \
  --replace \
  -d \
  timberio/vector:0.48.0-debian

podman container logs emulated-enshrouded-vector -f

podman pod stop --ignore emulated-enshrouded-logs
