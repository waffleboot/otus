variable image_id {
    type = string
    default = "fd81v7g3b2g481h03tsp"
}

variable ssh_user {
    type = string
    default = "almalinux"
}

variable nodes {
    type = map(object({
        nat = bool
        ip_address = string
    }))
    default = {
        bastion = {
            nat = true
            ip_address = "192.168.0.20"
        }
        nginx-1 = {
            nat = false
            ip_address = "192.168.0.21"
        }
        nginx-2 = {
            nat = false
            ip_address = "192.168.0.22"
        }
//      backend-1 = {
//          nat = false
//      }
//      backend-2 = {
//          nat = false
//      }
//      db = {
//          nat = false
//      }
    }
}
