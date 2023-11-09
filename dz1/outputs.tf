output "public_ip" {
  value = yandex_compute_instance.nginx.network_interface.0.nat_ip_address
}
