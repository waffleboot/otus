variable image_id {
    type = string
    default = "fd80hokdubc6pj50kvsd"
}

variable ssh_user {
    type = string
    default = "centos"
}

variable nginx_load_balancer {
    type = object({
        addr = string
        port = number
    })
    default = {
        addr = "192.168.0.254"
        port = 80
    }
}

variable postgresql_load_balancer {
    type = object({
        addr = string
        port = number
    })
    default = {
        addr = "192.168.0.251"
        port = 5432
    }
}

variable nodes {
    type = map(object({
        ip_address = optional(string,null)
    }))
    default = {
        bastion = {
        }
        nginx-1 = {
            ip_address = "192.168.0.21"
        }
        nginx-2 = {
            ip_address = "192.168.0.22"
        }
        backend-1 = {
        }
        backend-2 = {
        }
        db-1 = {
            ip_address = "192.168.0.41"
        }
        db-2 = {
            ip_address = "192.168.0.42"
        }
        db-3 = {
            ip_address = "192.168.0.43"
        }
    }
}
