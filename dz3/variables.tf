variable image_id {
    type = string
    default = "fd81v7g3b2g481h03tsp"
}

variable ssh_user {
    type = string
    default = "almalinux"
}

variable load_balancer {
    type = object({
        addr = string
        port = number
    })
    default = {
        addr = "192.168.0.254"
        port = 80
    }
}

variable nodes {
    type = map(object({
        ip_address = string
    }))
    default = {
        bastion = {
            ip_address = "192.168.0.10"
        }
        nginx-1 = {
            ip_address = "192.168.0.21"
        }
//      nginx-2 = {
//          ip_address = "192.168.0.22"
//      }
        backend-1 = {
            ip_address = "192.168.0.31"
        }
//      backend-2 = {
//          ip_address = "192.168.0.32"
//      }
        db = {
            ip_address = "192.168.0.41"
        }
    }
}
