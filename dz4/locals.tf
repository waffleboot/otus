locals {
  bastion = yandex_compute_instance.my-instance["bastion"].network_interface.0.nat_ip_address
}
