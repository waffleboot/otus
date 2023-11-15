output "server_ip" {
  value = yandex_compute_instance.server.network_interface.0.nat_ip_address
}

output "client_ip" {
  value = yandex_compute_instance.client.network_interface.0.ip_address
}
