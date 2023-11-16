variable image_id {
    type = string
    default = "fd80hokdubc6pj50kvsd"
}

variable user {
    type = string
    default = "centos"
}

variable nodes {
    type = map(object({
        nat = bool
    }))
    default = {
        server = {
            nat = true
        }
        node1 = {
            nat = false
        }
        node2 = {
            nat = false
        }
    }
}