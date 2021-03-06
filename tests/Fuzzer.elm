module Fuzzer exposing (..)

-- DO NOT EDIT
-- AUTOGENERATED BY THE ELM PROTOCOL BUFFER COMPILER
-- https://github.com/gomeet/elm-protobuf
-- source file: fuzzer.proto

import Protobuf exposing (..)

import Json.Decode as JD
import Json.Encode as JE


type alias Fuzz =
    { stringField : String -- 1
    , int32Field : Int -- 2
    , stringValueField : Maybe String -- 3
    , int32ValueField : Maybe Int -- 4
    , timestampField : Maybe Timestamp -- 5
    }


fuzzDecoder : JD.Decoder Fuzz
fuzzDecoder =
    JD.lazy <| \_ -> decode Fuzz
        |> required "stringField" JD.string ""
        |> required "int32Field" intDecoder 0
        |> optional "stringValueField" stringValueDecoder
        |> optional "int32ValueField" intValueDecoder
        |> optional "timestampField" timestampDecoder


fuzzEncoder : Fuzz -> JE.Value
fuzzEncoder v =
    JE.object <| List.filterMap identity <|
        [ (requiredFieldEncoder "stringField" JE.string "" v.stringField)
        , (requiredFieldEncoder "int32Field" JE.int 0 v.int32Field)
        , (optionalEncoder "stringValueField" stringValueEncoder v.stringValueField)
        , (optionalEncoder "int32ValueField" intValueEncoder v.int32ValueField)
        , (optionalEncoder "timestampField" timestampEncoder v.timestampField)
        ]
