terraform {
  required_providers {
    yandex = {
      source = "yandex-cloud/yandex"
    }
  }
  required_version = ">= 0.13"
}

provider "yandex" {
  zone = "ru-central1-a"
}

resource "yandex_vpc_network" "nginx" {}

resource "yandex_vpc_subnet" "nginx" {
  network_id     = "${yandex_vpc_network.nginx.id}"
  v4_cidr_blocks = ["192.168.0.0/16"]
  zone           = "ru-central1-a"
}

resource "yandex_compute_instance" "nginx" {
  resources {
    cores = 2
    memory = 4
  }
  boot_disk {
    initialize_params {
        image_id = "fd826honb8s0i1jtt6cg"
    }
  }
  network_interface {
    subnet_id = "${yandex_vpc_subnet.nginx.id}"
    nat = true
  }
  metadata = {
    ssh-keys = "ubuntu:${file("id_rsa.pub")}"
  }
}

resource "local_file" "inventory" {
    filename = "inventory.ini"
    content = templatefile("inventory.ini.tftpl", { ip_addr = yandex_compute_instance.nginx.network_interface.0.nat_ip_address })
}
