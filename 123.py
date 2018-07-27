from google.protobuf import json_format

import demo_pb2
from google.protobuf.json_format import MessageToJson

json_p = demo_pb2.Person()
json_format.Parse("""
{
  "email": "json@mail.ru", 
  "firstName": "asdasd"
}
""", json_p)
print(u"json_p = %s" % str(MessageToJson(json_p)))

p = demo_pb2.Person()
p.email = "asdasdasd"
p.first_name = "asdasd"
json_p = MessageToJson(p)
print(u"json_p = %s" % str(json_p))
