import json
from pprint import pprint
import binascii
from collections import defaultdict

import glog_pb2


def serialize_and_deserialize():
    init_msg = glog_pb2.LogMessage(
        level=glog_pb2.INFO,
        service_name='encode test',
        message='test message',
        timestamp=11187097077,
        trace_id='default',
    )

    print(f'> Initial Message:\n{init_msg}')
    # serialize
    binary_req = init_msg.SerializeToString()
    print(f'> After Serialization: {binary_req}')
    print('>> Binary form:', ' '.join(f'{byte:08b}' for byte in binary_req))
    # check in hex string
    hex_req_str = binary_req.hex()
    beautify_hex_str = ' '.join(
        hex_req_str[i:i + 2] for i in range(0, len(hex_req_str), 2))  # group into hex-represented bytes
    print(f'>> Hex Representation: {beautify_hex_str}')
    print('>>> Trying to decode the serialized message...')
    try_decode_proto_binary(binary_req)
    # deserialize
    recoverd_msg = glog_pb2.LogMessage()  # empty initialization
    recoverd_msg.ParseFromString(binary_req)
    print(f'> Deserialized Message:\n{recoverd_msg}')


def try_decode_proto_binary(binary_str):
    '''
    A super simple decoder based on the following link:
    https://protobuf.dev/programming-guides/encoding
    '''
    print('-----------------------------------------------------------------------')
    binary_str = ''.join(f'{byte:08b}' for byte in binary_str)
    ptr = 0
    result = defaultdict()  # store everything in a more organized fashion (field_number => payload)
    # in each loop, we decode one record (a key-value pair from the message proto)
    # NOTE: Protobuf do not encode keys, they can be inferred from the proto
    while ptr < len(binary_str):
        record = {}
        # check record tag - field number + wire type
        """
        "The “tag” of a record is encoded as a varint formed from the field number 
        and the wire type via the formula (field_number << 3) | wire_type. 
        In other words, after decoding the varint representing a field, 
        the low 3 bits tell us the wire type, and the rest of the integer tells us the field number."
        """
        # NOTE: we are lazy so we just pick the first 5 bits.
        # when field numbers become large, we should not do this. Decode a varint properly instead.
        record['field_number'] = int(binary_str[ptr:ptr + 5], 2)
        record['wire_type'] = int(binary_str[ptr + 5:ptr + 8], 2)
        ptr += 8
        # TODO
        if record['wire_type'] == -1:
            # ???: https://protobuf.dev/programming-guides/encoding/#structure
            record['wire_type_name'] = 'TODO'
            '''
            Hints:
              1. https://docs.python.org/3/library/struct.html#struct.unpack
              2. https://docs.python.org/3/library/struct.html#byte-order-size-and-alignment
              3. https://docs.python.org/3/library/struct.html#format-characters
            '''
            record['payload'] = 'TODO'
            ptr += -1
            result[record['field_number']] = record['payload']
        # for this demo we only implement LEN decoding for string
        if record['wire_type'] == 2:
            # WIRE: https://protobuf.dev/programming-guides/encoding/#structure
            record['wire_type_name'] = 'LEN'
            """
            https://protobuf.dev/programming-guides/encoding/#length-types
            "The LEN wire type has a dynamic length, 
            specified by a varint immediately after the tag, 
            which is followed by the payload as usual."
            """
            # check length (NOTE: we are lazy so we just grab the next byte)
            record['length'] = int(binary_str[ptr:ptr + 8], 2)
            ptr += 8
            # check payload (the next `length` bytes are UTF-8 encoded chars)
            record_length = record['length']
            payload_str = f'0b{binary_str[ptr:ptr + 8 * record_length]}'
            ptr += 8 * record['length']
            record['payload'] = binascii.unhexlify('%x' % int(payload_str, 2)).decode('utf-8')
            result[record['field_number']] = record['payload']
        print(f'>>> Record: {record}')
    print('>>> Final Result:')
    pprint(dict(result))
    print('-----------------------------------------------------------------------')


def compare_with_json():
    # JSON serialization
    json_req = {
        'level': glog_pb2.INFO,
        'service_name': 'encode test',
        'message': 'test message',
        'timestamp': 11187097077,
        'trace_id': 'default',
    }
    ser_json = json.dumps(json_req).encode('utf-8')
    print(f'> JSON serialized into {len(ser_json)} bytes: {ser_json}')
    # serialization via Protobuf
    proto_req = glog_pb2.LogMessage(**json_req)
    binary_req = proto_req.SerializeToString()
    print(f'> Protobuf serialized into {len(binary_req)} bytes: {binary_req}')


if __name__ == '__main__':
    # serialize and deserialize protobuf messages
    print('\n' + serialize_and_deserialize.__name__)
    serialize_and_deserialize()
    # compare protobuf with json
    print('\n' + compare_with_json.__name__)
    compare_with_json()
