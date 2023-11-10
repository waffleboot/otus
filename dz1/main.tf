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

resource yandex_vpc_network nginx {}

resource yandex_vpc_subnet nginx {
  network_id     = yandex_vpc_network.nginx.id
  v4_cidr_blocks = ["192.168.0.0/16"]
  zone           = "ru-central1-a"
}

resource yandex_compute_instance nginx {
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
    subnet_id = yandex_vpc_subnet.nginx.id
    nat = true
  }
  metadata = {
    ssh-keys = "ubuntu:${tls_private_key.key.public_key_openssh}"
  }
  provisioner remote-exec {
    inline = ["true"]
    connection {
      type = "ssh"
      user = "ubuntu"
      host = self.network_interface.0.nat_ip_address
      private_key = tls_private_key.key.private_key_pem
    }
  }
  provisioner local-exec {
    command = "ansible-playbook -u ubuntu -i '${self.network_interface.0.nat_ip_address},' --private-key id_rsa playbook.yml"
    environment = {
      ANSIBLE_HOST_KEY_CHECKING: "False"
    }
  }
  provisioner local-exec {
    command = "curl http://${self.network_interface.0.nat_ip_address}"
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
