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
    }))
    default = {
        bastion = {
            nat = true
        }
//      nginx-1 = {
//          nat = false
//      }
//      nginx-2 = {
//          nat = false
//      }
//      backend-1 = {
//          nat = false
//      }
//      backend-2 = {
//          nat = false
//      }
        db = {
            nat = false
        }
    }
}
