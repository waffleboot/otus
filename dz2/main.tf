terraform {
  required_providers {
    yandex = {
      source = "yandex-cloud/yandex"
    }
  }
  required_version = ">= 0.13"
}

provider yandex {
  zone = "ru-central1-a"
}

resource yandex_vpc_gateway dz2 {}

resource yandex_vpc_network dz2 {}

resource yandex_vpc_subnet dz2 {
  network_id     = yandex_vpc_network.dz2.id
  v4_cidr_blocks = ["192.168.0.0/16"]
  zone           = "ru-central1-a"
  route_table_id = yandex_vpc_route_table.dz2.id
}

resource yandex_vpc_route_table dz2 {
  network_id = yandex_vpc_network.dz2.id
  static_route {
    destination_prefix = "0.0.0.0/0"
    gateway_id         = yandex_vpc_gateway.dz2.id
  }
}

resource yandex_compute_instance server {
  resources {
    cores = 2
    memory = 4
  }
  boot_disk {
    initialize_params {
        image_id = var.image_id
    }
  }
  network_interface {
    subnet_id = yandex_vpc_subnet.dz2.id
    nat = true
  }
  hostname = "server"
  metadata = {
    ssh-keys = "${var.user}:${tls_private_key.key.public_key_openssh}"
  }
  /*provisioner remote-exec {
    inline = ["true"]
    connection {
      type = "ssh"
      user = var.user
      host = self.network_interface.0.nat_ip_address
      private_key = tls_private_key.key.private_key_pem
    }
  }
  provisioner local-exec {
    command = "ansible-playbook -i '${self.network_interface.0.nat_ip_address},' playbook.yml"
    environment = {
      ANSIBLE_REMOTE_USER: var.user
      ANSIBLE_PRIVATE_KEY_FILE: "id_rsa"
      ANSIBLE_HOST_KEY_CHECKING: "False"
    }
  }*/
}

resource yandex_compute_instance client {
  resources {
    cores = 2
    memory = 4
  }
  boot_disk {
    initialize_params {
        image_id = var.image_id
    }
  }
  network_interface {
    subnet_id = yandex_vpc_subnet.dz2.id
    nat = false
  }
  hostname = "client"
  metadata = {
    ssh-keys = "${var.user}:${tls_private_key.key.public_key_openssh}"
  }
}

resource tls_private_key key {
  algorithm = "RSA"
  rsa_bits = 4096
  provisioner local-exec {
    command = "echo '${self.private_key_pem}' > id_rsa"
  }
  provisioner local-exec {
    command = "chmod og-rwx id_rsa"
  }
}