locals {
  public_ip  = yandex_compute_instance.dz2["server"].network_interface.0.nat_ip_address
  clients    = [for v in yandex_compute_instance.dz2: v.fqdn if v.name != "server"]
  server     = yandex_compute_instance.dz2["server"].fqdn
}
