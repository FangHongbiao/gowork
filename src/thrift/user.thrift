namespace go mall.user

/**
 * Thrift lets you do typedefs to get pretty names for your types. Standard
 * C style here.
 */
typedef i32 int

struct UserInfo {
  1: string username,
  2: int credits,
  3: int level,
  4: string sex,
  5: string phone,
  6: string address
}


exception UserNotExist {
  2: string why
}

exception PermissionDeny  {
  2: string why
}

service UserManager {

   void ping(),

   UserInfo getUserInfo(1:int userId) throws (1:UserNotExist ex)
}
