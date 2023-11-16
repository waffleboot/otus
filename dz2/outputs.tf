output server_ip {
  value = local.server_ip
}

output client_ip {
  value = local.client_ip
}

output ssh {
  value = "ssh ${var.user}@${local.client_ip} -oProxyCommand=\"ssh ${var.user}@${local.server_ip} -i id_rsa -W %h:%p\" -i id_rsa"
}
