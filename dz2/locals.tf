locals {
  client_ip = yandex_compute_instance.client.network_interface.0.ip_address
  server_ip = yandex_compute_instance.server.network_interface.0.nat_ip_address
}
