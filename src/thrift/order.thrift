
namespace go mall.order

/**
 * Thrift lets you do typedefs to get pretty names for your types. Standard
 * C style here.
 */
typedef i32 int

/**
 * Ahh, now onto the cool part, defining a service. Services just need a name
 * and can optionally inherit from another service using the extends keyword.
 */
service OrderManager  {

   void ping(),

   bool createOrder(1:int goodsId, 2:int userId)

}


