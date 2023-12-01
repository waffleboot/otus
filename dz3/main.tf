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

resource yandex_vpc_gateway dz3 {}

resource yandex_vpc_network dz3 {}

resource yandex_vpc_subnet dz3 {
  network_id     = yandex_vpc_network.dz3.id
  v4_cidr_blocks = ["192.168.0.0/16"]
  zone           = "ru-central1-a"
  route_table_id = yandex_vpc_route_table.dz3.id
}

resource yandex_vpc_route_table dz3 {
  network_id = yandex_vpc_network.dz3.id
  static_route {
    destination_prefix = "0.0.0.0/0"
    gateway_id         = yandex_vpc_gateway.dz3.id
  }
}

resource yandex_compute_instance dz3 {
  for_each = var.nodes
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
    subnet_id = yandex_vpc_subnet.dz3.id
    nat = each.value.nat
  }
  name = each.key
  hostname = each.key
  metadata = {
    ssh-keys = "${var.ssh_user}:${tls_private_key.key.public_key_openssh}"
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

resource local_file inventory-ini {
  content = templatefile("inventory.tftpl",{
    bastion  = local.bastion
    ssh_user = var.ssh_user
  })
  filename = "inventory.ini"
  provisioner remote-exec {
    inline = ["true"]
    connection {
      type = "ssh"
      user = var.ssh_user
      host = local.bastion
      private_key = tls_private_key.key.private_key_pem
    }
  }
}
