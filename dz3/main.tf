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

resource yandex_vpc_network net {}

resource yandex_vpc_subnet subnet {
  network_id     = yandex_vpc_network.net.id
  v4_cidr_blocks = ["192.168.0.0/24"]
  zone           = "ru-central1-a"
  route_table_id = yandex_vpc_route_table.rt.id
}

resource yandex_vpc_gateway nat_gateway {}

resource yandex_vpc_route_table rt {
  network_id = yandex_vpc_network.net.id
  static_route {
    destination_prefix = "0.0.0.0/0"
    gateway_id         = yandex_vpc_gateway.nat_gateway.id
  }
}

resource yandex_lb_network_load_balancer load_balancer {
  type = "internal"
  listener {
    name = "nginx"
    port = 80
    internal_address_spec {
      subnet_id = yandex_vpc_subnet.subnet.id
      address = var.load_balancer_address
    }
  }
  attached_target_group {
    target_group_id = yandex_lb_target_group.target_group.id
    healthcheck {
      name = "http"
      http_options {
        port = 80
      }
    }
  }
}

resource yandex_lb_target_group target_group {
  target {
    subnet_id = yandex_vpc_subnet.subnet.id
    address = "192.168.0.21"
  }
  target {
    subnet_id = yandex_vpc_subnet.subnet.id
    address = "192.168.0.22"
  }
}

resource yandex_compute_instance my-instance {
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
    subnet_id = yandex_vpc_subnet.subnet.id
    ip_address = each.value.ip_address
    nat = each.key == "bastion"
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

output "bastion" {
  value = local.bastion
}
