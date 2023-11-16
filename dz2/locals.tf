locals {
  client_ips = [for v in yandex_compute_instance.dz2: v.network_interface.0.ip_address if v.name != "server"]
  server_ip  = yandex_compute_instance.dz2["server"].network_interface.0.nat_ip_address
  clients    = [for v in yandex_compute_instance.dz2: v.fqdn if v.name != "server"]
}
