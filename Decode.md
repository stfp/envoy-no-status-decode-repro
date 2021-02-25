# node

    HyperText Transfer Protocol 2
        Stream: HEADERS, Stream ID: 1, Length 103
            Length: 103
            Type: HEADERS (1)
            Flags: 0x05, End Headers, End Stream
                00.0 ..0. = Unused: 0x00
                ..0. .... = Priority: False
                .... 0... = Padded: False
                .... .1.. = End Headers: True
                .... ...1 = End Stream: True
            0... .... .... .... .... .... .... .... = Reserved: 0x0
            .000 0000 0000 0000 0000 0000 0000 0001 = Stream Identifier: 1
            [Pad Length: 0]
            Header Block Fragment: 40889acac8b21234da8f013540899acac8b5254207317fc3ffa5440a8f216e2a204e5b0b…
            [Header Length: 156]
            [Header Count: 3]
            Header: grpc-status: 5
                Name Length: 11
                Name: grpc-status
                Value Length: 1
                Value: 5
                [Unescaped: 5]
                Representation: Literal Header Field with Incremental Indexing - New Name
            Header: grpc-message: #%20node:%20here's%20an%20error%20message%20with%20spaces%20&%20some%20interesting%20characters!
                Name Length: 12
                Name: grpc-message
                Value Length: 96
                Value: #%20node:%20here's%20an%20error%20message%20with%20spaces%20&%20some%20interesting%20characters!
                [Unescaped: # node: here's an error message with spaces & some interesting characters!]
                Representation: Literal Header Field with Incremental Indexing - New Name
            Header: errorid: dummy
                Name Length: 7
                Name: errorid
                Value Length: 5
                Value: dummy
                [Unescaped: dummy]
                Representation: Literal Header Field with Incremental Indexing - New Name

# go

    Stream: HEADERS, Stream ID: 1, Length 87, 200 OK
        Length: 87
        Type: HEADERS (1)
        Flags: 0x05, End Headers, End Stream
            00.0 ..0. = Unused: 0x00
            ..0. .... = Priority: False
            .... 0... = Padded: False
            .... .1.. = End Headers: True
            .... ...1 = End Stream: True
        0... .... .... .... .... .... .... .... = Reserved: 0x0
        .000 0000 0000 0000 0000 0000 0000 0001 = Stream Identifier: 1
        [Pad Length: 0]
        Header Block Fragment: 885f8b1d75d0620d263d4c4d656440889acac8b21234da8f013540899acac8b525420731…
        [Header Length: 162]
        [Header Count: 4]
        Header: :status: 200 OK
            Name Length: 7
            Name: :status
            Value Length: 3
            Value: 200
            :status: 200
            [Unescaped: 200]
            Representation: Indexed Header Field
            Index: 8
        Header: content-type: application/grpc
            Name Length: 12
            Name: content-type
            Value Length: 16
            Value: application/grpc
            content-type: application/grpc
            [Unescaped: application/grpc]
            Representation: Literal Header Field with Incremental Indexing - Indexed Name
            Index: 31
        Header: grpc-status: 5
            Name Length: 11
            Name: grpc-status
            Value Length: 1
            Value: 5
            [Unescaped: 5]
            Representation: Literal Header Field with Incremental Indexing - New Name
        Header: grpc-message: # here's an error message with spaces & some interesting characters!
            Name Length: 12
            Name: grpc-message
            Value Length: 68
            Value: # here's an error message with spaces & some interesting characters!
            [Unescaped: # here's an error message with spaces & some interesting characters!]
            Representation: Literal Header Field with Incremental Indexing - New Name
