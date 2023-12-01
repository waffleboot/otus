locals {
  bastion = yandex_compute_instance.dz3["bastion"].network_interface.0.nat_ip_address
}
