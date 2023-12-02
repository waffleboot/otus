variable image_id {
    type = string
    default = "fd80hokdubc6pj50kvsd"
}

variable ssh_user {
    type = string
    default = "centos"
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
