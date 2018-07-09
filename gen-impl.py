#!/usr/bin/env python
# encoding: utf-8
import sys
from google.protobuf.compiler.plugin_pb2 import CodeGeneratorRequest, CodeGeneratorResponse
import google.protobuf.descriptor_pb2 as descriptor

FieldD = descriptor.FieldDescriptorProto

request = CodeGeneratorRequest()
request.ParseFromString(sys.stdin.read())
response = CodeGeneratorResponse()

for file in request.proto_file:
    out = ''
    out += str(FieldD)
    out += """package io.devision.gen;\n\n"""
    out += """import java.util.Date;\n"""
    out += """import javax.persistence.Column;\n"""
    out += """import javax.persistence.Entity;\n"""
    out += """import javax.persistence.Id;\n"""
    out += """import javax.persistence.Table;\n"""
    # for enum in file.enum_type:
    #     out += 'enum ' + enum.name + '\n'
    #     for value in enum.value:
    #         out += '\t' + value.name + '\n'

    type_name = None
    for message in file.message_type:
        type_name = message.name
        if type_name == 'Column':
            continue

        out += """\n@Entity\n"""
        out += """@Table(name = "{tbl_name}")\n""".format(tbl_name=message.name)
        out += """public class {type_name} """.format(type_name=type_name)
        out += """ {\n\n"""

        for field in message.field:
            out += """    @Column(name = "{fname}")\n""".format(fname=field.name)
            out += """    private {lang_type} {fname};\n""".format(fname=field.json_name, lang_type={
                "TYPE_STRING": "String",
                "TYPE_INT32": "Integer",
            }.get(str(field.type), "String"))
            out += str(field.options.__str__())
            for x in field.options:
                out += """!"""

            out += """\n"""

    out += """}\n\n"""

    f = response.file.add()
    f.name = type_name + '.java'
    f.content = out
sys.stdout.write(response.SerializeToString())
