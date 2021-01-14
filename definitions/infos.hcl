
info "object" "meta" "content-md5" {
  type         = "string"
  display_name = "ContentMD5"
}
info "object" "meta" "content-type" {
  type = "string"
}
info "object" "meta" "etag" {
  type         = "string"
  display_name = "ETag"
}
info "object" "meta" "id" {
  type    = "string"
  export  = true
  comment = "ID is the unique key in storage."
}
info "object" "meta" "mode" {
  type   = "ObjectMode"
  export = true
}
info "object" "meta" "part-id" {
  type    = "string"
  comment = "PartID is the part id of part object."
}
info "object" "meta" "path" {
  type    = "string"
  export  = true
  comment = "Path is either the absolute path or the relative path towards storage's WorkDir depends on user's input."
}
info "object" "meta" "size" {
  type = "int64"
}
info "object" "meta" "storage-class" {
  type = "string"
}
info "object" "meta" "target" {
  type    = "string"
  comment = "Target is the symlink target for link object."
}
info "object" "meta" "updated-at" {
  type = "time.Time"
}
info "storage" "meta" "location" {
  type = "string"
}
info "storage" "meta" "name" {
  type   = "string"
  export = true
}
info "storage" "meta" "work-dir" {
  type   = "string"
  export = true
}
info "storage" "statistic" "count" {
  type = "int64"
}
info "storage" "statistic" "size" {
  type = "int64"
}
