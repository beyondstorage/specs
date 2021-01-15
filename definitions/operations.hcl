
interface "appender" {
  description = "is the interface for Append related operations."

  op "create_append" {
    description = "will create an append object."
    params      = ["path"]
    results     = ["o"]
  }
  op "write_append" {
    description = "will append content to an append object."
    params      = ["o", "r", "size"]
    results     = ["n"]
  }
}
interface "blocker" {
  description = "is the interface for Block related operations."

  op "create_block" {
    description = "will create a new block object."
    params      = ["path"]
    results     = ["o"]
  }
  op "write_block" {
    description = "will write content to a block."
    params      = ["o", "r", "size", "bid"]
    results     = ["n"]
  }
  op "combine_block" {
    description = "will combine blocks into an object."
    params      = ["o", "bids"]
  }
  op "list_block" {
    description = "will list blocks belong to this object."
    params      = ["o"]
    results     = ["bi"]
  }
}
interface "copier" {
  description = "is the interface for Copy."

  op "copy" {
    description = "will copy an Object or multiple object in the service."
    params      = ["src", "dst"]
  }
}
interface "fetcher" {
  description = "is the interface for Fetch."

  op "fetch" {
    description = "will fetch from a given url to path."
    params      = ["path", "url"]
  }
}
interface "mover" {
  description = "is the interface for Move."

  op "move" {
    description = "will move an object in the service."
    params      = ["src", "dst"]
  }
}
interface "multiparter" {
  description = "is the interface for Multipart related operations."

  op "create_multipart" {
    description = "will create a new multipart."
    params      = ["path"]
    results     = ["o"]
  }
  op "write_multipart" {
    description = "will write content to a multipart."
    params      = ["o", "r", "size", "index"]
    results     = ["n"]
  }
  op "complete_multipart" {
    description = "will complete a multipart upload and construct an Object."
    params      = ["o", "parts"]
  }
  op "list_multipart" {
    description = "will list parts belong to this multipart."
    params      = ["o"]
    results     = ["pi"]
  }
}
interface "pager" {
  description = "is the interface for Page related operations which support random write."

  op "create_page" {
    description = "will create a new page object."
    params      = ["path"]
    results     = ["o"]
  }
  op "write_page" {
    description = "will write content to specific offset."
    params      = ["o", "r", "size", "offset"]
    results     = ["n"]
  }
}
interface "reacher" {
  description = "is the interface for Reach."

  op "reach" {
    description = "will provide a way, which can reach the object."
    params      = ["path"]
    results     = ["url"]
  }
}
interface "servicer" {
  description = "can maintain multipart storage services."

  op "create" {
    description = "will create a new storager instance."
    params      = ["name"]
    results     = ["store"]
  }
  op "delete" {
    description = "will delete a storager instance."
    params      = ["name"]
  }
  op "get" {
    description = "will get a valid storager instance for service."
    params      = ["name"]
    results     = ["store"]
  }
  op "list" {
    description = "will list all storager instances under this service."
    results     = ["sti"]
  }
}
interface "statistician" {
  description = "is the interface for Statistical."

  op "statistical" {
    description = "will count service's statistics, such as Size, Count."
    results     = ["statistic"]
  }
}
interface "storager" {
  description = "is the interface for storage service."

  op "delete" {
    description = "will delete an Object from service."
    params      = ["path"]
  }
  op "metadata" {
    description = "will return current storager metadata."
    results     = ["meta"]
  }
  op "list" {
    description = "will return list a specific path."
    params      = ["path"]
    pairs       = ["list_mode"]
    results     = ["oi"]
  }
  op "read" {
    description = "will read the file's data."
    params      = ["path", "w"]
    pairs       = ["size", "offset", "read_callback_func"]
    results     = ["n"]
  }
  op "stat" {
    description = "will stat a path to get info of an object."
    params      = ["path"]
    results     = ["o"]
  }
  op "write" {
    description = "will write data into a file."
    params      = ["path", "r", "size"]
    pairs       = ["storage_class", "content_type", "content_md5", "read_callback_func"]
    results     = ["n"]
  }
}

field "bi" {
  type = "*BlockIterator"
}
field "bid" {
  type = "string"
}
field "bids" {
  type = "[]string"
}
field "dst" {
  type = "string"
}
field "err" {
  type = "error"
}
field "index" {
  type = "int"
}
field "meta" {
  type = "*StorageMeta"
}
field "n" {
  type = "int64"
}
field "name" {
  type = "string"
}
field "o" {
  type = "*Object"
}
field "offset" {
  type = "int64"
}
field "oi" {
  type = "*ObjectIterator"
}
field "pairs" {
  type = "...Pair"
}
field "parts" {
  type = "[]*Part"
}
field "path" {
  type = "string"
}
field "pi" {
  type = "*PartIterator"
}
field "r" {
  type = "io.Reader"
}
field "seg" {
  type = "Segment"
}
field "si" {
  type = "*SegmentIterator"
}
field "size" {
  type = "int64"
}
field "src" {
  type = "string"
}
field "statistic" {
  type = "*StorageStatistic"
}
field "sti" {
  type = "*StoragerIterator"
}
field "store" {
  type = "Storager"
}
field "url" {
  type = "string"
}
field "w" {
  type = "io.Writer"
}
