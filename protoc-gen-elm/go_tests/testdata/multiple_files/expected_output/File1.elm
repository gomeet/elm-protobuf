module File1 exposing (..)

-- DO NOT EDIT
-- AUTOGENERATED BY THE ELM PROTOCOL BUFFER COMPILER
-- https://github.com/gomeet/elm-protobuf
-- source file: file1.proto

import Protobuf exposing (..)

import Json.Decode as JD
import Json.Encode as JE


type alias File1Message =
    { field : Bool -- 1
    }


file1MessageDecoder : JD.Decoder File1Message
file1MessageDecoder =
    JD.lazy <| \_ -> decode File1Message
        |> required "field" JD.bool False


file1MessageEncoder : File1Message -> JE.Value
file1MessageEncoder v =
    JE.object <| List.filterMap identity <|
        [ (requiredFieldEncoder "field" JE.bool False v.field)
        ]
